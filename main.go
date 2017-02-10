package main

import (
	"fmt"
	"net/http"

	"github.com/Csloan7597/conorsloan-uk-data/dataman"
)

var (
	config       Config
	siteDataRepo dataman.SiteDataRepository
	careerRepo   dataman.CareerRepository
	projectRepo  dataman.ProjectRepository
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func taglineHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{\"hello\": \"world\"}")
}

func aboutMeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{}")
}

func aboutMeImagesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{}")
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{}")
}

func projectsListHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{}")
}

func projectsGlanceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{}")
}

func cvHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{}")
}

func jobsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{}")
}

func jobsGlanceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{}")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{}")
}

func main() {
	fmt.Println("conorsloan.uk data-service. Starting...")
	config = NewConfig()
	fmt.Printf("Config: %v\n", config)

	siteDataRepo, err := dataman.NewJSONSiteDataRepository(config.DataPath)
	if err != nil {
		panic(err)
	}

	careerRepo, err := dataman.NewJSONCareerRepository(config.DataPath)
	if err != nil {
		panic(err)
	}

	projectRepo, err := dataman.NewJSONProjectRepository(config.DataPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data: %v %v %v\n", siteDataRepo, careerRepo, projectRepo)

	http.HandleFunc("/api/tagline", taglineHandler)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// tagline /GET
// Aboutme /GET
// Aboutme/images /GET
// projectList /GET
// projects /GET
// projects/glance
// CV /GET
// jobs  /GET
// jobs/glance
// Contact /POST
