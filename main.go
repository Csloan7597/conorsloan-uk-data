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
	link, err := siteDataRepo.GetCVLink()

	if err == nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, link)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err.Error())
	}
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

func techUsedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{}")
}

func glanceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{}")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "{}")
}

func main() {
	fmt.Println("conorsloan.uk data-service. Starting...")
	config = NewConfig()
	fmt.Printf("Config: %v\n", config)

	var err error

	siteDataRepo, err = dataman.NewJSONSiteDataRepository(config.DataPath)
	if err != nil {
		panic(err)
	}

	careerRepo, err = dataman.NewJSONCareerRepository(config.DataPath)
	if err != nil {
		panic(err)
	}

	projectRepo, err = dataman.NewJSONProjectRepository(config.DataPath)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data: %v %v %v\n", siteDataRepo, careerRepo, projectRepo)

	http.HandleFunc("/api/tagline", taglineHandler)
	http.HandleFunc("/api/aboutme", aboutMeHandler)
	http.HandleFunc("/api/glance", glanceHandler)

	http.HandleFunc("/api/project/list", projectsListHandler)
	http.HandleFunc("/api/project", projectsHandler)

	http.HandleFunc("/api/career/cv", cvHandler)
	http.HandleFunc("/api/career/jobs", jobsHandler)
	http.HandleFunc("/api/career/techused", techUsedHandler)

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// TODO: All the posts
// TODO: /contact
