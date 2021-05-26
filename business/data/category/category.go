package category

import (
	"context"
	"log"
	"time"

	"github.com/baybaraandrey/advertising/foundation/database"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"
)

//  Category ...
type Category struct {
	log *log.Logger
	db  *sqlx.DB
}

// New constructs a User for api access.
func New(log *log.Logger, db *sqlx.DB) Category {
	return Category{
		log: log,
		db:  db,
	}
}

// Create inserts a new category into the database.
func (c Category) Create(ctx context.Context, traceID string, nc NewCategory, now time.Time) (Info, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "internal.data.user.create")
	defer span.End()

	category := Info{
		ID:      uuid.New().String(),
		Name:    nc.Name,
		Created: now.UTC(),
		Updated: now.UTC(),
	}

	const q = `INSERT INTO categories
	(uuid, name, created, updated)
	VALUES ($1, $2, $3, $4)`

	c.log.Printf("%s : %s : query : %s", traceID, "category.Create",
		database.Log(q, category.ID, category.Name),
	)

	if _, err := c.db.ExecContext(ctx, q, category.ID, category.Name, category.Created, category.Updated); err != nil {
		return Info{}, errors.Wrap(err, "inserting category")
	}

	return category, nil

}

// Query retrieves a list of existing categories from the database.
func (c Category) Query(ctx context.Context, traceID string) ([]Info, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "business.data.category.query")
	defer span.End()

	const q = `SELECT * FROM categories`

	categories := []Info{}
	if err := c.db.SelectContext(ctx, &categories, q); err != nil {
		return nil, errors.Wrap(err, "selecting categories")
	}

	return categories, nil
}
