package session

type SessionCommon interface {
	DeleteSession(token string) error
	IsSessionExists(token string) (bool, error)
}
