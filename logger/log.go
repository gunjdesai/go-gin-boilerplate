package logger

type Logger struct {
}

func (l *Logger) init() {

}

func Initilize() *Logger {

	log := Logger{}
	log.init()

	return &log

}
