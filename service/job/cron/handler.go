package cron

type Handler interface {
	DoTask()
}
