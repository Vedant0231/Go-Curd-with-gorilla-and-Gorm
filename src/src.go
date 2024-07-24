package src

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

var DB *gorm.DB
var err error

const DNS = "root:netcore@tcp(localhost:3306)/netcore_db?charset=utf8mb4&parseTime=True&loc=Local"

func Initalmigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	DB.AutoMigrate(&User{})
	log.Println("Database migration complete")
}

func Getusers(w http.ResponseWriter, r *http.Request) {

	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)

}

func Getuser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var users User
	DB.Find(&users, params["id"])
	json.NewEncoder(w).Encode(users)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")

	fmt.Println("ssdddddddddddddddddddddddddddddddddddddddddddddd")
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func Updateuser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)

}

func Deleteuser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var user User
	DB.Delete(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	json.NewEncoder(w).Encode("this user is deleted")
}
