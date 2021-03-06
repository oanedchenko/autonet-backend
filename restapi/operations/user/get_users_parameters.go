// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUsersParams creates a new GetUsersParams object
// with the default values initialized.
func NewGetUsersParams() GetUsersParams {

	var (
		// initialize parameters with default values

		pageNumberDefault = int32(0)
		pageSizeDefault   = int32(10)
	)

	return GetUsersParams{
		PageNumber: &pageNumberDefault,

		PageSize: &pageSizeDefault,
	}
}

// GetUsersParams contains all the bound params for the get users operation
// typically these are obtained from a http.Request
//
// swagger:parameters getUsers
type GetUsersParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*department id filter
	  In: query
	*/
	DepartmentID *string
	/*page number
	  In: query
	  Default: 0
	*/
	PageNumber *int32
	/*page size
	  In: query
	  Default: 10
	*/
	PageSize *int32
	/*team id filter
	  In: query
	*/
	TeamID *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetUsersParams() beforehand.
func (o *GetUsersParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qDepartmentID, qhkDepartmentID, _ := qs.GetOK("departmentId")
	if err := o.bindDepartmentID(qDepartmentID, qhkDepartmentID, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageNumber, qhkPageNumber, _ := qs.GetOK("pageNumber")
	if err := o.bindPageNumber(qPageNumber, qhkPageNumber, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageSize, qhkPageSize, _ := qs.GetOK("pageSize")
	if err := o.bindPageSize(qPageSize, qhkPageSize, route.Formats); err != nil {
		res = append(res, err)
	}

	qTeamID, qhkTeamID, _ := qs.GetOK("teamId")
	if err := o.bindTeamID(qTeamID, qhkTeamID, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetUsersParams) bindDepartmentID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.DepartmentID = &raw

	return nil
}

func (o *GetUsersParams) bindPageNumber(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetUsersParams()
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("pageNumber", "query", "int32", raw)
	}
	o.PageNumber = &value

	return nil
}

func (o *GetUsersParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetUsersParams()
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("pageSize", "query", "int32", raw)
	}
	o.PageSize = &value

	return nil
}

func (o *GetUsersParams) bindTeamID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.TeamID = &raw

	return nil
}
