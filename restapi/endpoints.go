package restapi

import (
	"auto1.com/autonet-be/restapi/operations/department"
	"github.com/go-openapi/runtime/middleware"
)

func GetDepartments(params department.GetDepartmentsParams) middleware.Responder {
	// todo: real implementation
	//department.
	department.NewGetDepartmentsOK().WithPayload()
}
