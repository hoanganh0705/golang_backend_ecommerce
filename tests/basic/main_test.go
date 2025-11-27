package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddOne(t *testing.T) {
	// using normal test

	// var (
	// 	input  = 1
	// 	output = 2
	// )

	// actual := AddOne(input)
	// if actual != output {
	// 	t.Errorf("AddOne(%d) = %d; want %d", input, actual, output)
	// }

	// use testify
	assert.Equal(t, AddOne(2), 3, "AddOne(2) should be 3")
	assert.NotEqual(t, 2, 3)
	assert.Nil(t, nil, nil)
	fmt.Println("This line will be executed even if the above assert fails")

	// khi assert fail thì vẫn tiếp tục thực hiện các test phía dưới
}

func TestAddOne_Require(t *testing.T) {
	// use require
	require.Equal(t, 2, 3)

	fmt.Println("This line will not be executed if the above require fails")

	// khi require fail thì dừng luôn, không thực hiện các test phía dưới
}

// use this command to run the test coverage:  go test -coverprofile=coverage.out
// then use this command to see the coverage report in html:  go tool cover -html=coverage.out -o coverage.html

func TestAddOne2(t *testing.T) {

	var (
		input  = 1
		output = 2
	)

	actual := AddOne2(1)
	if actual != output {
		t.Errorf("AddOne(%d) = %d; want %d", input, actual, output)
	}
}
