package language

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/VxVxN/social_network/cmd/ajax_server/context"
	"github.com/VxVxN/social_network/internal/log"
)

func TestSetLanguage(t *testing.T) {
	type testStruct struct {
		req    requestLanguage
		failed bool
	}
	tests := []testStruct{
		{req: requestLanguage{Language: "RU"}},
		{req: requestLanguage{Language: "EN"}},
		{req: requestLanguage{Language: "ru"}, failed: true},
		{req: requestLanguage{Language: "en"}, failed: true},
		{req: requestLanguage{Language: ""}, failed: true},
		{req: requestLanguage{Language: "123"}, failed: true},
	}

	context := &context.Context{Log: log.Init("test.log", true)}

	for _, test := range tests {
		bodyJson, err := json.Marshal(test.req)
		if err != nil {
			t.Fatal(err)
		}
		body := bytes.NewBuffer(bodyJson)
		req, err := http.NewRequest("POST", "/ajax/language", body)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		resp := SetLanguage(rr, req, context)
		status := resp.Code

		if status != http.StatusOK && !test.failed {
			t.Errorf("handler returned wrong status code, got: %v, want: %v", status, http.StatusOK)
		}

		if status == http.StatusOK && test.failed {
			t.Errorf("handler returned wrong status code, got: %v, expected failed: true", status)
		}
	}
}
