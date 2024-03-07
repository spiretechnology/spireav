package metadata

import (
	"strconv"
)

// Meta is the overall metadata for a media file
type Meta struct {
	Format  FormatMeta   `json:"format"`
	Streams []StreamMeta `json:"streams"`
}

// FormatMeta is the metadata specific to the file format some media is contained in
type FormatMeta struct {
	Filename       string            `json:"filename"`
	NbStreams      int               `json:"nb_streams"`
	NbPrograms     int               `json:"nb_programs"`
	FormatName     string            `json:"format_name"`
	FormatLongName string            `json:"format_long_name"`
	StartTime      string            `json:"start_time"`
	Duration       string            `json:"duration"`
	Size           string            `json:"size"`
	BitRate        string            `json:"bit_rate"`
	ProbeScore     int               `json:"probe_score"`
	Tags           map[string]string `json:"tags"`
}

// StreamMeta is the metadata specific to a stream of data, with a specific codec
type StreamMeta struct {
	Index              int               `json:"index"`
	CodecName          string            `json:"codec_name"`
	CodecLongName      string            `json:"codec_long_name"`
	Profile            string            `json:"profile"`
	CodecType          string            `json:"codec_type"`
	CodecTagString     string            `json:"codec_tag_string"`
	CodecTag           string            `json:"codec_tag"`
	Width              int               `json:"width"`
	Height             int               `json:"height"`
	CodedWidth         int               `json:"coded_width"`
	CodedHeight        int               `json:"coded_height"`
	ClosedCaptions     int               `json:"closed_captions"`
	HasBFrames         int               `json:"has_b_frames"`
	SampleAspectRatio  string            `json:"sample_aspect_ratio"`
	DisplayAspectRatio string            `json:"display_aspect_ratio"`
	PixFmt             string            `json:"pix_fmt"`
	ColorSpace         string            `json:"color_space"`
	Level              int               `json:"level"`
	ChromaLocation     string            `json:"chroma_location"`
	Refs               int               `json:"refs"`
	IsAvc              string            `json:"is_avc"`
	NalLengthSize      string            `json:"nal_length_size"`
	RFrameRate         string            `json:"r_frame_rate"`
	AvgFrameRate       string            `json:"avg_frame_rate"`
	TimeBase           string            `json:"time_base"`
	StartPts           int               `json:"start_pts"`
	StartTime          string            `json:"start_time"`
	DurationTs         int               `json:"duration_ts"`
	Duration           string            `json:"duration"`
	BitRate            string            `json:"bit_rate"`
	BitsPerRawSample   string            `json:"bits_per_raw_sample"`
	NbFrames           string            `json:"nb_frames"`
	Disposition        StreamDisposition `json:"disposition"`
	Tags               map[string]string `json:"tags"`
}

type StreamDisposition struct {
	Default         int `json:"default"`
	Dub             int `json:"dub"`
	Original        int `json:"original"`
	Comment         int `json:"comment"`
	Lyrics          int `json:"lyrics"`
	Karaoke         int `json:"karaoke"`
	Forced          int `json:"forced"`
	HearingImpaired int `json:"hearing_impaired"`
	VisualImpaired  int `json:"visual_impaired"`
	CleanEffects    int `json:"clean_effects"`
	AttachedPic     int `json:"attached_pic"`
	TimedThumbnails int `json:"timed_thumbnails"`
}

// GetVideoStream gets the metadata for the video stream within the metadata
func (m *Meta) GetVideoStream() *StreamMeta {
	for _, v := range m.Streams {
		if v.CodecType == "video" {
			return &v
		}
	}
	return nil
}

// GetLengthFrames gets the duration of the media in frames
func (m *Meta) GetLengthFrames() int {
	video := m.GetVideoStream()
	if video == nil {
		return 0
	}
	return video.DurationTs
}

// GetFrameRate gets the frame rate of the video stream, represented as a float64 of frames per second
func (m *Meta) GetFrameRate() float64 {
	video := m.GetVideoStream()
	if video == nil {
		return 0
	}
	durationSec, err := strconv.ParseFloat(video.Duration, 64)
	if err != nil {
		return 0
	}
	return float64(video.DurationTs) / durationSec
}
