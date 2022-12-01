package main

import (
	"fmt"
	"github.com/Gebes/env-gen/cmd/internal/base"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

func main() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	log.SetFlags(0)
	cmd := &cobra.Command{Use: "env"}
	cmd.AddCommand(
		base.GenerateCmd(),
	)
	_ = cmd.Execute()
}
