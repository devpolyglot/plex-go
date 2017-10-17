package model

type SearchResultsEpisode struct {
	Children []struct {
		Children []struct {
			Children []struct {
				ElementType           string `json:"_elementType"`
				Container             string `json:"container"`
				Duration              int    `json:"duration"`
				File                  string `json:"file"`
				Has64bitOffsets       bool   `json:"has64bitOffsets"`
				HasThumbnail          string `json:"hasThumbnail"`
				ID                    string `json:"id"`
				Key                   string `json:"key"`
				OptimizedForStreaming bool   `json:"optimizedForStreaming"`
				Size                  int    `json:"size"`
				VideoProfile          string `json:"videoProfile"`
			} `json:"_children"`
			ElementType           string `json:"_elementType"`
			AspectRatio           string `json:"aspectRatio"`
			AudioChannels         int    `json:"audioChannels"`
			AudioCodec            string `json:"audioCodec"`
			Bitrate               int    `json:"bitrate"`
			Container             string `json:"container"`
			Duration              int    `json:"duration"`
			Has64bitOffsets       bool   `json:"has64bitOffsets"`
			Height                int    `json:"height"`
			ID                    int    `json:"id"`
			OptimizedForStreaming int    `json:"optimizedForStreaming"`
			VideoCodec            string `json:"videoCodec"`
			VideoFrameRate        string `json:"videoFrameRate"`
			VideoProfile          string `json:"videoProfile"`
			VideoResolution       string `json:"videoResolution"`
			Width                 int    `json:"width"`
		} `json:"_children"`
		ElementType           string `json:"_elementType"`
		AddedAt               int    `json:"addedAt"`
		ChapterSource         string `json:"chapterSource"`
		Duration              int    `json:"duration"`
		Index                 int    `json:"index"`
		Key                   string `json:"key"`
		LastViewedAt          int    `json:"lastViewedAt"`
		OriginallyAvailableAt string `json:"originallyAvailableAt"`
		ParentKey             string `json:"parentKey"`
		ParentRatingKey       int    `json:"parentRatingKey"`
		Rating                string `json:"rating"`
		RatingKey             int    `json:"ratingKey"`
		Summary               string `json:"summary"`
		Thumb                 string `json:"thumb"`
		Title                 string `json:"title"`
		Type                  string `json:"type"`
		UpdatedAt             int    `json:"updatedAt"`
		ViewCount             int    `json:"viewCount"`
		Year                  int    `json:"year"`
	} `json:"_children"`
	ElementType              string `json:"_elementType"`
	AllowSync                string `json:"allowSync"`
	Art                      string `json:"art"`
	Banner                   string `json:"banner"`
	GrandparentContentRating string `json:"grandparentContentRating"`
	GrandparentRatingKey     string `json:"grandparentRatingKey"`
	GrandparentStudio        string `json:"grandparentStudio"`
	GrandparentTheme         string `json:"grandparentTheme"`
	GrandparentThumb         string `json:"grandparentThumb"`
	GrandparentTitle         string `json:"grandparentTitle"`
	Identifier               string `json:"identifier"`
	Key                      string `json:"key"`
	LibrarySectionID         string `json:"librarySectionID"`
	LibrarySectionTitle      string `json:"librarySectionTitle"`
	LibrarySectionUUID       string `json:"librarySectionUUID"`
	MediaTagPrefix           string `json:"mediaTagPrefix"`
	MediaTagVersion          string `json:"mediaTagVersion"`
	Nocache                  string `json:"nocache"`
	ParentIndex              string `json:"parentIndex"`
	ParentTitle              string `json:"parentTitle"`
	ParentYear               string `json:"parentYear"`
	SortAsc                  string `json:"sortAsc"`
	Theme                    string `json:"theme"`
	Thumb                    string `json:"thumb"`
	Title1                   string `json:"title1"`
	Title2                   string `json:"title2"`
	ViewGroup                string `json:"viewGroup"`
	ViewMode                 string `json:"viewMode"`
}
