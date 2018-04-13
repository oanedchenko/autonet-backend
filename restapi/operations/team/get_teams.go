// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetTeamsHandlerFunc turns a function with the right signature into a get teams handler
type GetTeamsHandlerFunc func(GetTeamsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTeamsHandlerFunc) Handle(params GetTeamsParams) middleware.Responder {
	return fn(params)
}

// GetTeamsHandler interface for that can handle valid get teams params
type GetTeamsHandler interface {
	Handle(GetTeamsParams) middleware.Responder
}

// NewGetTeams creates a new http.Handler for the get teams operation
func NewGetTeams(ctx *middleware.Context, handler GetTeamsHandler) *GetTeams {
	return &GetTeams{Context: ctx, Handler: handler}
}

/*GetTeams swagger:route GET /team team getTeams

Get the list of teams, optionaly filtered by department

*/
type GetTeams struct {
	Context *middleware.Context
	Handler GetTeamsHandler
}

func (o *GetTeams) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetTeamsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
