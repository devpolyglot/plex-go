package plex

import "encoding/xml"

type SessionsService service

// TranscodeSessionsResponse is the result for transcode session endpoint /transcode/sessions
type TranscodeSessionsResponse struct {
	Children []struct {
		ElementType   string  `json:"_elementType"`
		AudioChannels int     `json:"audioChannels"`
		AudioCodec    string  `json:"audioCodec"`
		AudioDecision string  `json:"audioDecision"`
		Container     string  `json:"container"`
		Context       string  `json:"context"`
		Duration      int     `json:"duration"`
		Height        int     `json:"height"`
		Key           string  `json:"key"`
		Progress      float64 `json:"progress"`
		Protocol      string  `json:"protocol"`
		Remaining     int     `json:"remaining"`
		Speed         float64 `json:"speed"`
		Throttled     bool    `json:"throttled"`
		VideoCodec    string  `json:"videoCodec"`
		VideoDecision string  `json:"videoDecision"`
		Width         int     `json:"width"`
	} `json:"_children"`
	ElementType string `json:"_elementType"`
}

// CurrentSessionsVideo are current video sessions
type CurrentSessionsVideo struct {
	AddedAt               string `json:"addedAt" xml:"addedAt,attr"`
	Art                   string `json:"art" xml:"art,attr"`
	ChapterSource         string `json:"chapterSource" xml:"chapterSource,attr"`
	ContentRating         string `json:"contentRating" xml:"contentRating,attr"`
	Duration              int    `json:"duration" xml:"duration,attr"`
	GUID                  string `json:"guid" xml:"guid,attr"`
	Key                   string `json:"key" xml:"key,attr"`
	LibrarySectionID      string `json:"librarySectionID" xml:"librarySectionID,attr"`
	OriginallyAvailableAt string `json:"originallyAvailableAt" xml:"originallyAvailableAt,attr"`
	PrimaryExtraKey       string `json:"primaryExtraKey" xml:"primaryExtraKey,attr"`
	Rating                string `json:"rating" xml:"rating,attr"`
	RatingKey             string `json:"ratingKey" xml:"ratingKey,attr"`
	SessionKey            string `json:"sessionKey" xml:"sessionKey,attr"`
	Studio                string `json:"studio" xml:"studio,attr"`
	Summary               string `json:"summary" xml:"summary,attr"`
	Tagline               string `json:"tagline" xml:"tagline,attr"`
	Thumb                 string `json:"thumb" xml:"thumb,attr"`
	Title                 string `json:"title" xml:"title,attr"`
	TitleSort             string `json:"titleSort" xml:"titleSort,attr"`
	Type                  string `json:"type" xml:"type,attr"`
	UpdatedAt             string `json:"updatedAt" xml:"updatedAt,attr"`
	ViewOffset            int    `json:"viewOffset" xml:"viewOffset,attr"`
	Year                  string `json:"year" xml:"year,attr"`
	Media                 struct {
		AspectRatio           string `json:"aspectRatio" xml:"aspectRatio,attr"`
		AudioChannels         string `json:"audioChannels" xml:"audioChannels,attr"`
		AudioCodec            string `json:"audioCodec" xml:"audioCodec,attr"`
		AudioProfile          string `json:"audioProfile" xml:"audioProfile,attr"`
		Bitrate               string `json:"bitrate" xml:"bitrate,attr"`
		Container             string `json:"container" xml:"container,attr"`
		Duration              string `json:"duration" xml:"duration,attr"`
		Has64bitOffsets       string `json:"has64bitOffsets" xml:"has64bitOffsets,attr"`
		Height                string `json:"height" xml:"height,attr"`
		ID                    string `json:"id" xml:"id,attr"`
		OptimizedForStreaming string `json:"optimizedForStreaming" xml:"optimizedForStreaming,attr"`
		VideoCodec            string `json:"videoCodec" xml:"videoCodec,attr"`
		VideoFrameRate        string `json:"videoFrameRate" xml:"videoFrameRate,attr"`
		VideoProfile          string `json:"videoProfile" xml:"videoProfile,attr"`
		VideoResolution       string `json:"videoResolution" xml:"videoResolution,attr"`
		Width                 string `json:"width" xml:"width,attr"`
		Part                  struct {
			AudioProfile          string `json:"audioProfile" xml:"audioProfile,attr"`
			Container             string `json:"container" xml:"container,attr"`
			Duration              string `json:"duration" xml:"duration,attr"`
			File                  string `json:"file" xml:"file,attr"`
			Has64bitOffsets       string `json:"has64bitOffsets" xml:"has64bitOffsets,attr"`
			ID                    string `json:"id" xml:"id,attr"`
			Indexes               string `json:"indexes" xml:"indexes,attr"`
			Key                   string `json:"key" xml:"key,attr"`
			OptimizedForStreaming string `json:"optimizedForStreaming" xml:"optimizedForStreaming,attr"`
			Size                  string `json:"size" xml:"size,attr"`
			VideoProfile          string `json:"videoProfile" xml:"videoProfile,attr"`
			Stream                []struct {
				BitDepth           string `json:"bitDepth" xml:"bitDepth,attr"`
				Bitrate            string `json:"bitrate" xml:"bitrate,attr"`
				Cabac              string `json:"cabac" xml:"cabac,attr"`
				ChromaSubsampling  string `json:"chromaSubsampling" xml:"chromaSubsampling,attr"`
				Codec              string `json:"codec" xml:"codec,attr"`
				CodecID            string `json:"codecID" xml:"codecID,attr"`
				ColorRange         string `json:"colorRange" xml:"colorRange,attr"`
				ColorSpace         string `json:"colorSpace" xml:"colorSpace,attr"`
				Default            string `json:"default" xml:"default,attr"`
				Duration           string `json:"duration" xml:"duration,attr"`
				FrameRate          string `json:"frameRate" xml:"frameRate,attr"`
				FrameRateMode      string `json:"frameRateMode" xml:"frameRateMode,attr"`
				HasScalingMatrix   string `json:"hasScalingMatrix" xml:"hasScalingMatrix,attr"`
				Height             string `json:"height" xml:"height,attr"`
				ID                 string `json:"id" xml:"id,attr"`
				Index              string `json:"index" xml:"index,attr"`
				Level              string `json:"level" xml:"level,attr"`
				PixelFormat        string `json:"pixelFormat" xml:"pixelFormat,attr"`
				Profile            string `json:"profile" xml:"profile,attr"`
				RefFrames          string `json:"refFrames" xml:"refFrames,attr"`
				ScanType           string `json:"scanType" xml:"scanType,attr"`
				StreamIdentifier   string `json:"streamIdentifier" xml:"streamIdentifier,attr"`
				StreamType         string `json:"streamType" xml:"streamType,attr"`
				Width              string `json:"width" xml:"width,attr"`
				AudioChannelLayout string `json:"audioChannelLayout" xml:"audioChannelLayout,attr"`
				BitrateMode        string `json:"bitrateMode" xml:"bitrateMode,attr"`
				Channels           string `json:"channels" xml:"channels,attr"`
				Language           string `json:"language" xml:"language,attr"`
				LanguageCode       string `json:"languageCode" xml:"languageCode,attr"`
				SamplingRate       string `json:"samplingRate" xml:"samplingRate,attr"`
				Selected           string `json:"selected" xml:"selected,attr"`
				Format             string `json:"format" xml:"format,attr"`
				Key                string `json:"key" xml:"key,attr"`
			} `json:"Stream" xml:"Stream"`
		} `json:"Part" xml:"Part"`
	} `json:"Media" xml:"Media"`
	Genre []struct {
		Count string `json:"count" xml:"count,attr"`
		ID    string `json:"id" xml:"id,attr"`
		Tag   string `json:"tag" xml:"tag,attr"`
	} `json:"Genre" xml:"Genre"`
	Writer []struct {
		ID    string `json:"id" xml:"id,attr"`
		Tag   string `json:"tag" xml:"tag,attr"`
		Count string `json:"count" xml:"count,attr"`
	} `json:"Writer" xml:"Writer"`
	Director struct {
		Count string `json:"count" xml:"count,attr"`
		ID    string `json:"id" xml:"id,attr"`
		Tag   string `json:"tag" xml:"tag,attr"`
	} `json:"Director" xml:"Director"`
	Producer []struct {
		Count string `json:"count" xml:"count,attr"`
		ID    string `json:"id" xml:"id,attr"`
		Tag   string `json:"tag" xml:"tag,attr"`
	} `json:"Producer" xml:"Producer"`
	Country struct {
		Count string `json:"count" xml:"count,attr"`
		ID    string `json:"id" xml:"id,attr"`
		Tag   string `json:"tag" xml:"tag,attr"`
	} `json:"Country" xml:"Country"`
	Role []struct {
		Count string `json:"count" xml:"count,attr"`
		ID    string `json:"id" xml:"id,attr"`
		Role  string `json:"role" xml:"role,attr"`
		Tag   string `json:"tag" xml:"tag,attr"`
	} `json:"Role" xml:"Role"`
	Collection struct {
		Count string `json:"count" xml:"count,attr"`
		ID    string `json:"id" xml:"id,attr"`
		Tag   string `json:"tag" xml:"tag,attr"`
	} `json:"Collection" xml:"Collection"`
	Label struct {
		ID  string `json:"id" xml:"id,attr"`
		Tag string `json:"tag" xml:"tag,attr"`
	} `json:"Label" xml:"Label"`
	Field struct {
		Locked string `json:"locked" xml:"locked,attr"`
		Name   string `json:"name" xml:"name,attr"`
	} `json:"Field" xml:"Field"`
	User struct {
		ID    int    `json:"id" xml:"id,attr"`
		Title string `json:"title" xml:"title,attr"`
		Thumb string `json:"thumb" xml:"thumb,attr"`
	} `json:"User" xml:"User"`
	Player struct {
		Address           string `json:"address" xml:"address,attr"`
		Device            string `json:"device" xml:"device,attr"`
		MachineIdentifier string `json:"machineIdentifier" xml:"machineIdentifier,attr"`
		Model             string `json:"model" xml:"model,attr"`
		Platform          string `json:"platform" xml:"platform,attr"`
		PlatformVersion   string `json:"platformVersion" xml:"platformVersion,attr"`
		Product           string `json:"product" xml:"product,attr"`
		Profile           string `json:"profile" xml:"profile,attr"`
		State             string `json:"state" xml:"state,attr"`
		Title             string `json:"title" xml:"title,attr"`
		Vendor            string `json:"vendor" xml:"vendor,attr"`
		Version           string `json:"version" xml:"version,attr"`
	} `json:"Player" xml:"Player"`
	GrandparentArt       string `json:"grandparentArt" xml:"grandparentArt,attr"`
	GrandparentKey       string `json:"grandparentKey" xml:"grandparentKey,attr"`
	GrandparentRatingKey string `json:"grandparentRatingKey" xml:"grandparentRatingKey,attr"`
	GrandparentTheme     string `json:"grandparentTheme" xml:"grandparentTheme,attr"`
	GrandparentThumb     string `json:"grandparentThumb" xml:"grandparentThumb,attr"`
	GrandparentTitle     string `json:"grandparentTitle" xml:"grandparentTitle,attr"`
	Index                string `json:"index" xml:"index,attr"`
	LastViewedAt         string `json:"lastViewedAt" xml:"lastViewedAt,attr"`
	ParentIndex          string `json:"parentIndex" xml:"parentIndex,attr"`
	ParentKey            string `json:"parentKey" xml:"parentKey,attr"`
	ParentRatingKey      string `json:"parentRatingKey" xml:"parentRatingKey,attr"`
	ParentThumb          string `json:"parentThumb" xml:"parentThumb,attr"`
	ViewCount            string `json:"viewCount" xml:"viewCount,attr"`
	Session              struct {
		ID        string `json:"id" xml:"id,attr"`
		Bandwidth string `json:"bandwidth" xml:"bandwidth,attr"`
		Location  string `json:"location" xml:"location,attr"`
	} `json:"Session" xml:"Session"`
	TranscodeSession struct {
		Key           string `json:"key" xml:"key,attr"`
		Throttled     string `json:"throttled" xml:"throttled,attr"`
		Progress      string `json:"progress" xml:"progress,attr"`
		Speed         string `json:"speed" xml:"speed,attr"`
		Duration      string `json:"duration" xml:"duration,attr"`
		Remaining     string `json:"remaining" xml:"remaining,attr"`
		Context       string `json:"context" xml:"context,attr"`
		VideoDecision string `json:"videoDecision" xml:"videoDecision,attr"`
		AudioDecision string `json:"audioDecision" xml:"audioDecision,attr"`
		Protocol      string `json:"protocol" xml:"protocol,attr"`
		Container     string `json:"container" xml:"container,attr"`
		VideoCodec    string `json:"videoCodec" xml:"videoCodec,attr"`
		AudioCodec    string `json:"audioCodec" xml:"audioCodec,attr"`
		AudioChannels string `json:"audioChannels" xml:"audioChannels,attr"`
		Width         string `json:"width" xml:"width,attr"`
		Height        string `json:"height" xml:"height,attr"`
	} `json:"TranscodeSession" xml:"TranscodeSession"`
}

// CurrentSessions is xml because plex returns a dynamic type (string or number) for the duration field
type CurrentSessions struct {
	XMLName xml.Name               `xml:"MediaContainer"`
	Size    string                 `xml:"size,attr"`
	Video   []CurrentSessionsVideo `xml:"Video"`
	Track   []struct {
		AddedAt              string `xml:"addedAt,attr"`
		Art                  string `xml:"art,attr"`
		ChapterSource        string `xml:"chapterSource,attr"`
		Duration             int    `xml:"duration,attr"`
		GrandparentArt       string `xml:"grandparentArt,attr"`
		GrandparentKey       string `xml:"grandparentKey,attr"`
		GrandparentRatingKey string `xml:"grandparentRatingKey,attr"`
		GrandparentThumb     string `xml:"grandparentThumb,attr"`
		GrandparentTitle     string `xml:"grandparentTitle,attr"`
		GUID                 string `xml:"guid,attr"`
		Index                string `xml:"index,attr"`
		Key                  string `xml:"key,attr"`
		LastViewedAt         string `xml:"lastViewedAt,attr"`
		LibrarySectionID     string `xml:"librarySectionID,attr"`
		ParentIndex          string `xml:"parentIndex,attr"`
		ParentKey            string `xml:"parentKey,attr"`
		ParentRatingKey      string `xml:"parentRatingKey,attr"`
		ParentTitle          string `xml:"parentTitle,attr"`
		RatingKey            string `xml:"ratingKey,attr"`
		SessionKey           string `xml:"sessionKey,attr"`
		Summary              string `xml:"summary,attr"`
		Tagline              string `xml:"tagline,attr"`
		Thumb                string `xml:"thumb,attr"`
		Title                string `xml:"title,attr"`
		Type                 string `xml:"type,attr"`
		UpdatedAt            string `xml:"updatedAt,attr"`
		ViewCount            int    `xml:"viewCount,attr"`
		ViewOffset           int    `xml:"viewOffset,attr"`
		Media                struct {
			AudioChannels string `xml:"audioChannels,attr"`
			AudioCodec    string `xml:"audioCodec,attr"`
			Bitrate       string `xml:"bitrate,attr"`
			Container     string `xml:"container,attr"`
			Duration      string `xml:"duration,attr"`
			ID            string `xml:"id,attr"`
			Part          struct {
				Container string `xml:"container,attr"`
				Duration  string `xml:"duration,attr"`
				File      string `xml:"file,attr"`
				ID        string `xml:"id,attr"`
				Key       string `xml:"key,attr"`
				Size      string `xml:"size,attr"`
				Stream    []struct {
					AudioChannelLayout string `xml:"audioChannelLayout,attr"`
					Bitrate            string `xml:"bitrate,attr"`
					BitrateMode        string `xml:"bitrateMode,attr"`
					Channels           string `xml:"channels,attr"`
					Codec              string `xml:"codec,attr"`
					Duration           string `xml:"duration,attr"`
					ID                 string `xml:"id,attr"`
					Index              string `xml:"index,attr"`
					SamplingRate       string `xml:"samplingRate,attr"`
					Selected           string `xml:"selected,attr"`
					StreamType         string `xml:"streamType,attr"`
				} `xml:"Stream"`
			} `xml:"Part"`
		} `xml:"Media"`
		User struct {
			ID    int    `xml:"id,attr"`
			Title string `xml:"title,attr"`
			Thumb string `xml:"thumb,attr"`
		} `xml:"User"`
		Player struct {
			Address           string `xml:"address,attr"`
			Device            string `xml:"device,attr"`
			MachineIdentifier string `xml:"machineIdentifier,attr"`
			Model             string `xml:"model,attr"`
			Platform          string `xml:"platform,attr"`
			PlatformVersion   string `xml:"platformVersion,attr"`
			Product           string `xml:"product,attr"`
			Profile           string `xml:"profile,attr"`
			State             string `xml:"state,attr"`
			Title             string `xml:"title,attr"`
			Vendor            string `xml:"vendor,attr"`
			Version           string `xml:"version,attr"`
		} `xml:"Player"`
		TranscodeSession struct {
			Key           string `xml:"key,attr"`
			Throttled     string `xml:"throttled,attr"`
			Progress      string `xml:"progress,attr"`
			Speed         string `xml:"speed,attr"`
			Duration      string `xml:"duration,attr"`
			Remaining     string `xml:"remaining,attr"`
			Context       string `xml:"context,attr"`
			VideoDecision string `xml:"videoDecision,attr"`
			AudioDecision string `xml:"audioDecision,attr"`
			Protocol      string `xml:"protocol,attr"`
			Container     string `xml:"container,attr"`
			VideoCodec    string `xml:"videoCodec,attr"`
			AudioCodec    string `xml:"audioCodec,attr"`
			AudioChannels string `xml:"audioChannels,attr"`
			Width         string `xml:"width,attr"`
			Height        string `xml:"height,attr"`
		} `xml:"TranscodeSession"`
	} `xml:"Track"`
}
