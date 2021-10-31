
package Driver
import(
	"github.com/fatih/structs"
)


func IsStruct(i interface{}) bool{
	return structs.IsStruct(i)
}

func Name(i interface{}) string {
	return structs.Name(i)
}

func Fields(i interface{}) []string {
	return structs.Names(i);
}

func Values(i interface{}) []interface{}{
	return structs.Values(i)
}


func Get(i interface{}, key string) interface{} {

	if IsStruct(i) == false{
		return nil
	}
	m := structs.Map(i)
	value := m[key]
	return value
}
