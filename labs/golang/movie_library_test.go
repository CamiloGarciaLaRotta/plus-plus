package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	library := newMovieLibrary()
	toyStory := newMovie("Toy Story", Pixar, 1995)
	aBugsLife := newMovie("A Bug's Life", Pixar, 1998)
	fantasia := newMovie("Fantasia", Disney, 1940)
	library.add(toyStory)
	library.add(aBugsLife)
	library.add(fantasia)

	assert.Equal(t, 3, library.totalCount())
	assert.Equal(t, toyStory, library.all()[0])
	assert.Equal(t, aBugsLife, library.all()[1])
	assert.Equal(t, fantasia, library.all()[2])
}

func TestRemove(t *testing.T) {
	library := newMovieLibrary()
	toyStory := newMovie("Toy Story", Pixar, 1995)
	aBugsLife := newMovie("A Bug's Life", Pixar, 1998)
	fantasia := newMovie("Fantasia", Disney, 1940)
	library.add(toyStory)
	library.add(aBugsLife)
	library.add(fantasia)

	library.remove(aBugsLife)

	assert.Equal(t, 2, library.totalCount())

	library.each(func(movie movie) {
		assert.NotEqual(t, aBugsLife, movie)
	})
}

func TestFindBy(t *testing.T) {
	library := newMovieLibrary()
	toyStory := newMovie("Toy Story", Pixar, 1995)
	aBugsLife := newMovie("A Bug's Life", Pixar, 1998)
	fantasia := newMovie("Fantasia", Disney, 1940)
	library.add(toyStory)
	library.add(aBugsLife)
	library.add(fantasia)

	result := library.findBy(func(movie movie) bool {
		return movie.title == "Fantasia"
	})

	assert.Equal(t, fantasia, result)
}

func TestSortBy(t *testing.T) {
	library := newMovieLibrary()
	toyStory := newMovie("Toy Story", Pixar, 1995)
	aBugsLife := newMovie("A Bug's Life", Pixar, 1998)
	fantasia := newMovie("Fantasia", Disney, 1940)
	library.add(toyStory)
	library.add(aBugsLife)
	library.add(fantasia)

	result := library.sortBy(func(x movie, y movie) int {
		return strings.Compare(x.title, y.title)
	})

	assert.NotEqual(t, nil, result)
	assert.Equal(t, []*movie{aBugsLife, fantasia, toyStory}, result)
}
