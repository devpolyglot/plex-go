package plex

import (
	"fmt"
	"net/http"

	"github.com/Lorac/plex-go/model"
)

type LibraryService service

type LibrarySections struct {
	MediaContainer struct {
		Directory []model.Directory `json:"directory" xml:"Directory"`
	} `json:"mediacontainer" xml:"MediaContainer"`
}

type MediaMetadata struct {
	MediaContainer struct {
		Size            int    `json:"size" xml:"size,attr"`
		Identifier      string `json:"identifier" xml:"identifier,attr"`
		MediaTagPrefix  string `json:"mediaTagPrefix" xml:"mediaTagPrefix,attr"`
		MediaTagVersion int    `json:"mediaTagVersion" xml:"mediaTagVersion,attr"`
		Metadata        []struct {
			AllowSync             string  `json:"allowSync" xml:"allowSync,attr"`
			LibrarySectionID      int     `json:"librarySectionID" xml:"librarySectionID,attr"`
			LibrarySectionTitle   string  `json:"librarySectionTitle" xml:"librarySectionTitle,attr"`
			LibrarySectionUUID    string  `json:"librarySectionUUID" xml:"librarySectionUUID,attr"`
			RatingKey             string  `json:"ratingKey" xml:"ratingKey,attr"`
			Key                   string  `json:"key" xml:"key,attr"`
			ParentRatingKey       string  `json:"parentRatingKey" xml:"parentRatingKey,attr"`
			GrandparentRatingKey  string  `json:"grandparentRatingKey" xml:"grandparentRatingKey,attr"`
			GUID                  string  `json:"guid" xml:"guid,attr"`
			LibrarySectionKey     string  `json:"librarySectionKey"`
			Type                  string  `json:"type" xml:"type,attr"`
			Title                 string  `json:"title" xml:"title,attr"`
			GrandparentKey        string  `json:"grandparentKey" xml:"grandparentKey,attr"`
			ParentKey             string  `json:"parentKey" xml:"parentKey,attr"`
			GrandparentTitle      string  `json:"grandparentTitle" xml:"grandparentTitle,attr"`
			ParentTitle           string  `json:"parentTitle" xml:"parentTitle,attr"`
			ContentRating         string  `json:"contentRating" xml:"contentRating,attr"`
			Summary               string  `json:"summary" xml:"summary,attr"`
			Index                 int     `json:"index"`
			ParentIndex           int     `json:"parentIndex"`
			Rating                float32 `json:"rating"`
			Year                  int     `json:"year" xml:"year,attr"`
			Thumb                 string  `json:"thumb" xml:"thumb,attr"`
			Art                   string  `json:"art" xml:"art,attr"`
			ParentThumb           string  `json:"parentThumb"`
			GrandParentThumb      string  `json:"grandparentThumb"`
			GrandParentArt        string  `json:"grandparentArt"`
			Duration              int     `json:"duration" xml:"duration,attr"`
			OriginallyAvailableAt string  `json:"originallyAvailableAt" xml:"originallyAvailableAt,attr"`
			AddedAt               int     `json:"addedAt" xml:"addedAt,attr"`
			UpdatedAt             int     `json:"updatedAt" xml:"updatedAt,attr"`
			Media                 []struct {
				VideoResolution string  `json:"videoResolution" xml:"videoResolution,attr"`
				ID              int     `json:"id" xml:"id,attr"`
				Duration        int     `json:"duration" xml:"duration,attr"`
				Bitrate         int     `json:"bitrate" xml:"bitrate,attr"`
				Width           int     `json:"width" xml:"width,attr"`
				Height          int     `json:"height" xml:"height,attr"`
				AspectRatio     float32 `json:"aspectRatio" xml:"aspectRatio,attr"`
				AudioChannels   int     `json:"audioChannels" xml:"audioChannels,attr"`
				AudioCodec      string  `json:"audioCodec" xml:"audioCodec,attr"`
				VideoCodec      string  `json:"videoCodec" xml:"videoCodec,attr"`
				Container       string  `json:"container" xml:"container,attr"`
				VideoFrameRate  string  `json:"videoFrameRate" xml:"videoFrameRate,attr"`
				VideoProfile    string  `json:"videoProfile" xml:"videoProfile,attr"`
				Part            []struct {
					ID           int    `json:"id" xml:"id,attr"`
					Key          string `json:"key" xml:"key,attr"`
					Duration     int    `json:"duration" xml:"duration,attr"`
					File         string `json:"file" xml:"file,attr"`
					Size         int    `json:"size" xml:"size,attr"`
					Container    string `json:"container" xml:"container,attr"`
					VideoProfile string `json:"videoProfile" xml:"videoProfile,attr"`
					Stream       []struct {
						ID                 int     `json:"id" xml:"id,attr"`
						StreamType         int     `json:"streamType" xml:"streamType,attr"`
						Default            bool    `json:"default" xml:"default,attr"`
						Codec              string  `json:"codec" xml:"codec,attr"`
						Index              int     `json:"index" xml:"index,attr"`
						Bitrate            int     `json:"bitrate" xml:"bitrate,attr"`
						BitDepth           int     `json:"bitDepth" xml:"bitDepth,attr"`
						ChromaSubsampling  string  `json:"chromaSubsampling" xml:"chromaSubsampling,attr"`
						ColorRange         string  `json:"colorRange" xml:"colorRange,attr"`
						ColorSpace         string  `json:"colorSpace" xml:"colorSpace,attr"`
						FrameRate          float32 `json:"frameRate" xml:"frameRate,attr"`
						HasScalingMatrix   bool    `json:"hasScalingMatrix" xml:"hasScalingMatrix,attr"`
						Height             int     `json:"height" xml:"height,attr"`
						Level              int     `json:"level" xml:"level,attr"`
						Profile            string  `json:"profile" xml:"profile,attr"`
						RefFrames          int     `json:"refFrames" xml:"refFrames,attr"`
						ScanType           string  `json:"scanType" xml:"scanType,attr"`
						Width              int     `json:"width" xml:"width,attr"`
						Selected           bool    `json:"selected" xml:"selected,attr"`
						Channels           int     `json:"channels" xml:"channels,attr"`
						Language           string  `json:"language" xml:"language,attr"`
						LanguageCode       string  `json:"languageCode" xml:"languageCode,attr"`
						AudioChannelLayout string  `json:"audioChannelLayout" xml:"audioChannelLayout,attr"`
						SamplingRate       int     `json:"samplingRate" xml:"samplingRate,attr"`
						Title              string  `json:"title" xml:"title,attr"`
					} `json:"Stream" xml:"Stream"`
				} `json:"Part" xml:"Part"`
			} `json:"Media" xml:"Media"`
			Director []struct {
				ID     int    `json:"id"`
				Filter string `json:"filter"`
				Tag    string `json:"tag"`
			} `json:"Director"`
			Writer []struct {
				ID     int    `json:"id"`
				Filter string `json:"filter"`
				Tag    string `json:"tag"`
			} `json:"Writer"`
			Label struct {
				ID  int    `json:"id" xml:"id,attr"`
				Tag string `json:"tag" xml:"tag,attr"`
			} `json:"label" xml:"Label"`
			Field []struct {
				Name   string `json:"name" xml:"name,attr"`
				Locked string `json:"locked" xml:"locked,attr"`
			} `json:"field" xml:"Field"`
		} `json:"Metadata" xml:"metadata"`
	} `json:"MediaContainer" xml:"MediaContainer"`
}

type LibraryLabels struct {
	ElementType     string `json:"_elementType"`
	AllowSync       string `json:"allowSync"`
	Art             string `json:"art"`
	Content         string `json:"content"`
	Identifier      string `json:"identifier"`
	MediaTagPrefix  string `json:"mediaTagPrefix"`
	MediaTagVersion string `json:"mediaTagVersion"`
	Thumb           string `json:"thumb"`
	Title1          string `json:"title1"`
	Title2          string `json:"title2"`
	ViewGroup       string `json:"viewGroup"`
	ViewMode        string `json:"viewMode"`
	Children        []struct {
		ElementType string `json:"_elementType"`
		FastKey     string `json:"fastKey"`
		Key         string `json:"key"`
		Title       string `json:"title"`
	} `json:"children"`
}

func (l *LibraryService) ListSections() (*LibrarySections, *http.Response, error) {
	req, err := l.client.NewRequest("GET", "library/sections", nil)
	if err != nil {
		return nil, nil, err
	}

	librarySections := new(LibrarySections)
	resp, err := l.client.Do(req, librarySections)

	if err != nil {
		return nil, nil, err
	}

	return librarySections, resp, nil
}

func (l *LibraryService) ListLabels(sectionKey string) (*LibraryLabels, *http.Response, error) {
	u := fmt.Sprintf("library/sections/%s/labels", sectionKey)
	req, err := l.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	libraryLabels := new(LibraryLabels)
	resp, err := l.client.Do(req, libraryLabels)

	if err != nil {
		return nil, nil, err
	}

	return libraryLabels, resp, nil
}

func (l *LibraryService) ListOnDeck(sectionKey string) (*MediaMetadata, *http.Response, error) {
	u := fmt.Sprintf("library/sections/%s/onDeck", sectionKey)
	req, err := l.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	metadata := new(MediaMetadata)
	resp, err := l.client.Do(req, metadata)

	if err != nil {
		return nil, nil, err
	}

	return metadata, resp, nil
}
