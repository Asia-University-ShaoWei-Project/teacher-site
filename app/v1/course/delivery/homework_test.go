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
	homeworkFilePathFormat = `../../../../static/doc/%s/homework/%s`
)

func TestCreateHomework(t *testing.T) {
	r, route := NewServer()
	conf.Server.HomeworkPathFormat = homeworkFilePathFormat
	NewHandler(ctx, route, usecaseMock, conf)
	urlFormat = ApiUrl + `/%s/homework`

	testCases := []struct {
		desc      string
		courseId  string
		number    string
		fileTitle string
		upload    bool
		result    HttpStatusCode
	}{
		{
			desc:      "invalid course id",
			courseId:  mock.WordStr,
			number:    mock.WordStr,
			fileTitle: mock.WordStr,
			upload:    false,
			result:    http.StatusBadRequest,
		},
		{
			desc:      "empty course id(0)",
			courseId:  mock.EmptyStr,
			number:    mock.WordStr,
			fileTitle: mock.WordStr,
			upload:    false,
			result:    http.StatusBadRequest,
		},
		{
			desc:      "empty number",
			courseId:  mock.PkStr,
			number:    mock.EmptyStr,
			fileTitle: mock.WordStr,
			upload:    false,
			result:    http.StatusBadRequest,
		},
		{
			desc:      "empty fileTitle",
			courseId:  mock.PkStr,
			number:    mock.WordStr,
			fileTitle: mock.EmptyStr,
			upload:    false,
			result:    http.StatusBadRequest,
		},
		{
			desc:      "normal",
			courseId:  mock.PkStr,
			number:    mock.WordStr,
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
			writer.WriteField("number", tC.number)
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
func TestUpdateHomework(t *testing.T) {
	r, route := NewServer()
	conf.Server.HomeworkPathFormat = homeworkFilePathFormat
	NewHandler(ctx, route, usecaseMock, conf)

	urlFormat = ApiUrl + `/%s/homework/%s`

	testCases := []struct {
		desc       string
		courseId   string
		homeworkId string
		number     string
		fileTitle  string
		upload     bool
		result     HttpStatusCode
	}{
		{
			desc:       "invalid course id",
			courseId:   mock.WordStr,
			homeworkId: mock.PkStr,
			number:     mock.WordStr,
			fileTitle:  mock.WordStr,
			upload:     false,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "invalid homework id",
			courseId:   mock.PkStr,
			homeworkId: mock.WordStr,
			number:     mock.WordStr,
			fileTitle:  mock.WordStr,
			upload:     false,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "empty course id(0)",
			courseId:   mock.EmptyStr,
			homeworkId: mock.PkStr,
			number:     mock.WordStr,
			fileTitle:  mock.WordStr,
			upload:     false,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "empty homework id",
			courseId:   mock.PkStr,
			homeworkId: mock.EmptyStr,
			number:     mock.WordStr,
			fileTitle:  mock.WordStr,
			// Note: is not a bad request
			upload: false,
			result: http.StatusNotFound,
		},
		{
			desc:       "zero value of homework id",
			courseId:   mock.PkStr,
			homeworkId: mock.PkZeroStr,
			number:     mock.WordStr,
			fileTitle:  mock.WordStr,
			upload:     false,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "empty number",
			courseId:   mock.PkStr,
			homeworkId: mock.PkStr,
			number:     mock.EmptyStr,
			fileTitle:  mock.WordStr,
			upload:     false,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "empty fileTitle",
			courseId:   mock.PkStr,
			homeworkId: mock.PkStr,
			number:     mock.WordStr,
			fileTitle:  mock.EmptyStr,
			upload:     false,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "normal",
			courseId:   mock.PkStr,
			homeworkId: mock.PkStr,
			number:     mock.WordStr,
			fileTitle:  mock.WordStr,
			upload:     true,
			result:     http.StatusOK,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url = fmt.Sprintf(urlFormat, tC.courseId, tC.homeworkId)
			w = httptest.NewRecorder()

			body = &bytes.Buffer{}

			writer = multipart.NewWriter(body)
			writer.WriteField("number", tC.number)
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
func TestDeleteHomework(t *testing.T) {
	r, route := NewServer()
	NewHandler(ctx, route, usecaseMock, conf)
	urlFormat = ApiUrl + `/%s/homework/%s`

	testCases := []struct {
		desc       string
		courseId   string
		homeworkId string
		result     HttpStatusCode
	}{
		{
			desc:       "invalid course id",
			courseId:   mock.WordStr,
			homeworkId: mock.PkStr,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "invalid homework id",
			courseId:   mock.PkStr,
			homeworkId: mock.WordStr,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "empty course id(0)",
			courseId:   mock.EmptyStr,
			homeworkId: mock.PkStr,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "empty homework id",
			courseId:   mock.PkStr,
			homeworkId: mock.EmptyStr,
			// Note: is not a bad request
			result: http.StatusNotFound,
		},
		{
			desc:       "zero value of homework id",
			courseId:   mock.PkStr,
			homeworkId: mock.PkZeroStr,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "normal",
			courseId:   mock.PkStr,
			homeworkId: mock.PkStr,
			result:     http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url = fmt.Sprintf(urlFormat, tC.courseId, tC.homeworkId)
			w = httptest.NewRecorder()
			req, _ = http.NewRequest("DELETE", url, nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}
