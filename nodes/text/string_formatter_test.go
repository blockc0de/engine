package text

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringFormatter(t *testing.T) {
	var s string
	var formatter StringFormatter

	s = formatter.Format("Hello {0}, we are greeting you here: {1}! {2}", []interface{}{"Bob", "blockc0de", 2022})
	assert.Equal(t, s, "Hello Bob, we are greeting you here: blockc0de! 2022")

	s = formatter.Format("Hello {1}, we are greeting you here: {0}! {2}", []interface{}{"blockc0de", "Bob", 2022})
	assert.Equal(t, s, "Hello Bob, we are greeting you here: blockc0de! 2022")

	s = formatter.Format("Hello {hello}, we are greeting you here: {blockc0de}! {year}", map[string]interface{}{
		"hello":     "Bob",
		"blockc0de": "blockc0de",
		"year":      2022,
	})
	assert.Equal(t, s, "Hello Bob, we are greeting you here: blockc0de! 2022")
}
