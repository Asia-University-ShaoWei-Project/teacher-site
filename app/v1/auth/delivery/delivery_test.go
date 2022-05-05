package delivery

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"teacher-site/config"
	"teacher-site/mock"
	"teacher-site/mock/auth/usecase"
	"teacher-site/pkg/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	methodPost      = "POST"
	JsonContentType = "application/json"
)

var (
	ctx         = context.Background()
	r           = gin.Default()
	route       = r.Group("/:teacherDomain/api/v1/auth")
	usecaseMock = usecase.NewUsecase()
	conf        = config.New()
	ApiUrl      = mock.ApiUrl + "/auth"
)
var (
	data, dataFormat string
	// err                   error
	// body                  []byte
	req *http.Request
	w   *httptest.ResponseRecorder
)

type HttpStatusCode int

func TestLogin(t *testing.T) {
	NewHandler(ctx, route, usecaseMock, conf)
	dataFormat = `{"id":"%s", "password":"%s"}`
	url := ApiUrl + `/login`
	token, _ := util.GenerateJwt(conf.Jwt, mock.GetJwtRequest())

	testCases := []struct {
		desc         string
		userId       string
		userPassword string
		isLogged     bool
		result       HttpStatusCode
	}{
		{
			desc:         "logged in",
			userId:       mock.UserId,
			userPassword: mock.UserPassword,
			isLogged:     true,
			result:       http.StatusFound,
		},
		{
			desc:         "fail request data",
			userId:       mock.EmptyStr,
			userPassword: mock.EmptyStr,
			isLogged:     false,
			result:       http.StatusBadRequest,
		},
		{
			desc:         "normal",
			userId:       mock.UserId,
			userPassword: mock.UserPassword,
			isLogged:     false,
			result:       http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			w = httptest.NewRecorder()
			data = fmt.Sprintf(dataFormat, tC.userId, tC.userPassword)
			req, _ = http.NewRequest("POST", url, strings.NewReader(data))
			if tC.isLogged {
				setupHeader(req, JsonContentType, token)
			}
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}

func TestLogout(t *testing.T) {
	NewHandler(ctx, route, usecaseMock, conf)
	url := ApiUrl + "/logout"
	token, _ := util.GenerateJwt(conf.Jwt, mock.GetJwtRequest())

	testCases := []struct {
		desc       string
		authHeader string
		result     HttpStatusCode
	}{
		{
			desc:       "fail token",
			authHeader: mock.Unknown,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "real token",
			authHeader: token,
			result:     http.StatusNoContent,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			w = httptest.NewRecorder()
			req, _ = http.NewRequest("POST", url, nil)
			setupHeader(req, JsonContentType, tC.authHeader)
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}
func TestRegister(t *testing.T) {
	NewHandler(ctx, route, usecaseMock, conf)
	url := ApiUrl + "/register"

	testCases := []struct {
		desc   string
		data   string
		result HttpStatusCode
	}{
		{
			desc:   "empty id",
			data:   `{"id":"", "password":"password", "domain":"domain", "email":"email", "nameZh":"name"}`,
			result: http.StatusBadRequest,
		},
		{
			desc:   "empty password",
			data:   `{"id":"id", "password":"", "domain":"domain", "email":"email", "nameZh":"name"}`,
			result: http.StatusBadRequest,
		},
		{
			desc:   "empty domain",
			data:   `{"id":"id", "password":"password", "domain":"", "email":"email", "nameZh":"name"}`,
			result: http.StatusBadRequest,
		},
		{
			desc:   "empty email",
			data:   `{"id":"id", "password":"password", "domain":"domain", "email":"", "nameZh":"name"}`,
			result: http.StatusBadRequest,
		},
		{
			desc:   "empty name",
			data:   `{"id":"id", "password":"password", "domain":"domain", "email":"email", "nameZh":""}`,
			result: http.StatusBadRequest,
		},
		{
			desc:   "normal",
			data:   `{"id":"id", "password":"password", "domain":"domain", "email":"email", "nameZh":"name"}`,
			result: http.StatusCreated,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			w = httptest.NewRecorder()
			req, _ = http.NewRequest("POST", url, strings.NewReader(tC.data))
			req.Header.Add("Content-Type", "application/json")
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}

func setupHeader(req *http.Request, contentType, authToken string) {
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", "Bearer "+authToken)
}
