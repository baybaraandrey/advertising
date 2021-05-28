package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/baybaraandrey/advertising/business/data/advert"
	"github.com/baybaraandrey/advertising/business/sys/validate"
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

	offset := r.URL.Query().Get("offset")
	offsetNumber, err := strconv.Atoi(offset)
	if err != nil {
		return validate.NewRequestError(fmt.Errorf("invalid offset format: %s", offset), http.StatusBadRequest)
	}

	limit := r.URL.Query().Get("limit")
	limitNumber, err := strconv.Atoi(limit)
	if err != nil {
		return validate.NewRequestError(fmt.Errorf("invalid limit format: %s", limit), http.StatusBadRequest)
	}
	r.URL.Query().Del("offset")
	r.URL.Query().Del("limit")

	filters := map[string][]string(r.URL.Query())
	adverts, err := ag.advert.Query(ctx, v.TraceID, limitNumber, offsetNumber, filters)
	if err != nil {
		return err
	}

	return web.Respond(ctx, w, adverts, http.StatusOK)
}
