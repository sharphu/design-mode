package prototype

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestKeywords_Clone(t *testing.T) {
	updateAt, _ := time.Parse("2006", "2020")
	words := Keywords{
		"testA": &Keyword{
			Word:      "testA",
			Visit:     1,
			UpdatedAt: &updateAt,
		},
		"testB": &Keyword{
			Word:      "testB",
			Visit:     2,
			UpdatedAt: &updateAt,
		},
		"testC": &Keyword{
			Word:      "testC",
			Visit:     3,
			UpdatedAt: &updateAt,
		},
	}

	now := time.Now()
	updatedWords := []*Keyword{
		{
			Word:      "testB",
			Visit:     10,
			UpdatedAt: &now,
		},
	}

	got := words.Clone(updatedWords)

	assert.Equal(t, words["testA"], got["testA"])
	assert.NotEqual(t, words["testB"], got["testB"])
	assert.Equal(t, updatedWords[0].Word, got["testB"].Word)
	assert.Equal(t, updatedWords[0].Visit, got["testB"].Visit)
	assert.True(t, got["testB"].UpdatedAt.Equal(*updatedWords[0].UpdatedAt), "UpdatedAt fields should be equal")
	assert.Equal(t, words["testC"], got["testC"])
}
