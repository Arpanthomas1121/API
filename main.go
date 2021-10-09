package main

import (
	//"fmt"
	"net/http"
	"encoding/json"
  "log"
  "time"
)

type UserData struct{
	ID string `json:"ID"`
	Name string `json:"Name"`
	Email string `json:"Email"`
	Password string `json:"Password"`
}

type PostData struct{
	
	ID string `json:"ID"`
	Caption string `json:"caption"`
	Image_url string `json:"url"`
	Stamp time.Time `json:"time"`

}

type Data []UData 		//array for user data
type postData []PData  //array for Post Data

func userData(w http.ResponseWriter, r *http.Request){
	data:= Data{
		UData{ID:"123", Name:"Art", Email:"sdasads@sdfsdf.com", Password:"dsdfsgd"},
	            }
	json.NewEncoder(w).Encode(data)
}

func postsData(w http.ResponseWriter, r *http.Request){
	now:=time.Now()
	p_data:=postDaata{
		PData{ID:"123", Caption:"Hey, it on-air", Image_url:"abcdsdf",Stamp:now},
	}
	json.NewEncoder(w).Encode(p_data)
}


func homepage(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Instagram</h1>"))
}


func HandlereRequest(){
	http.HandleFunc("/", homepage)
	http.HandleFunc("/users", userData)
	http.HandleFunc("/posts", postsData)
	
	log.Fatal(http.ListenAndServe(":3999", nil))
}


func main(){
	HandlereRequest()
}



