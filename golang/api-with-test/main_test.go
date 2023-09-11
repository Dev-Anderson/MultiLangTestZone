package main

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestGetUser_CookieMatching(t *testing.T) {
	apitest.New().
		Handler(newApp().Router).
		Get("/user/1234").
		Expect(t).
		Cookies(apitest.NewCookie("TomsFavouriteDring").
			Value("Beer").
			Path("/")).
		Status(http.StatusOK).
		End()
}

func TestGetUser_Success(t *testing.T) {
	apitest.New().
		Handler(newApp().Router).
		Get("/user/1234").
		Expect(t).
		Body(`{"id": "1234", "name": "Anderson"}`).
		Status(http.StatusOK).
		End()
}

func TestGetUser_Sucess_JsonPath(t *testing.T) {
	apitest.New().
		Handler(newApp().Router).
		Get("/user/1234").
		Expect(t).
		Assert(jsonpath.Equal(`$.id`, "1234")).
		Status(http.StatusOK).
		End()
}

func TestGetUser_NotFound(t *testing.T) {
	apitest.New().
		Handler(newApp().Router).
		Get("/user/12344").
		Expect(t).
		Status(http.StatusNotFound).
		End()
}
