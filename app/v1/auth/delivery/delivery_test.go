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
	route       = r.Group("/:teacher_domain/api/v1/auth")
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
		userID       string
		userPassword string
		isLogged     bool
		result       HttpStatusCode
	}{
		{
			desc:         "logged in",
			userID:       mock.UserID,
			userPassword: mock.UserPassword,
			isLogged:     true,
			result:       http.StatusFound,
		},
		{
			desc:         "fail request data",
			userID:       mock.EmptyStr,
			userPassword: mock.EmptyStr,
			isLogged:     false,
			result:       http.StatusBadRequest,
		},
		{
			desc:         "normal",
			userID:       mock.UserID,
			userPassword: mock.UserPassword,
			isLogged:     false,
			result:       http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			w = httptest.NewRecorder()
			data = fmt.Sprintf(dataFormat, tC.userID, tC.userPassword)
			req, _ = http.NewRequest(methodPost, url, strings.NewReader(data))
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
			req, _ = http.NewRequest(methodPost, url, nil)
			setupHeader(req, JsonContentType, tC.authHeader)
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}

func setupHeader(req *http.Request, contentType, authToken string) {
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", "Bearer "+authToken)
}
