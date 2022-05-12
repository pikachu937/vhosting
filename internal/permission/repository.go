package permission

type PermRepository interface {
	PermCommon

	CreatePermission(permission Permission) error
}
