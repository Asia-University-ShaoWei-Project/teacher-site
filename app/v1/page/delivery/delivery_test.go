package delivery

import (
	"context"
	"net/http"
	"net/http/httptest"
	"teacher-site/config"
	"teacher-site/mock"
	"teacher-site/mock/page/usecase"
	"testing"

	"github.com/gin-gonic/gin"
)

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
	req *http.Request
	w   *httptest.ResponseRecorder
)

type HttpStatusCode int

// todo: get teacher list by api, test negative digit
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
