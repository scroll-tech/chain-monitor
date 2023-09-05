package controller

// WatcherAPI watcher api, relate to l1watcher and l2watcher.
type WatcherAPI interface {
	IsReady() bool
	StartNumber() uint64
}

// MonitorAPI monitor public api, used by l1watcher and l2watcher.
type MonitorAPI interface {
	SlackNotify(msg string)
}
