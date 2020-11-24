package test

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/mandatorySuicide/golang-code-quality/src/app"
	"github.com/mandatorySuicide/golang-code-quality/src/utility"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	testServer := httptest.NewServer(app.SetupServer())
	defer testServer.Close()
	resp, err := http.Get(fmt.Sprintf("%s/ping", testServer.URL))
	defer resp.Body.Close()
	assert.Nil(t, err, fmt.Sprintf("Error should be nil"))
	assert.Equal(t, resp.StatusCode, 200, "Response code should be 200")
	body := utility.ReadHttpBody(resp.Body)
	// ***** without new line test fails, ironic!!! *****
	fmt.Println(body)
	// todo : I need to find a way to parse json and validate them easily
	// e.g contract testing
	data := []byte(body)
	message, messageErr := jsonparser.GetString(data, "message")
	assert.Nil(t, messageErr, "'message' key should be present in the JSON")
	assert.Equal(t, "pong", message, "Message should be pong")
}
