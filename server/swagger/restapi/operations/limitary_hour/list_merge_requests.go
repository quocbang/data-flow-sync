// Code generated by go-swagger; DO NOT EDIT.

package limitary_hour

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/quocbang/data-flow-sync/server/swagger/models"
)

// ListMergeRequestsHandlerFunc turns a function with the right signature into a list merge requests handler
type ListMergeRequestsHandlerFunc func(ListMergeRequestsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ListMergeRequestsHandlerFunc) Handle(params ListMergeRequestsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ListMergeRequestsHandler interface for that can handle valid list merge requests params
type ListMergeRequestsHandler interface {
	Handle(ListMergeRequestsParams, *models.Principal) middleware.Responder
}

// NewListMergeRequests creates a new http.Handler for the list merge requests operation
func NewListMergeRequests(ctx *middleware.Context, handler ListMergeRequestsHandler) *ListMergeRequests {
	return &ListMergeRequests{Context: ctx, Handler: handler}
}

/* ListMergeRequests swagger:route GET /limitary-hour/merge-request limitary-hour listMergeRequests

get limitary hour merge requests

*/
type ListMergeRequests struct {
	Context *middleware.Context
	Handler ListMergeRequestsHandler
}

func (o *ListMergeRequests) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListMergeRequestsParams()
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

// ListMergeRequestsOKBody list merge requests o k body
//
// swagger:model ListMergeRequestsOKBody
type ListMergeRequestsOKBody struct {

	// data
	Data []*models.MergeRequest `json:"data"`
}

// Validate validates this list merge requests o k body
func (o *ListMergeRequestsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListMergeRequestsOKBody) validateData(formats strfmt.Registry) error {
	if swag.IsZero(o.Data) { // not required
		return nil
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listMergeRequestsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this list merge requests o k body based on the context it is used
func (o *ListMergeRequestsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ListMergeRequestsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {
			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("listMergeRequestsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ListMergeRequestsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ListMergeRequestsOKBody) UnmarshalBinary(b []byte) error {
	var res ListMergeRequestsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}