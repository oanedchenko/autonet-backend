package dao

import (
	"os"
	"github.com/arangodb/go-driver/http"
	"github.com/arangodb/go-driver"
	log "github.com/sirupsen/logrus"
	"context"
)

const (
	ENV_DB_URL  = "DB_URL"
	ENV_DB_NAME = "DB_NAME"
	ENV_DB_USER = "DB_USER"
	DB_PASSWORD = "DB_PASSWORD"

	COL_USERS       = "users"
	COL_TEAMS       = "teams"
	COL_DEPARTMENTS = "departments"

	ZERO = int32(0)
)

var (
	dbUrl      = os.Getenv(ENV_DB_URL)
	dbName     = os.Getenv(ENV_DB_NAME)
	dbUser     = os.Getenv(ENV_DB_USER)
	dbPassword = os.Getenv(DB_PASSWORD)

	db       driver.Database
	colUsers driver.Collection
	colTeams driver.Collection
	colDeps  driver.Collection
)

type UserData struct {
	Id             string    `json:"_key,omitempty"`
	FirstName      string    `json:"firstName,omitempty"`
	LastName       string    `json:"lastName,omitempty"`
	Avatar         string    `json:"avatar,omitempty"`
	DepartmentID   string    `json:"departmentId,omitempty"`
	DepartmentName string    `json:"departmentName,omitempty"`
	TeamID         string    `json:"teamId,omitempty"`
	TeamName       string    `json:"teamName,omitempty"`
	JobTitle       string    `json:"jobTitle,omitempty"`
	Contacts       []Contact `json:"contacts"`
}

type Contact struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

type TeamData struct {
	Id   string `json:"_key,omitempty"`
	Name string `json:"name,omitempty"`
}

type DepartmentData struct {
	Id   string `json:"_key,omitempty"`
	Name string `json:"name,omitempty"`
}

type Dao interface {
	GetUsersPage(departmentId string, teamId string, offset int, limit int) (*[]UserData, *int32, error)
	GetTeamsPage(departmentId string, offset int, limit int) (*[]TeamData, *int32, error)
	GetDepartmentsPage(offset int, limit int) (*[]DepartmentData, *int32, error)
}

func init() {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{dbUrl},
	})

	if err != nil {
		panic(err)
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(dbUser, dbPassword),
	})
	if err != nil {
		panic(err)
	}

	db, err = c.Database(nil, dbName)
	if err != nil {
		panic(err)
	}

	colUsers, err = db.Collection(nil, COL_USERS)
	if err != nil {
		panic(err)
	}

	colTeams, err = db.Collection(nil, COL_TEAMS)
	if err != nil {
		panic(err)
	}

	colDeps, err = db.Collection(nil, COL_DEPARTMENTS)
	if err != nil {
		panic(err)
	}
}

func NewDao() *dao {
	return &dao{}
}

type dao struct{}

func (d *dao) GetUsersPage(departmentId string, teamId string, offset int, limit int) (*[]UserData, *int32, error) {
	log.Debugf("Getting users page: offset %v, limit %v", offset, limit)
	var users = make([]UserData, 0, limit)
	vars := map[string]interface{}{
		"limit":        limit,
		"offset":       offset,
		"departmentId": departmentId,
		"teamId":       teamId,
	}
	ctx := driver.WithQueryCount(context.Background())
	cursor, err := db.Query(ctx, `FOR u IN users
		FILTER (@departmentId == "" || u.departmentId == @departmentId)
			|| (@teamId == "" || u.teamId == @teamId)
		SORT u.lastName, u.firstName 
		LIMIT @offset, @limit 
		RETURN u`, vars)

	if err != nil {
		z := ZERO
		log.Errorf("Fail to query users page: %+v", err)
		return nil, &z, err
	}

	defer cursor.Close()

	for {
		var ud UserData
		_, err := cursor.ReadDocument(ctx, &ud)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			panic(err)
		}
		users = append(users, ud)
	}

	vars = map[string]interface{}{
		"departmentId": departmentId,
		"teamId":       teamId,
	}
	countCursor, err := db.Query(context.Background(), `FOR u IN users
		FILTER (@departmentId == "" || u.departmentId == @departmentId)
			|| (@teamId == "" || u.teamId == @teamId)
        COLLECT WITH COUNT INTO length
        RETURN length`, vars)

	if err != nil {
		z := ZERO
		log.Errorf("Fail to query users count: %+v", err)
		return nil, &z, err
	}
	defer countCursor.Close()

	var count int32
	_, err = countCursor.ReadDocument(ctx, &count)
	if err != nil {
		z := ZERO
		log.Errorf("Fail to read count from cursor: %+v", err)
		return nil, &z, err
	}
	return &users, &count, err
}

func (d *dao) GetTeamsPage(departmentId string, offset int, limit int) (*[]TeamData, *int32, error) {
	log.Debugf("Getting teams page: offset %v, limit %v", offset, limit)
	var teams = make([]TeamData, 0, limit)
	vars := map[string]interface{}{
		"limit":        limit,
		"offset":       offset,
		"departmentId": departmentId,
	}
	ctx := driver.WithQueryCount(context.Background())
	// todo: replace departmentId with some nice graph query!
	cursor, err := db.Query(ctx, `
		FOR t IN teams
		FILTER @departmentId == "" || t.departmentId == @departmentId  
		SORT t.name 
		LIMIT @offset, @limit 
		RETURN t`, vars)

	if err != nil {
		z := ZERO
		log.Errorf("Fail to query teams page: %+v", err)
		return nil, &z, err
	}

	defer cursor.Close()

	for {
		var td TeamData
		_, err := cursor.ReadDocument(ctx, &td)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			panic(err)
		}
		teams = append(teams, td)
	}

	vars = map[string]interface{}{
		"departmentId": departmentId,
	}
	// todo: add department creteria!!!!!!!!!!!!!
	countCursor, err := db.Query(context.Background(), `FOR t IN teams
		FILTER @departmentId == "" || t.departmentId == @departmentId
        COLLECT WITH COUNT INTO length
        RETURN length`, vars)

	if err != nil {
		z := ZERO
		log.Errorf("Fail to query teams count: %+v", err)
		return nil, &z, err
	}
	defer countCursor.Close()

	var count int32
	_, err = countCursor.ReadDocument(ctx, &count)
	if err != nil {
		z := ZERO
		log.Errorf("Fail to read count from cursor: %+v", err)
		return nil, &z, err
	}
	return &teams, &count, err
}

func (d *dao) GetDepartmentsPage(offset int, limit int) (*[]DepartmentData, *int32, error) {
	log.Debugf("Getting users page: offset %v, limit %v", offset, limit)
	var departments = make([]DepartmentData, 0, limit)
	vars := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
	}
	ctx := driver.WithQueryCount(context.Background())
	cursor, err := db.Query(ctx, "FOR d IN departments "+
		"SORT d.name "+
		"LIMIT @offset, @limit "+
		"RETURN d", vars)

	if err != nil {
		z := ZERO
		log.Errorf("Fail to query departments page: %+v", err)
		return nil, &z, err
	}

	defer cursor.Close()

	for {
		var dd DepartmentData
		_, err := cursor.ReadDocument(ctx, &dd)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			panic(err)
		}
		departments = append(departments, dd)
	}

	countCursor, err := db.Query(context.Background(), `FOR d IN departments
        COLLECT WITH COUNT INTO length
        RETURN length`, nil)

	if err != nil {
		z := ZERO
		log.Errorf("Fail to query departments count: %+v", err)
		return nil, &z, err
	}
	defer countCursor.Close()

	var count int32
	_, err = countCursor.ReadDocument(ctx, &count)
	if err != nil {
		z := ZERO
		log.Errorf("Fail to read count from cursor: %+v", err)
		return nil, &z, err
	}
	return &departments, &count, err
}
