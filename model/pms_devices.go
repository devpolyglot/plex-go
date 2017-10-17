package model

type PMSDevices struct {
	Name                 string `json:"name" xml:"name,attr"`
	Product              string `json:"product" xml:"product,attr"`
	ProductVersion       string `json:"productVersion" xml:"productVersion,attr"`
	Platform             string `json:"platform" xml:"platform,attr"`
	PlatformVersion      string `json:"platformVersion" xml:"platformVersion,attr"`
	Device               int    `json:"device" xml:"device,attr"`
	ClientIdentifier     string `json:"clientIdentifier" xml:"clientIdentifier,attr"`
	CreatedAt            string `json:"createdAt" xml:"createdAt,attr"`
	LastSeenAt           string `json:"lastSeenAt" xml:"lastSeenAt,attr"`
	Provides             string `json:"provides" xml:"provides,attr"`
	Owned                string `json:"owned" xml:"owned,attr"`
	AccessToken          string `json:"accessToken" xml:"accessToken,attr"`
	HTTPSRequired        string `json:"httpsRequired" xml:"httpsRequired,attr"`
	Synced               string `json:"synced" xml:"synced,attr"`
	PublicAddressMatches string `json:"publicAddressMatches" xml:"publicAddressMatches,attr"`
	Presence             string `json:"presence" xml:"presence,attr"`
	Connection           []struct {
		Protocol string `json:"protocol" xml:"protocol,attr"`
		Address  string `json:"address" xml:"address,attr"`
		Port     string `json:"port" xml:"port,attr"`
		URI      string `json:"uri" xml:"uri,attr"`
		Local    int    `json:"local" xml:"local,attr"`
	} `json:"connection" xml:"Connection"`
}
