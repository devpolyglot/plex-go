package model

type SearchResults struct {
	MediaContainer struct {
		Metadata []struct {
			Genre []struct {
				Tag string `json:"tag"`
			} `json:"Genre"`
			Role []struct {
				Tag string `json:"tag"`
			} `json:"Role"`
			AddedAt               int64   `json:"addedAt"`
			AllowSync             bool    `json:"allowSync"`
			Art                   string  `json:"art"`
			Banner                string  `json:"banner"`
			ChildCount            int64   `json:"childCount"`
			ContentRating         string  `json:"contentRating"`
			Duration              int64   `json:"duration"`
			Index                 int64   `json:"index"`
			Key                   string  `json:"key"`
			LastViewedAt          int64   `json:"lastViewedAt"`
			LeafCount             int64   `json:"leafCount"`
			LibrarySectionID      int64   `json:"librarySectionID"`
			LibrarySectionTitle   string  `json:"librarySectionTitle"`
			LibrarySectionUUID    string  `json:"librarySectionUUID"`
			OriginallyAvailableAt string  `json:"originallyAvailableAt"`
			Personal              bool    `json:"personal"`
			Rating                float64 `json:"rating"`
			RatingKey             string  `json:"ratingKey"`
			SourceTitle           string  `json:"sourceTitle"`
			Studio                string  `json:"studio"`
			Summary               string  `json:"summary"`
			Theme                 string  `json:"theme"`
			Thumb                 string  `json:"thumb"`
			Title                 string  `json:"title"`
			Type                  string  `json:"type"`
			UpdatedAt             int64   `json:"updatedAt"`
			ViewCount             int64   `json:"viewCount"`
			ViewedLeafCount       int64   `json:"viewedLeafCount"`
			Year                  int64   `json:"year"`
		} `json:"Metadata"`
		Provider []struct {
			Key   string `json:"key"`
			Title string `json:"title"`
			Type  string `json:"type"`
		} `json:"Provider"`
		Identifier      string `json:"identifier"`
		MediaTagPrefix  string `json:"mediaTagPrefix"`
		MediaTagVersion int64  `json:"mediaTagVersion"`
		Size            int64  `json:"size"`
	} `json:"MediaContainer"`
}
