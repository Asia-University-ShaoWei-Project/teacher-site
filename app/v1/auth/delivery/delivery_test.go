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
	url, data, dataFormat string
	// err                   error
	// body                  []byte
	req *http.Request
	w   *httptest.ResponseRecorder
)

type HttpStatusCode int

func TestLogin(t *testing.T) {
	NewAuthHandler(ctx, route, usecaseMock, conf.Jwt)
	dataFormat = `{"id":"%s", "password":"%s"}`
	url = ApiUrl + `/login`
	token, _ := util.GenerateJwt(conf.Jwt, mock.UserID)

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

// // todo: teacher_domain
// // todo: none last_modified error handle
// func TestGet(t *testing.T) {
// 	NewInfoHandler(ctx, route, usecaseMock, conf)

// 	testCases := []struct {
// 		desc   string
// 		uri    string
// 		result HttpStatusCode
// 	}{
// 		{
// 			desc:   "empty last_modified",
// 			uri:    mock.EmptyStr,
// 			result: http.StatusOK,
// 		},
// 		{
// 			desc:   "have last_modified",
// 			uri:    `?last_modified=` + repository.CurrLastModidied,
// 			result: http.StatusOK,
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			url := ApiUrl + `/bulletin` + tC.uri
// 			w := httptest.NewRecorder()
// 			req, _ := http.NewRequest("GET", url, nil)
// 			r.ServeHTTP(w, req)
// 			defer w.Result().Body.Close()
// 			body, err = ioutil.ReadAll(w.Body)
// 			// w.Result().StatusCode
// 			fmt.Println("📝 Body:", string(body))
// 		})
// 	}
// }
// func TestUpdate(t *testing.T) {
// 	NewInfoHandler(ctx, route, usecaseMock, conf)
// 	data := `{"content":""}`
// 	urlFormat := `/%v/bulletin/%v`
// 	token, _ := util.GenerateJwt(conf.Jwt, mock.UserID)

// 	testCases := []struct {
// 		desc       string
// 		token      string
// 		infoID     interface{}
// 		bulletinID interface{}
// 		data       string
// 		result     HttpStatusCode
// 	}{
// 		{
// 			desc:       "unauthorized",
// 			token:      mock.EmptyStr,
// 			infoID:     mock.NumPK,
// 			bulletinID: mock.NumPK,
// 			data:       data,
// 			result:     http.StatusUnauthorized,
// 		},
// 		{
// 			desc:       "fail info id",
// 			token:      token,
// 			infoID:     mock.StrWord,
// 			bulletinID: mock.NumPK,
// 			data:       data,
// 			result:     http.StatusBadRequest,
// 		},
// 		{
// 			desc:       "fail bulletin id",
// 			token:      token,
// 			infoID:     mock.NumPK,
// 			bulletinID: mock.StrWord,
// 			data:       data,
// 			result:     http.StatusBadRequest,
// 		},
// 		{
// 			desc:       "empty bulletin content",
// 			token:      token,
// 			infoID:     mock.NumPK,
// 			bulletinID: mock.NumPK,
// 			data:       mock.EmptyJson,
// 			// todo: concert the binding
// 			result: http.StatusBadRequest,
// 		},
// 		{
// 			desc:       "normal",
// 			token:      token,
// 			infoID:     mock.NumPK,
// 			bulletinID: mock.NumPK,
// 			data:       data,
// 			result:     http.StatusOK,
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			url := ApiUrl + fmt.Sprintf(urlFormat, tC.infoID, tC.bulletinID)
// 			w := httptest.NewRecorder()
// 			req, _ := http.NewRequest("PUT", url, strings.NewReader(tC.data))
// 			setupHeader(req, JsonContentType, tC.token)
// 			r.ServeHTTP(w, req)
// 			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
// 		})
// 	}
// }
// func TestDelete(t *testing.T) {
// 	NewInfoHandler(ctx, route, usecaseMock, conf)
// 	urlFormat := `/%v/bulletin/%v`
// 	token, _ := util.GenerateJwt(conf.Jwt, mock.UserID)
// 	testCases := []struct {
// 		desc       string
// 		token      string
// 		infoID     interface{}
// 		bulletinID interface{}
// 		result     HttpStatusCode
// 	}{
// 		{
// 			desc:       "unauthorized",
// 			token:      mock.EmptyStr,
// 			infoID:     mock.NumPK,
// 			bulletinID: mock.NumPK,
// 			result:     http.StatusUnauthorized,
// 		},
// 		{
// 			desc:       "fail info id",
// 			token:      token,
// 			infoID:     mock.StrWord,
// 			bulletinID: mock.NumPK,
// 			result:     http.StatusBadRequest,
// 		},
// 		{
// 			desc:       "fail bulletin id",
// 			token:      token,
// 			infoID:     mock.NumPK,
// 			bulletinID: mock.StrWord,
// 			result:     http.StatusBadRequest,
// 		},
// 		{
// 			desc:       "empty bulletin content",
// 			token:      token,
// 			infoID:     mock.NumPK,
// 			bulletinID: mock.NumPK,
// 			// todo: concert the binding
// 			result: http.StatusBadRequest,
// 		},
// 		{
// 			desc:       "normal",
// 			token:      token,
// 			infoID:     mock.NumPK,
// 			bulletinID: mock.NumPK,
// 			result:     http.StatusOK,
// 		},
// 	}
// 	for _, tC := range testCases {
// 		t.Run(tC.desc, func(t *testing.T) {
// 			url := ApiUrl + fmt.Sprintf(urlFormat, tC.infoID, tC.bulletinID)
// 			w := httptest.NewRecorder()
// 			req, _ := http.NewRequest("DELETE", url, nil)
// 			setupHeader(req, JsonContentType, tC.token)
// 			r.ServeHTTP(w, req)
// 			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
// 		})
// 	}
// }

func setupHeader(req *http.Request, contentType, authToken string) {
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", "Bearer "+authToken)
}
