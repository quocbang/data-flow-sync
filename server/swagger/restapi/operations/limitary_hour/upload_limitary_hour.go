// Code generated by go-swagger; DO NOT EDIT.

package limitary_hour

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/quocbang/data-flow-sync/server/swagger/models"
)

// UploadLimitaryHourHandlerFunc turns a function with the right signature into a upload limitary hour handler
type UploadLimitaryHourHandlerFunc func(UploadLimitaryHourParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn UploadLimitaryHourHandlerFunc) Handle(params UploadLimitaryHourParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// UploadLimitaryHourHandler interface for that can handle valid upload limitary hour params
type UploadLimitaryHourHandler interface {
	Handle(UploadLimitaryHourParams, *models.Principal) middleware.Responder
}

// NewUploadLimitaryHour creates a new http.Handler for the upload limitary hour operation
func NewUploadLimitaryHour(ctx *middleware.Context, handler UploadLimitaryHourHandler) *UploadLimitaryHour {
	return &UploadLimitaryHour{Context: ctx, Handler: handler}
}

/* UploadLimitaryHour swagger:route POST /limitary-hour/upload/{mergeRequestID} limitary-hour uploadLimitaryHour

upload limitary hour

*/
type UploadLimitaryHour struct {
	Context *middleware.Context
	Handler UploadLimitaryHourHandler
}

func (o *UploadLimitaryHour) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUploadLimitaryHourParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}