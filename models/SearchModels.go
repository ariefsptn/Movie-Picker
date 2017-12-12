package models

import (
	"fmt"
    "log"
	"database/sql"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)



func IsAuthentic( minimum int, mintahun int, maxtahun int, genre1 string, genre2 string, genre3 string) (result []*Film) {
	var x []*Film
	var f Film
	//var w http.ResponseWriter


	//open connection
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/cinema_management_system")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	//check if member exists
	//var isAvailable bool
	newQuery := "%"+genre1+"%"
	tglMin := strconv.Itoa(mintahun)+"-01-01"
	//tglMax := strconv.Itoa(maxtahun)+"-01-01"
	/*rows, err := db.Query("SELECT * FROM film where Rating >= \" + minimum + \" and Genre like \"" + newQuery + "\" and Tanggal_Tayang >= \"" + tglMin + "\" and Tanggal_Tayang <= \"" + tglMax + "\"  ")*/
	rows, err := db.Query("SELECT * FROM film where Rating >= ? and Tanggal_Tayang >= \"" + tglMin + "\"  and Genre like \"" + newQuery + "\" ",minimum)
	if err!= nil {
		log.Fatal(err)
	}

	defer rows.Close()
        for rows.Next() {

                err := rows.Scan(&f.ID_Film, &f.Judul_Film, &f.Durasi, &f.Tanggal_Tayang, &f.Rating, &f.Genre, &f.Sinopsis)
                if err != nil{
                        log.Fatal(err)
                } 
                var y = new(Film)
		        y.ID_Film = f.ID_Film 
		        y.Judul_Film = f.Judul_Film
		        y.Durasi = f.Durasi
		        y.Tanggal_Tayang = f.Tanggal_Tayang
		        y.Rating = f.Rating
		        y.Genre = f.Genre
		        y.Sinopsis = f.Sinopsis
		        x= append(x,y)
                
        }
        err=rows.Err()
        if err != nil {
                log.Fatal(err)
        }
    return x
}



func PrintArray(result []*Film) {

	var i int
	for i=0;i<len(result);i++ { 
      fmt.Print(result[i].ID_Film)
       fmt.Print("")
      fmt.Printf(result[i].Judul_Film)
      fmt.Print(result[i].Durasi)
      fmt.Print("")
      fmt.Printf(result[i].Tanggal_Tayang)
      fmt.Print(result[i].Rating)
       fmt.Print("")
      fmt.Printf(result[i].Genre)
      fmt.Printf(result[i].Sinopsis)
    } 
}