package main

import (
	"services/display"
	"net/http"
	"github.com/gorilla/mux"
	"controllers/homectrl"
	"services/database"
	_ "github.com/go-sql-driver/mysql"
	"controllers/photoctrl"
)

func main(){

	display.LoadTemplates()

	database.Connect("mysql","golang")
	defer database.DB.Close()
	var r = mux.NewRouter()

	r.HandleFunc("/",homectrl.Index )
	r.HandleFunc("/editpic",photoctrl.Index)
	r.HandleFunc("/editpic/duzenle/{ID}",photoctrl.Update)
	r.HandleFunc("/editpic/delete",photoctrl.DeletePhoto)
	r.HandleFunc("/anasayfa",homectrl.Index ).Methods("GET")
	r.HandleFunc("/anasayfa", homectrl.IndexPost).Methods("POST")

	http.Handle("/",r)

	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("static"))))
	http.Handle("/public/",http.StripPrefix("/public/",http.FileServer(http.Dir("public"))))


	http.ListenAndServe(":8080",nil)
}
