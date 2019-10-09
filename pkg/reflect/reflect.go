package reflect

import(
	"reflect"
)

var errorInterface = reflect.TypeOf((*error)(nil)).Elem()

// IsValueAnError checks if Value is of type error
func IsValueAnError(value reflect.Value) bool {
	if !value.IsValid() {
		return false
	}
	return value.Type().Implements(errorInterface)
}

// ValueToError converts Value to error
func ValueToError(value reflect.Value) error {
	err, ok := value.Interface().(error)
	if !ok {
		return nil
	}
	return err
}
