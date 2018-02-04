package plex

import (
	"fmt"
	"net/http"

	"github.com/Lorac/plex-go/model"
)

type LibraryService service

type LibrarySections struct {
	MediaContainer struct {
		Directory []model.Directory `json:"directory" xml:"Directory"`
	} `json:"mediacontainer" xml:"MediaContainer"`
}

type LibraryLabels struct {
	ElementType     string `json:"_elementType"`
	AllowSync       string `json:"allowSync"`
	Art             string `json:"art"`
	Content         string `json:"content"`
	Identifier      string `json:"identifier"`
	MediaTagPrefix  string `json:"mediaTagPrefix"`
	MediaTagVersion string `json:"mediaTagVersion"`
	Thumb           string `json:"thumb"`
	Title1          string `json:"title1"`
	Title2          string `json:"title2"`
	ViewGroup       string `json:"viewGroup"`
	ViewMode        string `json:"viewMode"`
	Children        []struct {
		ElementType string `json:"_elementType"`
		FastKey     string `json:"fastKey"`
		Key         string `json:"key"`
		Title       string `json:"title"`
	} `json:"children"`
}

func (l *LibraryService) ListSections() (*LibrarySections, *http.Response, error) {
	req, err := l.client.NewRequest("GET", "library/sections", nil)
	if err != nil {
		return nil, nil, err
	}

	librarySections := &LibrarySections{}
	resp, err := l.client.Do(req, librarySections)

	if err != nil {
		return nil, nil, err
	}

	return librarySections, resp, nil
}

func (l *LibraryService) ListLabels(sectionKey string) (*LibraryLabels, *http.Response, error) {
	u := fmt.Sprintf("library/sections/%s/labels", sectionKey)
	req, err := l.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	libraryLabels := &LibraryLabels{}
	resp, err := l.client.Do(req, libraryLabels)

	if err != nil {
		return nil, nil, err
	}

	return libraryLabels, resp, nil
}
