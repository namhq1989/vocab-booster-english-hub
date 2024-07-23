package infrastructure

import (
	"database/sql"
	"time"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/table"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/infrastructure/mapping"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type UserExerciseInteractedHistoryRepository struct {
	db *database.Database
}

func NewUserExerciseInteractedHistoryRepository(db *database.Database) UserExerciseInteractedHistoryRepository {
	r := UserExerciseInteractedHistoryRepository{
		db: db,
	}
	return r
}

func (r UserExerciseInteractedHistoryRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (UserExerciseInteractedHistoryRepository) getTable() *table.UserExerciseInteractedHistoriesTable {
	return table.UserExerciseInteractedHistories
}

func (r UserExerciseInteractedHistoryRepository) UpsertUserExerciseInteractedHistory(ctx *appcontext.AppContext, history domain.UserExerciseInteractedHistory) error {
	mapper := mapping.UserExerciseInteractedHistoryMapper{}
	doc, err := mapper.FromDomainToModel(history)
	if err != nil {
		return err
	}

	stmt := r.getTable().INSERT(r.getTable().AllColumns).
		MODEL(doc).
		ON_CONFLICT(r.getTable().UserID, r.getTable().UserID, r.getTable().ExerciseID, r.getTable().Date).DO_NOTHING()

	_, err = stmt.ExecContext(ctx.Context(), r.getDB())
	return err
}

func (r UserExerciseInteractedHistoryRepository) AggregateUserExercisesInTimeRange(ctx *appcontext.AppContext, userID, timezone string, from, to time.Time) ([]domain.UserAggregatedExercise, error) {
	stmt := postgres.RawStatement(
		`WITH RECURSIVE date_series AS (
						SELECT 
							$from_date::DATE AS date
						UNION ALL
						SELECT 
						 	(date + INTERVAL '1 day')::DATE
						FROM 
							date_series
						WHERE 
							date + INTERVAL '1 day' <= $to_date::DATE
					)
					SELECT 
						TO_CHAR(date_series.date, 'DD/MM') AS "user_aggregated_exercise.date",
						COALESCE(COUNT(ueih.exercise_id), 0) AS "user_aggregated_exercise.exercise"
					FROM 
						date_series
					LEFT JOIN 
						user_exercise_interacted_histories ueih
					ON 
						DATE_TRUNC('day', ueih.date AT TIME ZONE $timezone) = date_series.date
						AND ueih.user_id = $userID::text
					GROUP BY 
						date_series.date
					ORDER BY 
						date_series.date;`,
		postgres.RawArgs{
			"$from_date": manipulation.ToSQLDateFrom(from, timezone),
			"$to_date":   manipulation.ToSQLDateTo(to, timezone),
			"$timezone":  timezone,
			"$userID":    userID,
		},
	)

	var (
		result = make([]domain.UserAggregatedExercise, 0)
		docs   = make([]mapping.UserAggregatedExercise, 0)
	)
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &docs); err != nil {
		if r.db.IsNoRowsError(err) {
			return result, nil
		}
		return result, err
	}

	var (
		mapper = mapping.UserAggregatedExerciseMapper{}
	)
	for _, doc := range docs {
		uae, _ := mapper.FromModelToDomain(doc)
		result = append(result, *uae)
	}
	return result, nil
}
