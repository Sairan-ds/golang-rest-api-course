// build e2e

package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/api/comment")
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 200, resp.StatusCode())
}

func TestPostComment(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{"slug": "/", "author" : "12345", "body" : "Hello World"`).
		Post("http://localhost:8080/api/comment")
	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
}
