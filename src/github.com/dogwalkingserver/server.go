package main

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/dogwalkingserver/controllers"
	"github.com/dogwalkingserver/services"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}

func main() {
	//instantiate new router
	r := httprouter.New()

	//create controller instances
	rc := controllers.NewRegistrationController(getSession())

	hc := controllers.NewHealthCheckController()

	lc := controllers.NewLoginController(getSession())
	pc := controllers.NewProfileController(getSession())

	//another way of creating an instance.
	//test := &controllers.HealthCheck{}

	//health HealthCheck
	r.GET("/health", hc.GetHealth)

	//Login
	r.POST("/login", lc.Login)

	//add user
	r.POST("/register", rc.AddUser)

	r.POST("/getprofile", services.Validate(pc.GetProfile))

	handler := cors.Default().Handler(r)

	//Fire up the server
	http.ListenAndServe("localhost:3000", handler)
}
