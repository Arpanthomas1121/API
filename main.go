package main

import (
	"fmt"
	"net/http" 	
	"log"
	"encoding/json"	
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

type Data []UserData 		
type postData []PostData  



func userData(w http.ResponseWriter, r *http.Request){
	n1:="Art"
	n2:="Test"
	data:=Data{
		UserData{ID:"123", Name:n1, Email:"Example@mail.in", Password:"abc123@"},UserData{ID:"456", Name:n2, Email:"Test@gmail.com", Password:"Test@"},
	}
	json.NewEncoder(w).Encode(data)

	keys, ok := r.URL.Query()["id"]
        if !ok || len(keys[0]) < 1 {
            log.Println("Url Param 'id' is missing")
            return
        }
    	key := keys[0]
		fmt.Println(key)
		switch k:=key; k{
		case "1":
			fmt.Println("Password: 1234@")
			w.Write([]byte(n1))
		case "2":
			fmt.Println("Password: Test@")
			w.Write([]byte(n2))

		}
		
}

func postsData(w http.ResponseWriter, r *http.Request){
	now:=time.Now()
	id1,id2:=001,002
	c1:="it is my First Post"
	c2:="the second Post"
	p_data:=postDaata{
		PData{ID:"123", Caption:c1, Image_url:"abcd.com",Stamp:now},PData{ID:"456", Caption:c2, Image_url:"efgh.com",Stamp:now},
	}
	json.NewEncoder(w).Encode(p_data)

	keys, ok := r.URL.Query()["id"]
        if !ok || len(keys[0]) < 1 {
            log.Println("Url Param 'id' is missing")
            return
        }
    	key := keys[0]
		fmt.Println(key)
		switch k:=key; k{
		case "1":
			fmt.Println(id1)
			w.Write([]byte(c1))
		case "2":
			fmt.Println(id2)
			w.Write([]byte(c2))

		}
}

func homepage(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Instagram API ~arpan</h1>"))
}

func HandlereRequest(){
	http.HandleFunc("/", homepage)
	http.HandleFunc("/users/", userData)
	http.HandleFunc("/posts/", postsData)
	log.Fatal(http.ListenAndServe(":3399", nil))
}
func main(){
	HandlereRequest()
}
