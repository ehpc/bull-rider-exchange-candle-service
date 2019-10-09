package reflect

import(
	"testing"
	"errors"
	"reflect"

	"github.com/stretchr/testify/assert"
)

func TestIsValueAnError(t *testing.T) {
	t.Run("test real error", func (t *testing.T) {
		value := reflect.ValueOf(errors.New("test"))
		assert.True(t, IsValueAnError(value))
	})
	t.Run("test string", func (t *testing.T) {
		value := reflect.ValueOf("test")
		assert.False(t, IsValueAnError(value))
	})
	t.Run("test slice", func (t *testing.T) {
		value := reflect.ValueOf([]string{"test"})
		assert.False(t, IsValueAnError(value))
	})
	t.Run("test nil", func (t *testing.T) {
		value := reflect.ValueOf(nil)
		assert.False(t, IsValueAnError(value))
	})
	t.Run("test slice of errors", func (t *testing.T) {
		value := reflect.ValueOf([]error{errors.New("test")})
		assert.False(t, IsValueAnError(value))
	})
}

func TestValueToError(t *testing.T) {
	t.Run("test real error", func (t *testing.T) {
		value := reflect.ValueOf(errors.New("test"))
		assert.Error(t, ValueToError(value))
	})
	t.Run("test not error", func (t *testing.T) {
		value := reflect.ValueOf("test")
		assert.Nil(t, ValueToError(value))
	})
}