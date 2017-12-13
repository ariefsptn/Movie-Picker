package controllers

import (
	
	"Movie-Picker/models"
    
)

type JsonController struct {
	GeneralExtendedController
}


func (this *JsonController) Get() {

	
	var x []*models.Film
	
	session := this.StartSession()
	
	auth := Authentication{}

	a := session.Get("Minimum").(int)
	b := session.Get("Mintahun").(int)
	c := session.Get("Maxtahun").(int)
	d := session.Get("Genre1").(string)
	e := session.Get("Genre2").(string)
	f := session.Get("Genre3").(string)
	auth.Minimum = a
	auth.Mintahun = b
	auth.Maxtahun = c
	auth.Genre1 = d
	auth.Genre2 = e
	auth.Genre3 = f
	
	x = models.SearchMovie(auth.Minimum, auth.Mintahun, auth.Maxtahun, auth.Genre1, auth.Genre2, auth.Genre3)
	

	
	
	
	
	//Mengambil film ke n dari SearchController/SearchAgainController
	n := session.Get("N").(int)

	this.Data["title"] = "Result"
	this.Data["isi1"] = x[n].Judul_Film
	this.Data["isi2"] = x[n].Durasi
	this.Data["isi3"] = x[n].Tanggal_Tayang
	this.Data["isi4"] = x[n].Rating
	this.Data["isi5"] = x[n].Sinopsis
	this.Data["json"]=x[n]
	
	//Mencetak dalam bentuk Json
	this.ServeJSON()
	//this.TplName = "search-result.tpl"
	
}



