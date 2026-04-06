package main

import (
	_ "embed"

	"github.com/dcssoftware/bafoeg-manager/src/cmd"

	_ "github.com/dcssoftware/bafoeg-manager/swagger-docs"
)

func main() {
	cmd.Execute()
}
