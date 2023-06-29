package mxf2raw

import "encoding/xml"

type Mxf2RawResult struct {
	XMLName     xml.Name    `xml:"bmx"`
	Created     string      `xml:"created,attr"`
	Version     string      `xml:"version,attr"`
	Application Application `xml:"application"`
	Files       Files       `xml:"files"`
	Clip        Clip        `xml:"clip"`
	LogMessages LogMessages `xml:"log_messages"`
}

type Application struct {
	XMLName        xml.Name `xml:"application"`
	Name           string   `xml:"name"`
	LibName        string   `xml:"lib_name"`
	LibVersion     string   `xml:"lib_version"`
	BuildTimestamp string   `xml:"build_timestamp"`
	ScmVersion     string   `xml:"scm_version"`
}

type Files struct {
	XMLName xml.Name `xml:"files"`
	Size    int      `xml:"size,attr"`
	Files   []File   `xml:"file"`
}

type File struct {
	XMLName            xml.Name           `xml:"file"`
	Index              int                `xml:"index,attr"`
	FileURI            string             `xml:"file_uri"`
	MxfVersion         string             `xml:"mxf_version"`
	OpLabel            string             `xml:"op_label"`
	FrameWrapped       bool               `xml:"frame_wrapped"`
	Identifications    Identifications    `xml:"identifications"`
	MaterialPackages   MaterialPackages   `xml:"material_packages"`
	FileSourcePackages FileSourcePackages `xml:"file_source_packages"`
}

type OpLabel struct {
	XMLName           xml.Name `xml:"op_label"`
	MultiSourceClip   bool     `xml:"multi_source_clip"`
	MultiEssenceTrack bool     `xml:"multi_essence_track"`
	Label             string   `xml:"label"`
	OpPattern         string   `xml:",chardata"`
}

type Identifications struct {
	XMLName         xml.Name         `xml:"identifications"`
	Size            int              `xml:"size,attr"`
	Identifications []Identification `xml:"identification"`
}

type Identification struct {
	XMLName        xml.Name `xml:"identification"`
	Index          int      `xml:"index,attr"`
	GenerationUID  string   `xml:"generation_uid"`
	CompanyName    string   `xml:"company_name"`
	ProductName    string   `xml:"product_name"`
	ProductVersion string   `xml:"product_version"`
	VersionString  string   `xml:"version_string"`
	ProductUID     string   `xml:"product_uid"`
	ModifiedDate   string   `xml:"modified_date"`
	ToolkitVersion string   `xml:"toolkit_version"`
	Platform       string   `xml:"platform"`
}

type MaterialPackages struct {
	XMLName          xml.Name          `xml:"material_packages"`
	Size             int               `xml:"size,attr"`
	MaterialPackages []MaterialPackage `xml:"material_package"`
}

type MaterialPackage struct {
	XMLName      xml.Name `xml:"material_package"`
	Index        int      `xml:"index,attr"`
	PackageUID   string   `xml:"package_uid"`
	Name         string   `xml:"name"`
	CreationDate string   `xml:"creation_date"`
	ModifiedDate string   `xml:"modified_date"`
}

type FileSourcePackages struct {
	XMLName            xml.Name            `xml:"file_source_packages"`
	Size               int                 `xml:"size,attr"`
	FileSourcePackages []FileSourcePackage `xml:"file_source_package"`
}

type FileSourcePackage struct {
	XMLName      xml.Name `xml:"file_source_package"`
	Index        int      `xml:"index,attr"`
	PackageUID   string   `xml:"package_uid"`
	Name         string   `xml:"name"`
	CreationDate string   `xml:"creation_date"`
	ModifiedDate string   `xml:"modified_date"`
}

type Clip struct {
	XMLName        xml.Name       `xml:"clip"`
	Name           string         `xml:"name"`
	EditRate       string         `xml:"edit_rate"`
	Duration       Duration       `xml:"duration"`
	StartTimecodes StartTimecodes `xml:"start_timecodes"`
	Tracks         Tracks         `xml:"tracks"`
}

type Duration struct {
	XMLName  xml.Name `xml:"duration"`
	Count    int      `xml:"count,attr"`
	Timecode string   `xml:",chardata"`
}

type StartTimecodes struct {
	XMLName        xml.Name       `xml:"start_timecodes"`
	PhysicalSource PhysicalSource `xml:"physical_source"`
}

type PhysicalSource struct {
	XMLName  xml.Name `xml:"physical_source"`
	Name     string   `xml:"name,attr"`
	Rate     string   `xml:"rate,attr"`
	Timecode string   `xml:",chardata"`
}

type Tracks struct {
	XMLName xml.Name `xml:"tracks"`
	Size    int      `xml:"size,attr"`
	Tracks  []Track  `xml:"track"`
}

type Track struct {
	XMLName         xml.Name        `xml:"track"`
	Index           int             `xml:"index,attr"`
	EssenceKind     string          `xml:"essence_kind"`
	EssenceType     string          `xml:"essence_type"`
	ECLabel         string          `xml:"ec_label"`
	EditRate        string          `xml:"edit_rate"`
	Duration        Duration        `xml:"duration"`
	Packages        Packages        `xml:"packages"`
	SoundDescriptor SoundDescriptor `xml:"sound_descriptor"`
}

type Packages struct {
	XMLName  xml.Name  `xml:"packages"`
	Size     int       `xml:"size,attr"`
	Packages []Package `xml:"package"`
}

type Package struct {
	XMLName    xml.Name   `xml:"package"`
	Index      int        `xml:"index,attr"`
	Material   Material   `xml:"material"`
	FileSource FileSource `xml:"file_source"`
}

type Material struct {
	XMLName     xml.Name `xml:"material"`
	PackageUID  string   `xml:"package_uid"`
	TrackID     int      `xml:"track_id"`
	TrackNumber int      `xml:"track_number"`
}

type FileSource struct {
	XMLName     xml.Name `xml:"file_source"`
	PackageUID  string   `xml:"package_uid"`
	TrackID     int      `xml:"track_id"`
	TrackNumber string   `xml:"track_number"`
	FileURI     string   `xml:"file_uri"`
}

type SoundDescriptor struct {
	XMLName       xml.Name `xml:"sound_descriptor"`
	SamplingRate  string   `xml:"sampling_rate"`
	BitsPerSample int      `xml:"bits_per_sample"`
	BlockAlign    int      `xml:"block_align"`
	ChannelCount  int      `xml:"channel_count"`
}

type LogMessages struct {
	XMLName  xml.Name  `xml:"log_messages"`
	Size     int       `xml:"size,attr"`
	Warnings []Warning `xml:"warning"`
}

type Warning struct {
	XMLName xml.Name `xml:"warning"`
	Source  string   `xml:"source,attr"`
	Message string   `xml:",chardata"`
}
