package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CrunchyData/pg_featureserv/api"
)

func TestRootHandler(t *testing.T) {
	resp := execHTTPRequest(t, "/", handleRootJSON)
	body, _ := ioutil.ReadAll(resp.Body)

	// Check the response body
	//var v map[string]interface{}
	var v api.RootInfo
	json.Unmarshal(body, &v)

	fmt.Println("Response ==>")
	fmt.Println(v.Title)
	fmt.Println(v.Description)
	/*
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	*/
}

func TestCollectionMissing(t *testing.T) {
	execHTTPRequestStatus(t, "/collections/missing", handleCollection, http.StatusNotFound)

	fmt.Println("Response ==>")

	/*
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	*/
}

func execHTTPRequest(t *testing.T, url string,
	handler http.HandlerFunc) *httptest.ResponseRecorder {
	return execHTTPRequestStatus(t, url, handler, http.StatusOK)
}

func execHTTPRequestStatus(t *testing.T, url string,
	handler http.HandlerFunc,
	statusExpected int) *httptest.ResponseRecorder {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h := http.HandlerFunc(handler)
	h.ServeHTTP(rr, req)

	// Check the status code
	fmt.Println("Status:", rr.Code)
	if status := rr.Code; status != statusExpected {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, statusExpected)
	}
	return rr
}
