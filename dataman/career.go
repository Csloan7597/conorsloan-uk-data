package dataman

import (
	"encoding/json"
	"sync"
)

// Job Model representing a job in my work history
type Job struct {
	ID               string   `json:"id"`
	Company          string   `json:"company"`
	EmployerIcon     string   `json:"employerIcon"`
	StartDate        string   `json:"startDate"`
	EndDate          string   `json:"endDate"`
	JobTitle         string   `json:"jobTitle"`
	Location         string   `json:"location"`
	Responsibilities []string `json:"responsibilities"`
	Achievements     []string `json:"achievements"`
	TechUsed         []string `json:"techUsed"`
}

// TechUsed represents a technology i have used and a description of where
type TechUsed struct {
	Title       string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type,omitempty"`
}

// CareerRepository A place to store and retrieve persisted Career info
type CareerRepository interface {
	GetJobs() ([]Job, error)
	SaveJobs([]Job) error
	GetTechUsed() ([]TechUsed, error)
	SaveTechUsed([]TechUsed) error
}

// JSONCareerRepository Implementation of CareerRepository storing things in files
type JSONCareerRepository struct {
	jobsStore     DataStore
	techUsedStore DataStore
	jobs          []Job
	techUsed      []TechUsed
	lock          sync.RWMutex
}

// GetJobs This implementation is eagerly loaded, so just returns
func (repo *JSONCareerRepository) GetJobs() ([]Job, error) {
	repo.lock.RLock()
	defer repo.lock.RUnlock()
	return repo.jobs, nil
}

// SaveJobs save the specified jobs to file
func (repo *JSONCareerRepository) SaveJobs(jobs []Job) error {
	repo.lock.Lock()
	defer repo.lock.Unlock()

	linesToWrite := make([][]byte, len(jobs))

	for i, job := range jobs {
		json, err := json.Marshal(job)
		if err != nil {
			return err
		}

		linesToWrite[i] = json
	}

	if err := repo.jobsStore.Empty(); err != nil {
		return err
	}

	if err := repo.jobsStore.WriteLines(linesToWrite); err != nil {
		return err
	}

	// Update the current state of this object to what was just written
	// TODO: Arguably should just read but this assumes the write succeeded
	repo.jobs = jobs

	return nil
}

// GetTechUsed This implementation is eagerly loaded, so just returns
func (repo *JSONCareerRepository) GetTechUsed() ([]TechUsed, error) {
	repo.lock.RLock()
	defer repo.lock.RUnlock()
	return repo.techUsed, nil
}

// SaveTechUsed save the specified tech used to file
func (repo *JSONCareerRepository) SaveTechUsed(techUsed []TechUsed) error {
	repo.lock.Lock()
	defer repo.lock.Unlock()

	linesToWrite := make([][]byte, len(techUsed))

	for i, techUsed := range techUsed {
		json, err := json.Marshal(techUsed)
		if err != nil {
			return err
		}

		linesToWrite[i] = json
	}

	if err := repo.techUsedStore.Empty(); err != nil {
		return err
	}

	if err := repo.techUsedStore.WriteLines(linesToWrite); err != nil {
		return err
	}

	// Update the current state of this object to what was just written
	// TODO: Arguably should just read but this assumes the write succeeded
	repo.techUsed = techUsed

	return nil

}

// Loads this repository from file(s)
func (repo *JSONCareerRepository) load() error {

	repo.lock.Lock()
	defer repo.lock.Unlock()

	// Load Jobs
	var jobs []Job

	jobLines, err := repo.jobsStore.ReadLines()
	if err != nil {
		return err
	}

	for _, jobLine := range jobLines {
		var job Job
		if jsonerr := json.Unmarshal(jobLine, &job); jsonerr != nil {
			return jsonerr
		}
		jobs = append(jobs, job)
	}

	// Load Tech Used

	var techsUsed []TechUsed

	techUsedLines, err := repo.jobsStore.ReadLines()
	if err != nil {
		return err
	}

	for _, techUsedLine := range techUsedLines {
		var techUsed TechUsed
		if err := json.Unmarshal(techUsedLine, &techUsed); err != nil {
			return err
		}
		techsUsed = append(techsUsed, techUsed)
	}

	// Assuming all that went well
	repo.jobs = jobs
	repo.techUsed = techsUsed

	return nil
}

// NewJSONCareerRepository New instance of a JSONProjectRepository using default file from config
func NewJSONCareerRepository(dataDirPath string) (*JSONCareerRepository, error) {
	return NewJSONCareerRepositoryUsingFiles(dataDirPath+"/jobs.data", dataDirPath+"/techUsed.data")
}

// NewJSONCareerRepositoryUsingFiles New instance of JSONProjectRepository using the specified file
func NewJSONCareerRepositoryUsingFiles(jobsPath string, techUsedPath string) (*JSONCareerRepository, error) {

	repo := JSONCareerRepository{
		jobsStore: &OSFileDataStore{
			path: jobsPath,
		},
		techUsedStore: &OSFileDataStore{
			path: techUsedPath,
		},
	}

	if err := repo.load(); err != nil {
		return nil, err
	}

	return &repo, nil
}
