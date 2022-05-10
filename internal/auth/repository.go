package auth

type AuthRepository interface {
	AuthCommon

	UpdateNamepassLastLogin(username, token string) error
}
