package examples

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorldServer(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(HelloWorldHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		log.Fatal(err)
	}

	output, err := ioutil.ReadAll(res.Body)

	assert.Equal(t, 200, res.StatusCode)
	assert.Equal(t, []byte("Hello world!"), output)
}

func TestHelloWorldServer_DisallowPost(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(HelloWorldHandler))
	defer ts.Close()

	res, err := http.Post(ts.URL, "", nil)
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, 400, res.StatusCode)
}
