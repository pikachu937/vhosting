package logging

type LogCommon interface {
	CreateLogRecord(log *Log) error
}
