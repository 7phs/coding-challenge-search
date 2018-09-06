package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/7phs/coding-challenge-search/config"
	"github.com/stretchr/testify/assert"
	"github.com/verdverm/frisby"
)

func TestHealthCheck(t *testing.T) {
	defer testGinMode()()

	srv := httptest.NewServer(DefaultRouter(&config.Config{}))
	defer srv.Close()

	f := frisby.Create("Test successful a health check").
		Get(srv.URL + "/health/check").
		Send().
		ExpectStatus(200).
		ExpectContent("ok")

	assert.Empty(t, f.Errs, "%s", f.Errs)
}
