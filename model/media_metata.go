package model

type MediaMetadata struct {
	MediaContainer struct {
		Size            string `json:"size" xml:"size,attr,omitempty"`
		Identifier      string `json:"identifier" xml:"identifier,attr,omitempty"`
		MediaTagPrefix  string `json:"mediaTagPrefix" xml:"mediaTagPrefix,attr,omitempty"`
		MediaTagVersion string `json:"mediaTagVersion" xml:"mediaTagVersion,attr,omitempty"`
		Metadata        []struct {
			AllowSync             string `json:"allowSync" xml:"allowSync,attr,omitempty"`
			LibrarySectionID      string `json:"librarySectionID" xml:"librarySectionID,attr,omitempty"`
			LibrarySectionTitle   string `json:"librarySectionTitle" xml:"librarySectionTitle,attr,omitempty"`
			LibrarySectionUUID    string `json:"librarySectionUUID" xml:"librarySectionUUID,attr,omitempty"`
			RatingKey             string `json:"ratingKey" xml:"ratingKey,attr,omitempty"`
			Key                   string `json:"key" xml:"key,attr,omitempty"`
			ParentRatingKey       string `json:"parentRatingKey" xml:"parentRatingKey,attr,omitempty"`
			GrandparentRatingKey  string `json:"grandparentRatingKey" xml:"grandparentRatingKey,attr,omitempty"`
			GUID                  string `json:"guid" xml:"guid,attr,omitempty"`
			LibrarySectionKey     string `json:"librarySectionKey,omitempty"`
			Type                  string `json:"type" xml:"type,attr,omitempty"`
			Title                 string `json:"title" xml:"title,attr,omitempty"`
			GrandparentKey        string `json:"grandparentKey" xml:"grandparentKey,attr,omitempty"`
			ParentKey             string `json:"parentKey" xml:"parentKey,attr,omitempty"`
			GrandparentTitle      string `json:"grandparentTitle" xml:"grandparentTitle,attr,omitempty"`
			ParentTitle           string `json:"parentTitle" xml:"parentTitle,attr,omitempty"`
			ContentRating         string `json:"contentRating" xml:"contentRating,attr,omitempty"`
			Summary               string `json:"summary" xml:"summary,attr,omitempty"`
			Index                 int    `json:"index,omitempty"`
			ParentIndex           int    `json:"parentIndex,omitempty"`
			Rating                int    `json:"rating,omitempty"`
			Year                  string `json:"year" xml:"year,attr,omitempty"`
			Thumb                 string `json:"thumb" xml:"thumb,attr,omitempty"`
			Art                   string `json:"art" xml:"art,attr,omitempty"`
			ParentThumb           string `json:"parentThumb,omitempty"`
			GrandParentThumb      string `json:"grandparentThumb,omitempty"`
			GrandParentArt        string `json:"grandparentArt,omitempty"`
			Duration              int    `json:"duration" xml:"duration,attr,omitempty"`
			OriginallyAvailableAt string `json:"originallyAvailableAt" xml:"originallyAvailableAt,attr,omitempty"`
			AddedAt               string `json:"addedAt" xml:"addedAt,attr,omitempty"`
			UpdatedAt             string `json:"updatedAt" xml:"updatedAt,attr,omitempty"`
			Media                 struct {
				VideoResolution string  `json:"videoResolution" xml:"videoResolution,attr,omitempty"`
				ID              string  `json:"id" xml:"id,attr,omitempty"`
				Duration        int     `json:"duration" xml:"duration,attr,omitempty"`
				Bitrate         int     `json:"bitrate" xml:"bitrate,attr,omitempty"`
				Width           int     `json:"width" xml:"width,attr,omitempty"`
				Height          int     `json:"height" xml:"height,attr,omitempty"`
				AspectRatio     float32 `json:"aspectRatio" xml:"aspectRatio,attr,omitempty"`
				AudioChannels   int     `json:"audioChannels" xml:"audioChannels,attr,omitempty"`
				AudioCodec      string  `json:"audioCodec" xml:"audioCodec,attr,omitempty"`
				VideoCodec      string  `json:"videoCodec" xml:"videoCodec,attr,omitempty"`
				Container       string  `json:"container" xml:"container,attr,omitempty"`
				VideoFrameRate  string  `json:"videoFrameRate" xml:"videoFrameRate,attr,omitempty"`
				VideoProfile    string  `json:"videoProfile" xml:"videoProfile,attr,omitempty"`
				Part            []struct {
					ID           int    `json:"id" xml:"id,attr,omitempty"`
					Key          string `json:"key" xml:"key,attr,omitempty"`
					Duration     int    `json:"duration" xml:"duration,attr,omitempty"`
					File         string `json:"file" xml:"file,attr,omitempty"`
					Size         string `json:"size" xml:"size,attr,omitempty"`
					Container    string `json:"container" xml:"container,attr,omitempty"`
					VideoProfile string `json:"videoProfile" xml:"videoProfile,attr,omitempty"`
					Stream       []struct {
						ID                 int    `json:"id" xml:"id,attr,omitempty"`
						StreamType         int    `json:"streamType" xml:"streamType,attr,omitempty"`
						Default            bool   `json:"default" xml:"default,attr,omitempty"`
						Codec              string `json:"codec" xml:"codec,attr,omitempty"`
						Index              int    `json:"index" xml:"index,attr,omitempty"`
						Bitrate            int    `json:"bitrate" xml:"bitrate,attr,omitempty"`
						BitDepth           int    `json:"bitDepth" xml:"bitDepth,attr,omitempty"`
						ChromaSubsampling  string `json:"chromaSubsampling" xml:"chromaSubsampling,attr,omitempty"`
						ColorRange         string `json:"colorRange" xml:"colorRange,attr,omitempty"`
						ColorSpace         string `json:"colorSpace" xml:"colorSpace,attr,omitempty"`
						FrameRate          string `json:"frameRate" xml:"frameRate,attr,omitempty"`
						HasScalingMatrix   bool   `json:"hasScalingMatrix" xml:"hasScalingMatrix,attr,omitempty"`
						Height             int    `json:"height" xml:"height,attr,omitempty"`
						Level              int    `json:"level" xml:"level,attr,omitempty"`
						Profile            string `json:"profile" xml:"profile,attr,omitempty"`
						RefFrames          int    `json:"refFrames" xml:"refFrames,attr,omitempty"`
						ScanType           string `json:"scanType" xml:"scanType,attr,omitempty"`
						Width              int    `json:"width" xml:"width,attr,omitempty"`
						Selected           bool   `json:"selected" xml:"selected,attr,omitempty"`
						Channels           int    `json:"channels" xml:"channels,attr,omitempty"`
						AudioChannelLayout string `json:"audioChannelLayout" xml:"audioChannelLayout,attr,omitempty"`
						SamplingRate       int    `json:"samplingRate" xml:"samplingRate,attr,omitempty"`
					} `json:"stream" xml:"Stream,omitempty"`
				} `json:"part" xml:"Part,omitempty"`
			} `json:"media" xml:"Media,omitempty"`
			Director []struct {
				ID     string `json:"id,omitempty"`
				Filter string `json:"filter,omitempty"`
				Tag    string `json:"tag,omitempty"`
			} `json:"Director,omitempty"`
			Writer []struct {
				ID     string `json:"id,omitempty"`
				Filter string `json:"filter,omitempty"`
				Tag    string `json:"tag,omitempty"`
			} `json:"Writer,omitempty"`
			Label struct {
				ID  string `json:"id" xml:"id,attr,omitempty"`
				Tag string `json:"tag" xml:"tag,attr,omitempty"`
			} `json:"label" xml:"Label,omitempty"`
			Field []struct {
				Name   string `json:"name" xml:"name,attr,omitempty"`
				Locked string `json:"locked" xml:"locked,attr,omitempty"`
			} `json:"field" xml:"Field,omitempty"`
		} `json:"Metadata" xml:"metadata,omitempty"`
	} `json:"MediaContainer" xml:"MediaContainer,omitempty"`
}
