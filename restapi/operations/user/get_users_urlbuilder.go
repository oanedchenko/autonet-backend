// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetUsersURL generates an URL for the get users operation
type GetUsersURL struct {
	DepartmentID *string
	PageNumber   *int32
	PageSize     *int32
	TeamID       *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetUsersURL) WithBasePath(bp string) *GetUsersURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetUsersURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetUsersURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/user"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var departmentID string
	if o.DepartmentID != nil {
		departmentID = *o.DepartmentID
	}
	if departmentID != "" {
		qs.Set("departmentId", departmentID)
	}

	var pageNumber string
	if o.PageNumber != nil {
		pageNumber = swag.FormatInt32(*o.PageNumber)
	}
	if pageNumber != "" {
		qs.Set("pageNumber", pageNumber)
	}

	var pageSize string
	if o.PageSize != nil {
		pageSize = swag.FormatInt32(*o.PageSize)
	}
	if pageSize != "" {
		qs.Set("pageSize", pageSize)
	}

	var teamID string
	if o.TeamID != nil {
		teamID = *o.TeamID
	}
	if teamID != "" {
		qs.Set("teamId", teamID)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetUsersURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetUsersURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetUsersURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetUsersURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetUsersURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetUsersURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
