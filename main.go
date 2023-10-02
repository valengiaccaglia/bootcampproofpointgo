package main

import (
	"html/template"
	"log"
	"math/rand" 
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func passwordGenerator(w http.ResponseWriter, r *http.Request) {
	contraseña := PassWordGenerator()
	log.Println(contraseña)

	w.Header().Set("Content-Type", "text/plain")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(contraseña))
}

func main() {
	http.HandleFunc("/password", passwordGenerator)
	http.HandleFunc("/", homePage)

	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

func PassWordGenerator() string {
	caracteres := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&()-_=+;:,.<>?/"
	password := make([]byte, 16)

	var i int
	for i = 0; i < 16; i++ {
		index := rand.Intn(len(caracteres))
		password[i] = caracteres[index]
	}

	return string(password)
}
