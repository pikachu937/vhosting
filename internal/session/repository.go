package session

type SessRepository interface {
	SessCommon

	CreateSession(session Session) error
}
