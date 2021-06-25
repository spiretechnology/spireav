package meta

import (
	"encoding/json"
	"os/exec"
)

type Meta struct {
	Format  FormatMeta   `json:"format"`
	Streams []StreamMeta `json:"streams"`
}

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

func GetMetadata(
	filename string,
	ffprobePath string,
) (*Meta, error) {

	// If there is no probe path, use the default
	if len(ffprobePath) == 0 {
		ffprobePath = "ffprobe"
	}

	// Create the command
	cmd := exec.Command(ffprobePath, "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", filename)

	// Run the command
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON output
	var meta Meta
	if err := json.Unmarshal(output, &meta); err != nil {
		return nil, err
	}

	// Return the metadata
	return &meta, nil

}