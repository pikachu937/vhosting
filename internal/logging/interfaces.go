package logging

type LogCommon interface {
	CreateLogRecord(log *Log) error
}

type LogUseCase interface {
	LogCommon
}

type LogRepository interface {
	LogCommon
}
