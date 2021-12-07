package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	library := NewMovieLibrary()
	toyStory := NewMovie("Toy Story", Pixar, 1995)
	aBugsLife := NewMovie("A Bug's Life", Pixar, 1998)
	fantasia := NewMovie("Fantasia", Disney, 1940)
	library.Add(toyStory)
	library.Add(aBugsLife)
	library.Add(fantasia)

	assert.Equal(t, 3, library.TotalCount())
	assert.Equal(t, toyStory, library.All()[0])
	assert.Equal(t, aBugsLife, library.All()[1])
	assert.Equal(t, fantasia, library.All()[2])
}

func TestRemove(t *testing.T) {
	library := NewMovieLibrary()
	toyStory := NewMovie("Toy Story", Pixar, 1995)
	aBugsLife := NewMovie("A Bug's Life", Pixar, 1998)
	fantasia := NewMovie("Fantasia", Disney, 1940)
	library.Add(toyStory)
	library.Add(aBugsLife)
	library.Add(fantasia)

	library.Remove(aBugsLife)

	assert.Equal(t, 2, library.TotalCount())

	library.Each(func(movie Movie) {
		assert.NotEqual(t, aBugsLife, movie)
	})
}

func TestFindBy(t *testing.T) {
	library := NewMovieLibrary()
	toyStory := NewMovie("Toy Story", Pixar, 1995)
	aBugsLife := NewMovie("A Bug's Life", Pixar, 1998)
	fantasia := NewMovie("Fantasia", Disney, 1940)
	library.Add(toyStory)
	library.Add(aBugsLife)
	library.Add(fantasia)

	result := library.FindBy(func(movie Movie) bool {
		return movie.Title == "Fantasia"
	})

	assert.Equal(t, fantasia, result)
}

func TestSortBy(t *testing.T) {
	library := NewMovieLibrary()
	toyStory := NewMovie("Toy Story", Pixar, 1995)
	aBugsLife := NewMovie("A Bug's Life", Pixar, 1998)
	fantasia := NewMovie("Fantasia", Disney, 1940)
	library.Add(toyStory)
	library.Add(aBugsLife)
	library.Add(fantasia)

	result := library.SortBy(func(x Movie, y Movie) int {
		return strings.Compare(x.Title, y.Title)
	})

	assert.NotEqual(t, nil, result)
	assert.Equal(t, []*Movie{aBugsLife, fantasia, toyStory}, result)
}
