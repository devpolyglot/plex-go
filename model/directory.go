package model

type Directory struct {
	Agent      string      `json:"agent" xml:"agent,attr"`
	AllowSync  bool        `json:"allowSync" xml:"allowSync,attr"`
	Art        string      `json:"art" xml:"art,attr"`
	Composite  string      `json:"composite" xml:"composite,attr"`
	CreatedAt  int         `json:"createdAt" xml:"createdAt,attr"`
	Filter     bool        `json:"filters" xml:"filters,attr"`
	Key        string      `json:"key" xml:"key,attr"`
	Language   string      `json:"language" xml:"language,attr"`
	Location   *[]Location `json:"location" xml:"location,attr"`
	Name       string      `json:"name" xml:"name,attr"`
	Refreshing bool        `json:"refreshing" xml:"refreshing,attr"`
	Scanner    string      `json:"scanner" xml:"scanner,attr"`
	Thumb      string      `json:"thumb" xml:"thumb,attr"`
	Title      string      `json:"title" xml:"title,attr"`
	Type       string      `json:"type" xml:"type,attr"`
	UpdatedAt  int         `json:"updatedAt" xml:"updatedAt,attr"`
	UUID       string      `json:"uuid" xml:"uuid,attr"`
}
