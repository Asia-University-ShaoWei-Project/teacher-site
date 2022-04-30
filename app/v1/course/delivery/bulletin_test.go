package delivery

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"teacher-site/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateBulletin(t *testing.T) {
	r, route := NewServer()
	NewHandler(ctx, route, usecaseMock, conf)
	urlFormat = ApiUrl + `/%s/bulletin`
	dataFormat := `{"content":"%s"}`

	testCases := []struct {
		desc     string
		courseId string
		data     string
		result   HttpStatusCode
	}{
		{
			desc:     "invalid course id",
			courseId: mock.WordStr,
			data:     fmt.Sprintf(dataFormat, mock.WordStr),
			result:   http.StatusBadRequest,
		},
		{
			desc:     "empty course id",
			courseId: mock.EmptyStr,
			data:     fmt.Sprintf(dataFormat, mock.WordStr),
			result:   http.StatusBadRequest,
		},
		{
			desc:     "empty content",
			courseId: mock.PkStr,
			data:     fmt.Sprintf(dataFormat, mock.EmptyStr),
			result:   http.StatusBadRequest,
		},
		{
			desc:     "normal",
			courseId: mock.PkStr,
			data:     fmt.Sprintf(dataFormat, mock.WordStr),
			result:   http.StatusCreated,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url = fmt.Sprintf(urlFormat, tC.courseId)
			w = httptest.NewRecorder()
			req, _ = http.NewRequest("POST", url, strings.NewReader(tC.data))
			req.Header.Add("Content-Type", jsonContentType)
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}
func TestUpdateBulletin(t *testing.T) {
	r, route := NewServer()
	NewHandler(ctx, route, usecaseMock, conf)
	urlFormat = ApiUrl + `/%s/bulletin/%s`
	dataFormat := `{"content":"%s"}`

	testCases := []struct {
		desc       string
		courseId   string
		bulletinId string
		data       string
		result     HttpStatusCode
	}{
		{
			desc:       "invalid course id",
			courseId:   mock.WordStr,
			bulletinId: mock.PkStr,
			data:       fmt.Sprintf(dataFormat, mock.WordStr),
			result:     http.StatusBadRequest,
		},
		{
			desc:       "invalid bulletin id",
			courseId:   mock.PkStr,
			bulletinId: mock.WordStr,
			data:       fmt.Sprintf(dataFormat, mock.WordStr),
			result:     http.StatusBadRequest,
		},
		{
			desc:       "empty course id(0)",
			courseId:   mock.EmptyStr,
			bulletinId: mock.PkStr,
			data:       fmt.Sprintf(dataFormat, mock.WordStr),
			result:     http.StatusBadRequest,
		},
		{
			desc:       "empty bulletin id",
			courseId:   mock.PkStr,
			bulletinId: mock.EmptyStr,
			data:       fmt.Sprintf(dataFormat, mock.WordStr),
			// Note: is not a bad request
			result: http.StatusNotFound,
		},
		{
			desc:       "zero value of bulletin id",
			courseId:   mock.PkStr,
			bulletinId: mock.PkZeroStr,
			data:       fmt.Sprintf(dataFormat, mock.WordStr),
			result:     http.StatusBadRequest,
		},
		{
			desc:       "empty content",
			courseId:   mock.PkStr,
			bulletinId: mock.PkStr,
			data:       fmt.Sprintf(dataFormat, mock.EmptyStr),
			result:     http.StatusBadRequest,
		},
		{
			desc:       "normal",
			courseId:   mock.PkStr,
			bulletinId: mock.PkStr,
			data:       fmt.Sprintf(dataFormat, mock.WordStr),
			result:     http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url = fmt.Sprintf(urlFormat, tC.courseId, tC.bulletinId)
			w = httptest.NewRecorder()
			req, _ = http.NewRequest("PUT", url, strings.NewReader(tC.data))
			req.Header.Add("Content-Type", jsonContentType)
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}

func TestDeleteBulletin(t *testing.T) {
	r, route := NewServer()
	NewHandler(ctx, route, usecaseMock, conf)
	urlFormat = ApiUrl + `/%s/bulletin/%s`

	testCases := []struct {
		desc       string
		courseId   string
		bulletinId string
		result     HttpStatusCode
	}{
		{
			desc:       "invalid course id",
			courseId:   mock.WordStr,
			bulletinId: mock.PkStr,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "invalid homework id",
			courseId:   mock.PkStr,
			bulletinId: mock.WordStr,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "empty course id(0)",
			courseId:   mock.EmptyStr,
			bulletinId: mock.PkStr,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "empty homework id",
			courseId:   mock.PkStr,
			bulletinId: mock.EmptyStr,
			// Note: is not a bad request
			result: http.StatusNotFound,
		},
		{
			desc:       "zero value of homework id",
			courseId:   mock.PkStr,
			bulletinId: mock.PkZeroStr,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "normal",
			courseId:   mock.PkStr,
			bulletinId: mock.PkStr,
			result:     http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url = fmt.Sprintf(urlFormat, tC.courseId, tC.bulletinId)
			w = httptest.NewRecorder()
			req, _ = http.NewRequest("DELETE", url, nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}
