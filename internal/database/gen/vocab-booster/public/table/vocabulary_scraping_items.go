//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var VocabularyScrapingItems = newVocabularyScrapingItemsTable("public", "vocabulary_scraping_items", "")

type vocabularyScrapingItemsTable struct {
	postgres.Table

	// Columns
	ID        postgres.ColumnString
	Term      postgres.ColumnString
	CreatedAt postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type VocabularyScrapingItemsTable struct {
	vocabularyScrapingItemsTable

	EXCLUDED vocabularyScrapingItemsTable
}

// AS creates new VocabularyScrapingItemsTable with assigned alias
func (a VocabularyScrapingItemsTable) AS(alias string) *VocabularyScrapingItemsTable {
	return newVocabularyScrapingItemsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new VocabularyScrapingItemsTable with assigned schema name
func (a VocabularyScrapingItemsTable) FromSchema(schemaName string) *VocabularyScrapingItemsTable {
	return newVocabularyScrapingItemsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new VocabularyScrapingItemsTable with assigned table prefix
func (a VocabularyScrapingItemsTable) WithPrefix(prefix string) *VocabularyScrapingItemsTable {
	return newVocabularyScrapingItemsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new VocabularyScrapingItemsTable with assigned table suffix
func (a VocabularyScrapingItemsTable) WithSuffix(suffix string) *VocabularyScrapingItemsTable {
	return newVocabularyScrapingItemsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newVocabularyScrapingItemsTable(schemaName, tableName, alias string) *VocabularyScrapingItemsTable {
	return &VocabularyScrapingItemsTable{
		vocabularyScrapingItemsTable: newVocabularyScrapingItemsTableImpl(schemaName, tableName, alias),
		EXCLUDED:                     newVocabularyScrapingItemsTableImpl("", "excluded", ""),
	}
}

func newVocabularyScrapingItemsTableImpl(schemaName, tableName, alias string) vocabularyScrapingItemsTable {
	var (
		IDColumn        = postgres.StringColumn("id")
		TermColumn      = postgres.StringColumn("term")
		CreatedAtColumn = postgres.TimestampzColumn("created_at")
		allColumns      = postgres.ColumnList{IDColumn, TermColumn, CreatedAtColumn}
		mutableColumns  = postgres.ColumnList{TermColumn, CreatedAtColumn}
	)

	return vocabularyScrapingItemsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		Term:      TermColumn,
		CreatedAt: CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
