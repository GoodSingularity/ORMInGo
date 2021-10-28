package Interfaces
import(
	"github.com/fatih/structs"
	"fmt"
)

//Constuctor
func New(i interface{}) interface{} {

	if IsStruct(i) == false{
		return nil
	}
	
	{
	    fields := Fields(i)
	    err := 0
	    for j := 0; j<len(fields); j++ {
	    	field := fields[j]
	    	value := Values(i)[j]
		    if value != "" {
				fmt.Println(field+ " is initialized")
			}else {
				fmt.Println(field +" is not initialized")
				err = err +1
			}
	    }
	    if err > 0	{
	    	fmt.Println("You couldnt create instance of " + Name(i))
	    	return nil
	    }
    }
	fmt.Println("You've created instance of " + Name(i))
	return i
}

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
