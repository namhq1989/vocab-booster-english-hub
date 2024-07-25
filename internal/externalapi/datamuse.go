package externalapi

import (
	"fmt"
	"strings"
	"sync"
	"unicode"

	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type DatamuseSearchTermResult struct {
	Definitions []DatamuseTermDefinition
	Frequency   float64
	Ipa         string
	Synonyms    []string
	Antonyms    []string
}

type DatamuseTermDefinition struct {
	Pos        string
	Definition string
}

func (ea ExternalAPI) SearchTermWitDatamuse(ctx *appcontext.AppContext, term string) (*DatamuseSearchTermResult, error) {
	var (
		result = &DatamuseSearchTermResult{}
		err    error
		wg     sync.WaitGroup
	)

	wg.Add(3)

	go func() {
		defer wg.Done()
		if e := ea.searchTermDefinitionWitDatamuse(ctx, term, result); e != nil {
			err = e
		}
	}()

	go func() {
		defer wg.Done()
		if e := ea.searchTermSynonymsWitDatamuse(ctx, term, result); e != nil {
			err = e
		}
	}()

	go func() {
		defer wg.Done()
		if e := ea.searchTermAntonymsWitDatamuse(ctx, term, result); e != nil {
			err = e
		}
	}()

	wg.Wait()

	return result, err
}

type datamuseApiSearchTermResult struct {
	Word        string   `json:"word"`
	Definitions []string `json:"defs"`
	Tags        []string `json:"tags"`
}

func uncapitalizeDefinition(s string) string {
	// trim spaces and the period at the end
	s = strings.TrimSpace(s)
	s = strings.TrimSuffix(s, ".")

	// find the position of the first character after the annotation
	closingParenIndex := strings.LastIndex(s, ")")
	if closingParenIndex != -1 && closingParenIndex < len(s)-1 {
		annotation := s[:closingParenIndex+1]
		definition := strings.TrimSpace(s[closingParenIndex+1:])
		if len(definition) > 0 {
			for i, v := range definition {
				return annotation + " " + string(unicode.ToLower(v)) + definition[i+1:]
			}
		}
	} else {
		// no annotation, uncapitalize the whole string
		definition := strings.TrimSpace(s)
		if len(definition) > 0 {
			for i, v := range definition {
				return string(unicode.ToLower(v)) + definition[i+1:]
			}
		}
	}
	return s // if the string is empty, return it as is
}

func (ea ExternalAPI) searchTermDefinitionWitDatamuse(ctx *appcontext.AppContext, term string, result *DatamuseSearchTermResult) error {
	var apiResults []datamuseApiSearchTermResult

	_, err := ea.datamuse.R().
		SetQueryParams(map[string]string{
			"sp":  term,
			"qe":  "sp",
			"md":  "drf",
			"ipa": "1",
			"max": "1",
		}).
		SetResult(&apiResults).
		Get("/words")

	if err != nil {
		ctx.Logger().Error("[externalapi] error when searching term with datamuse", err, appcontext.Fields{"term": term})
		return err
	}

	if len(apiResults) == 0 {
		ctx.Logger().Error("[externalapi] datamuse api search result is empty", err, appcontext.Fields{"term": term})
		return nil
	}

	var apiResult = apiResults[0]

	if apiResult.Word != term {
		ctx.Logger().Error("[externalapi] result's word doesn't match term", err, appcontext.Fields{"term": term, "word": apiResult.Word})
		return nil
	}

	if len(apiResult.Definitions) == 0 {
		ctx.Logger().Error("[externalapi] result's definitions is empty", err, appcontext.Fields{"term": term})
		return nil
	}

	for _, def := range apiResult.Definitions {
		parts := strings.Split(def, "\t")
		if len(parts) == 2 {
			pos := parts[0]
			definition := strings.TrimSpace(parts[1])
			definition = strings.TrimSuffix(definition, ".")
			definition = uncapitalizeDefinition(definition)
			result.Definitions = append(result.Definitions, DatamuseTermDefinition{
				Pos:        pos,
				Definition: definition,
			})
		}
	}

	// Process tags
	for _, tag := range apiResult.Tags {
		if strings.HasPrefix(tag, "ipa_pron:") {
			result.Ipa = strings.Split(tag, ":")[1]
		} else if strings.HasPrefix(tag, "f:") {
			_, _ = fmt.Sscanf(strings.Split(tag, ":")[1], "%f", &result.Frequency)
		}
	}

	return nil
}

type datamuseApiSearchSemanticRelationsResult struct {
	Word string `json:"word"`
}

func (ea ExternalAPI) searchTermSynonymsWitDatamuse(ctx *appcontext.AppContext, term string, result *DatamuseSearchTermResult) error {
	var synonyms []datamuseApiSearchSemanticRelationsResult

	_, err := ea.datamuse.R().
		SetQueryParams(map[string]string{
			"rel_syn": term,
			"max":     "5",
		}).
		SetResult(&synonyms).
		Get("/words")

	if err != nil {
		ctx.Logger().Error("[externalapi] error when searching synonyms with datamuse", err, appcontext.Fields{"term": term})
		return err
	}

	if len(synonyms) == 0 {
		ctx.Logger().Error("[externalapi] datamuse api search synonyms result is empty", err, appcontext.Fields{"term": term})
		return nil
	}

	for _, item := range synonyms {
		result.Synonyms = append(result.Synonyms, item.Word)
	}

	return nil
}

func (ea ExternalAPI) searchTermAntonymsWitDatamuse(ctx *appcontext.AppContext, term string, result *DatamuseSearchTermResult) error {
	var antonyms []datamuseApiSearchSemanticRelationsResult

	_, err := ea.datamuse.R().
		SetQueryParams(map[string]string{
			"rel_ant": term,
			"max":     "5",
		}).
		SetResult(&antonyms).
		Get("/words")

	if err != nil {
		ctx.Logger().Error("[externalapi] error when searching antonyms with datamuse", err, appcontext.Fields{"term": term})
		return err
	}

	if len(antonyms) == 0 {
		ctx.Logger().Error("[externalapi] datamuse api search antonyms result is empty", err, appcontext.Fields{"term": term})
		return nil
	}

	for _, item := range antonyms {
		result.Antonyms = append(result.Antonyms, item.Word)
	}

	return nil
}
