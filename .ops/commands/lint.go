package commands

import (
	"log"

	"lesiw.io/cmdio/sys"
)

func (Ops) Lint() {
	var rnr = sys.Runner().WithEnv(map[string]string{
		"PWD":	"./",
	})
	defer rnr.Close()

	err := rnr.Run("golangci-lint", "run")
	if err != nil {
		log.Fatal(err)
	}
}