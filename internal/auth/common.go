package auth

type AuthCommon interface {
	GetNamepass(namepass Namepass) error
	UpdateNamepassPassword(namepass Namepass) error
	IsNamepassExists(username, passwordHash string) (bool, error)
}
