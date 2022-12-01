package main

import (
	"github.com/Gebes/env-gen/cmd/internal/base"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	log.SetFlags(0)
	cmd := &cobra.Command{Use: "env"}
	cmd.AddCommand(
		base.GenerateCmd(),
	)
	_ = cmd.Execute()
}
