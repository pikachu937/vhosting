package permission

type PermCommon interface {
	IsPermissionExists(id int) (bool, error)
}
