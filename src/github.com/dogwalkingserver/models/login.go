package models

type (
	//Login : new user detail goes here
	Login struct {
		Username string `json:"username" bson:"username"`
		Password string `json:"password" bson:"password"`
	}
)
