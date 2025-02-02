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

var CommunitySentenceLikes = newCommunitySentenceLikesTable("public", "community_sentence_likes", "")

type communitySentenceLikesTable struct {
	postgres.Table

	// Columns
	ID         postgres.ColumnString
	UserID     postgres.ColumnString
	SentenceID postgres.ColumnString
	CreatedAt  postgres.ColumnTimestampz

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type CommunitySentenceLikesTable struct {
	communitySentenceLikesTable

	EXCLUDED communitySentenceLikesTable
}

// AS creates new CommunitySentenceLikesTable with assigned alias
func (a CommunitySentenceLikesTable) AS(alias string) *CommunitySentenceLikesTable {
	return newCommunitySentenceLikesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new CommunitySentenceLikesTable with assigned schema name
func (a CommunitySentenceLikesTable) FromSchema(schemaName string) *CommunitySentenceLikesTable {
	return newCommunitySentenceLikesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new CommunitySentenceLikesTable with assigned table prefix
func (a CommunitySentenceLikesTable) WithPrefix(prefix string) *CommunitySentenceLikesTable {
	return newCommunitySentenceLikesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new CommunitySentenceLikesTable with assigned table suffix
func (a CommunitySentenceLikesTable) WithSuffix(suffix string) *CommunitySentenceLikesTable {
	return newCommunitySentenceLikesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newCommunitySentenceLikesTable(schemaName, tableName, alias string) *CommunitySentenceLikesTable {
	return &CommunitySentenceLikesTable{
		communitySentenceLikesTable: newCommunitySentenceLikesTableImpl(schemaName, tableName, alias),
		EXCLUDED:                    newCommunitySentenceLikesTableImpl("", "excluded", ""),
	}
}

func newCommunitySentenceLikesTableImpl(schemaName, tableName, alias string) communitySentenceLikesTable {
	var (
		IDColumn         = postgres.StringColumn("id")
		UserIDColumn     = postgres.StringColumn("user_id")
		SentenceIDColumn = postgres.StringColumn("sentence_id")
		CreatedAtColumn  = postgres.TimestampzColumn("created_at")
		allColumns       = postgres.ColumnList{IDColumn, UserIDColumn, SentenceIDColumn, CreatedAtColumn}
		mutableColumns   = postgres.ColumnList{IDColumn, CreatedAtColumn}
	)

	return communitySentenceLikesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:         IDColumn,
		UserID:     UserIDColumn,
		SentenceID: SentenceIDColumn,
		CreatedAt:  CreatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
