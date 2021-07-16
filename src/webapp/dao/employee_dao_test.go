package dao

import (
	"net/http"
	"testing"
)

func TestGetEmployee(t *testing.T) {
	emp, err := GetEmployee(999)
	if emp != nil {
		t.Error("employee_id with 999 is not accepted")
	}
	if err == nil {
		t.Error("Excepting error for empployee_id 999")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Error("Excepting 404 status when employee_id is 999")
	}
}
