package route

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"teacher-site/cache"
	"teacher-site/database"
	"teacher-site/log"
	"teacher-site/model"
	"teacher-site/service"
	"testing"

	"github.com/gin-gonic/gin"
)

var (
	apiVersion = "v1"
	domain     = "teacher_domain"
	url        = fmt.Sprintf("/%s/%s", apiVersion, domain)
	db         = database.NewSqlite("../database")
	c          = cache.NewCache()
	ctx        = context.Background()
	// logger     = log.NewLog(ctx)
	logger = log.NewLogrus(ctx)
	srv    = service.NewService(db, c, logger)
	cfg    = &model.Config{
		JWTSecure: []byte(`secure`),
		// jwtSecure: []byte(os.Getenv(`secure`)),
	}
	r = gin.Default()
)

func TestUpdateInfo(t *testing.T) {
	SetupRoute(ctx, r, srv, cfg)
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
func TestInit(t *testing.T) {
	w := setupHTTP("GET", url+"/init", nil)
	data := &gin.H{
		"data": model.Init{},
	}
	err := json.Unmarshal(w.Body.Bytes(), data)
	if err != nil {
		t.Fatal(err)
	}
	j, _ := json.MarshalIndent(data, "", " ")
	srv.Debug(string(j))
}

func TestVerifyToken(t *testing.T) {

	SetupRoute(ctx, r, srv, cfg)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", url+"/edit/test_token", nil)
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmUiOjE2NDg3MjQ3MTEsImlzVXNlciI6InJpa2tpIn0.Emr4wb5s-JTLiqe8gimFDycEl2J0a3YUK6QTv8Ybvvo"
	req.Header.Add("Authorization", "Bearer "+token)
	r.ServeHTTP(w, req)
}
func setupHTTP(method, _url string, body io.Reader) *httptest.ResponseRecorder {
	SetupRoute(ctx, r, srv, cfg)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, _url, body)
	r.ServeHTTP(w, req)
	return w
}

// auth.POST("/login", Login(ctx, srv))
// auth.POST("/logout", Logout(ctx, srv))
// auth.POST("/getToken", func(c *gin.Context) {
// 	token := sessions.Default(c).Get("token")
// 	c.String(200, "token:%s", token)
// })
