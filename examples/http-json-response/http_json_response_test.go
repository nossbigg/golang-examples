package examples

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPJSONResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(EndpointRequestHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	output, err := ioutil.ReadAll(res.Body)

	outputItem := Item{}
	expectedItem := Item{
		Name:  "some-item",
		Price: 123,
	}
	json.Unmarshal(output, &outputItem)

	assert.Equal(t, "application/json", res.Header.Get("Content-Type"))
	assert.Equal(t, expectedItem, outputItem)
}
