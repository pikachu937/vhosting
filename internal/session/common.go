package session

type SessCommon interface {
	DeleteSession(token string) error
	IsSessionExists(token string) (bool, error)
}
