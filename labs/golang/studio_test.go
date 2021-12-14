package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStudioString(t *testing.T) {
	assert.Equal(t, "Castle Rock", CastleRock)
	assert.Equal(t, "Disney", Disney)
	assert.Equal(t, "Miramax Films", MiramaxFilms)
	assert.Equal(t, "Pixar", Pixar)
	assert.Equal(t, "Regency Enterprises", RegencyEnterprises)
}
