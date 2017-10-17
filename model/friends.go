package model

type Friends struct {
	ID                        int    `json:"id" xml:"id,attr"`
	Title                     string `json:"title" xml:"title,attr"`
	Thumb                     string `json:"thumb" xml:"thumb,attr"`
	Protected                 string `json:"protected" xml:"protected,attr"`
	Home                      string `json:"home" xml:"home,attr"`
	AllowSync                 string `json:"allowSync" xml:"allowSync,attr"`
	AllowCameraUpload         string `json:"allowCameraUpload" xml:"allowCameraUpload,attr"`
	AllowChannels             string `json:"allowChannels" xml:"allowChannels,attr"`
	FilterAll                 string `json:"filterAll" xml:"filterAll,attr"`
	FilterMovies              string `json:"filterMovies" xml:"filterMovies,attr"`
	FilterMusic               string `json:"filterMusic" xml:"filterMusic,attr"`
	FilterPhotos              string `json:"filterPhotos" xml:"filterPhotos,attr"`
	FilterTelevision          string `json:"filterTelevision" xml:"filterTelevision,attr"`
	Restricted                string `json:"restricted" xml:"restricted,attr"`
	Username                  string `json:"username" xml:"username,attr"`
	Email                     string `json:"email" xml:"email,attr"`
	RecommendationsPlaylistID string `json:"recommendationsPlaylistId" xml:"recommendationsPlaylistId,attr"`
	Server                    struct {
		ID                string `json:"id" xml:"id,attr"`
		ServerID          string `json:"serverId" xml:"serverId,attr"`
		MachineIdentifier string `json:"machineIdentifier" xml:"machineIdentifier,attr"`
		Name              string `json:"name" xml:"name,attr"`
		LastSeenAt        string `json:"lastSeenAt" xml:"lastSeenAt,attr"`
		NumLibraries      string `json:"numLibraries" xml:"numLibraries,attr"`
		AllLibraries      string `json:"allLibraries" xml:"allLibraries,attr"`
		Owned             string `json:"owned" xml:"owned,attr"`
		Pending           string `json:"pending" xml:"pending,attr"`
	} `xml:"server"`
}
