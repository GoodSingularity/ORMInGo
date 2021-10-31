package main
import(
	cfg "Project/Config"
	"fmt"
	db "Project/Driver"
	"log"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/mitchellh/colorstring"
	. "Project/Models"
)

func main(){
	var status = cfg.SetConnection()
	fmt.Println(status)
	id := uuid.New()
	uuids := id.String()
	cmp := Company{uuids, "Int",[]string{"kamil", "lukasz", "tycjan"}, 10}
	db.New(cmp)
  	all := db.All(Company{})

    prettyJSON, err := json.MarshalIndent(all, "", "    ")
    if err != nil {
        log.Fatal("Failed to generate json", err)
    }
    colorstring.Println("[yellow]" +string(prettyJSON)+"\n")
    find := db.Find(Company{},"3abdf70c-6273-4644-b048-dbbc50d95094")

    prettyjSON, err := json.MarshalIndent(find, "", "    ")
    if err != nil {
        log.Fatal("Failed to generate json", err)
    }
    rows := string(prettyjSON)
    colorstring.Println("[yellow]" +rows+"\n")
	ap := db.Get(find, "Name")
	newid := uuid.New()
	uid := newid.String()
	c := Company{uid, ap.(string), []string{}, 100}
	fmt.Println(c)
	zs := db.Set(find, "Name", "Text")
	
    j, err := json.MarshalIndent(zs, "", "    ")
    if err != nil {
        log.Fatal("Failed to generate json", err)
    }
    js := string(j)
    colorstring.Println("[yellow]" +js+"\n")
	
}
