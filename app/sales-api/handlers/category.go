package handlers

import (
	"context"
	"net/http"

	"github.com/baybaraandrey/advertising/business/data/category"
	"github.com/baybaraandrey/advertising/foundation/web"
	"github.com/pkg/errors"

	"go.opentelemetry.io/otel/trace"
)

type categoryGroup struct {
	category category.Category
}

func (cg categoryGroup) create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "handlers.category.list")
	defer span.End()

	v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return web.NewShutdownError("web value missing from context")
	}

	var nc category.NewCategory
	if err := web.Decode(r, &nc); err != nil {
		return errors.Wrap(err, "")
	}

	c, err := cg.category.Create(ctx, v.TraceID, nc, v.Now)
	if err != nil {
		return errors.Wrapf(err, "Category: %+v", &c)
	}

	return web.Respond(ctx, w, c, http.StatusCreated)
}

func (cg categoryGroup) query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "handlers.category.list")
	defer span.End()

	v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return web.NewShutdownError("web value missing from context")
	}

	categories, err := cg.category.Query(ctx, v.TraceID)
	if err != nil {
		return err
	}

	return web.Respond(ctx, w, categories, http.StatusOK)
}
