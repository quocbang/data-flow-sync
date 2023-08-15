package roles

import (
	"github.com/quocbang/data-flow-sync/server/utils/function"
)

var funcPermission map[function.FuncName]Roles

func HasPermission(f function.FuncName, role Roles) bool {
	if r, ok := funcPermission[f]; ok {
		return r == role
	}
	return false
}

func init() {
	funcPermission = map[function.FuncName]Roles{
		function.FuncName_UPLOAD_LIMITARY_HOUR: Roles_LEADER,
	}
}
