package delivery

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"teacher-site/config"
	"teacher-site/mock"
	"teacher-site/mock/page/usecase"
	"testing"

	"github.com/gin-gonic/gin"
)

// todo: TestLogin

const (
	methodGet       = "GET"
	methodPost      = "POST"
	JsonContentType = "application/json"
)

var (
	ctx         = context.Background()
	r           = gin.Default()
	route       = r.Group("/")
	usecaseMock = usecase.NewUsecase()
	conf        = config.New()
)
var (
	url string
	req *http.Request
	w   *httptest.ResponseRecorder
)

type HttpStatusCode int

func TestTeacherList(t *testing.T) {
	NewHandler(ctx, route, usecaseMock, conf)
	urlFormat := `/page/%s`
	testCases := []struct {
		desc   string
		page   string
		result HttpStatusCode
	}{
		{
			desc:   "unknown page",
			page:   mock.Unknown,
			result: http.StatusBadRequest,
		},
		{
			desc:   "negative digit",
			page:   "-1",
			result: http.StatusBadRequest,
		},
		{
			desc:   "normal",
			page:   "2",
			result: http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			w = httptest.NewRecorder()
			url = fmt.Sprintf(urlFormat, tC.page)
			req, _ = http.NewRequest(methodGet, url, nil)
			r.ServeHTTP(w, req)
		})
	}
}

func TestHome(t *testing.T) {
	NewHandler(ctx, route, usecaseMock, conf)
	testCases := []struct {
		desc          string
		teacherDomain string
		result        HttpStatusCode
	}{
		{
			desc:          "unknown teacher domain",
			teacherDomain: mock.Unknown,
			result:        http.StatusNotFound,
		},
		{
			desc:          "normal",
			teacherDomain: mock.TeacherDomain,
			result:        http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			w = httptest.NewRecorder()
			url := "/" + tC.teacherDomain
			req, _ = http.NewRequest(methodGet, url, nil)
			r.ServeHTTP(w, req)
		})
	}
}
