package handlers

import (
	"context"
	"net/http"

	"github.com/baybaraandrey/advertising/business/data/advert"
	"github.com/baybaraandrey/advertising/foundation/web"
	"go.opentelemetry.io/otel/trace"
)

type advertGroup struct {
	advert advert.Advert
}

func (ag advertGroup) query(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "handlers.user.list")
	defer span.End()

	v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return web.NewShutdownError("web value missing from context")
	}

	adverts, err := ag.advert.Query(ctx, v.TraceID)
	if err != nil {
		return err
	}

	return web.Respond(ctx, w, adverts, http.StatusOK)
}
