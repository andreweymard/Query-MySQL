package main

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main(){

	db, err := sql.Open("myysql", "<USERNAME>:<PASSWORD>@tcp(<IP>:<PORT>)/<DB>")
	errhandle(err)
	defer db.Close()

	Club := "FC Bayern München"
	rows, err := db.Query("SELECT Name FROM data WHERE Club=?", Club)
	errhandle(err)
	defer rows.Close()

	for rows.Next() {
        var name string
        	if err := rows.Scan(&name); err != nil {
                log.Fatal(err)
        	}
        fmt.Printf("%s is %s\n", name, Club)
	}
	errhandle(err)
}

func errhandle(err error){
	if err != nil {
		panic(err.Error())
	}
}
