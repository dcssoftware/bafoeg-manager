package xdomeaconsts

type AussonderungsArt string

const (
	Aussonderungs_Art_Archivwürdigung AussonderungsArt = "A" // Das Schriftgutobjekt ist archivwürdig.
	Aussonderungs_Art_Bewertung       AussonderungsArt = "B" // Das Schriftgutobjekt ist zum Bewerten markiert.
	Aussonderungs_Art_Vernichten      AussonderungsArt = "V" // Das Schriftgutobjekt ist zum Vernichten markiert.
)
