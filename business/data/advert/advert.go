package advert

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/baybaraandrey/advertising/foundation/database"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"
)

// Advert ...
type Advert struct {
	log *log.Logger
	db  *sqlx.DB
}

// New ...
func New(log *log.Logger, db *sqlx.DB) Advert {
	return Advert{
		log: log,
		db:  db,
	}
}

// Query retrieves a list of existing categories from the database.
func (a Advert) Query(ctx context.Context, traceID string, limit int, offset int, filters map[string][]string) ([]AdvertInfo, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "business.data.advert.query")
	defer span.End()

	allowedFilters := map[string]string{
		"category_uuid": "categories.uuid",
		"user_uuid":     "users.uuid",
		"is_active":     "adverts.is_active",
	}

	data := struct {
		Limit  int `db:"limit"`
		Offset int `db:"offset"`
	}{
		Limit:  limit,
		Offset: offset,
	}

	filterString := database.BuildFilterString(filters, allowedFilters)
	q := fmt.Sprintf(`
	SELECT 
		adverts.*,
		users.uuid "user.uuid",
		users.name "user.name",
		users.email "user.email",
		users.phone "user.phone",
		users.roles "user.roles",
		users.created "user.created",
		users.updated "user.updated",
	
		categories.uuid "category.uuid",
		categories.name "category.name",
		categories.created "category.created",
		categories.updated "category.updated"
	FROM adverts 
		INNER JOIN users ON adverts.user_uuid = users.uuid 
		INNER JOIN categories ON adverts.category_uuid = categories.uuid
	%s
	ORDER BY adverts.uuid ASC
	LIMIT :limit OFFSET :offset;`, filterString)

	log.Printf("%s : %s : query : %s", traceID, "advert.Query",
		database.Log(q),
	)

	adverts := []AdvertInfo{}
	if err := database.NamedQuerySlice(ctx, a.db, q, data, &adverts); err != nil {
		return nil, errors.Wrap(err, "selecting categories")
	}

	return adverts, nil
}

// Create creates new advert
func (a Advert) Create(ctx context.Context, traceID string, na NewAdvert, now time.Time) (Info, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "business.data.advert.create")
	defer span.End()

	adv := Info{
		ID:          uuid.New().String(),
		UserID:      na.UserID,
		CategoryID:  na.CategoryID,
		Title:       na.Title,
		Description: na.Description,
		Location:    na.Location,
		Price:       na.Price,
		IsActive:    true,
		Created:     now.UTC(),
		Updated:     now.UTC(),
	}

	const q = `INSERT INTO adverts
	(uuid, user_uuid, category_uuid, title, description, location, price, is_active, created, updated)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	log.Printf("%s : %s : query : %s", traceID, "advert.Create",
		database.Log(q, adv.ID, adv.UserID, adv.CategoryID, adv.Title, adv.Description, adv.Location, adv.Price, adv.IsActive, adv.Created, adv.Updated),
	)

	if _, err := a.db.ExecContext(ctx, q, adv.ID, adv.UserID, adv.CategoryID, adv.Title, adv.Description, adv.Location, adv.Price, adv.IsActive, adv.Created, adv.Updated); err != nil {
		return Info{}, errors.Wrap(err, "creating advert")
	}

	return adv, nil
}

// Query retrieves a total count of adverts from the database.
func (a Advert) TotalActive(ctx context.Context, traceID string) (int, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "business.data.advert.total")
	defer span.End()

	q := fmt.Sprintf(`
	SELECT 
		count(*) as c
	FROM adverts 
		INNER JOIN users ON adverts.user_uuid = users.uuid 
		INNER JOIN categories ON adverts.category_uuid = categories.uuid
	WHERE adverts.is_active IN ('true')
	`)

	log.Printf("%s : %s : query : %s", traceID, "advert.Total",
		database.Log(q),
	)

	count := struct {
		Count int `db:"c"`
	}{}

	if err := a.db.Get(&count, q); err != nil {
		return 0, errors.Wrap(err, "selecting adverts total count")
	}

	return count.Count, nil
}
