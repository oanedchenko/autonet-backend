// Code generated by go-swagger; DO NOT EDIT.

package team

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "auto1.com/autonet-be/models"
)

// GetTeamsOKCode is the HTTP code returned for type GetTeamsOK
const GetTeamsOKCode int = 200

/*GetTeamsOK OK

swagger:response getTeamsOK
*/
type GetTeamsOK struct {

	/*
	  In: Body
	*/
	Payload *models.TeamPage `json:"body,omitempty"`
}

// NewGetTeamsOK creates GetTeamsOK with default headers values
func NewGetTeamsOK() *GetTeamsOK {

	return &GetTeamsOK{}
}

// WithPayload adds the payload to the get teams o k response
func (o *GetTeamsOK) WithPayload(payload *models.TeamPage) *GetTeamsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get teams o k response
func (o *GetTeamsOK) SetPayload(payload *models.TeamPage) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTeamsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTeamsUnauthorizedCode is the HTTP code returned for type GetTeamsUnauthorized
const GetTeamsUnauthorizedCode int = 401

/*GetTeamsUnauthorized Unauthorized

swagger:response getTeamsUnauthorized
*/
type GetTeamsUnauthorized struct {
}

// NewGetTeamsUnauthorized creates GetTeamsUnauthorized with default headers values
func NewGetTeamsUnauthorized() *GetTeamsUnauthorized {

	return &GetTeamsUnauthorized{}
}

// WriteResponse to the client
func (o *GetTeamsUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// GetTeamsForbiddenCode is the HTTP code returned for type GetTeamsForbidden
const GetTeamsForbiddenCode int = 403

/*GetTeamsForbidden Forbidden

swagger:response getTeamsForbidden
*/
type GetTeamsForbidden struct {
}

// NewGetTeamsForbidden creates GetTeamsForbidden with default headers values
func NewGetTeamsForbidden() *GetTeamsForbidden {

	return &GetTeamsForbidden{}
}

// WriteResponse to the client
func (o *GetTeamsForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}
