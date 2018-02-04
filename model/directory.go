package model

type Directory struct {
	Agent      string     `json:"agen,omitempty"`
	AllowSync  bool       `json:"allowSync,omitempty"`
	Art        string     `json:"art,omitempty"`
	Composite  string     `json:"composite,omitempty"`
	CreatedAt  int        `json:"createdAt,omitempty"`
	Filter     bool       `json:"filters,omitempty"`
	Key        string     `json:"key,omitempty"`
	Language   string     `json:"language,omitempty"`
	Location   []Location `json:"location,omitempty"`
	Refreshing bool       `json:"refreshing,omitempty"`
	Scanner    string     `json:"scanner,omitempty"`
	Thumb      string     `json:"thumb,omitempty"`
	Title      string     `json:"title,omitempty"`
	Type       string     `json:"type,omitempty"`
	UpdatedAt  int        `json:"updatedAt,omitempty"`
	UUID       string     `json:"uuid,omitempty"`
}
