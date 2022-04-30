package delivery

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"teacher-site/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	slideFilePathFormat = `../../../../static/doc/%s/slide/%s`
)

func TestCreateSlide(t *testing.T) {
	r, route := NewServer()
	conf.Server.SlidePathFormat = slideFilePathFormat
	NewHandler(ctx, route, usecaseMock, conf)
	urlFormat = ApiUrl + `/%s/slide`

	testCases := []struct {
		desc      string
		courseId  string
		chapter   string
		fileTitle string
		upload    bool
		result    HttpStatusCode
	}{
		{
			desc:      "invalid course id",
			courseId:  mock.WordStr,
			chapter:   mock.WordStr,
			fileTitle: mock.WordStr,
			upload:    false,
			result:    http.StatusBadRequest,
		},
		{
			desc:      "empty course id(0)",
			courseId:  mock.EmptyStr,
			chapter:   mock.WordStr,
			fileTitle: mock.WordStr,
			upload:    false,
			result:    http.StatusBadRequest,
		},
		{
			desc:      "empty chapter",
			courseId:  mock.PkStr,
			chapter:   mock.EmptyStr,
			fileTitle: mock.WordStr,
			upload:    false,
			result:    http.StatusBadRequest,
		},
		{
			desc:      "empty fileTitle",
			courseId:  mock.PkStr,
			chapter:   mock.WordStr,
			fileTitle: mock.EmptyStr,
			upload:    false,
			result:    http.StatusBadRequest,
		},
		{
			desc:      "normal",
			courseId:  mock.PkStr,
			chapter:   mock.WordStr,
			fileTitle: mock.WordStr,
			upload:    true,
			result:    http.StatusCreated,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			w = httptest.NewRecorder()
			url = fmt.Sprintf(urlFormat, tC.courseId)
			body = &bytes.Buffer{}

			writer = multipart.NewWriter(body)
			writer.WriteField("chapter", tC.chapter)
			writer.WriteField("fileTitle", tC.fileTitle)

			if tC.upload {
				part, err := writer.CreateFormFile("file", filepath.Base(path))
				if err != nil {
					t.Error(err)
				}
				file, err := os.Open(path)
				if err != nil {
					t.Error(err)
				}
				_, err = io.Copy(part, file)
				if err != nil {
					t.Error(err)
				}
				file.Close()
			}

			writer.Close()

			req, _ = http.NewRequest("POST", url, body)
			req.Header.Add("Content-Type", writer.FormDataContentType())

			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}
func TestUpdateSlide(t *testing.T) {
	r, route := NewServer()
	conf.Server.SlidePathFormat = slideFilePathFormat
	NewHandler(ctx, route, usecaseMock, conf)

	urlFormat = ApiUrl + `/%s/slide/%s`

	testCases := []struct {
		desc      string
		courseId  string
		slideId   string
		chapter   string
		fileTitle string
		upload    bool
		result    HttpStatusCode
	}{
		{
			desc:      "invalid course id",
			courseId:  mock.WordStr,
			slideId:   mock.PkStr,
			chapter:   mock.WordStr,
			fileTitle: mock.WordStr,
			upload:    false,
			result:    http.StatusBadRequest,
		},
		{
			desc:      "invalid slide id",
			courseId:  mock.PkStr,
			slideId:   mock.WordStr,
			chapter:   mock.WordStr,
			fileTitle: mock.WordStr,
			upload:    false,
			result:    http.StatusBadRequest,
		},
		{
			desc:      "empty course id(0)",
			courseId:  mock.EmptyStr,
			slideId:   mock.PkStr,
			chapter:   mock.WordStr,
			fileTitle: mock.WordStr,
			upload:    false,
			result:    http.StatusBadRequest,
		},
		{
			desc:      "empty slide id",
			courseId:  mock.PkStr,
			slideId:   mock.EmptyStr,
			chapter:   mock.WordStr,
			fileTitle: mock.WordStr,
			// Note: is not a bad request
			upload: false,
			result: http.StatusNotFound,
		},
		{
			desc:      "zero value of slide id",
			courseId:  mock.PkStr,
			slideId:   mock.PkZeroStr,
			chapter:   mock.WordStr,
			fileTitle: mock.WordStr,
			upload:    false,
			result:    http.StatusBadRequest,
		},
		{
			desc:      "empty chapter",
			courseId:  mock.PkStr,
			slideId:   mock.PkStr,
			chapter:   mock.EmptyStr,
			fileTitle: mock.WordStr,
			upload:    false,
			result:    http.StatusBadRequest,
		},
		{
			desc:      "empty fileTitle",
			courseId:  mock.PkStr,
			slideId:   mock.PkStr,
			chapter:   mock.WordStr,
			fileTitle: mock.EmptyStr,
			upload:    false,
			result:    http.StatusBadRequest,
		},
		{
			desc:      "normal",
			courseId:  mock.PkStr,
			slideId:   mock.PkStr,
			chapter:   mock.WordStr,
			fileTitle: mock.WordStr,
			upload:    true,
			result:    http.StatusOK,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url = fmt.Sprintf(urlFormat, tC.courseId, tC.slideId)
			w = httptest.NewRecorder()

			body = &bytes.Buffer{}

			writer = multipart.NewWriter(body)
			writer.WriteField("chapter", tC.chapter)
			writer.WriteField("fileTitle", tC.fileTitle)

			if tC.upload {
				part, err := writer.CreateFormFile("file", filepath.Base(path))
				if err != nil {
					t.Error(err)
				}
				file, err := os.Open(path)
				if err != nil {
					t.Error(err)
				}
				_, err = io.Copy(part, file)
				if err != nil {
					t.Error(err)
				}
				file.Close()
			}

			writer.Close()

			req, _ = http.NewRequest("PUT", url, body)
			req.Header.Add("Content-Type", writer.FormDataContentType())
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}
func TestDeleteSlide(t *testing.T) {
	r, route := NewServer()
	NewHandler(ctx, route, usecaseMock, conf)
	urlFormat = ApiUrl + `/%s/slide/%s`

	testCases := []struct {
		desc     string
		courseId string
		slideId  string
		result   HttpStatusCode
	}{
		{
			desc:     "invalid course id",
			courseId: mock.WordStr,
			slideId:  mock.PkStr,
			result:   http.StatusBadRequest,
		},
		{
			desc:     "invalid slide id",
			courseId: mock.PkStr,
			slideId:  mock.WordStr,
			result:   http.StatusBadRequest,
		},
		{
			desc:     "empty course id(0)",
			courseId: mock.EmptyStr,
			slideId:  mock.PkStr,
			result:   http.StatusBadRequest,
		},
		{
			desc:     "empty slide id",
			courseId: mock.PkStr,
			slideId:  mock.EmptyStr,
			// Note: is not a bad request
			result: http.StatusNotFound,
		},
		{
			desc:     "zero value of slide id",
			courseId: mock.PkStr,
			slideId:  mock.PkZeroStr,
			result:   http.StatusBadRequest,
		},
		{
			desc:     "normal",
			courseId: mock.PkStr,
			slideId:  mock.PkStr,
			result:   http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url = fmt.Sprintf(urlFormat, tC.courseId, tC.slideId)
			w = httptest.NewRecorder()
			req, _ = http.NewRequest("DELETE", url, nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}
