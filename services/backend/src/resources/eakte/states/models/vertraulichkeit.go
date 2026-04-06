package models

type VertraulichkeitEnum string

const (
	VertraulichkeitIntern      VertraulichkeitEnum = "INTERN"
	VertraulichkeitVertraulich VertraulichkeitEnum = "VERTRAULICH"
	VertraulichkeitOffen       VertraulichkeitEnum = "OFFEN"
)

func (e VertraulichkeitEnum) String() string {
	return string(e)
}
