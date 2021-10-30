package Config

import (
//  "database/sql"
	"fmt"
  _ "github.com/lib/pq"
          "database/sql"
  	"github.com/bengadbois/flippytext"

)

var Conn *sql.DB

func ConnectionInfo() string {
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  flippytext.New().Write(psqlInfo)
  return psqlInfo
}

func SetConnection() string{
	db, err := sql.Open("postgres", ConnectionInfo())
	if err != nil {
		return "Failed"
	}
	Conn = db
	return "Succeed"
}

func CloseConnection(){
	defer Conn.Close()
}
