package main

import (
	"flag"
	"fmt"

	"github.com/raismaulana/blogP/application"
	"github.com/raismaulana/blogP/application/registry"
)

func main() {
	appMap := map[string]func() application.RegistryContract{
		"usingdb": registry.NewUsingdb(),
	}

	flag.Parse()

	app, exist := appMap[flag.Arg(0)]
	if exist {
		application.Run(app())
	} else {
		fmt.Println("You may try this app name:")
		for appName := range appMap {
			fmt.Printf("%s\n", appName)
		}
	}

}
