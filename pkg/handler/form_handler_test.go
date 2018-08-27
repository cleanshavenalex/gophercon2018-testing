package handler_test

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cleanshavenalex/gophercon2018-testing/pkg/handler"
)

func Test_U_FormHandler_Template_Error(t *testing.T) {

	// pass invalid hex strings
	req := httptest.NewRequest("POST", "/form", strings.NewReader("%zzzzz"))

	// set the header `Content-Type` to `application/x-www-form-urlencoded`
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// create a new httptest.NewRecorder
	res := httptest.NewRecorder()

	// Call the FormHandler
	handler.FormHandler(res, req)
	// Test to see the the response code is 500
	if res.Code != 500 {
		t.Errorf("Expected Response Error Code to be 500, received %d", res.Code)
	}
	// test the body is `Oops!`
	if res.Body.String() != "Oops!" {
		t.Errorf("Expected Response Body to be 'Oops!', received %s", res.Body)
	}

}
