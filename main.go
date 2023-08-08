package main

import (
	"employeeleave/delivery"
)

func main() {
	// CLI
	// delivery.NewConsole().Run()
	// REST API
	delivery.NewServer().Run()
}
