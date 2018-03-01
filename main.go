package main

import (
	"github.com/berryhill/mine-daemon/services"
)

func main () {

	go services.StartLogs()

	for {}
}
