package dataman

import (
	"encoding/json"
	"sync"
)

// Project Represents a Project I have worked on
type Project struct {
	ID               string   `json:"id"`
	Name             string   `json:"name"`
	ShortDescription string   `json:"shortDescription"`
	LongDescription  string   `json:"longDescription"`
	LargeIcon        string   `json:"largeIcon"`
	SmallIcon        string   `json:"smallIcon"`
	TechUsed         []string `json:"techUsed"`
	Repository       string   `json:"repoLink"`
}

// ProjectRepository A place to store and retrieve persisted Project info
type ProjectRepository interface {
	GetProjects() ([]Project, error)
	Save(projects []Project) error
}

// JSONProjectRepository An implementation of ProjectRepository using json
// & local file
type JSONProjectRepository struct {
	projectStore DataStore
	projects     []Project
	lock         *sync.RWMutex
}

// Loads this repository from file
func (repo *JSONProjectRepository) load() error {

	repo.lock.Lock()
	defer repo.lock.Unlock()

	var projects []Project
	var p Project

	lines, err := repo.projectStore.ReadLines()
	if err != nil {
		return err
	}

	for _, line := range lines {
		if err := json.Unmarshal(line, &p); err != nil {
			return err
		}
		projects = append(projects, p)
	}

	// Assuming that went well...
	repo.projects = projects

	return nil
}

// GetProjects get the projects in this repo.
// In this case they have been eagerly loaded until it needs to be more complex.
func (repo *JSONProjectRepository) GetProjects() ([]Project, error) {
	repo.lock.RLock()
	defer repo.lock.RUnlock()
	return repo.projects, nil
}

// Save save the specified projects, overwriting what was previously there.
// Currently not expected to be used by the API, more for data management,
// So no reason to make this thread safe or anything right?
func (repo *JSONProjectRepository) Save(projects []Project) error {
	repo.lock.Lock()
	defer repo.lock.Unlock()

	linesToWrite := make([][]byte, len(projects))

	for i, project := range projects {
		json, err := json.Marshal(project)
		if err != nil {
			return err
		}

		linesToWrite[i] = json
	}

	if err := repo.projectStore.Empty(); err != nil {
		return err
	}

	if err := repo.projectStore.WriteLines(linesToWrite); err != nil {
		return err
	}

	// Update the current state of this object to what was just written
	// TODO: Arguably should just read but this assumes the write succeeded
	repo.projects = projects

	return nil
}

// NewJSONProjectRepository New instance of a JSONProjectRepository using default file from config
func NewJSONProjectRepository() (*JSONProjectRepository, error) {
	return NewJSONProjectRepositoryUsingFile("TODOGETMEFROMCONFIG")
}

// NewJSONProjectRepositoryUsingFile New instance of JSONProjectRepository using the specified file
func NewJSONProjectRepositoryUsingFile(path string) (*JSONProjectRepository, error) {

	dataStore := OSFileDataStore{
		path: path,
	}

	repo := JSONProjectRepository{
		projectStore: &dataStore,
	}

	if err := repo.load(); err != nil {
		return nil, err
	}

	return &repo, nil
}
