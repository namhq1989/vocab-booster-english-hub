package main

import (
	"os"

	"github.com/go-jet/jet/v2/generator/metadata"
	"github.com/go-jet/jet/v2/generator/postgres"
	"github.com/go-jet/jet/v2/generator/template"
	postgres2 "github.com/go-jet/jet/v2/postgres"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
)

func main() {
	conn, err := parsePostgresConnectionString(os.Getenv("POSTGRES_CONN"))
	if err != nil {
		panic(err)
	}

	if err = postgres.Generate(
		"internal/database/gen",
		*conn,
		template.Default(postgres2.Dialect).
			UseSchema(func(schema metadata.Schema) template.Schema {
				return template.DefaultSchema(schema).
					UseModel(template.DefaultModel().
						UseTable(func(table metadata.Table) template.TableModel {
							return template.DefaultTableModel(table).
								UseField(func(column metadata.Column) template.TableModelField {
									field := template.DefaultTableModelField(column)

									if schema.Name == "public" {
										switch table.Name {
										case "vocabularies":
											field = vocabularies(field, column)
										case "vocabulary_examples":
											field = vocabularyExamples(field, column)
										case "exercises":
											field = exercises(field, column)
										case "community_sentences":
											field = communitySentence(field, column)
										case "community_sentence_drafts":
											field = communitySentenceDraft(field, column)
										}
									}

									return field
								})
						}),
					)
			}),
	); err != nil {
		panic(err)
	}
}

func vocabularies(field template.TableModelField, column metadata.Column) template.TableModelField {
	switch column.Name {
	case "parts_of_speech", "synonyms", "antonyms":
		field.Type = template.NewType(database.StringArray{})
	}

	return field
}

func vocabularyExamples(field template.TableModelField, _ metadata.Column) template.TableModelField {
	return field
}

func exercises(field template.TableModelField, column metadata.Column) template.TableModelField {
	if column.Name == "options" {
		field.Type = template.NewType(database.StringArray{})
	}

	return field
}

func communitySentence(field template.TableModelField, column metadata.Column) template.TableModelField {
	if column.Name == "required_vocabulary" {
		field.Type = template.NewType(database.StringArray{})
	}

	return field
}

func communitySentenceDraft(field template.TableModelField, column metadata.Column) template.TableModelField {
	if column.Name == "required_vocabulary" {
		field.Type = template.NewType(database.StringArray{})
	}

	return field
}
