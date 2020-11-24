package main

import (
	"fmt"

	"github.com/AllinChen/MyCfg/mycfg"
)

func main() {

	fmt.Println(mycfg.Read("test.cfg", "="))

}
