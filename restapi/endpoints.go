package restapi

import (
	"auto1.com/autonet-be/restapi/operations/department"
	"github.com/go-openapi/runtime/middleware"
	"auto1.com/autonet-be/dao"
	"auto1.com/autonet-be/models"
	"net/http"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"auto1.com/autonet-be/restapi/operations/team"
	"auto1.com/autonet-be/restapi/operations/user"
)

func GetDepartmentsPage(params department.GetDepartmentsParams) middleware.Responder {

	d := dao.NewDao()

	data, count, err := d.GetDepartmentsPage(pageParams(params.PageNumber, params.PageSize))

	if err != nil || data == nil {
		return newInternalServerError()
	}

	page := models.DepartmentPage{
		PageNumber: *params.PageNumber,
		TotalPages: totalPages(*count, *params.PageSize),
		Items:      make([]*models.Department, len(*data)),
	}

	for idx, dep := range *data {
		page.Items[idx] = &models.Department{
			ID:   strfmt.UUID(dep.Id),
			Name: dep.Name,
		}
	}

	return department.NewGetDepartmentsOK().WithPayload(&page)
}

func GetTeamsPage(params team.GetTeamsParams) middleware.Responder {

	d := dao.NewDao()

	offset, limit := pageParams(params.PageNumber, params.PageSize)
	data, count, err := d.GetTeamsPage(uuidPtrToStr(params.DepartmentID), offset, limit)

	if err != nil || data == nil {
		return newInternalServerError()
	}

	page := models.TeamPage{
		PageNumber: *params.PageNumber,
		TotalPages: totalPages(*count, *params.PageSize),
		Items:      make([]*models.Team, len(*data)),
	}

	for idx, td := range *data {
		page.Items[idx] = &models.Team{
			ID:   strfmt.UUID(td.Id),
			Name: td.Name,
		}
	}

	return team.NewGetTeamsOK().WithPayload(&page)
}

func GetUsersPage(params user.GetUsersParams) middleware.Responder {

	d := dao.NewDao()

	offset, limit := pageParams(params.PageNumber, params.PageSize)
	data, count, err := d.GetUsersPage(ptrToStr(params.DepartmentID), ptrToStr(params.DepartmentID), offset, limit)

	if err != nil || data == nil {
		return newInternalServerError()
	}

	page := models.UserPage{
		PageNumber: *params.PageNumber,
		TotalPages: totalPages(*count, *params.PageSize),
		Items:      make([]*models.User, len(*data)),
	}

	for idx, ud := range *data {
		page.Items[idx] = &models.User{
			ID:             strfmt.UUID(ud.Id),
			FirstName:      ud.FirstName,
			LastName:       ud.LastName,
			Avatar:         strfmt.UUID(ud.Avatar),
			DepartmentID:   strfmt.UUID(ud.DepartmentID),
			DepartmentName: ud.DepartmentName,
			TeamID:         strfmt.UUID(ud.TeamID),
			TeamName:       ud.TeamName,
			JobTitle:       ud.JobTitle,
			Contacts:       buildContactsModel(ud.Contacts),
		}
	}

	return user.NewGetUsersOK().WithPayload(&page)
}

func ptrToStr(p *string) string {
	if p == nil {
		return ""
	}

	return string(*p)
}

func uuidPtrToStr(up *strfmt.UUID) string {
	if up == nil {
		return ""
	}

	return (*up).String()
}

func buildContactsModel(contacts []dao.Contact) []*models.Contact {
	mc := make([]*models.Contact, len(contacts))
	for idx, c := range contacts {
		mc[idx] = &models.Contact{
			Type:  c.Type,
			Value: c.Value,
		}
	}
	return mc
}

func pageParams(pageNumber *int32, pageSize *int32) (offset int, limit int) {
	limit = int(*pageSize)
	offset = int(*pageNumber) * limit
	return
}

func totalPages(count int32, pageSize int32) int32 {

	if count <= 0 {
		return 0
	}

	fullPages := count / pageSize
	if count%pageSize > 0 {
		fullPages ++
	}

	return fullPages
}

type internalServerErrorResponder struct{}

func (ise *internalServerErrorResponder) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}

func newInternalServerError() *internalServerErrorResponder {
	return &internalServerErrorResponder{}
}
