package dao

import "testing"

func TestDao_GetDepartmentsPage(t *testing.T) {

	d := NewDao()

	data, count, err := d.GetDepartmentsPage(0, 10)

	if err != nil {
		t.Errorf("Failed to query departments data: %+v", err)
	}

	if *count == 0 {
		t.Errorf("Fail to query departments count")
	}

	if data == nil || len(*data) == 0 {
		t.Errorf("No departments data received")
	}
}

func TestDao_GetTeamsPage(t *testing.T) {
	d := NewDao()

	// empty department id:
	data, count, err := d.GetTeamsPage("", 0, 10)

	if err != nil {
		t.Errorf("Failed to query teams data: %+v", err)
	}

	if *count == 0 {
		t.Errorf("Fail to query teams count")
	}

	if data == nil || len(*data) == 0 {
		t.Errorf("No teams data received")
	}

	// existing department id:
	data, count, err = d.GetTeamsPage("8c2ebd13-9ed5-4840-b9b7-5aaab8cf0413", 0, 10)

	if err != nil {
		t.Errorf("Failed to query teams data: %+v", err)
	}

	if *count == 0 {
		t.Errorf("Fail to query teams count")
	}

	if data == nil || len(*data) == 0 {
		t.Errorf("No teams data received")
	}
}

func TestDao_GetUsersPage(t *testing.T) {
	d := NewDao()

	data, count, err := d.GetUsersPage("8c2ebd13-9ed5-4840-b9b7-5aaab8cf0413", "c21975a6-c97c-4ef3-ace3-9570df9839ce", 0, 10)

	if err != nil {
		t.Errorf("Failed to query users data: %+v", err)
	}

	if *count == 0 {
		t.Errorf("Fail to query users count")
	}

	if data == nil || len(*data) == 0 {
		t.Errorf("No users data received")
	}
}