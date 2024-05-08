package constants

import (
	"os"
	"syscall"
)

var SignalsToShutdown = []os.Signal{syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM, os.Interrupt}

const FieldDatabase = "database"
