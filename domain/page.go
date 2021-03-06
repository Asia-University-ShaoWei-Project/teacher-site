package domain

import "context"

type PageUsecase interface {
	TeacherList(ctx context.Context, req *TeacherListRequest) (TeacherListResponse, error)
	// TeacherListByApi(ctx context.Context, req *TeacherListRequest) (TeacherListResponse, error)
	Home(ctx context.Context, req *HomeRequest) (HomeResponse, error)
	Login(ctx context.Context, userId, token string) error
}
type PageDbRepository interface {
	GetTeachers(ctx context.Context, limit, offset int) ([]TeacherResponse, error)
	GetTeacherByDomain(ctx context.Context, teacherDomain string) (Teachers, error)
	CheckAuthByIdAndToken(ctx context.Context, userId, token string) error
}
type PageCacheRepository interface {
}
type TeacherListRequest struct {
	Page uint `uri:"pageNumber"`
}

func (t *TeacherListRequest) SetToFirstPage() {
	t.Page = 1
}

// func (t *TeacherListRequest) SetPage(num int) {
// 	t.Page = num
// }

type TeacherListResponse struct {
	Teachers []TeacherResponse
}
type TeacherResponse struct {
	Domain string
	NameZh string
	NameUs string
}

type HomeRequest struct {
	Domain string `uri:"teacherDomain" binding:"required"`
}
type HomeResponse struct {
	Domain    string
	Email     string `json:"email"`
	NameZh    string `json:"name_zh"`
	NameUs    string `json:"name_us"`
	Office    string `json:"office"`
	Call      string `json:"call"`
	Education string `json:"education"`
}
