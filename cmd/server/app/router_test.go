package app

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geisonbiazus/templatemanager/internal/support/assert"
	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
)

func TestMux(t *testing.T) {
	type fixture struct {
		server *httptest.Server
	}

	setup := func() *fixture {
		server := httptest.NewServer(Mux("../../../" + templatemanager.DefaultTemplatePath))

		return &fixture{
			server: server,
		}
	}

	t.Run("POST /templates/render", func(t *testing.T) {
		f := setup()
		res := doPost(f.server, "/v1/templates/render", `{"template": {"body": {"type":"Page"}}}`)
		expected := `{"html":"\u003c!DOCTYPE html\u003e\n\u003chtml\u003e\n\u003chead\u003e\n\u003cmeta charset=\"UTF-8\"\u003e\n\u003ctitle\u003e\u003c/title\u003e\n\u003c/head\u003e\n\u003cbody\u003e\n  \n\u003c/body\u003e\n\u003c/html\u003e\n"}` + "\n"
		assertResponse(t, res, http.StatusOK, expected)
	})

	t.Run("POST /templates", func(t *testing.T) {
		t.SkipNow()
		f := setup()
		res := doPost(f.server, "/v1/templates", `{"template": {"body": {"type":"Page"}}}`)
		expected := `{"template": {"id": "1", "body": {"type":"Page"}}}` + "\n"
		assertResponse(t, res, http.StatusCreated, expected)
	})
}

func doPost(s *httptest.Server, path, body string) *http.Response {
	buffer := bytes.NewBufferString(body)
	res, _ := http.Post(s.URL+path, "application/json", buffer)
	return res
}

func assertResponse(t *testing.T, res *http.Response, status int, expectedBody string) {
	t.Helper()
	assert.Equal(t, status, res.StatusCode)
	resBody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, expectedBody, string(resBody))
}
