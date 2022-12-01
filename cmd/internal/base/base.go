package base

import (
	"fmt"
	"github.com/Gebes/env-gen/pkg/gen"
	"github.com/spf13/cobra"
	"strings"
)

// GenerateCmd returns the generate command for ent/c packages.
func GenerateCmd(postRun ...func(*gen.Config)) *cobra.Command {
	var (
		cfg gen.Config
		cmd = &cobra.Command{
			Use:   "generate [flags]",
			Short: "generate env code from the .env file",
			Example: examples(
				"env generate ./ent/schema",
				"ent generate github.com/a8m/x",
			),
			Run: func(cmd *cobra.Command, path []string) {
				err := gen.Generate(cfg)
				if err != nil {
					fmt.Println("Could not generate config:", err)
				}
			},
		}
	)
	cmd.Flags().StringVar(&cfg.Env, "env", "./.env", "environment file the config should be based on")
	cmd.Flags().StringVar(&cfg.Output, "output", "./pkg/env/env.go", "output file")
	cmd.Flags().StringVar(&cfg.PackageName, "package", "env", "name of the package")
	cmd.Flags().BoolVar(&cfg.GodotEnvEnabled, "godotenv", true, "enables godotenv initial load")
	cmd.Flags().BoolVar(&cfg.GodotEnvLoggingEnabled, "godotenv-logging", false, "logs an error when godotenv fails")
	cmd.Flags().BoolVar(&cfg.ExitOnParseError, "exit-on-parse-error", true, "exits when a variable couldn't be parsed")
	cmd.Flags().BoolVar(&cfg.LogParseError, "log-parse-error", true, "logs when a variable couldn't be parsed")

	return cmd
}

// examples formats the given examples to the cli.
func examples(ex ...string) string {
	for i := range ex {
		ex[i] = "  " + ex[i] // indent each row with 2 spaces.
	}
	return strings.Join(ex, "\n")
}
