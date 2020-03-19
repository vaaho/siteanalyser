package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExtractCsvColumn_1(t *testing.T) {
	val0 := ExtractCsvColumn("a;bb;ccc", 0)
	assert.Equal(t, "a", val0)

	val1 := ExtractCsvColumn("a;bb;ccc", 1)
	assert.Equal(t, "bb", val1)

	val2 := ExtractCsvColumn("a;bb;ccc", 2)
	assert.Equal(t, "ccc", val2)
}

func TestExtractCsvColumn_2(t *testing.T) {
	val3 := ExtractCsvColumn("a;bb;ccc", 3)
	assert.Equal(t, "", val3)

	val4 := ExtractCsvColumn("a;bb;ccc", 4)
	assert.Equal(t, "", val4)
}

func TestExtractCsvColumn_3(t *testing.T) {
	val := ExtractCsvColumn("a; bb  ; ccc", 1)
	assert.Equal(t, "bb", val)
}

func TestExtractCsvColumn_4(t *testing.T) {
	val := ExtractCsvColumn("a;\" bb \";ccc", 1)
	assert.Equal(t, " bb ", val)
}
