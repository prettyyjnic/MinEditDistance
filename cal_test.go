package MinEditDistance

import "testing"
import "github.com/bmizerany/assert"

func TestMinEditDistance(t *testing.T) {
	distance0 := CalMinEditDistance([]rune("123"), []rune("123"))
	assert.Equal(t, distance0, 0)

	distance3 := CalMinEditDistance([]rune("sot"), []rune("stop"))
	assert.Equal(t, distance3, 3)

	distance8 := CalMinEditDistance([]rune("intention"), []rune("execution"))
	assert.Equal(t, distance8, 8)
}
