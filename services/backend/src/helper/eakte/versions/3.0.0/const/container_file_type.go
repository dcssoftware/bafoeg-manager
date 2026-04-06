package xdomeaconsts

type ContainerFileType string

const (
	ContainerFileTypeZip    ContainerFileType = "zip"
	ContainerFileTypeXdomea ContainerFileType = "xdomea"
)

func (fileType ContainerFileType) GetFileEnding() string {
	switch fileType {
	case ContainerFileTypeXdomea:
		return "xdomea"
	default:
		return "zip"
	}
}

func (ftype ContainerFileType) GetNameLowercase() string {
	return string(ftype)
}
