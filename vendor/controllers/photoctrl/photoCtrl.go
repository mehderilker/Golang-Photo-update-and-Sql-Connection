package photoctrl

import (
	"net/http"
	"models"
	"services/menuservice"
	"services/photoservice"
	"services/display"
	"github.com/gorilla/mux"
	"fmt"
	"strconv"
	"os"
)

func Index(w http.ResponseWriter , r *http.Request){

	var data models.PhotoPage
	data.Page.Title = "Photo"
	data.Page.Description ="Photo"
	data.Page.Category=menuservice.GetMenuList()
	data.PhotoList = photoservice.GetList()


	display.View(w,r,"Photos",data)

	}


func Update(w http.ResponseWriter , r *http.Request){
		vars := mux.Vars(r)
		IDstr := vars["ID"]
 		ID , err := strconv.Atoi(IDstr)
 		if err !=nil {
 			fmt.Println(err)
 		}
		var data models.PhotoUpdate
		data.Page.Title = "Photo"
		data.Page.Description ="Photo"
		data.Page.Category=menuservice.GetMenuList()
		data.PhotoInfo=photoservice.GetPhotoInfo(ID)

		display.View(w,r,"photoEdit",data)
}

func DeletePhoto(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	var idStr = r.FormValue("ID")
	var id,_ = strconv.Atoi(idStr)
	var photoInfo = photoservice.GetPhotoInfo(id)
	var response = photoservice.DeleteInfo(id)
	if response.Status{
		if photoInfo.Pname.String != ""{
			os.Remove("./public/pics/"+photoInfo.Pname.String)
		}
	}
	display.Json(w,response)
}
