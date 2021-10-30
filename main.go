package main
import(
	cfg "Config"
	"fmt"
	"Interfaces"
	"log"
	"encoding/json"
)

func main(){
	var status = cfg.SetConnection()
	fmt.Println(status)
	cmp := Interfaces.Company{"Int",[]string{"kamil", "lukasz", "tycjan"}, 10}
	Interfaces.New(cmp)
  	all := Interfaces.All(Interfaces.Company{})

    prettyJSON, err := json.MarshalIndent(all, "", "    ")
    if err != nil {
        log.Fatal("Failed to generate json", err)
    }
    fmt.Printf("%s\n", string(prettyJSON))
}
