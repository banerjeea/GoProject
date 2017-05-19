package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dogwalkingserver/models"
	"github.com/dogwalkingserver/services"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//LoginController :
type LoginController struct {
	session *mgo.Session
}

//NewLoginController :
func NewLoginController(s *mgo.Session) *LoginController {
	return &LoginController{s}
}

//Login :
func (lc LoginController) Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	login := models.Login{}
	reg := models.Registration{}

	json.NewDecoder(r.Body).Decode(&login)

	uid := login.Username
	pwd := []byte(login.Password)

	//fetch pwd hash
	lc.session.DB("PetWalk").C("Users").Find(bson.M{"username": uid}).One(&reg)

	hashedPwd := []byte(reg.Password)

	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword(hashedPwd, pwd)

	if err != nil {
		json.NewEncoder(w).Encode(false)
		return
	}

	token := services.SetToken(reg.Username)
	json.NewEncoder(w).Encode(token)

}
