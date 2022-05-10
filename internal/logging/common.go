package logging

type LoggingCommon interface {
	CreateLogRecord(log *Log) error
}
