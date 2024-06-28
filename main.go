package main

import (
	"fmt"

	"github.com/lkmproject/bigtwo/core"
	"github.com/lkmproject/bigtwo/helper"
)

func main() {
	core.NewApp()
	fmt.Println(helper.GetName())
}
