package handler

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/7phs/coding-challenge-search/config"
	"github.com/7phs/coding-challenge-search/model"
	"github.com/stretchr/testify/assert"
	"github.com/verdverm/frisby"
)

type testSearchParam struct {
	name       string
	searchTerm string
	lat        string
	long       string
	status     int
	result     model.ItemsList
	err        bool
}

func (o *testSearchParam) Param() string {
	params := url.Values{}
	for _, p := range []struct {
		name  string
		value string
	}{
		{name: "searchTerm", value: o.searchTerm},
		{name: "lat", value: o.lat},
		{name: "long", value: o.long},
	} {
		if len(p.value) > 0 {
			params.Add(p.name, p.value)
		}
	}

	return params.Encode()
}

func (o *testSearchParam) Response(body io.ReadCloser) (resp SearchHandlerResponse) {
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}

	json.Unmarshal(data, &resp)

	return
}

func TestSearch(t *testing.T) {
	var (
		count          = 100
		mockDataSource = model.NewMockSearchDataSource(count)
	)

	defer testGinMode()()

	defer model.MockNewSearchKeywords(200)()

	defer func() func() {
		preSearchModel := model.SearchModel

		model.SearchModel = model.NewSearch(mockDataSource)

		return func() {
			model.SearchModel = preSearchModel
		}
	}()()

	srv := httptest.NewServer(DefaultRouter(&config.Config{}))
	defer srv.Close()

	testSuites := []*testSearchParam{
		{
			name:   "empty params",
			status: 422,
			err:    true,
		},
		{
			name:   "invalid location",
			lat:    "inv",
			long:   "inv",
			status: 400,
			err:    true,
		},
		{
			name:       "general",
			searchTerm: "hello",
			result:     mockDataSource.Items[:20],
			status:     200,
		},
	}

	for i, test := range testSuites {
		f := frisby.Create("search: " + test.name).
			Get(srv.URL + "/search?" + test.Param()).
			Send().
			ExpectStatus(test.status)

		resp := test.Response(f.Resp.Body)

		assert.Empty(t, f.Errs, "%d: %s", i+1, f.Errs)
		if test.err {
			assert.NotNil(t, resp.RespError.Errors)
		} else {
			assert.Nil(t, resp.RespError.Errors)
		}

		assert.Equal(t, test.result, resp.Data)
	}
}
