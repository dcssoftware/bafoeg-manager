package xdomeaconsts

type MediumTyp string

const (
	Medium_Typ_Elektronisch MediumTyp = "001" // Das Schriftgutobjekt liegt ausschließlich in elektronischer Form vor.
	Medium_Typ_Hybrid       MediumTyp = "002" // Das Schriftgutobjekt liegt teilweise in elektronischer Form und teilweise als Papier vor.
	Medium_Typ_Papier       MediumTyp = "003" // Das Schriftgutobjekt liegt ausschließlich als Papier vor.
)
