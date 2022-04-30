package delivery

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"teacher-site/mock"
	"teacher-site/mock/course/usecase"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	JsonContentType = "application/json"
)

var (
	ctx = context.Background()

	// logger      = log.NewLogrus(ctx)
	conf        = mock.Conf
	usecaseMock = usecase.NewUsecase()
	ApiUrl      = mock.ApiUrl + "/course"
	// file path
	path                = "../../../../mock/file.txt"
	slideFilePathFormat = `../../../../static/doc/%s/slide/%s`
)
var (
	url, urlFormat string
	req            *http.Request
	w              *httptest.ResponseRecorder
	// form
	body   *bytes.Buffer
	writer *multipart.Writer
)

type HttpStatusCode int

func NewServer() (*gin.Engine, *gin.RouterGroup) {
	r := gin.Default()
	route := r.Group("/:teacherDomain/api/v1/course")
	return r, route
}
func TestCreate(t *testing.T) {
	r, route := NewServer()
	NewHandler(ctx, route, usecaseMock, conf)
	dataFormat := `{"nameZh":"%s", "nameUs":"%s"}`

	testCases := []struct {
		desc   string
		data   string
		result HttpStatusCode
	}{
		{
			desc:   "empty value",
			data:   fmt.Sprintf(dataFormat, mock.EmptyStr, mock.EmptyStr),
			result: http.StatusBadRequest,
		},
		{
			desc:   "normal",
			data:   fmt.Sprintf(dataFormat, mock.WordStr, mock.EmptyStr),
			result: http.StatusCreated,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			w = httptest.NewRecorder()
			req, _ = http.NewRequest("POST", ApiUrl, strings.NewReader(tC.data))
			req.Header.Add("Content-Type", JsonContentType)
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}

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
			req.Header.Add("Content-Type", JsonContentType)
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}

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
func TestCreateHomework(t *testing.T) {
	r, route := NewServer()
	NewHandler(ctx, route, usecaseMock, conf)
	urlFormat = ApiUrl + `/%s/homework`
	dataFormat := `{"number":"%s","fileTitle":"%s"}`

	testCases := []struct {
		desc     string
		courseId string
		data     string
		result   HttpStatusCode
	}{
		{
			desc:     "invalid course id",
			courseId: mock.WordStr,
			data:     fmt.Sprintf(dataFormat, mock.WordStr, mock.WordStr),
			result:   http.StatusBadRequest,
		},
		{
			desc:     "empty course id(0)",
			courseId: mock.EmptyStr,
			data:     fmt.Sprintf(dataFormat, mock.WordStr, mock.WordStr),
			result:   http.StatusBadRequest,
		},
		{
			desc:     "empty number",
			courseId: mock.PkStr,
			data:     fmt.Sprintf(dataFormat, mock.EmptyStr, mock.WordStr),
			result:   http.StatusBadRequest,
		},
		{
			desc:     "empty fileTitle",
			courseId: mock.PkStr,
			data:     fmt.Sprintf(dataFormat, mock.WordStr, mock.EmptyStr),
			result:   http.StatusBadRequest,
		},
		{
			desc:     "normal",
			courseId: mock.PkStr,
			data:     fmt.Sprintf(dataFormat, mock.WordStr, mock.WordStr),
			result:   http.StatusCreated,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url = fmt.Sprintf(urlFormat, tC.courseId)
			w = httptest.NewRecorder()
			req, _ = http.NewRequest("POST", url, strings.NewReader(tC.data))
			req.Header.Add("Content-Type", JsonContentType)
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}

// ===== GET =====
func TestGetContent(t *testing.T) {
	r, route := NewServer()
	NewHandler(ctx, route, usecaseMock, conf)
	urlFormat := ApiUrl + "/%s"

	testCases := []struct {
		desc     string
		courseId string
		result   HttpStatusCode
	}{
		{
			desc:     "not digit value of the course id",
			courseId: mock.WordStr,
			result:   http.StatusBadRequest,
		},
		{
			desc:     "negative digit",
			courseId: mock.NegativePkStr,
			result:   http.StatusBadRequest,
		},
		{
			desc:     "normal",
			courseId: mock.PkStr,
			result:   http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url = fmt.Sprintf(urlFormat, tC.courseId)
			w = httptest.NewRecorder()
			req, _ = http.NewRequest("GET", url, nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}

// ===== Update =====
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
			req.Header.Add("Content-Type", JsonContentType)
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
func TestUpdateHomework(t *testing.T) {
	r, route := NewServer()
	NewHandler(ctx, route, usecaseMock, conf)
	urlFormat = ApiUrl + `/%s/homework/%s`
	dataFormat := `{"number":"%s", "fileTitle":"%s"}`

	testCases := []struct {
		desc       string
		courseId   string
		homeworkId string
		data       string
		result     HttpStatusCode
	}{
		{
			desc:       "invalid course id",
			courseId:   mock.WordStr,
			homeworkId: mock.PkStr,
			data:       fmt.Sprintf(dataFormat, mock.WordStr, mock.WordStr),
			result:     http.StatusBadRequest,
		},
		{
			desc:       "invalid homework id",
			courseId:   mock.PkStr,
			homeworkId: mock.WordStr,
			data:       fmt.Sprintf(dataFormat, mock.WordStr, mock.WordStr),
			result:     http.StatusBadRequest,
		},
		{
			desc:       "empty course id(0)",
			courseId:   mock.EmptyStr,
			homeworkId: mock.PkStr,
			data:       fmt.Sprintf(dataFormat, mock.WordStr, mock.WordStr),
			result:     http.StatusBadRequest,
		},
		{
			desc:       "empty homework id",
			courseId:   mock.PkStr,
			homeworkId: mock.EmptyStr,
			data:       fmt.Sprintf(dataFormat, mock.WordStr, mock.WordStr),
			// Note: is not a bad request
			result: http.StatusNotFound,
		},
		{
			desc:       "zero value of homework id",
			courseId:   mock.PkStr,
			homeworkId: mock.PkZeroStr,
			data:       fmt.Sprintf(dataFormat, mock.WordStr, mock.WordStr),
			result:     http.StatusBadRequest,
		},
		{
			desc:       "empty number",
			courseId:   mock.PkStr,
			homeworkId: mock.PkStr,
			data:       fmt.Sprintf(dataFormat, mock.EmptyStr, mock.WordStr),
			result:     http.StatusBadRequest,
		},
		{
			desc:       "empty fileTitle",
			courseId:   mock.PkStr,
			homeworkId: mock.PkStr,
			data:       fmt.Sprintf(dataFormat, mock.WordStr, mock.EmptyStr),
			result:     http.StatusBadRequest,
		},
		{
			desc:       "normal",
			courseId:   mock.PkStr,
			homeworkId: mock.PkStr,
			data:       fmt.Sprintf(dataFormat, mock.WordStr, mock.WordStr),
			result:     http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url = fmt.Sprintf(urlFormat, tC.courseId, tC.homeworkId)
			w = httptest.NewRecorder()
			req, _ = http.NewRequest("PUT", url, strings.NewReader(tC.data))
			req.Header.Add("Content-Type", JsonContentType)
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}

// Delete
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
