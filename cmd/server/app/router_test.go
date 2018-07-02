package app

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/geisonbiazus/addrvrf/assert"
	"github.com/geisonbiazus/templatemanager/internal/templatemanager"
)

func TestMux(t *testing.T) {
	server := httptest.NewServer(Mux("../../../" + templatemanager.DefaultTemplatePath))

	body := bytes.NewBufferString(`{"template": {"type":"Page"}}`)
	res, _ := http.Post(server.URL+"/render_by_json", "application/json", body)

	assert.Equal(t, http.StatusOK, res.StatusCode)

	expected := `{"html":"\u003c!DOCTYPE html\u003e\n\u003chtml\u003e\n\u003chead\u003e\n\u003cmeta charset=\"UTF-8\"\u003e\n\u003ctitle\u003e\u003c/title\u003e\n\u003c/head\u003e\n\u003cbody\u003e\n  \n\u003c/body\u003e\n\u003c/html\u003e\n"}` + "\n"
	resBody, _ := ioutil.ReadAll(res.Body)
	assert.Equal(t, expected, string(resBody))
}
