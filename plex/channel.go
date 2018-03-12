package plex

import "net/http"

// ChannelService ...
type ChannelService service

type Channels struct {
	MediaContainer struct {
		Size      int        `json:"size"`
		Content   string     `json:"content"`
		Directory *[]Channel `json:"Directory"`
	} `json:"MediaContainer"`
}

type Channel struct {
	Art            string `json:"art"`
	Platforms      string `json:"platforms"`
	Protocols      string `json:"protocols"`
	Thumb          string `json:"thumb"`
	Type           string `json:"type"`
	Key            string `json:"key"`
	Title          string `json:"title"`
	Share          int    `json:"share"`
	Identifier     string `json:"identifier"`
	LastAccessedAt int    `json:"lastAccessedAt"`
	Content        bool   `json:"content"`
}

func (s *ChannelService) ListRecentlyViewed() (*Channels, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "channels/recentlyViewed", nil)
	if err != nil {
		return nil, nil, err
	}

	channels := new(Channels)
	resp, err := s.client.Do(req, channels)

	if err != nil {
		return nil, nil, err
	}

	return channels, resp, nil
}

func (s *ChannelService) ListAll() (*Channels, *http.Response, error) {
	req, err := s.client.NewRequest("GET", "channels/all", nil)
	if err != nil {
		return nil, nil, err
	}

	channels := new(Channels)
	resp, err := s.client.Do(req, channels)

	if err != nil {
		return nil, nil, err
	}

	return channels, resp, nil
}
