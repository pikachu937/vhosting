package session

type SessionRepository interface {
	SessionCommon

	CreateSession(sess Session) error
}
