package logger

type itophook interface {
	Call(l *Clog)
}
