package model

import "encoding/xml"

type ServerInfo struct {
	XMLName           xml.Name `json:"mediaContainer" xml:"MediaContainer"`
	FriendlyName      string   `json:"friendlyName" xml:"friendlyName,attr"`
	Identifier        string   `json:"identifier" xml:"identifier,attr"`
	MachineIdentifier string   `json:"machineIdentifier" xml:"machineIdentifier,attr"`
	Size              int      `json:"size" xml:"size,attr"`
	Server            []struct {
		AccessToken       string `json:"accessToken" xml:"accessToken,attr"`
		Name              string `json:"name" xml:"name,attr"`
		Address           string `json:"address" xml:"address,attr"`
		Port              string `json:"port" xml:"port,attr"`
		Version           string `json:"version" xml:"version,attr"`
		Scheme            string `json:"scheme" xml:"scheme,attr"`
		Host              string `json:"host" xml:"host,attr"`
		LocalAddresses    string `json:"localAddresses" xml:"localAddresses,attr"`
		MachineIdentifier string `json:"machineIdentifier" xml:"machineIdentifier,attr"`
		CreatedAt         string `json:"createdAt" xml:"createdAt,attr"`
		UpdatedAt         string `json:"updatedAt" xml:"updatedAt,attr"`
		Owned             string `json:"owned" xml:"owned,attr"`
		Synced            string `json:"synced" xml:"synced,attr"`
	} `json:"server" xml:"server"`
}
