package main
import(
	cfg "Config"
	"fmt"
	"Interfaces"
	"log"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/mitchellh/colorstring"
)

func main(){
	var status = cfg.SetConnection()
	fmt.Println(status)
	id := uuid.New()
	uuid := id.String()
	cmp := Interfaces.Company{uuid, "Int",[]string{"kamil", "lukasz", "tycjan"}, 10}
	Interfaces.New(cmp)
  	all := Interfaces.All(Interfaces.Company{})

    prettyJSON, err := json.MarshalIndent(all, "", "    ")
    if err != nil {
        log.Fatal("Failed to generate json", err)
    }
    colorstring.Println("[yellow]" +string(prettyJSON)+"\n")
    find := Interfaces.Find(Interfaces.Company{},"3abdf70c-6273-4644-b048-dbbc50d95094")

    prettyjSON, err := json.MarshalIndent(find, "", "    ")
    if err != nil {
        log.Fatal("Failed to generate json", err)
    }
    colorstring.Println("[yellow]" +string(prettyjSON)+"\n")
}
