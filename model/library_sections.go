package model

type LibrarySections struct {
	MediaContainer struct {
		Directory []Directory `json:"directory"`
	} `json:"mediacontainer"`
}
