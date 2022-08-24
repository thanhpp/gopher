package bootstrap

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitTerminateSignals() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	select {
	case <-sigChan:
		return
	}
}
