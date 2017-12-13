package models

import (
	"fmt"
    "log"
	"database/sql"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)



func SearchMovie( minimum int, mintahun int, maxtahun int, genre1 string, genre2 string, genre3 string) (result []*Film) {
	var x []*Film
	var f Film
	

	//open connection
	db, err := sql.Open("mysql", "root:@tcp(167.205.67.251:3306)/cinema_management_system")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	
	//mengubah masukan genre ke dalam format %genre%
	newQuery := "%"+genre1+"%"
	newQuery2 := "%"+genre2+"%"
	newQuery3 := "%"+genre3+"%"


	//mengubah masukan tanggal menjadi string dengan format tglMin/tglMax-01-01
	tglMin := strconv.Itoa(mintahun)+"-01-01"
	tglMax := strconv.Itoa(maxtahun)+"-01-01"

	rows, err := db.Query("Select * from film where Rating=10")

	
		if (mintahun==0) {
			if (maxtahun==0) {

				rows, err = db.Query("SELECT * FROM film where Rating >= ? and Genre like \"" + newQuery + "\" and Genre like \"" + newQuery2 + "\" and Genre like \"" + newQuery3 + "\"",minimum)

			} else if (maxtahun>mintahun) {

				rows, err = db.Query("SELECT * FROM film where Rating >= ? and Tanggal_Tayang <= \"" + tglMax + "\" and Genre like \"" + newQuery + "\" and Genre like \"" + newQuery2 + "\" and Genre like \"" + newQuery3 + "\" ",minimum)

			}

		} else if (mintahun>0) {

			if (maxtahun==0) {

				rows, err = db.Query("SELECT * FROM film where Rating >= ? and Tanggal_Tayang >= \"" + tglMin + "\"  and Genre like \"" + newQuery + "\" and Genre like \"" + newQuery2 + "\" and Genre like \"" + newQuery3 + "\"",minimum)

			} else {

				rows, err = db.Query("SELECT * FROM film where Rating >= ? and Tanggal_Tayang >= \"" + tglMin + "\"  and Tanggal_Tayang <= \"" + tglMax + "\" and Genre like \"" + newQuery + "\" and Genre like \"" + newQuery2 + "\" and Genre like \"" + newQuery3 + "\" ",minimum)
			}

		} 
	
	


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
    //this.serveJSON()
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