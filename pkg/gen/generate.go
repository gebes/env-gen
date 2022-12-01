package gen

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

var titleFormatter = cases.Title(language.English)

type Field struct {
	EnvName  string
	CodeName string
	Type     string
}
type Data struct {
	Fields []Field
	Config Config
}

func Generate(config Config) error {
	envMap, err := godotenv.Read(config.Env)
	if err != nil {
		return fmt.Errorf("could not parse .env file: %w", err)
	}
	tmpl, err := template.ParseFiles("./pkg/gen/templates/env.tmpl")
	if err != nil {
		return fmt.Errorf("parse template: %w", err)
	}

	var fields []Field
	for key, value := range envMap {
		field := Field{
			EnvName:  key,
			CodeName: toVariableName(key),
			Type:     "string",
		}

		switch {
		case isInt(value):
			field.Type = "int"
			break
		case isBool(value):
			field.Type = "bool"
			break
		}

		fields = append(fields, field)
	}

	sort.Slice(fields, func(i, j int) bool {
		return strings.Compare(fields[i].CodeName, fields[j].CodeName) < 0
	})

	file, err := os.Create(config.Output)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}

	err = tmpl.Execute(file, Data{
		fields,
		config,
	})
	if err != nil {
		return fmt.Errorf("execute template: %w", err)
	}
	err = file.Close()
	if err != nil {
		return fmt.Errorf("file close: %w", err)
	}

	return nil
}

func toVariableName(envKey string) string {
	split := strings.Split(envKey, "_")
	for i := range split {
		split[i] = titleFormatter.String(split[i])
	}
	return strings.Join(split, "")
}

func isInt(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}

func isBool(value string) bool {
	_, err := strconv.ParseBool(value)
	return err == nil
}
