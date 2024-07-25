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

var Exercises = newExercisesTable("public", "exercises", "")

type exercisesTable struct {
	postgres.Table

	// Columns
	ID                  postgres.ColumnString
	VocabularyExampleID postgres.ColumnString
	Audio               postgres.ColumnString
	Level               postgres.ColumnString
	Content             postgres.ColumnString
	Vocabulary          postgres.ColumnString
	CorrectAnswer       postgres.ColumnString
	Options             postgres.ColumnString
	CreatedAt           postgres.ColumnTimestampz
	Frequency           postgres.ColumnFloat

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ExercisesTable struct {
	exercisesTable

	EXCLUDED exercisesTable
}

// AS creates new ExercisesTable with assigned alias
func (a ExercisesTable) AS(alias string) *ExercisesTable {
	return newExercisesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ExercisesTable with assigned schema name
func (a ExercisesTable) FromSchema(schemaName string) *ExercisesTable {
	return newExercisesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ExercisesTable with assigned table prefix
func (a ExercisesTable) WithPrefix(prefix string) *ExercisesTable {
	return newExercisesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ExercisesTable with assigned table suffix
func (a ExercisesTable) WithSuffix(suffix string) *ExercisesTable {
	return newExercisesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newExercisesTable(schemaName, tableName, alias string) *ExercisesTable {
	return &ExercisesTable{
		exercisesTable: newExercisesTableImpl(schemaName, tableName, alias),
		EXCLUDED:       newExercisesTableImpl("", "excluded", ""),
	}
}

func newExercisesTableImpl(schemaName, tableName, alias string) exercisesTable {
	var (
		IDColumn                  = postgres.StringColumn("id")
		VocabularyExampleIDColumn = postgres.StringColumn("vocabulary_example_id")
		AudioColumn               = postgres.StringColumn("audio")
		LevelColumn               = postgres.StringColumn("level")
		ContentColumn             = postgres.StringColumn("content")
		VocabularyColumn          = postgres.StringColumn("vocabulary")
		CorrectAnswerColumn       = postgres.StringColumn("correct_answer")
		OptionsColumn             = postgres.StringColumn("options")
		CreatedAtColumn           = postgres.TimestampzColumn("created_at")
		FrequencyColumn           = postgres.FloatColumn("frequency")
		allColumns                = postgres.ColumnList{IDColumn, VocabularyExampleIDColumn, AudioColumn, LevelColumn, ContentColumn, VocabularyColumn, CorrectAnswerColumn, OptionsColumn, CreatedAtColumn, FrequencyColumn}
		mutableColumns            = postgres.ColumnList{VocabularyExampleIDColumn, AudioColumn, LevelColumn, ContentColumn, VocabularyColumn, CorrectAnswerColumn, OptionsColumn, CreatedAtColumn, FrequencyColumn}
	)

	return exercisesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:                  IDColumn,
		VocabularyExampleID: VocabularyExampleIDColumn,
		Audio:               AudioColumn,
		Level:               LevelColumn,
		Content:             ContentColumn,
		Vocabulary:          VocabularyColumn,
		CorrectAnswer:       CorrectAnswerColumn,
		Options:             OptionsColumn,
		CreatedAt:           CreatedAtColumn,
		Frequency:           FrequencyColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
