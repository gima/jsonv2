package govalid_test

import (
	"testing"

	v "github.com/gima/govalid/v1"
)

const (
	PASS = true
	FAIL = false
)

func test(t *testing.T, title string, expectPass bool, validator v.Validator, data interface{}) {
	path, err := validator.Validate(data)

	if err == nil {

		if expectPass {
			return
		} else {
			t.Fatalf("'%s' failed (%s). Path: %s", title, "nil error, but error expected", path)
		}

	} else if err != nil {

		if !expectPass {
			return
		} else {
			t.Fatalf("'%s' failed (%s). Path: %s", title, err, path)
		}

	}
}
