package main

import (
	"fmt"
	"net/http"

	"hkjn.me/jax/jaxhelper"

	"github.com/gorilla/mux"
)

// HomeHandler writes the results for the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	j1 := jaxhelper.Jax{
		Happiness: 30,
		Name:      "Jaxieboy",
		Writer:    w,
	}
	j1.Say("Hello hello!")
	j1.Say("Jax wants some food!")
	j1.Feed(jaxhelper.Banana)
	j1.Feed(jaxhelper.Banana)
	j1.Feed(jaxhelper.Banana)

	j2 := jaxhelper.Jax{
		Happiness: 90,
		Name:      "CutieJax",
		Writer:    w,
	}
	j2.Say("Oh!")
	j2.Say("A new friend")
	j2.Feed(jaxhelper.Chicken)
	j2.Feed(jaxhelper.Chicken)

	j1.HowAreYou()
	j2.HowAreYou()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	//r.HandleFunc("/products", ProductsHandler)
	//r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
	err := http.ListenAndServe(":8080", nil)
	fmt.Printf("FATAL: %v\n", err)
}
