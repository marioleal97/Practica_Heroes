package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Biography struct {
	FullName string `json:"fullName"`
}

type Powerstats struct {
	Intelligence int `json:"intelligence"`
	Strength     int `json:"strength"`
	Speed        int `json:"speed"`
	Durability   int `json:"durability"`
	Power        int `json:"power"`
	Combat       int `json:"combat"`
}

type Images struct {
	Xs string `json:"xs"`
	Sm string `json:"sm"`
	Md string `json:"md"`
	Lg string `json:"lg"`
}

type Superhero struct {
	Name       string     `json:"name"`
	Biography  Biography  `json:"biography"`
	Powerstats Powerstats `json:"powerstats"`
	Images     Images     `json:"images"`
}

var superheroes = []Superhero{
	{
		Name: "Wolverine",
		Biography: Biography{
			FullName: "John Logan",
		},
		Powerstats: Powerstats{
			Intelligence: 63,
			Strength:     32,
			Speed:        50,
			Durability:   100,
			Power:        89,
			Combat:       100,
		},
		Images: Images{
			Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/717-wolverine.jpg",
			Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/717-wolverine.jpg",
			Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/717-wolverine.jpg",
			Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/717-wolverine.jpg",
		},
	},
	{
		Name: "Spiderman",
		Biography: Biography{
			FullName: "Peter Parker",
		},
		Powerstats: Powerstats{
			Intelligence: 90,
			Strength:     55,
			Speed:        67,
			Durability:   75,
			Power:        74,
			Combat:       85,
		},
		Images: Images{
			Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/620-spider-man.jpg",
			Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/620-spider-man.jpg",
			Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/620-spider-man.jpg",
			Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/620-spider-man.jpg",
		},
	},
	{
		Name: "Iron Man",
		Biography: Biography{
			FullName: "Tony Stark",
		},
		Powerstats: Powerstats{
			Intelligence: 100,
			Strength:     85,
			Speed:        58,
			Durability:   85,
			Power:        100,
			Combat:       64,
		},
		Images: Images{
			Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/346-iron-man.jpg",
			Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/346-iron-man.jpg",
			Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/346-iron-man.jpg",
			Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/346-iron-man.jpg",
		},
	},
	{
		Name: "Batman",
		Biography: Biography{
			FullName: "Bruce Wayne",
		},
		Powerstats: Powerstats{
			Intelligence: 81,
			Strength:     40,
			Speed:        29,
			Durability:   55,
			Power:        63,
			Combat:       90,
		},
		Images: Images{
			Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/70-batman.jpg",
			Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/70-batman.jpg",
			Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/70-batman.jpg",
			Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/70-batman.jpg",
		},
	},
	{
		Name: "Superman",
		Biography: Biography{
			FullName: "Clark Kent",
		},
		Powerstats: Powerstats{
			Intelligence: 94,
			Strength:     100,
			Speed:        100,
			Durability:   100,
			Power:        100,
			Combat:       85,
		},
		Images: Images{
			Xs: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/xs/644-superman.jpg",
			Sm: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/sm/644-superman.jpg",
			Md: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/md/644-superman.jpg",
			Lg: "https://cdn.rawgit.com/akabab/superhero-api/0.2.0/api/images/lg/644-superman.jpg",
		},
	},
}

func getSuperhero(w http.ResponseWriter, r *http.Request) {
	heroName := r.URL.Query().Get("hero")
	if heroName == "" {
		http.Error(w, "Hero name is required", http.StatusBadRequest)
		return
	}

	for _, hero := range superheroes {
		if strings.ToLower(hero.Name) == strings.ToLower(heroName) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(hero)
			return
		}
	}

	http.Error(w, "Hero not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/api/superhero", getSuperhero)
	log.Println("Server running at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
