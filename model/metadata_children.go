package model

type MetadataChildren struct {
	MediaContainer struct {
		Directory []struct {
			Key             string `json:"key"`
			LeafCount       int64  `json:"leafCount"`
			Thumb           string `json:"thumb"`
			Title           string `json:"title"`
			ViewedLeafCount int64  `json:"viewedLeafCount"`
		} `json:"Directory"`
		Metadata []struct {
			AddedAt         int64  `json:"addedAt"`
			Art             string `json:"art"`
			Index           int64  `json:"index"`
			Key             string `json:"key"`
			LastViewedAt    int64  `json:"lastViewedAt"`
			LeafCount       int64  `json:"leafCount"`
			ParentIndex     int64  `json:"parentIndex"`
			ParentKey       string `json:"parentKey"`
			ParentRatingKey string `json:"parentRatingKey"`
			ParentTheme     string `json:"parentTheme"`
			ParentThumb     string `json:"parentThumb"`
			ParentTitle     string `json:"parentTitle"`
			RatingKey       string `json:"ratingKey"`
			Summary         string `json:"summary"`
			Thumb           string `json:"thumb"`
			Title           string `json:"title"`
			Type            string `json:"type"`
			UpdatedAt       int64  `json:"updatedAt"`
			ViewCount       int64  `json:"viewCount"`
			ViewedLeafCount int64  `json:"viewedLeafCount"`
		} `json:"Metadata"`
		AllowSync           bool   `json:"allowSync"`
		Art                 string `json:"art"`
		Banner              string `json:"banner"`
		Identifier          string `json:"identifier"`
		Key                 string `json:"key"`
		LibrarySectionID    int64  `json:"librarySectionID"`
		LibrarySectionTitle string `json:"librarySectionTitle"`
		LibrarySectionUUID  string `json:"librarySectionUUID"`
		MediaTagPrefix      string `json:"mediaTagPrefix"`
		MediaTagVersion     int64  `json:"mediaTagVersion"`
		Nocache             bool   `json:"nocache"`
		ParentIndex         int64  `json:"parentIndex"`
		ParentTitle         string `json:"parentTitle"`
		ParentYear          int64  `json:"parentYear"`
		Size                int64  `json:"size"`
		Summary             string `json:"summary"`
		Theme               string `json:"theme"`
		Thumb               string `json:"thumb"`
		Title1              string `json:"title1"`
		Title2              string `json:"title2"`
		ViewGroup           string `json:"viewGroup"`
		ViewMode            int64  `json:"viewMode"`
	} `json:"MediaContainer"`
}
