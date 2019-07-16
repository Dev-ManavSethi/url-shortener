package controllers

import (
	"fmt"
	"github.com/Dev-ManavSethi/url-shortener/models"
	"github.com/Dev-ManavSethi/url-shortener/utils"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:

		err := models.Templates.ExecuteTemplate(w, "home.html", nil)
		utils.HandleErr(err, "Error executing template home.html method GET", "")
		if err != nil {

		}

		break

	case http.MethodPost:

		err := r.ParseForm()
		utils.HandleErr(err, "Error parsing POST form", "")
		if err != nil {

		}

		link := r.FormValue("link")
		prefferedName := r.FormValue("name")

		link = "http://"+link

		if prefferedName == "" {
			//create custom name
		}
		//expiry := r.FormValue("expiry")




	utils.SetMapValue(prefferedName, link)
	err2:= utils.TakeMapBackup()
	if err2!=nil{

		fmt.Fprintln(w, "Error creating short URL ", err2.Error())

	} else{

		//success
		var Pair models.Pair
		Pair.Key = prefferedName
		Pair.Value= link
		err := models.Templates.ExecuteTemplate(w, "success.html", Pair)
		if err!=nil{

		}
	}

		break
	}

}

func Redirect(w http.ResponseWriter, r *http.Request){
	url := r.URL.Path[1:]

	if url==""{
		w.WriteHeader(http.StatusBadRequest)
	}

	link := models.Map[url]

	if link==""{
w.WriteHeader(http.StatusNotFound)

	} else {


		http.Redirect(w, r, link, http.StatusSeeOther)
	}
}

func Success(w http.ResponseWriter, r *http.Request){
	fmt.Println("s")
}

func Search(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query().Get("q")

	if _, ok := models.Map[query]; ok{
		fmt.Fprint(w, "e")
	}  else {
		fmt.Fprint(w, "dne")
	}

}

func AllLinks(w http.ResponseWriter, r *http.Request){

	var Pairs []models.Pair
	for name, link := range models.Map{

		var Pair models.Pair
		Pair.Key = name
		Pair.Value = link

		Pairs=append(Pairs, Pair)
	}

err2 := models.Templates.ExecuteTemplate(w, "all.html", Pairs)
if err2!=nil{

}



}