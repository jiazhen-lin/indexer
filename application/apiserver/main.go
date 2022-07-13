package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

var (
	port = ":8081"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	fmt.Println("done...")
}
