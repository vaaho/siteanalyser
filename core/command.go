package core

type Command int

const (
	Unknown Command = iota
	Import
	Analyse
	UpdateAnalyse
	Export
	Stats
	Help
)

var commandNames = []string{
	"unknown",
	"import",
	"analyse",
	"update-analyse",
	"export",
	"stats",
	"help",
}

func (c Command) String() string {
	return commandNames[c]
}

func ParseCommand(input string) Command {
	for i, name := range commandNames {
		if name == input {
			return Command(i)
		}
	}
	return Unknown
}
