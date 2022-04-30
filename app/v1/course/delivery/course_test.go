package delivery

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"teacher-site/mock"
	"teacher-site/mock/course/usecase"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	jsonContentType = "application/json"
)

var (
	ctx = context.Background()

	// logger      = log.NewLogrus(ctx)
	conf        = mock.Conf
	usecaseMock = usecase.NewUsecase()
	ApiUrl      = mock.ApiUrl + "/course"
	// file path
	path = "../../../../mock/file.txt"
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
			req.Header.Add("Content-Type", jsonContentType)
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
