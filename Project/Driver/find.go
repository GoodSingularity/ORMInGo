package Driver

import(
	cfg "Project/Config"
	"github.com/mitchellh/colorstring"
	"reflect"
)

func Find(u interface{}, id string) interface{}{
	db := cfg.Conn
    name := Name(u)
    query := "SELECT * FROM " + name + " WHERE " + "id = '"+id+"';"
    colorstring.Println("[blue]" + query)
	rows, _ := db.Query(query)

    for rows.Next() {

         t := reflect.TypeOf(u)
         val := reflect.New(t).Interface()

		 scan := Scanning(val)
         errScan := rows.Scan(scan...)

         if errScan != nil {
             //proper err handling
         }
    return val
    }
    return nil
}
