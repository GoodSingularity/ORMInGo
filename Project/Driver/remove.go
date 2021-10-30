package Driver

import(
	cfg "Project/Config"
	"github.com/mitchellh/colorstring"
)

func Remove(u interface{}, id string) string{
	db := cfg.Conn
    name := Name(u)
    {
	    findQuery := "SELECT * FROM " + name + " WHERE " + "id = '"+id+"';"
	    colorstring.Println("[blue]" + findQuery)
		_, err := db.Query(findQuery)
		if err != nil{
			return "NOT FOUND"
		}
	}
	query := "DELETE FROM " + name + " WHERE id = '" + id + "';"
	db.Exec(query)
	return "Deleted"
}
