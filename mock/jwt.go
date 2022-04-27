package mock

import "teacher-site/domain"

func GetJwtRequest() *domain.JwtInfoRequest {
	return &domain.JwtInfoRequest{
		UserId: UserId,
		Domain: TeacherDomain,
	}
}
