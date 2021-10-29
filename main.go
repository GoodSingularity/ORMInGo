package main
import(
	cfg "Config"
	"fmt"
	"Interfaces"
	"encoding/json"
	"log"

)

func main(){
	var status = cfg.SetConnection()
	fmt.Println(status)
	cmp := Interfaces.Company{"Int", 10}
	Interfaces.New(cmp)
  	var cmps Interfaces.Company
  	jsonInfo, _ := json.Marshal(Interfaces.All(cmps))
	log.Printf("jsonInfo: %s\n", jsonInfo)
}
