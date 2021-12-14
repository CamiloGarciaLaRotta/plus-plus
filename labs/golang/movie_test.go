package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMovie(t *testing.T) {
	toyStory := newMovie("Toy Story", Pixar, 1995)

	assert.Equal(t, "Toy Story", toyStory.title)
	assert.Equal(t, 1995, toyStory.yearPublished)
	assert.Equal(t, Pixar, toyStory.studio)
}

func TestEquality(t *testing.T) {
	toyStory := newMovie("Toy Story", Pixar, 1995)
	aBugsLife := newMovie("A Bug's Life", Pixar, 1998)

	assert.Equal(t, aBugsLife, aBugsLife)
	assert.Equal(t, toyStory, newMovie("Toy Story", Pixar, 1995))
	assert.Equal(t, toyStory, toyStory)
	assert.NotEqual(t, toyStory, aBugsLife)
}

func TestCompare(t *testing.T) {
	toyStory := newMovie("Toy Story", Pixar, 1995)
	aBugsLife := newMovie("A Bug's Life", Pixar, 1998)

	assert.Equal(t, 0, aBugsLife.compareTo(aBugsLife))
	assert.Equal(t, -1, aBugsLife.compareTo(toyStory))
	assert.Equal(t, 1, toyStory.compareTo(aBugsLife))
}
