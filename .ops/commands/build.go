package commands

import (
	"log"
	"os"
	"path/filepath"

	"lesiw.io/cmdio/sys"
)

func (Ops) Build() {
var rnr = sys.Runner().WithEnv(map[string]string{
        "GOOS":   "linux",
        "GOARCH": "amd64",
    })
    defer rnr.Close()

    // Build the binary for AWS Lambda (Linux)
    output := "main"
    err := rnr.Run("go", "build", "-buildvcs=false", "-o", output, ".")
    if err != nil {
        log.Fatal(err)
    }

    // Ensure the build directory exists
    buildDir := "build"
    if err := os.MkdirAll(buildDir, 0755); err != nil {
        log.Fatalf("Failed to create build directory: %v", err)
    }

    // Zip the binary for Lambda deployment
    zipPath := filepath.Join(buildDir, "runner-metrics-lambda.zip")
    err = rnr.Run("zip", "-j", zipPath, output)
    if err != nil {
        log.Fatal(err)
    }

    // Optionally, clean up the binary after zipping
    err = os.Remove(output)
    if err != nil {
        log.Printf("Warning: could not remove binary: %v", err)
    }
}
