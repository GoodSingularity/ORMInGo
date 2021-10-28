package main
import(
	cfg "Config"
	"fmt"
)

func main(){
	var status = cfg.SetConnection()
	fmt.Println(status)
}
