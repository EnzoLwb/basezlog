package cuslog


func (l *logger) Debug(args ...interface{}) {
	l.log.Debug(args)
}
func (l *logger) Debugf(format string, args ...interface{}) {
	l.log.Debugf(format,args)
}

func (l *logger) Info(args ...interface{}) {
	l.log.Info(args)
}
func (l *logger) Infof(format string, args ...interface{}) {
	l.log.Infof(format,args)
}

// std logger
func Debug(args ...interface{}) {
	std.Debug(args)
}

func Debugf(format string, args ...interface{}) {
	std.Debugf(format,args)
}


func Info(args ...interface{}) {
	std.Info(args)
}

func Infof(format string, args ...interface{}) {
	std.Infof(format,args)
}