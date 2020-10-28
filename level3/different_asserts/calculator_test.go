package different_asserts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	assert := assert.New(t)

	calculator := Calculator{}
	assert.True(calculator.Add(0) == 0, "calculator.Add(0) should be equal 0")

	assert.True(calculator.Add(5) == 5, "calculator.Add(5) should be equal 5")

	assert.Equal(calculator.Add(8), float32(13), "calculator.Add(8) should be equal 13")
}

func TestDivide(t *testing.T) {

	assert := assert.New(t)

	calculator := Calculator{9}

	result, err := calculator.Divide(3)
	assert.NotEqual(result, float32(4), "calculator.Divide(3) should be different than 4")
	assert.Nil(err, "calculator.Add(0) should not return an error")

	_, err = calculator.Divide(0)
	assert.NotNil(err, "calculator.Add(0) should return an error")

}

func BenchmarkDevide(b *testing.B) {

	calculator := Calculator{9}

	calculator.Divide(3)
}
