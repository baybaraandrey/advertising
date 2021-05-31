package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/baybaraandrey/advertising/business/data/advert"
	"github.com/baybaraandrey/advertising/business/sys/validate"
	"github.com/baybaraandrey/advertising/foundation/web"
	"github.com/pkg/errors"
	"go.opentelemetry.io/otel/trace"
)

type advertGroup struct {
	advert advert.Advert
}

// @Summary get adverts
// @Description get all adverts
// @Produce  json
// @Success 200 {object} paginatedLimitOffsetAdvertResponse
// @Router /v1/adverts/ [get]
// @Tags advert
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

	total, err := ag.advert.TotalActive(ctx, v.TraceID)
	if err != nil {
		return err
	}

	data := paginatedLimitOffsetAdvertResponse{
		PaginatedLimitOffsetResponse: &web.PaginatedLimitOffsetResponse{
			Limit:   limitNumber,
			Offset:  offsetNumber,
			Records: len(adverts),
			Total:   total,
		},
		Results: adverts,
	}

	return web.Respond(ctx, w, data, http.StatusOK)
}

// @Summary create an advert
// @Description create an advert
// @Accept  json
// @Produce  json
// @Param query body advert.NewAdvert true "create an advert"
// @Success 201 {object} advert.Info
// @Router /v1/adverts/ [post]
// @Tags advert
func (ag advertGroup) create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "handlers.advert.create")
	defer span.End()

	v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return web.NewShutdownError("web value missing from context")
	}

	var na advert.NewAdvert
	if err := web.Decode(r, &na); err != nil {
		return errors.Wrap(err, "")
	}

	adv, err := ag.advert.Create(ctx, v.TraceID, na, v.Now)
	if err != nil {
		return errors.Wrapf(err, "Advert: %+v", &adv)
	}

	return web.Respond(ctx, w, adv, http.StatusCreated)
}

// @Summary get an advert by id
// @Descriptionget an advert by id
// @Accept  json
// @Produce  json
// @Success 200 {object} advert.AdvertInfo
// @Param id path int true "Advert ID"
// @Router /v1/adverts/{id}/ [get]
// @Tags advert
func (ag advertGroup) queryByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "handlers.advert.retrieve")
	defer span.End()

	v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return web.NewShutdownError("web value missing from context")
	}

	id := web.Param(r, "id")
	usr, err := ag.advert.QueryByID(ctx, v.TraceID, id)
	if err != nil {
		switch err {
		case advert.ErrInvalidID:
			return web.NewRequestError(err, http.StatusBadRequest)
		case advert.ErrNotFound:
			return web.NewRequestError(err, http.StatusNotFound)
		default:
			return errors.Wrapf(err, "Id: %s", id)
		}
	}

	return web.Respond(ctx, w, usr, http.StatusOK)
}

// @Summary deactivate an advert
// @Description deactivate an advert
// @Accept  json
// @Produce  json
// @Success 200 {object} advert.AdvertInfo
// @Param id path int true "Advert ID"
// @Router /v1/adverts/{id}/deactivate/ [post]
// @Tags advert
func (ag advertGroup) deactivate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "handlers.advert.deactivate")
	defer span.End()

	v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return web.NewShutdownError("web value missing from context")
	}

	id := web.Param(r, "id")

	adv, err := ag.advert.Deactivate(ctx, v.TraceID, id)
	if err != nil {
		return errors.Wrapf(err, "Advert: %+v", &adv)
	}

	return web.Respond(ctx, w, adv, http.StatusOK)
}

// @Summary activate an advert
// @Description activate an advert
// @Accept  json
// @Produce  json
// @Success 200 {object} advert.AdvertInfo
// @Param id path int true "Advert ID"
// @Router /v1/adverts/{id}/activate/ [post]
// @Tags advert
func (ag advertGroup) activate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ctx, span := trace.SpanFromContext(ctx).Tracer().Start(ctx, "handlers.advert.activate")
	defer span.End()

	v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return web.NewShutdownError("web value missing from context")
	}

	id := web.Param(r, "id")

	adv, err := ag.advert.Activate(ctx, v.TraceID, id)
	if err != nil {
		return errors.Wrapf(err, "Advert: %+v", &adv)
	}

	return web.Respond(ctx, w, adv, http.StatusOK)
}

type paginatedLimitOffsetAdvertResponse struct {
	*web.PaginatedLimitOffsetResponse
	Results []advert.AdvertInfo `json:"results"`
}
