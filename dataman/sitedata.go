package dataman

import (
	"encoding/json"
	"sync"
)

// GlanceItem a feature of the home page of my site
type GlanceItem struct {
	ID           string   `json:"id,omitempty"`
	Title        string   `json:"title"`
	StartDate    string   `json:"startDate,omitempty"`
	EndDate      string   `json:"endDate,omitempty"`
	Content      []string `json:"content"`
	RelatedImage string   `json:"relatedImage,omitempty"`
}

// ProjectListing Link to a project
type ProjectListing struct {
	ProjectID   string `json:"id"`
	ProjectName string `json:"projectName"`
	Path        string `json:"path"`
}

// AboutMeData Data for About Me, including main content and a set of images
type AboutMeData struct {
	Content []string `json:"content"`
	Images  []string `json:"images"`
}

// SiteDataRepository Place to store site navigation and site-related content
type SiteDataRepository interface {
	GetGlanceItems() ([]GlanceItem, error)
	SaveGlanceItems([]GlanceItem) error
	GetProjectListings() ([]ProjectListing, error)
	SaveProjectListings([]ProjectListing) error
	GetCVLink() (string, error)
	SaveCVLink(string) error
	GetTagLine() (string, error)
	SaveTagLine(string) error
	GetAboutMeData() (AboutMeData, error)
	SaveAboutMeData(AboutMeData) error
}

// JSONSiteDataRepository implementation of SiteDataRepository using JSON files
type JSONSiteDataRepository struct {
	glanceItemsStore    DataStore
	projectListingStore DataStore
	cvLinkStore         DataStore
	tagLineStore        DataStore
	aboutMeDataStore    DataStore

	glanceItems     []GlanceItem
	projectListings []ProjectListing
	cvLink          string
	tagLine         string
	aboutMeData     AboutMeData

	lock sync.RWMutex
}

func (repo *JSONSiteDataRepository) load() error {
	repo.lock.Lock()
	defer repo.lock.Unlock()

	// Load glanceItems

	var glanceItems []GlanceItem

	glanceData, err := repo.glanceItemsStore.ReadLines()
	if err != nil {
		return err
	}

	for _, lineData := range glanceData {
		var glanceItem GlanceItem

		if jsonerr := json.Unmarshal(lineData, &glanceItem); jsonerr != nil {
			return jsonerr
		}

		glanceItems = append(glanceItems, glanceItem)
	}

	// Load projectListings

	var projects []ProjectListing

	projectData, err := repo.projectListingStore.ReadLines()
	if err != nil {
		return err
	}

	for _, projectLine := range projectData {
		var project ProjectListing

		if jsonerr := json.Unmarshal(projectLine, &project); jsonerr != nil {
			return jsonerr
		}
		projects = append(projects, project)
	}

	// Load cvLink

	linkData, err := repo.cvLinkStore.Read()
	if err != nil {
		return err
	}

	link := string(linkData)

	// Load tagLine

	tagLineData, err := repo.tagLineStore.Read()
	if err != nil {
		return err
	}

	tagLine := string(tagLineData)

	// Load aboutMeData

	var aboutMeData AboutMeData

	aboutMeRawData, err := repo.aboutMeDataStore.Read()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(aboutMeRawData, &aboutMeData); err != nil {
		return err
	}

	// Assuming all that went well...
	repo.glanceItems = glanceItems
	repo.projectListings = projects
	repo.cvLink = link
	repo.aboutMeData = aboutMeData
	repo.tagLine = tagLine

	return nil
}

// GetGlanceItems This implementation is eagerly loaded, so just returns
func (repo *JSONSiteDataRepository) GetGlanceItems() ([]GlanceItem, error) {
	repo.lock.RLock()
	defer repo.lock.RUnlock()
	return repo.glanceItems, nil
}

// SaveGlanceItems Saves to file
func (repo *JSONSiteDataRepository) SaveGlanceItems(glanceItems []GlanceItem) error {
	repo.lock.Lock()
	defer repo.lock.Unlock()

	linesToWrite := make([][]byte, len(glanceItems))

	for i, glanceItem := range glanceItems {
		json, err := json.Marshal(glanceItem)
		if err != nil {
			return err
		}

		linesToWrite[i] = json
	}

	if err := repo.glanceItemsStore.Empty(); err != nil {
		return err
	}

	if err := repo.glanceItemsStore.WriteLines(linesToWrite); err != nil {
		return err
	}

	// Update the current state of this object to what was just written
	// TODO: Arguably should just read but this assumes the write succeeded
	repo.glanceItems = glanceItems

	return nil
}

// GetProjectListings This implementation is eagerly loaded, so just returns
func (repo *JSONSiteDataRepository) GetProjectListings() ([]ProjectListing, error) {
	repo.lock.RLock()
	defer repo.lock.RUnlock()
	return repo.projectListings, nil
}

// SaveProjectListings Saves to file
func (repo *JSONSiteDataRepository) SaveProjectListings(projects []ProjectListing) error {
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

	if err := repo.projectListingStore.Empty(); err != nil {
		return err
	}

	if err := repo.projectListingStore.WriteLines(linesToWrite); err != nil {
		return err
	}

	// Update the current state of this object to what was just written
	// TODO: Arguably should just read but this assumes the write succeeded
	repo.projectListings = projects

	return nil
}

// GetCVLink This implementation is eagerly loaded, so just returns
func (repo *JSONSiteDataRepository) GetCVLink() (string, error) {
	repo.lock.RLock()
	defer repo.lock.RUnlock()
	return repo.cvLink, nil
}

// SaveCVLink Saves to file
func (repo *JSONSiteDataRepository) SaveCVLink(link string) error {
	repo.lock.Lock()
	defer repo.lock.Unlock()

	toWrite := []byte(link)

	if err := repo.cvLinkStore.Empty(); err != nil {
		return err
	}

	if err := repo.cvLinkStore.Write(toWrite); err != nil {
		return err
	}

	repo.cvLink = link

	return nil
}

// GetTagLine This implementation is eagerly loaded, so just returns
func (repo *JSONSiteDataRepository) GetTagLine() (string, error) {
	repo.lock.RLock()
	defer repo.lock.RUnlock()
	return repo.tagLine, nil
}

// SaveCVLink Saves to file
func (repo *JSONSiteDataRepository) SaveTagLine(line string) error {
	repo.lock.Lock()
	defer repo.lock.Unlock()

	toWrite := []byte(line)

	if err := repo.tagLineStore.Empty(); err != nil {
		return err
	}

	if err := repo.tagLineStore.Write(toWrite); err != nil {
		return err
	}

	repo.tagLine = line

	return nil
}

// GetAboutMeData This implementation is eagerly loaded, so just returns
func (repo *JSONSiteDataRepository) GetAboutMeData() (AboutMeData, error) {
	repo.lock.RLock()
	defer repo.lock.RUnlock()
	return repo.aboutMeData, nil
}

// SaveAboutMeData Saves to file
func (repo *JSONSiteDataRepository) SaveAboutMeData(data AboutMeData) error {
	repo.lock.Lock()
	defer repo.lock.Unlock()

	json, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := repo.aboutMeDataStore.Empty(); err != nil {
		return err
	}

	if err := repo.aboutMeDataStore.Write(json); err != nil {
		return err
	}

	repo.aboutMeData = data

	return nil
}

// NewJSONSiteDataRepository New instance of a NewJSONSiteDataRepository using default files from config
func NewJSONSiteDataRepository(dataDirPath string) (*JSONSiteDataRepository, error) {
	return NewJSONSiteDataRepositoryUsingFiles(dataDirPath+"/glance.data", dataDirPath+"/projectList.data",
		dataDirPath+"/cv.data", dataDirPath+"/aboutMe.data", dataDirPath+"/tagLine.data")
}

// NewJSONSiteDataRepositoryUsingFiles New instance of JSONSiteDataRepository using the specified files
func NewJSONSiteDataRepositoryUsingFiles(glancePath string, projectPath string,
	cvLinkPath string, aboutMePath string, tagLinePath string) (*JSONSiteDataRepository, error) {

	repo := JSONSiteDataRepository{
		aboutMeDataStore: &OSFileDataStore{
			path: aboutMePath,
		},
		cvLinkStore: &OSFileDataStore{
			path: cvLinkPath,
		},
		glanceItemsStore: &OSFileDataStore{
			path: glancePath,
		},
		projectListingStore: &OSFileDataStore{
			path: projectPath,
		},
		tagLineStore: &OSFileDataStore{
			path: tagLinePath,
		},
	}

	if err := repo.load(); err != nil {
		return nil, err
	}

	return &repo, nil
}
