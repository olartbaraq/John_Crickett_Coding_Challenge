package tests

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimulateServerDown(t *testing.T) {
	ArrayOfTrueFalse := []bool{true, false}
	index := rand.Intn(len(ArrayOfTrueFalse))
	assert.Len(t, ArrayOfTrueFalse, 2)
	assert.NotEmpty(t, index)

}

func TestJoinToSingleSlash(t *testing.T) {

}
