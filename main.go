package main
import(
	cfg "Config"
	"fmt"
	"Interfaces"
)

func main(){
	var status = cfg.SetConnection()
	fmt.Println(status)
	type Company struct{
		Name string
		Employees []string
		Power int
	}
	cmp := Company{"Int",[]string{"kamil", "luk"}, 10}
	Interfaces.New(cmp)
}
