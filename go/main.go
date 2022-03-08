package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
	"strings"
)

type ViewData struct {
	Etudiants []struct {
		Nom      string `json:"Nom"`
		Prenom   string `json:"Prenom"`
		Email    string `json:"Email"`
    Photo    string `json:"Photo"`
		Github   string `json:"Github,omitempty"`
		Linkedin string `json:"Linkedin,omitempty"`
	} `json:"Etudiants"`
  Intervenants []struct {
		Nom      string `json:"Nom"`
		Prenom   string `json:"Prenom"`
		Email    string `json:"Email"`
    Photo    string `json:"Photo"`
	} `json:"Intervenants"`
}

type Profil struct {
	Nom      string `json:"Nom"`
	Prenom   string `json:"Prenom"`
	Email    string `json:"Email"`
	Photo    string `json:"Photo"`
	Github   string `json:"Github,omitempty"`
	Linkedin string `json:"Linkedin,omitempty"`
}

func loadAPI() ViewData {
	vd := ViewData{}

	url := "https://raw.githubusercontent.com/Nimajjj/groupie-tracker/main/API/etudiant.json"

	httpClient := http.Client{
		Timeout: time.Second * 2, // define timeout
	}

	//create request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "API AT test <3")

	//make api call
	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	//parse response
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	jsonErr := json.Unmarshal(body, &vd)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return vd
}


func main() {
	viewData := loadAPI()

	fmt.Println("\nStarting server -> localhost:80")

	indexTemplate := template.Must(template.ParseFiles("../src/index.html"))
  studentTemplate := template.Must(template.ParseFiles("../src/pageperso.html"))
  enseignantsTemplate := template.Must(template.ParseFiles("../src/pageprof.html"))
	profilTemplate := template.Must(template.ParseFiles("../src/profil.html"))

	cssFolder := http.FileServer(http.Dir("../css"))
	http.Handle("/css/", http.StripPrefix("/css/", cssFolder))
  imgFolder := http.FileServer(http.Dir("../img"))
	http.Handle("/img/", http.StripPrefix("/img/", imgFolder))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		search := r.FormValue("searchBar")
		if search != "" {
			filteredViewData := ViewData{}
			for _, student := range viewData.Etudiants {
				if (strings.Contains(strings.ToLower(student.Nom), strings.ToLower(search)) || strings.Contains(strings.ToLower(student.Prenom), strings.ToLower(search))) {
					filteredViewData.Etudiants = append(filteredViewData.Etudiants, student)
				}
			}
			indexTemplate.Execute(w, filteredViewData)
		} else {
			indexTemplate.Execute(w, viewData)
		}
	})

  http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		studentTemplate.Execute(w, viewData)
	})

  http.HandleFunc("/enseignants", func(w http.ResponseWriter, r *http.Request) {
		enseignantsTemplate.Execute(w, nil)
	})

	http.HandleFunc("/profil/", func(w http.ResponseWriter, r *http.Request) {
		profil := Profil{}
		id := strings.ReplaceAll(r.URL.Path, "localhost/profil/", "")
		id = strings.ReplaceAll(r.URL.Path, "/profil/", "")	// idk why but I need this to work
		for _, student := range viewData.Etudiants {
			if student.Prenom == id {
				profil = student
				break
			}
		}

		profilTemplate.Execute(w, profil)
	})

	http.ListenAndServe(":80", nil)
}
