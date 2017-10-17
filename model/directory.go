package model

type Directory struct {
	Location   []Location `json:"location"`
	Agent      string     `json:"agent"`
	AllowSync  bool       `json:"allowSync"`
	Art        string     `json:"art"`
	Composite  string     `json:"composite"`
	CreatedAt  int        `json:"createdAt"`
	Filter     bool       `json:"filters"`
	Key        string     `json:"key"`
	Language   string     `json:"language"`
	Refreshing bool       `json:"refreshing"`
	Scanner    string     `json:"scanner"`
	Thumb      string     `json:"thumb"`
	Title      string     `json:"title"`
	Type       string     `json:"type"`
	UpdatedAt  int        `json:"updatedAt"`
	UUID       string     `json:"uuid"`
}
