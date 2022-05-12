package permission

type PermCommon interface {
	CreatePermission(permission Permission) error
	GetPermission(id int) (*Permission, error)
	GetAllPermissions() (map[int]*Permission, error)
	PartiallyUpdatePermission(permission *Permission) error
	DeletePermission(id int) error

	IsPermissionExists(id int) (bool, error)
}
