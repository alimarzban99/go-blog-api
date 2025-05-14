package logging

type Logger interface {
	Init()
	Info(sub SubCategory, msg string)
	Error(sub SubCategory, msg string)
	Fatal(sub SubCategory, msg string)
}

func NewLogger() Logger {
	return newZapLogger()
}
