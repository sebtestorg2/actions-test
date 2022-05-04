//go:generate go run gen.go

package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Hello world")
	logrus.Info("yo")
}
