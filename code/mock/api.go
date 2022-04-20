package mock

import "fmt"

const (
	ApiVersion  = "v1"
	ApiAuthPath = "/auth/edit"
)

var (
	ApiURL = fmt.Sprintf("/api/%s/%s", ApiVersion, TeacherDomain)
)
