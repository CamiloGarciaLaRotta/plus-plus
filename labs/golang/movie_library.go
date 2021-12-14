package main

type movieLibrary struct{}

func newMovieLibrary() *movieLibrary {
	return &movieLibrary{}
}

func (library *movieLibrary) add(movie *movie) {
}

func (library *movieLibrary) remove(movie *movie) {
}

func (library *movieLibrary) totalCount() int {
	return 0
}

func (library *movieLibrary) all() []movie {
	return []movie{}
}

func (library *movieLibrary) each(visitor func(movie)) {
}

func (library *movieLibrary) findBy(predicate func(movie) bool) *movie {
	return nil
}

func (library *movieLibrary) sortBy(comparer func(movie, movie) int) []movie {
	return []movie{}
}
