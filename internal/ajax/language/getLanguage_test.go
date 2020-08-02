package language

import (
	"net/http"
	"testing"

	"github.com/VxVxN/social_network/cmd/ajax_server/context"
	cnfg "github.com/VxVxN/social_network/internal/config"
	"github.com/VxVxN/social_network/internal/log"
)

func TestGetLanguage(t *testing.T) {
	type testStruct struct {
		cookie        string
		isEmptyCookie bool
		expected      string
		failed        bool
	}
	tests := []testStruct{
		{cookie: "RU", expected: "RU"},
		{cookie: "EN", expected: "EN"},
		{cookie: "ru", failed: true},
		{cookie: "en", failed: true},
		{cookie: "", failed: true},
		{cookie: "123", failed: true},
		{isEmptyCookie: true, expected: cnfg.Config.DefaultLanguage},
	}

	context := &context.Context{Log: log.Init("test.log", true)}

	for _, test := range tests {
		req, err := http.NewRequest("GET", "/ajax/language", nil)
		if err != nil {
			t.Fatal(err)
		}

		cookie := http.Cookie{Name: "language", Value: test.cookie}

		if !test.isEmptyCookie {
			req.AddCookie(&cookie)
		}

		var w http.ResponseWriter
		resp := GetLanguage(w, req, context)
		status := resp.Code

		if status != http.StatusOK && !test.failed {
			t.Errorf("handler returned wrong status code, got: %v, want: %v", status, http.StatusOK)
		}

		if status == http.StatusOK && test.failed {
			t.Errorf("handler returned wrong status code, got: %v, expected failed: true", status)
		}

		if resp.Data != test.expected && !test.failed {
			t.Errorf("handler returned unexpected body, input: %v, got: %v, want %v", test.cookie, resp.Data, test.expected)
		}
	}
}
