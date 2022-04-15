package route

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"teacher-site/cache"
	"teacher-site/database"
	"teacher-site/logsrv"
	"teacher-site/mock"
	"teacher-site/model"
	"teacher-site/service"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	methodGet  = "GET"
	methodPOST = "POST"
)

var (
	cacheConf = model.NewMockCacheConfig()
	_cache    = cache.NewCache(cacheConf)
	db        = database.NewSqlite("../database", logger)
	logger    = logsrv.NewLogrus(ctx)
	ctx       = context.Background()
	conf      = model.NewMockServiceConfig()
	srv       = service.NewService(db, _cache, logger, conf)
	r         = gin.Default()
)

var (
	err           error
	url           string
	reqData       string
	reqDataformat string
	req           *http.Request
	w             *httptest.ResponseRecorder
	body          []byte
)

func TestInit(t *testing.T) {
	SetupRoute(ctx, r, srv)
	w := httptest.NewRecorder()
	url = mock.ApiURL + "/init"
	req, _ := http.NewRequest(methodGet, url, nil)
	r.ServeHTTP(w, req)
	defer w.Result().Body.Close()
	body, err = ioutil.ReadAll(w.Body)
	fmt.Println(string(body))
}

func TestUpdateInfo(t *testing.T) {
	SetupRoute(ctx, r, srv)
	// path := "/edit/info"
	// info := &model.BindInfo{ID: 3, Info: "test route with update info"}
	// json, _ := json.Marshal(info)
	// data = string(json)
	tC := []struct {
		desc       string
		statusPath string
		result     http.ConnState
	}{
		{
			desc:       "had been jwt",
			statusPath: "login",
			result:     http.StatusOK,
		},
		{
			desc:       "haven't jwt",
			statusPath: "logout",
			result:     http.StatusUnauthorized,
		},
	}
	for _, v := range tC {
		t.Run(v.desc, func(t *testing.T) {

			// r.POST("/edit/info", UpdateInfo(ctx, srv, data))
			// http.NewRequest("POST", "/auth/"+v.statusPath, strings.NewReader(``))
			// req, _ := http.NewRequest("POST", "/auth/getToken", strings.NewReader(``))
			// // req, _ := http.NewRequest("GET", path, strings.NewReader(`{"id": "1","name": "joe"}`))
			// w := httptest.NewRecorder()
			// r.ServeHTTP(w, req)
			// t.Logf("status: %d", w.Code)
			// t.Logf("response: %s", w.Body.String())
		})
	}
}

// todo: login instead of verify token
func TestVerifyToken(t *testing.T) {
	SetupRoute(ctx, r, srv)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", url+"/edit/test_token", nil)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmUiOjE2NDg3MjQ3MTEsImlzVXNlciI6InJpa2tpIn0.Emr4wb5s-JTLiqe8gimFDycEl2J0a3YUK6QTv8Ybvvo"
	req.Header.Add("Authorization", "Bearer "+token)
	r.ServeHTTP(w, req)
}

func TestLogin(t *testing.T) {
	reqDataformat = `{"user_id":"%s", "user_password":"%s"}`
	testCases := []struct {
		desc         string
		userID       string
		userPassword string
		result       int
	}{
		{
			desc:         "Empty fields",
			userID:       "",
			userPassword: mock.UserPassword,
			result:       http.StatusBadRequest,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			reqData = fmt.Sprintf(reqDataformat, tC.userID, tC.userPassword)
			url = mock.ApiURL + `/auth/login`

			req, err = http.NewRequest(methodPOST, url, strings.NewReader(reqData))
			if err != nil {
				t.Fatal(err)
			}
			req.Header.Add("Content-Type", "application/json")
			w = httptest.NewRecorder()
			r.ServeHTTP(w, req)
			assert.Equal(t, tC.result, w.Result().StatusCode, "not match")
		})
	}
}
