package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dogwalkingserver/models"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//RegistrationController :
type RegistrationController struct {
	session *mgo.Session
}

//NewRegistrationController :
func NewRegistrationController(s *mgo.Session) *RegistrationController {
	return &RegistrationController{s}
}

//AddUser : type of RegistrationController.
func (rc RegistrationController) AddUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	reg := models.Registration{}

	//get request object and add values
	//to registration model
	json.NewDecoder(r.Body).Decode(&reg)

	password := []byte(reg.Password)
	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))
	reg.Password = string(hashedPassword)

	// Comparing the password with the hash
	//err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	//fmt.Println(err) // nil means it is a match

	reg.ID = bson.NewObjectId()

	// Write the user to mongo
	rc.session.DB("PetWalk").C("Users").Insert(reg)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode("User added")
}
