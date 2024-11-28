package main

type perm string

const (
	Read  perm = "read"
	Write perm = "write"
	Exec  perm = "execute"
)

var Admin = "admin"
var User = perm("user")

func checkPermission(p perm) {
	// check the permission
}

func main() {
	checkPermission(perm(Admin))
	checkPermission(User)
}
