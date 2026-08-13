package logrus

func (entry *Entry) LogRaw(level Level, args ...any) {
	entry.logArgs(level, false, args...)
}

func (entry *Entry) LogfRaw(level Level, format string, args ...any) {
	entry.logf(level, false, format, args...)
}

func (entry *Entry) LoglnRaw(level Level, args ...any) {
	entry.logln(level, false, args...)
}
