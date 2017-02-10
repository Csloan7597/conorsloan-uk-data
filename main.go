package main

import (
	"encoding/json"
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
	if link, err := siteDataRepo.GetTagLine(); err == nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, link)
	} else {
		writeError(w, err)
	}
}

func aboutMeHandler(w http.ResponseWriter, r *http.Request) {
	aboutMe, err := siteDataRepo.GetAboutMeData()
	if err != nil {
		writeError(w, err)
	} else {
		writeJSONOrError(w, aboutMe)
	}
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := projectRepo.GetProjects()
	if err != nil {
		writeError(w, err)
	} else {
		writeJSONOrError(w, projects)
	}
}

func projectsListHandler(w http.ResponseWriter, r *http.Request) {
	projects, err := siteDataRepo.GetProjectListings()
	if err != nil {
		writeError(w, err)
	} else {
		writeJSONOrError(w, projects)
	}
}

func cvHandler(w http.ResponseWriter, r *http.Request) {
	if link, err := siteDataRepo.GetCVLink(); err == nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, link)
	} else {
		writeError(w, err)
	}
}

func jobsHandler(w http.ResponseWriter, r *http.Request) {
	jobs, err := careerRepo.GetJobs()
	if err != nil {
		writeError(w, err)
	} else {
		writeJSONOrError(w, jobs)
	}
}

func techUsedHandler(w http.ResponseWriter, r *http.Request) {
	techsUsed, err := careerRepo.GetTechUsed()
	if err != nil {
		writeError(w, err)
	} else {
		writeJSONOrError(w, techsUsed)
	}
}

func glanceHandler(w http.ResponseWriter, r *http.Request) {
	glances, err := siteDataRepo.GetGlanceItems()
	if err != nil {
		writeError(w, err)
	} else {
		writeJSONOrError(w, glances)
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	writeError(w, fmt.Errorf("NOT YET IMPLEMENTED"))
}

func writeJSONOrError(w http.ResponseWriter, toMarshal interface{}) {
	json, err := json.Marshal(toMarshal)
	if err != nil {
		writeError(w, err)
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, string(json))
	}
}

func writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "{\"error\": \"%s\"}", err.Error())
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

	fmt.Printf("Loaded up, listening on port %v\n", config.Port)

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
