package model

import "encoding/xml"

type MediaMetadata struct {
	XMLName             xml.Name `json:"MediaContainer" xml:"MediaContainer"`
	Size                string   `json:"size" xml:"size,attr"`
	AllowSync           string   `json:"allowSync" xml:"allowSync,attr"`
	Identifier          string   `json:"identifier" xml:"identifier,attr"`
	LibrarySectionID    string   `json:"librarySectionID" xml:"librarySectionID,attr"`
	LibrarySectionTitle string   `json:"librarySectionTitle" xml:"librarySectionTitle,attr"`
	LibrarySectionUUID  string   `json:"librarySectionUUID" xml:"librarySectionUUID,attr"`
	MediaTagPrefix      string   `json:"mediaTagPrefix" xml:"mediaTagPrefix,attr"`
	MediaTagVersion     string   `json:"mediaTagVersion" xml:"mediaTagVersion,attr"`
	Video               struct {
		RatingKey            string `json:"ratingKey" xml:"ratingKey,attr"`
		Key                  string `json:"key" xml:"key,attr"`
		GrandparentTitle     string `json:"grandparentTitle" xml:"grandparentTitle,attr"`
		GrandparentRatingKey string `json:"grandparentRatingKey" xml:"grandparentRatingKey,attr"`
		ParentRatingKey      string `json:"parentRatingKey" xml:"parentRatingKey,attr"`
		ParentYear           string `json:"parentYear" xml:"parentYear,attr"`
		ParentTitle          string `json:"parentTitle" xml:"parentTitle,attr"`
		GUID                 string `json:"guid" xml:"guid,attr"`
		LibrarySectionID     string `json:"librarySectionID" xml:"librarySectionID,attr"`
		Type                 string `json:"type" xml:"type,attr"`
		Title                string `json:"title" xml:"title,attr"`
		Summary              string `json:"summary" xml:"summary,attr"`
		ViewCount            string `json:"viewCount" xml:"viewCount,attr"`
		LastViewedAt         string `json:"lastViewedAt" xml:"lastViewedAt,attr"`
		Year                 string `json:"year" xml:"year,attr"`
		Thumb                string `json:"thumb" xml:"thumb,attr"`
		Art                  string `json:"art" xml:"art,attr"`
		Duration             string `json:"duration" xml:"duration,attr"`
		AddedAt              string `json:"addedAt" xml:"addedAt,attr"`
		UpdatedAt            string `json:"updatedAt" xml:"updatedAt,attr"`
		ChapterSource        string `json:"chapterSource" xml:"chapterSource,attr"`
		Media                struct {
			VideoResolution string `json:"videoResolution" xml:"videoResolution,attr"`
			ID              string `json:"id" xml:"id,attr"`
			Duration        string `json:"duration" xml:"duration,attr"`
			Bitrate         string `json:"bitrate" xml:"bitrate,attr"`
			Width           string `json:"width" xml:"width,attr"`
			Height          string `json:"height" xml:"height,attr"`
			AspectRatio     string `json:"aspectRatio" xml:"aspectRatio,attr"`
			AudioChannels   string `json:"audioChannels" xml:"audioChannels,attr"`
			AudioCodec      string `json:"audioCodec" xml:"audioCodec,attr"`
			VideoCodec      string `json:"videoCodec" xml:"videoCodec,attr"`
			Container       string `json:"container" xml:"container,attr"`
			VideoFrameRate  string `json:"videoFrameRate" xml:"videoFrameRate,attr"`
			VideoProfile    string `json:"videoProfile" xml:"videoProfile,attr"`
			Part            struct {
				ID           string `json:"id" xml:"id,attr"`
				Key          string `json:"key" xml:"key,attr"`
				Duration     string `json:"duration" xml:"duration,attr"`
				File         string `json:"file" xml:"file,attr"`
				Size         string `json:"size" xml:"size,attr"`
				Container    string `json:"container" xml:"container,attr"`
				VideoProfile string `json:"videoProfile" xml:"videoProfile,attr"`
				Stream       []struct {
					ID                 string `json:"id" xml:"id,attr"`
					StreamType         string `json:"streamType" xml:"streamType,attr"`
					Default            string `json:"default" xml:"default,attr"`
					Codec              string `json:"codec" xml:"codec,attr"`
					Index              string `json:"index" xml:"index,attr"`
					Bitrate            string `json:"bitrate" xml:"bitrate,attr"`
					BitDepth           string `json:"bitDepth" xml:"bitDepth,attr"`
					Cabac              string `json:"cabac" xml:"cabac,attr"`
					ChromaSubsampling  string `json:"chromaSubsampling" xml:"chromaSubsampling,attr"`
					CodecID            string `json:"codecID" xml:"codecID,attr"`
					ColorRange         string `json:"colorRange" xml:"colorRange,attr"`
					ColorSpace         string `json:"colorSpace" xml:"colorSpace,attr"`
					Duration           string `json:"duration" xml:"duration,attr"`
					FrameRate          string `json:"frameRate" xml:"frameRate,attr"`
					FrameRateMode      string `json:"frameRateMode" xml:"frameRateMode,attr"`
					HasScalingMatrix   string `json:"hasScalingMatrix" xml:"hasScalingMatrix,attr"`
					HeaderStripping    string `json:"headerStripping" xml:"headerStripping,attr"`
					Height             string `json:"height" xml:"height,attr"`
					Level              string `json:"level" xml:"level,attr"`
					PixelFormat        string `json:"pixelFormat" xml:"pixelFormat,attr"`
					Profile            string `json:"profile" xml:"profile,attr"`
					RefFrames          string `json:"refFrames" xml:"refFrames,attr"`
					ScanType           string `json:"scanType" xml:"scanType,attr"`
					Width              string `json:"width" xml:"width,attr"`
					Selected           string `json:"selected" xml:"selected,attr"`
					Channels           string `json:"channels" xml:"channels,attr"`
					AudioChannelLayout string `json:"audioChannelLayout" xml:"audioChannelLayout,attr"`
					BitrateMode        string `json:"bitrateMode" xml:"bitrateMode,attr"`
					DialogNorm         string `json:"dialogNorm" xml:"dialogNorm,attr"`
					SamplingRate       string `json:"samplingRate" xml:"samplingRate,attr"`
				} `json:"stream" xml:"Stream"`
			} `json:"part" xml:"Part"`
		} `json:"media" xml:"Media"`
		Label struct {
			ID  string `json:"id" xml:"id,attr"`
			Tag string `json:"tag" xml:"tag,attr"`
		} `json:"label" xml:"Label"`
		Field []struct {
			Name   string `json:"name" xml:"name,attr"`
			Locked string `json:"locked" xml:"locked,attr"`
		} `json:"field" xml:"Field"`
	} `json:"video" xml:"Video"`
	Directory struct {
		RatingKey             string `json:"ratingKey" xml:"ratingKey,attr"`
		Key                   string `json:"key" xml:"key,attr"`
		GUID                  string `json:"guid" xml:"guid,attr"`
		LibrarySectionID      string `json:"librarySectionID" xml:"librarySectionID,attr"`
		ParentTitle           string `json:"parentTitle" xml:"parentTitle,attr"`
		ParentYear            string `json:"parentYear" xml:"parentYear,attr"`
		Studio                string `json:"studio" xml:"studio,attr"`
		Type                  string `json:"type" xml:"type,attr"`
		Title                 string `json:"title" xml:"title,attr"`
		TitleSort             string `json:"titleSort" xml:"titleSort,attr"`
		ContentRating         string `json:"contentRating" xml:"contentRating,attr"`
		Summary               string `json:"summary" xml:"summary,attr"`
		Index                 string `json:"index" xml:"index,attr"`
		Rating                string `json:"rating" xml:"rating,attr"`
		ViewCount             string `json:"viewCount" xml:"viewCount,attr"`
		LastViewedAt          string `json:"lastViewedAt" xml:"lastViewedAt,attr"`
		Year                  string `json:"year" xml:"year,attr"`
		Thumb                 string `json:"thumb" xml:"thumb,attr"`
		Art                   string `json:"art" xml:"art,attr"`
		Banner                string `json:"banner" xml:"banner,attr"`
		Theme                 string `json:"theme" xml:"theme,attr"`
		Duration              string `json:"duration" xml:"duration,attr"`
		OriginallyAvailableAt string `json:"originallyAvailableAt" xml:"originallyAvailableAt,attr"`
		LeafCount             string `json:"leafCount" xml:"leafCount,attr"`
		ViewedLeafCount       string `json:"viewedLeafCount" xml:"viewedLeafCount,attr"`
		ChildCount            string `json:"childCount" xml:"childCount,attr"`
		AddedAt               string `json:"addedAt" xml:"addedAt,attr"`
		UpdatedAt             string `json:"updatedAt" xml:"updatedAt,attr"`
		Genre                 []struct {
			ID  string `json:"id" xml:"id,attr"`
			Tag string `json:"tag" xml:"tag,attr"`
		} `json:"genre" xml:"Genre"`
		Role []struct {
			ID    string `json:"id" xml:"id,attr"`
			Tag   string `json:"tag" xml:"tag,attr"`
			Role  string `json:"role" xml:"role,attr"`
			Thumb string `json:"thumb" xml:"thumb,attr"`
		} `json:"role" xml:"Role"`
		Field struct {
			Name   string `json:"name" xml:"name,attr"`
			Locked string `json:"locked" xml:"locked,attr"`
		} `json:"field" xml:"Field"`
		Location string `json:"location" xml:"Location"`
	} `json:"directory" xml:"Directory"`
}
