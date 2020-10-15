package handlers

import (
	"math/rand"
	"net/http"
	"reitechsub/pkg/authentication"
	"reitechsub/pkg/utilsdb"
	"strings"
	"time"
)

func generateRandomString() string {
	rand.Seed(time.Now().Unix())
	var output strings.Builder
	charset := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP"
	length := 100
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charset))
		randomChar := charset[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}

func (h *Handler) HandleSignup(w http.ResponseWriter, r *http.Request) {
	// Check if data has been entered
	if r.Method != http.MethodPost {
		h.Tmpl.Execute(w, nil)
		return
	}

	// Take in form values
	userIn := utilsdb.User{
		Username: r.FormValue("UsernameInput"),
		Password: r.FormValue("PasswordInput"),
	}

	// Encrypt the password and generate the token
	password, err := authentication.EncryptPassword(userIn.Password)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	userIn.Password = password // password is the hash actually
	userIn.Token = generateRandomString()

	// Add the new user
	err = utilsdb.CreateUser(userIn)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// Redirect
	http.Redirect(w, r, "http://0.0.0.0:80/login", http.StatusSeeOther)
}
