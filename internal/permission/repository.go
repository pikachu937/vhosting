package permission

type PermRepository interface {
	PermCommon

	CreatePermission(perm Permission) error
}
