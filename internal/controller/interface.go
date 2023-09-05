package controller

type WatcherAPI interface {
	IsReady() bool
	StartNumber() uint64
}

type MonitorAPI interface {
	SlackNotify(msg string)
}
