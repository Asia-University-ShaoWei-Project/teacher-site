package mock

import "fmt"

const (
	ApiVersion = "v1"
)

var (
	ApiTeacherDomainFormatUrl = `/%s/api/` + ApiVersion
	ApiUrl                    = fmt.Sprintf("/%s/api/%s", TeacherDomain, ApiVersion)
)
