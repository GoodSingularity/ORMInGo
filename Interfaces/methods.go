
package Interfaces
import(
	"github.com/fatih/structs"
	"fmt"
	 pq "github.com/lib/pq"
	"strings"
	cfg "Config"
	"reflect"
)

	type Company struct{
		ID   string
		Name string
		Arr []string
		Power int
	}

func CurlyBrackets(t []interface{}) string {
	s := make([]string, len(t))

	for i, v := range t {
	   s[i] = fmt.Sprint(v)
	}
	replace := map[string]string{
	    "[": "{",
	    "]": "}",
		" ": ",",
	}
	str := strings.Join(s,"' '")
	for s, r := range replace {
	    str = strings.Replace(str, s, r, -1)
	}
	
	return "('"+str+"');"
}

//Constuctor
func New(i interface{}) interface{} {

	if IsStruct(i) == false{
		return nil
	}
	
	{
	    fields := Fields(i)
	    values := Values(i)
	    err := 0
	    for j := 0; j<len(fields); j++ {
	    	field := fields[j]
	    	value := values[j]
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
    name := Name(i)
	fmt.Println("You've created instance of " + name)
	v := CurlyBrackets(Values(i))
	fmt.Println(v)
	db := cfg.Conn
	query := "INSERT INTO " + name + " VALUES " + v
	fmt.Println(query)
	_, err := db.Exec(query)
	fmt.Println(err)
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

func All(u interface{}) []interface{}{
	db := cfg.Conn
    name := Name(u)
    query := "SELECT * FROM " + name + ";"
	rows, _ := db.Query(query)

  	var data []interface{}
    for rows.Next() {

         t := reflect.TypeOf(u)
         val := reflect.New(t).Interface()

		 scan := Scanning(val)
         errScan := rows.Scan(scan...)

         if errScan != nil {
             //proper err handling
         }
         data = append(data, val)
    }
    return data
}

func Scanning(u interface{}) []interface{} {
   val := reflect.ValueOf(u).Elem()
   v := make([]interface{}, val.NumField())
   for i := 0; i < val.NumField(); i++ {
   
      valueField := val.Field(i)
      v[i] = valueField.Addr().Interface()
	  
	  if fmt.Sprintf("%v", v[i]) == "&[]" {
	  	v[i] =  pq.Array(*&v[i])
	  }
   }
   return v
}


func Find(u interface{}, id string) []interface{}{
	db := cfg.Conn
    name := Name(u)
    query := "SELECT * FROM " + name + " WHERE " + "id = '"+id+"';"
    fmt.Println(query)
	rows, _ := db.Query(query)

  	var data []interface{}
    for rows.Next() {

         t := reflect.TypeOf(u)
         val := reflect.New(t).Interface()

		 scan := Scanning(val)
         errScan := rows.Scan(scan...)

         if errScan != nil {
             //proper err handling
         }
         data = append(data, val)
    }
    return data
}
