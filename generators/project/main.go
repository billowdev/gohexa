package main

import (
	"flag"
	"fmt"

	"github.com/rapidstellar/gohexa/generators/utils"
)

func main() {
	projectName := flag.String("name", "", "The name of the project")
	templateName := flag.String("template", "hexagonal", "The name of the template (default: hexagonal)")
	flag.Parse()

	if *projectName == "" {
		fmt.Println("Please provide a project name using the -name flag.")
		return
	}


	utils.CreateProject(*projectName, *templateName)
}
