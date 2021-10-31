package Migrate
import(
    "os"   
   "fmt"
      "io/fs"
   "path/filepath"
       "io/ioutil"
	cfg "Project/Config"

	
)

func Data(path string) string {
	file, err := ioutil.ReadFile(path)
    if err != nil {
        fmt.Println("Cannot read the file")
        os.Exit(1)
    }
    return string(file)
}

func Find(root, ext string) []string {
   var a []string
   filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
      if e != nil { return e }
      if filepath.Ext(d.Name()) == ext {
         a = append(a, s)
      }
      return nil
   })
   return a
}

func AutoMigrate(path string) error {
   db := cfg.Conn
   for _, s := range Find(path, ".migration") {
     path := "./"+s
	 query := Data(path)
	 db.Exec(query)
   }
   return nil
}
