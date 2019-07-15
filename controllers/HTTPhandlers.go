package controllers

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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




	if result.Val() == "OK"{
		f, err := os.OpenFile("keys.keys", os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
			f.Close()
			return
		}


		defer f.Close()

	_, err2 := f.Write([]byte(prefferedName + "\n"))
	if err2!=nil{
		panic(err2)
		f.Close()
		return
	}

var Pair models.Pair
	Pair.Key = prefferedName
	Pair.Value = link
	err3 := models.Templates.ExecuteTemplate(w, "success.html", Pair )
	if err3!=nil{

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

	result := models.RedisClient.Get(url)

	link := result.Val()

	if link ==""{

		w.WriteHeader(http.StatusNotFound)
	}

	//fmt.Fprintln(w, "Redirecting you to: " + link)
	//fmt.Fprintln(w, "")
	//fmt.Fprintln(w, "")
	//fmt.Fprintln(w, "Short link created by trim-link.herokuapp.com")
	//
	//time.Sleep(5*time.Second)

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

func AllLinks(w http.ResponseWriter, r *http.Request){

	f , err := os.Open("keys.keys")
	if err!=nil{
		panic(err)
		f.Close()
		return
	}

	reader := bufio.NewReader(f)
	contents , _ := ioutil.ReadAll(reader)
	keys := strings.Split(string(contents), "\n")


	var Pairs []models.Pair

	for _, key := range keys{


result := models.RedisClient.Get(key)

var Pair models.Pair
Pair.Key = key
Pair.Value = result.Val()

Pairs = append(Pairs, Pair)

}


err2 := models.Templates.ExecuteTemplate(w, "all.html", Pairs)
if err2!=nil{

}



}