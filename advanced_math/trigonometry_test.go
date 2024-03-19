package advanced_math

import (
	"github.com/stretchr/testify/assert"
	"github.com/tiwarishubham635/go-calculator/math"
	"testing"
)

func TestSin(t *testing.T) {
	angle, err := math.FloatDivision(PI, 2)
	assert.Nil(t, err)
	sine := Sin(angle)
	assert.Equal(t, sine, 1.0)

	angle, err = math.FloatDivision(PI, 6)
	assert.Nil(t, err)
	sine = Sin(angle)
	assert.Equal(t, sine, 0.5)
}

func TestCos(t *testing.T) {
	angle, err := math.FloatDivision(PI, 2)
	assert.Nil(t, err)
	cosine := Cos(angle)
	assert.Equal(t, cosine, 0.0)

	cosine = Cos(0.0)
	assert.Equal(t, cosine, 1.0)
}

func TestTan(t *testing.T) {
	angle, err := math.FloatDivision(PI, 4)
	assert.Nil(t, err)
	tan, err := Tan(angle)
	assert.Nil(t, err)
	assert.Equal(t, tan, 1.0)

	tan, err = Tan(0.0)
	assert.Nil(t, err)
	assert.Equal(t, tan, 0.0)
}
