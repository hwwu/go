package main

import (
	"errors"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	assert := assert.New(t)

	assert.Error(errors.New("123123"), "123123123")
	// assert equality
	assert.Equal(123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(123, 456, "they should not be equal")

	er := assert.NotEqual(123, 123, "they should not be equal")

	fmt.Println(11111, er)

	type People struct {
		Value string
	}

	var object People
	// assert for nil (good for errors)
	assert.Nil(object)

	// assert for not nil (good when you expect something)
	if assert.NotNil(object) {

		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal("Something", object.Value)
	}
}
