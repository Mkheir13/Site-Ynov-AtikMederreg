package main

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
		Nom    string `json:"Nom"`
		Prenom string `json:"Prenom"`
		Email  string `json:"Email"`
		Photo  string `json:"Photo"`
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
