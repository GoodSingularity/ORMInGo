package Driver
import(
	"reflect"
)

func Set(i interface{}, field string, value interface{}) interface{}{
	reflect.ValueOf(i).Elem().FieldByName(field).Set(reflect.ValueOf(value))
	return i
}

