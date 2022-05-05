package delivery

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"teacher-site/mock"
	"teacher-site/mock/info/repository"
	"teacher-site/mock/info/usecase"
	"teacher-site/pkg/util"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	JsonContentType = "application/json"
)

var (
	ctx   = context.Background()
	r     = gin.Default()
	route = r.Group("/:teacherDomain/api/v1/info")
	// logger      = log.NewLogrus(ctx)
	usecaseMock = usecase.NewUsecase()
	conf        = mock.Conf
	ApiUrl      = mock.ApiUrl + "/info"
)
var (
	url, data string
	err       error
	body      []byte
	req       *http.Request
	w         *httptest.ResponseRecorder
)

type HttpStatusCode int

func TestCreate(t *testing.T) {
	NewHandler(ctx, route, usecaseMock, conf)
	data := `{"content":""}`
	urlFormat := `/%v/bulletin`
	token, _ := util.GenerateJwt(conf.Jwt, mock.GetJwtRequest())

	testCases := []struct {
		desc   string
		token  string
		infoId interface{}
		data   string
		result HttpStatusCode
	}{
		{
			desc:   "unauthorized",
			token:  mock.EmptyStr,
			infoId: mock.NumPk,
			data:   data,
			result: http.StatusUnauthorized,
		},
		{
			desc:   "fail info id",
			token:  token,
			infoId: mock.NewMsg(),
			data:   data,
			result: http.StatusBadRequest,
		},
		{
			desc:   "empty bulletin content",
			token:  token,
			infoId: mock.NumPk,
			data:   mock.EmptyJson,
			result: http.StatusBadRequest,
		},
		{
			desc:   "normal",
			token:  token,
			infoId: mock.NumPk,
			data:   data,
			result: http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url = ApiUrl + fmt.Sprintf(urlFormat, tC.infoId)
			w = httptest.NewRecorder()
			req, _ = http.NewRequest("POST", url, strings.NewReader(tC.data))
			setupHeader(req, JsonContentType, tC.token)
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}

// todo: teacherDomain
// todo: none last_modified error handle
func TestGet(t *testing.T) {
	NewHandler(ctx, route, usecaseMock, conf)

	testCases := []struct {
		desc   string
		uri    string
		result HttpStatusCode
	}{
		{
			desc:   "empty last_modified",
			uri:    mock.EmptyStr,
			result: http.StatusOK,
		},
		{
			desc:   "have last_modified",
			uri:    `?last_modified=` + repository.CurrLastModidied,
			result: http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url = ApiUrl + `/bulletin` + tC.uri
			w = httptest.NewRecorder()
			req, _ = http.NewRequest("GET", url, nil)
			r.ServeHTTP(w, req)
			defer w.Result().Body.Close()
			body, err = ioutil.ReadAll(w.Body)
			// w.Result().StatusCode
			fmt.Println("📝 Body:", string(body))
		})
	}
}
func TestUpdate(t *testing.T) {
	NewHandler(ctx, route, usecaseMock, conf)
	dataFormat := `{"content":"%s"}`
	urlFormat := `/%s/bulletin/%s`
	token, _ := util.GenerateJwt(conf.Jwt, mock.GetJwtRequest())

	testCases := []struct {
		desc       string
		token      string
		infoId     string
		bulletinId string
		data       string
		result     HttpStatusCode
	}{
		{
			desc:       "unauthorized",
			token:      mock.EmptyStr,
			infoId:     mock.PkStr,
			bulletinId: mock.PkStr,
			data:       mock.WordStr,
			result:     http.StatusUnauthorized,
		},
		{
			desc:       "fail info id",
			token:      token,
			infoId:     mock.WordStr,
			bulletinId: mock.PkStr,
			data:       mock.WordStr,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "fail bulletin id",
			token:      token,
			infoId:     mock.PkStr,
			bulletinId: mock.WordStr,
			data:       mock.WordStr,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "empty bulletin content",
			token:      token,
			infoId:     mock.PkStr,
			bulletinId: mock.PkStr,
			data:       mock.EmptyStr,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "normal",
			token:      token,
			infoId:     mock.PkStr,
			bulletinId: mock.PkStr,
			data:       mock.WordStr,
			result:     http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url = ApiUrl + fmt.Sprintf(urlFormat, tC.infoId, tC.bulletinId)
			data = fmt.Sprintf(dataFormat, tC.data)
			w = httptest.NewRecorder()

			req, _ = http.NewRequest("PUT", url, strings.NewReader(data))
			setupHeader(req, JsonContentType, tC.token)
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Code))
		})
	}
}
func TestDelete(t *testing.T) {
	NewHandler(ctx, route, usecaseMock, conf)
	urlFormat := `/%v/bulletin/%v`
	token, _ := util.GenerateJwt(conf.Jwt, mock.GetJwtRequest())
	testCases := []struct {
		desc       string
		token      string
		infoId     interface{}
		bulletinId interface{}
		result     HttpStatusCode
	}{
		{
			desc:       "unauthorized",
			token:      mock.EmptyStr,
			infoId:     mock.NumPk,
			bulletinId: mock.NumPk,
			result:     http.StatusUnauthorized,
		},
		{
			desc:       "fail info id",
			token:      token,
			infoId:     mock.WordStr,
			bulletinId: mock.NumPk,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "fail bulletin id",
			token:      token,
			infoId:     mock.NumPk,
			bulletinId: mock.WordStr,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "empty bulletin content",
			token:      token,
			infoId:     mock.NumPk,
			bulletinId: mock.NumPk,
			result:     http.StatusBadRequest,
		},
		{
			desc:       "normal",
			token:      token,
			infoId:     mock.NumPk,
			bulletinId: mock.NumPk,
			result:     http.StatusOK,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			url = ApiUrl + fmt.Sprintf(urlFormat, tC.infoId, tC.bulletinId)
			w = httptest.NewRecorder()
			req, _ = http.NewRequest("DELETE", url, nil)
			setupHeader(req, JsonContentType, tC.token)
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, HttpStatusCode(w.Result().StatusCode))
		})
	}
}

func setupHeader(req *http.Request, contentType, authToken string) {
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("Authorization", "Bearer "+authToken)
}
