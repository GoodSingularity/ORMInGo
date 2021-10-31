
package Driver
import(
	"github.com/fatih/structs"
)

func Get(i interface{}, key string) interface{} {

	if IsStruct(i) == false{
		return nil
	}
	m := structs.Map(i)
	value := m[key]
	return value
}

