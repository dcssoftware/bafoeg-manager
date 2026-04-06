package consts

type BafögType string

const (
	BafögTypeSchueler BafögType = "SCHUELERBAFOEG"
	BafögTypeStudent  BafögType = "STUDIERENDENBAFOEG"
)

func (b BafögType) String() string {
	return string(b)
}
