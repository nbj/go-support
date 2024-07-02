package Feature

import (
	"fmt"
	"github.com/nbj/go-support/Support"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func Test_implements_function(t *testing.T) {
	// Arrange
	tests := []struct {
		Name     string
		Model    any
		Expected bool
	}{
		{"WithInterface", &TestWithInterface{}, true},
		{"WithoutInterface", &TestWithoutInterface{}, false},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("Test: %s", test.Name), func(t *testing.T) {
			t.Parallel()

			// Act
			implementsInterface := Support.Implements[TestInterface](test.Model)

			// Assert
			require.Equal(t, test.Expected, implementsInterface)
		})
	}
}

func Test_cast_function(t *testing.T) {
	// Arrange
	tests := []struct {
		Name     string
		Model    any
		Panicked bool
	}{
		{"CanCast", &TestWithInterface{}, false},
		{"CannotCast", &TestWithoutInterface{}, true},
	}

	for _, test := range tests {
		test := test
		t.Run(fmt.Sprintf("Test: %s", test.Name), func(t *testing.T) {
			panicked := false

			defer func() {
				if recover() != nil {
					panicked = true
				}
			}()

			// Act
			cast := Support.Cast[TestInterface](test.Model)

			// Assert
			assert.Equal(t, "*Feature.TestInterface", reflect.TypeOf(&cast).String())
			assert.Equal(t, panicked, test.Panicked)
		})
	}

}

type TestInterface interface {
	TestFunction()
}

type TestWithInterface struct{}

func (test *TestWithInterface) TestFunction() {}

type TestWithoutInterface struct{}
