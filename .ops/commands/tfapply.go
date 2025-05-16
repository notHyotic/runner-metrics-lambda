package commands

import (
	"log"

	"lesiw.io/cmdio/sys"
)

func (Ops) Tfapply() {
	var rnr = sys.Runner().WithEnv(map[string]string{
		"PWD": "./terraform",
	})
	var err error
	defer rnr.Close()

	err = rnr.Run("terraform", "init")
	if err != nil {
		log.Fatal(err)
	}

	err = rnr.Run("terraform", "validate")
	if err != nil {
		log.Fatal(err)
	}

	err = rnr.Run("terraform", "apply", "-auto-approve")
	if err != nil {
		log.Fatal(err)
	}
}