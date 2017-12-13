package controllers

import (
	
	"Movie-Picker/models"
    "math/rand"

    "encoding/gob"
)

type SearchController struct {
	GeneralExtendedController
}

type Authentication struct {
	Minimum int `form:"minimum"`
	Maxtahun int `form:"maxtahun"`
	Mintahun int `form:"mintahun"`
	Genre1 string `form:"genre1"`
	Genre2 string `form:"genre1"`
	Genre3 string `form:"genre1"`
}



func (this *SearchController) Post() {
	
	session := this.StartSession()
	auth := Authentication{}
	gob.Register(Authentication{})
	var x []*models.Film
	err := this.ParseForm(&auth)

	if err != nil {
		this.Data["title"] = "Form parsing failed"
	} else {
		
		session.Set("Minimum",auth.Minimum)
		if (auth.Mintahun>=1940) {
		session.Set("Mintahun",auth.Mintahun)
		} else {
			auth.Mintahun=1940
			session.Set("Mintahun",auth.Mintahun)
		}

		if (auth.Maxtahun<=2020) {
		session.Set("Maxtahun",auth.Maxtahun)
		} else {
			auth.Mintahun=2020
			session.Set("Maxtahun",auth.Maxtahun)
		}

		session.Set("Genre1",auth.Genre1)

		if (auth.Genre2!="") {
		session.Set("Genre2",auth.Genre2)
		} else {
			auth.Genre2 = "nil"
			session.Set("Genre2",auth.Genre2)
		} 

		if (auth.Genre3!="") {
			session.Set("Genre3",auth.Genre3)
		} else {
			auth.Genre3= "nil"
			session.Set("Genre3",auth.Genre3)
		}
		x = models.SearchMovie(auth.Minimum, auth.Mintahun, auth.Maxtahun, auth.Genre1, auth.Genre2, auth.Genre3)
		
	}
	
	
	
	

	if (x == nil) {
		
		this.TplName = "search-result-nil.tpl"
	} else {
	

	n := rand.Int() % (len(x)-1)
	session.Set("N",n)
	this.Data["title"] = "Result"
	this.Data["isi1"] = x[n].Judul_Film
	this.Data["isi2"] = x[n].Durasi
	this.Data["isi3"] = x[n].Tanggal_Tayang
	this.Data["isi4"] = x[n].Rating
	this.Data["isi5"] = x[n].Sinopsis

	//this.Data["json"]=x
	//this.ServeJSON()
	this.TplName = "search-result.tpl"
	}
}



