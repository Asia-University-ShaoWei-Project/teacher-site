package mock

import "fmt"

const (
	ApiVersion = "v1"
)

var (
	ApiUrl = fmt.Sprintf("/%s/api/%s", TeacherDomain, ApiVersion)
)
