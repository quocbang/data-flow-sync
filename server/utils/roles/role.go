package roles

import (
	"github.com/quocbang/data-flow-sync/server/utils/function"
)

var funcPermission map[function.FuncName][]Roles

func HasPermission(f function.FuncName, role Roles) bool {
	if rs, ok := funcPermission[f]; ok {
		for _, r := range rs {
			if r == role {
				return true
			}
		}
	}
	return false
}

func init() {
	funcPermission = map[function.FuncName][]Roles{
		function.FuncName_UPLOAD_LIMITARY_HOUR: {
			Roles_LEADER,
		},
		function.FuncName_CREATE_STATION_MERGE_REQUEST: {
			Roles_USER,
			Roles_LEADER,
		},
		function.FuncName_GET_STATION_MERGE_REQUEST: {
			Roles_USER,
			Roles_LEADER,
			Roles_ADMINISTRATOR,
			Roles_UNSPECIFIED,
		},
	}
}
