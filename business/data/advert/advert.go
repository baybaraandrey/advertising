package advert

import (
	"context"
	"fmt"
	"log"

	"github.com/baybaraandrey/advertising/foundation/database"
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
