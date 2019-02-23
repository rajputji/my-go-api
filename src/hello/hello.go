package main
 
import (
    "fmt"
    "log"
    "net/http"
)
 
type user struct {
    Id   int
    Name string
    DOB string
    friends []string
}

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
    fmt.Fprintln(w, "Add New User")
}
 
func ViewUser(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "view User")
}
func ViewAll(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "view All User")
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