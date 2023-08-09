package main

import (
	"employeeleave/delivery"

	_ "github.com/lib/pq"
)

func main() {
	delivery.NewServer().Run()
}
