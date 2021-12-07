package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStudioString(t *testing.T) {
	t.Skip()

	assert.Equal(t, "Castle Rock", CastleRock.String())
	assert.Equal(t, "Disney", Disney.String())
	assert.Equal(t, "Miramax Films", MiramaxFilms.String())
	assert.Equal(t, "Pixar", Pixar.String())
	assert.Equal(t, "Regency Enterprises", RegencyEnterprises.String())
}
