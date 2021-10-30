package Driver

import(
	"reflect"
	"fmt"
	"strings"
	 pq "github.com/lib/pq"

)

func RemodelQuery(t []interface{}) string {
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
