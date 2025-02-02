package infrastructure

import (
	"database/sql"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/model"
	"github.com/namhq1989/vocab-booster-english-hub/internal/database/gen/vocab-booster/public/table"
	"github.com/namhq1989/vocab-booster-english-hub/internal/utils/manipulation"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/domain"
	"github.com/namhq1989/vocab-booster-english-hub/pkg/exercise/infrastructure/mapping"
	"github.com/namhq1989/vocab-booster-utilities/appcontext"
)

type UserExerciseStatusRepository struct {
	db *database.Database
}

func NewUserExerciseStatusRepository(db *database.Database) UserExerciseStatusRepository {
	r := UserExerciseStatusRepository{
		db: db,
	}
	return r
}

func (r UserExerciseStatusRepository) getDB() *sql.DB {
	return r.db.GetDB()
}

func (UserExerciseStatusRepository) getTable() *table.UserExerciseStatusesTable {
	return table.UserExerciseStatuses
}

func (r UserExerciseStatusRepository) CreateUserExerciseStatus(ctx *appcontext.AppContext, status domain.UserExerciseStatus) error {
	mapper := mapping.UserExerciseStatusMapper{}
	doc, err := mapper.FromDomainToModel(status)
	if err != nil {
		return err
	}

	stmt := r.getTable().INSERT(
		r.getTable().AllColumns,
	).
		MODEL(doc)

	_, err = stmt.ExecContext(ctx.Context(), r.getDB())
	return err
}

func (r UserExerciseStatusRepository) UpdateUserExerciseStatus(ctx *appcontext.AppContext, status domain.UserExerciseStatus) error {
	mapper := mapping.UserExerciseStatusMapper{}
	doc, err := mapper.FromDomainToModel(status)
	if err != nil {
		return err
	}

	stmt := r.getTable().UPDATE(
		r.getTable().AllColumns,
	).
		MODEL(doc).
		WHERE(
			r.getTable().ID.EQ(postgres.String(doc.ID)),
		)

	_, err = stmt.ExecContext(ctx.Context(), r.getDB())
	return err
}

func (r UserExerciseStatusRepository) FindUserExerciseStatus(ctx *appcontext.AppContext, exerciseID, userID string) (*domain.UserExerciseStatus, error) {
	stmt := postgres.SELECT(
		r.getTable().AllColumns,
	).
		FROM(r.getTable()).
		WHERE(
			r.getTable().ExerciseID.EQ(postgres.String(exerciseID)).
				AND(r.getTable().UserID.EQ(postgres.String(userID))),
		)

	var doc model.UserExerciseStatuses
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &doc); err != nil {
		if r.db.IsNoRowsError(err) {
			return nil, nil
		}
		return nil, err
	}

	var (
		mapper    = mapping.UserExerciseStatusMapper{}
		result, _ = mapper.FromModelToDomain(doc)
	)
	return result, nil
}

func (r UserExerciseStatusRepository) CountUserReadyToReviewExercises(ctx *appcontext.AppContext, userID, timezone string) (int64, error) {
	var now = manipulation.Now(timezone)

	stmt := postgres.SELECT(
		postgres.COUNT(r.getTable().ID).AS("count_result.total"),
	).
		FROM(r.getTable()).
		WHERE(
			r.getTable().UserID.EQ(postgres.String(userID)).
				AND(r.getTable().NextReviewAt.LT(postgres.TimestampzT(now))),
		)

	var result = database.CountResult{}
	err := stmt.QueryContext(ctx.Context(), r.getDB(), &result)
	return result.Total, err
}

func (r UserExerciseStatusRepository) FindUserReadyToReviewExercises(ctx *appcontext.AppContext, filter domain.UserExerciseFilter) ([]domain.UserExercise, error) {
	var (
		ues = r.getTable().AS("ues")
		e   = table.Exercises.AS("e")
		now = manipulation.Now(filter.Timezone)
	)

	whereCond := postgres.AND(ues.UserID.EQ(postgres.String(filter.UserID)))
	whereCond = whereCond.AND(ues.NextReviewAt.LT(postgres.TimestampzT(now)))

	stmt := postgres.SELECT(
		e.ID, e.Level, e.Audio, e.Vocabulary, e.Content, e.CorrectAnswer, e.Options,
		ues.CorrectStreak, ues.IsFavorite, ues.IsMastered, ues.NextReviewAt,
	).
		FROM(
			ues.LEFT_JOIN(e, ues.ExerciseID.EQ(e.ID)),
		).
		WHERE(whereCond).
		ORDER_BY(
			ues.NextReviewAt,
		).
		LIMIT(filter.NumOfExercises)

	return r.queryExercisesWithStmt(ctx, stmt, filter.Lang)
}

func (r UserExerciseStatusRepository) FindUserFavoriteExercises(ctx *appcontext.AppContext, filter domain.UserFavoriteExerciseFilter) ([]domain.UserExercise, error) {
	var (
		ues = r.getTable().AS("ues")
		e   = table.Exercises.AS("e")
		now = manipulation.NowUTC()
	)

	whereCond := postgres.AND(ues.UserID.EQ(postgres.String(filter.UserID)))
	whereCond = whereCond.AND(ues.IsFavorite.EQ(postgres.Bool(true)))
	whereCond = whereCond.AND(ues.UpdatedAt.LT(postgres.TimestampzT(now)))

	stmt := postgres.SELECT(
		e.ID, e.Level, e.Audio, e.Vocabulary, e.Content, e.CorrectAnswer, e.Options,
		ues.CorrectStreak, ues.IsFavorite, ues.IsMastered, ues.NextReviewAt,
	).
		FROM(
			ues.LEFT_JOIN(e, ues.ExerciseID.EQ(e.ID)),
		).
		WHERE(whereCond).
		ORDER_BY(
			ues.NextReviewAt,
		).
		LIMIT(filter.NumOfExercises)

	return r.queryExercisesWithStmt(ctx, stmt, filter.Lang)
}

func (r UserExerciseStatusRepository) queryExercisesWithStmt(ctx *appcontext.AppContext, stmt postgres.SelectStatement, lang string) ([]domain.UserExercise, error) {
	var (
		docs   = make([]mapping.UserExercise, 0)
		result = make([]domain.UserExercise, 0)
	)
	if err := stmt.QueryContext(ctx.Context(), r.getDB(), &docs); err != nil {
		if r.db.IsNoRowsError(err) {
			return result, nil
		}
		return result, err
	}

	var (
		mapper = mapping.UserExerciseMapper{}
	)
	for _, doc := range docs {
		ue, _ := mapper.FromModelToDomain(doc, lang)
		result = append(result, *ue)
	}
	return result, nil
}

type statsResult struct {
	Mastered      int64 `json:"mastered"`
	ReadyToReview int64 `json:"ready_to_review"`
}

func (r UserExerciseStatusRepository) FindUserStats(ctx *appcontext.AppContext, userID, tz string) (*domain.UserStats, error) {
	stmt := postgres.RawStatement(
		`SELECT
    				COUNT(CASE ues.is_mastered WHEN TRUE::boolean THEN 1 END) AS "stats_result.mastered",
    				COUNT(CASE WHEN ues.next_review_at < $ts::timestamp THEN 1 END) AS "stats_result.ready_to_review"
				  FROM public.user_exercise_statuses AS ues
				  WHERE ues.user_id = $userID::text;`,
		postgres.RawArgs{
			"$ts":     manipulation.ToSQLTimestamp(manipulation.Now(tz), tz),
			"$userID": userID,
		},
	)

	var result = statsResult{}
	err := stmt.QueryContext(ctx.Context(), r.getDB(), &result)
	if err != nil {
		return nil, err
	}

	return &domain.UserStats{
		Mastered:         int(result.Mastered),
		WaitingForReview: int(result.ReadyToReview),
	}, nil
}
