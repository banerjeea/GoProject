package controllers

import (
	"encoding/json"
	"net/http"

	"fmt"

	"github.com/dogwalkingserver/models"
	"github.com/julienschmidt/httprouter"
	mgo "gopkg.in/mgo.v2"
)

type Key int

const MyKey Key = 0

//ProfileController :
type ProfileController struct {
	session *mgo.Session
}

//NewProfileController :
func NewProfileController(s *mgo.Session) *ProfileController {
	return &ProfileController{s}
}

//GetProfile :
func (pc ProfileController) GetProfile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Println("from profile c")
	blah := r.Context()
	claims, ok := r.Context().Value(MyKey).(models.Claims)
	fmt.Println(blah.Err())

	json.NewEncoder(w).Encode(claims)

	if !ok {
		fmt.Println("Context is nil in profile controller")
		//http.NotFound(w, r)
		return

	}

}
