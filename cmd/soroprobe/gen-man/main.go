// Command gen-man generates man pages from the soroprobe CLI help text.
//
// It is a development helper, not a released binary. The Makefile calls it
// through `make man`, and `.goreleaser.yaml` bundles the generated pages with
// each release archive. Generated pages are not committed to the repository;
// regenerate them whenever the CLI help text changes and include the new files
// in the next release.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra/doc"

	"github.com/soroworks/soroprobe/cmd/soroprobe"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Fprintf(os.Stderr, "usage: gen-man\n")
		os.Exit(1)
	}

	root, _ := soroprobe.NewRootCmd()

	if err := doc.GenManTree(root, nil, "docs/man"); err != nil {
		fmt.Fprintf(os.Stderr, "gen-man: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("generated man pages in docs/man")
}
