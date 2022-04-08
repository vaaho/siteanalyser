package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAction(t *testing.T) {
	var c Command

	c = ParseCommand("import")
	assert.Equal(t, Import, c)

	c = ParseCommand("analyse")
	assert.Equal(t, Analyse, c)

	c = ParseCommand("update-analyse")
	assert.Equal(t, UpdateAnalyse, c)

	c = ParseCommand("export")
	assert.Equal(t, Export, c)

	c = ParseCommand("stats")
	assert.Equal(t, Stats, c)

	c = ParseCommand("help")
	assert.Equal(t, Help, c)

	c = ParseCommand("-asdfasdf1234!")
	assert.Equal(t, Unknown, c)

	c = ParseCommand("")
	assert.Equal(t, Unknown, c)
}

func TestActionString(t *testing.T) {
	var s string

	s = Unknown.String()
	assert.Equal(t, "unknown", s)

	s = Import.String()
	assert.Equal(t, "import", s)

	s = Analyse.String()
	assert.Equal(t, "analyse", s)

	s = UpdateAnalyse.String()
	assert.Equal(t, "update-analyse", s)

	s = Export.String()
	assert.Equal(t, "export", s)

	s = Stats.String()
	assert.Equal(t, "stats", s)

	s = Help.String()
	assert.Equal(t, "help", s)
}
