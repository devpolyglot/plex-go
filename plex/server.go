package plex

import (
	"net/http"
)

// ServerService ...
type ServerService service

type Servers struct {
	MediaContainer struct {
		Size   int       `json:"size"`
		Server *[]Server `json:"Server"`
	} `json:"MediaContainer"`
}

type Server struct {
	Name              string `json:"name"`
	Host              string `json:"host"`
	Address           string `json:"address"`
	Port              int    `json:"port"`
	MachineIdentifier string `json:"machineIdentifier"`
	Version           string `json:"version"`
}

func (s *ServerService) List() (*Servers, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "servers", nil)
	if err != nil {
		return nil, nil, err
	}

	servers := new(Servers)
	resp, err := s.client.Do(req, servers)

	if err != nil {
		return nil, nil, err
	}

	return servers, resp, nil
}
