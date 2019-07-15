package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Dev-ManavSethi/url-shortener/models"
	"github.com/Dev-ManavSethi/url-shortener/utils"
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

	result:= 	models.RedisClient.Set(prefferedName, link, 10000*time.Hour)


		fmt.Fprint(w, result.Val())

	time.Sleep(3*time.Second)
http.Redirect(w,r, "/success", http.StatusSeeOther)


		break
	}

}

func Redirect(w http.ResponseWriter, r *http.Request){
	url := r.URL.Path[1:]

	if url==""{
		w.WriteHeader(http.StatusBadRequest)
	}

	result := models.RedisClient.Get(url)

	link := result.Val()

	if link ==""{

		w.WriteHeader(http.StatusNotFound)
	}

	http.Redirect(w, r, link, http.StatusSeeOther)
}

func Success(w http.ResponseWriter, r *http.Request){
	fmt.Println("s")
}

func Search(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query().Get("q")

	result := models.RedisClient.Get(query)

	if result.Val() != ""{
		fmt.Fprint(w, "e")
	} else {
		fmt.Fprint(w, "dne")
	}
}