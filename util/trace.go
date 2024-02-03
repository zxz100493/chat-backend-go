package util

import (
	"chat-go/infrastructurre/log"
	"time"
)

func TraceAction(action string) func() {
	start := time.Now()
	log.Logger.Info("before "+action, log.String("start_time", start.String()))

	return func() {
		log.Logger.Info("after "+action, log.String("used_time", time.Since(start).String()))
	}
}
