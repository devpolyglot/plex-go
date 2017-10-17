package model

type ServerSections struct {
	ID    int    `json:"id" xml:"id,attr"`
	Key   string `json:"key" xml:"key,attr"`
	Type  string `json:"type" xml:"type,attr"`
	Title string `json:"title" xml:"title,attr"`
}
