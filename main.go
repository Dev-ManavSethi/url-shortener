package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/Dev-ManavSethi/url-shortener/controllers"
	"github.com/Dev-ManavSethi/url-shortener/models"
	"github.com/Dev-ManavSethi/url-shortener/utils"
)

func init() {

	models.DummyError = nil
	models.Templates, models.DummyError = template.ParseGlob("templates/*")
	utils.HandleErr(models.DummyError, "Error parsing glob from templates", "Parsed glob from templates")

	err := godotenv.Load(".env")
	utils.HandleErr(err, "Error loading env vaiables", "Loaded env variables")

	models.DummyError = nil
	models.RedisClient, models.DummyError = utils.ConnectToRedis()
	utils.HandleErr(models.DummyError, "Error connecting to redis DB", "Connected to Redis db")

	models.Router = mux.NewRouter()

}

func main() {

models.Router.HandleFunc("/", controllers.Home)
models.Router.HandleFunc("/success", controllers.Success).Methods("GET")
models.Router.HandleFunc("/search", controllers.Search).Methods("GET")

	models.Router.HandleFunc("/{slug:[a-zA-Z0-9]+}", controllers.Redirect).Methods("GET")

http.Handle("/", models.Router)


	err := http.ListenAndServe(":"+os.Getenv("PORT"), models.Router)
	utils.HandleErr(err, "Error starting HTTP server", "")

}
