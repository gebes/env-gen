package gen

type Config struct {
	Env    string
	Output string

	PackageName      string
	LogParseError    bool
	ExitOnParseError bool

	GodotEnvEnabled        bool
	GodotEnvLoggingEnabled bool

	StringTypes bool
	IntTypes    bool
	BoolTypes   bool
}
