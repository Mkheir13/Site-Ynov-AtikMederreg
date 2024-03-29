package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Starting server :\n\thttp://localhost/")
	viewData := LoadAPI("etudiant.json")

	indexTemplate := template.Must(template.ParseFiles("../src/index.html"))
	studentTemplate := template.Must(template.ParseFiles("../src/pageperso.html"))
	enseignantsTemplate := template.Must(template.ParseFiles("../src/pageprof.html"))
	profilTemplate := template.Must(template.ParseFiles("../src/profil.html"))
	noAPITemplate := template.Must(template.ParseFiles("../src/static/noAPI.html"))

	apiFuckedUp := false
	if len(viewData.Etudiants) == 0 || len(viewData.Intervenants) == 0 {
		apiFuckedUp = true
	}

	cssFolder := http.FileServer(http.Dir("../css"))
	http.Handle("/css/", http.StripPrefix("/css/", cssFolder))
	imgFolder := http.FileServer(http.Dir("../img"))
	http.Handle("/img/", http.StripPrefix("/img/", imgFolder))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if apiFuckedUp {
			http.Redirect(w, r, "/noAPI", 303)
			return
		}
		search := r.FormValue("searchBar")
		if search != "" {
			filteredViewData := ViewData{}
			for _, student := range viewData.Etudiants {
				if strings.Contains(strings.ToLower(student.Nom), strings.ToLower(search)) || strings.Contains(strings.ToLower(student.Prenom), strings.ToLower(search)) {
					filteredViewData.Etudiants = append(filteredViewData.Etudiants, student)
				}
			}
			indexTemplate.Execute(w, filteredViewData)
		} else {
			indexTemplate.Execute(w, viewData)
		}
	})

	http.HandleFunc("/students", func(w http.ResponseWriter, r *http.Request) {
		if apiFuckedUp {
			http.Redirect(w, r, "/noAPI", 303)
			return
		}
		studentTemplate.Execute(w, viewData)
	})

	http.HandleFunc("/noAPI", func(w http.ResponseWriter, r *http.Request) {
		noAPITemplate.Execute(w, nil)
	})

	http.HandleFunc("/enseignants", func(w http.ResponseWriter, r *http.Request) {
		if apiFuckedUp {
			http.Redirect(w, r, "/noAPI", 303)
			return
		}
		enseignantsTemplate.Execute(w, viewData)
	})

	http.HandleFunc("/profil/", func(w http.ResponseWriter, r *http.Request) {
		if apiFuckedUp {
			http.Redirect(w, r, "/noAPI", 303)
			return
		}
		profil := Profil{}
		id := strings.ReplaceAll(r.URL.Path, "localhost/profil/", "")
		id = strings.ReplaceAll(r.URL.Path, "/profil/", "") // idk why but I need this to work
		for _, student := range viewData.Etudiants {
			if student.Prenom == id {
				profil = student
				break
			}
		}

		profilTemplate.Execute(w, profil)
	})

	http.ListenAndServe(":3030", nil)
}
