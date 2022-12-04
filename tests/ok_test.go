package tests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/aliforever/go-httpjson"
)

func TestOk(t *testing.T) {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		httpjson.Ok(writer, "", "Welcome!")
	})
	go http.ListenAndServe(":80", nil)

	resp, err := http.Get("http://localhost")
	if err != nil {
		t.Fatal("Error sending http request", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal("Error reading response", err)
		return
	}

	expected := fmt.Sprintf(`{"status_code":%d,"status_name":"OK","data":"Welcome!"}`, http.StatusOK)
	value := string(body)

	if expected != value {
		t.Fatal(fmt.Sprintf("Test Failed!\nNot the expected outcome!\nExpected: %s\nResult: %s", expected, value))
		return
	}

	expectedResultCode := http.StatusOK
	valueResultCode := resp.StatusCode

	if expectedResultCode != valueResultCode {
		t.Fatal(fmt.Sprintf("Test Failed\nNot the expected Result Code.\nExpected :%d\nValue: %d", expectedResultCode, valueResultCode))
	}
}
