package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"dbconn"
	"structs"

	log "github.com/gookit/color"
	"github.com/gorilla/mux"
)

//Register To register user
func Register(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Red.Println("error taking body", err)
	}
	var reg structs.User
	err = json.Unmarshal(b, &reg)
	if err != nil {
		log.Red.Println("Error in unmarshal reg data", err)
	}
	db, err := dbconn.OpenConnection()
	if err != nil {
		log.Red.Println("error in database connection", err)
	}
	if !db.HasTable(&reg) {
		db.Debug().AutoMigrate(&reg)
	}
	db.Create(&reg)
}

//Login API
func Login(w http.ResponseWriter, r *http.Request) {
	login, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error taking body", err)
	}
	var loginVar structs.Login
	var user structs.User
	err = json.Unmarshal(login, &loginVar)
	if err != nil {
		log.Red.Println("Error in unmarshaling login data", err)
	}
	fmt.Println(loginVar)
	db, err := dbconn.OpenConnection()
	if err != nil {
		fmt.Println("error in database connection", err)
	}
	db.Where(&structs.User{Email: loginVar.Email, Password: loginVar.Password}).First(&user)
	fmt.Println("user", user.Email, loginVar.Email)
	if user.Email == loginVar.Email {
		fmt.Println("che")
	} else {
		fmt.Println("nathi user")
	}
}

//GetUsers Get all users API
func GetUsers(w http.ResponseWriter, r *http.Request) {
	var user structs.User
	db, err := dbconn.OpenConnection()
	if err != nil {
		fmt.Println("error in database connection", err)
	}
	db.Find(&user)
	log.Info.Println(user)
	userBytes, err := json.Marshal(user)
	if err != nil {
		log.Red.Println("error in marshaling data", err)
	}
	log.Info.Prompt(string(userBytes))
	log.Yellow.Println(string(userBytes))
	w.Write(userBytes)
}

//UpdateUser update API
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Red.Println("error taking body", err)
	}
	var user structs.User
	var up structs.User
	err = json.Unmarshal(b, &up)
	if err != nil {
		log.Error.Println(err)
	}

	db, err := dbconn.OpenConnection()
	db.First(&user)
	user.FirstName = up.FirstName
	user.LastName = up.LastName
	user.Email = up.Email
	user.Password = up.Password
	fmt.Println(user)
	db.Model(&user).Updates(&user)
}

//DeleteUser Api
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var user structs.User
	user.Email = eventID
	fmt.Println(user)
	db, err := dbconn.OpenConnection()
	if err != nil {
		log.Error.Println(err)
	}
	db.Unscoped().Delete(&user)
}
