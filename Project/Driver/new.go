package Driver

import(
	cfg "Project/Config"
	"github.com/mitchellh/colorstring"

)

func New(i interface{}) interface{} {

	if IsStruct(i) == false{
		return nil
	}
	
	{
	    fields := Fields(i)
	    values := Values(i)
	    err := 0
	    for j := 0; j<len(fields); j++ {
	    	value := values[j]
		    if value == "" {
				err = err +1
			}
	    }
	    if err > 0	{
	    	return nil
	    }
    }
    name := Name(i)
	v := RemodelQuery(Values(i))
	db := cfg.Conn
	query := "INSERT INTO " + name + " VALUES " + v
	colorstring.Println("[green]"+query)
	db.Exec(query)
	return i
}
