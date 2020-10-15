package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"reitechsub/handlers"
	"time"

	"github.com/gorilla/mux"
)

func parseTemplate(tmplPath string) *template.Template {
	return template.Must(template.ParseFiles(tmplPath))
}

func main() {
	log.Println("Initialsing...")

	// Setting up the directory for the static assests
	var dir string
	flag.StringVar(&dir, "dir", ".static", "Directory for assets")
	flag.Parse()

	// Using Gorilla/Mux
	r := mux.NewRouter()

	// Setup the templates
	tmplLoginPage := parseTemplate("./templates/login.html")
	tmplSignupPage := parseTemplate("./templates/signup.html")
	tmplMasterDash := parseTemplate("./templates/master.html")

	// Creating Handlers
	LoginPageHandler := handlers.NewHandler(tmplLoginPage)
	SignupPageHandler := handlers.NewHandler(tmplSignupPage)
	MasterDashHandler := handlers.NewHandler(tmplMasterDash)

	// Set-Up the static server
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	// r.HandleFunc("/", handlers.)
	// r.HandleFunc("/home", func(w http.ResponseWriter, r http.Request))
	r.HandleFunc("/login", LoginPageHandler.StaticLogin)
	r.HandleFunc("/u", LoginPageHandler.HandleLogin)
	r.HandleFunc("/signup", SignupPageHandler.HandleSignup)
	r.HandleFunc("/u/{userhandle}/dashboard/master", MasterDashHandler.AuthUser(MasterDashHandler.HandleMasterDash))
	// r.HandleFunc("/u/{userhandler/dashboard/drone}", func(w http.ResponseWriter, r http.Request))

	// Custom Setting
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:80",
		// Nobody likes Slow Loris Attacks
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  90 * time.Second,
	}

	log.Println("Serving on 0.0.0.0 at port 80")

	// Serve
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
