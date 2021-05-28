package advert

import (
	"context"
	"log"

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
func (c Advert) Query(ctx context.Context, traceID string) ([]Info, error) {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "business.data.advert.query")
	defer span.End()

	const q = `
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
	ORDER BY  adverts.uuid ASC
	;`

	adverts := []Info{}
	if err := c.db.SelectContext(ctx, &adverts, q); err != nil {
		return nil, errors.Wrap(err, "selecting categories")
	}

	return adverts, nil
}
