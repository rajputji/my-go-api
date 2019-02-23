package main
 
import (
    "fmt"
    "log"
    "encoding/json"
	"io/ioutil"
    "net/http"
)

type user struct {
    Id   int
    Name string
    DOB string
    friends []int
}

var users []user

func main() {	
	http.HandleFunc("/",Index)
    http.HandleFunc("/adduser",AddUser)
    http.HandleFunc("/viewuser",ViewUser)
    http.HandleFunc("/viewall",ViewAll) 
    http.HandleFunc("/edituser",EditUser)
	http.HandleFunc("/addfriend",AddFriend)
	http.HandleFunc("/getfriends",GetFriends)

    log.Fatal(http.ListenAndServe(":8080", nil)) 
}

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Appointy Tech Task - 3")
}
 
func AddUser(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var u user
	err = json.Unmarshal(b, &u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	users = append(users,u)
    fmt.Fprintln(w, "New User Added",u.Name)
}
 
func ViewUser(w http.ResponseWriter, r *http.Request) {
    b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var u string
	err = json.Unmarshal(b, &u)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var us user
	for _, v := range users {
    if v.Name == u {
    	us = v
    	}
	}

	output, err := json.Marshal(us)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
}


func ViewAll(w http.ResponseWriter, r *http.Request) {
	for _, v := range users {
    {

	output, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)
    }
	}
}
func EditUser(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Edit User")
}
func AddFriend(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Add Friend")
}
func GetFriends(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "view All Friends")
}