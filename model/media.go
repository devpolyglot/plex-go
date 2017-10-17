package model

type Media struct {
	Part            []Part  `json:"part"`
	AspectRatio     float64 `json:"aspectRatio"`
	AudioChannels   int     `json:"audioChannels"`
	AudioCodec      string  `json:"audioCodec"`
	Bitrate         int     `json:"bitrate"`
	Container       string  `json:"container"`
	Duration        int     `json:"duration"`
	Height          int     `json:"height"`
	ID              int     `json:"id"`
	VideoCodec      string  `json:"videoCodec"`
	VideoFrameRate  string  `json:"videoFrameRate"`
	VideoProfile    string  `json:"videoProfile"`
	VideoResolution string  `json:"videoResolution"`
	Width           int     `json:"width"`
}
