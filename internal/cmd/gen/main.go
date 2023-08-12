package main

import (
	"github.com/becheran/smock/smock"
)

//go:generate go run ./
func main() {
	smock.GenerateMocks()
}
