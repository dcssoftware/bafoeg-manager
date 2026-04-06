package xdomeaconsts

type DatenaustauschArt string

const (
	Datenaustausch_Art_SharedDrive      DatenaustauschArt = "001" // Freigegebenes Verzeichnis
	Datenaustausch_Art_DatenbankTabelle DatenaustauschArt = "002" // Datenbanktabelle
	Datenaustausch_Art_WebDavHttps      DatenaustauschArt = "003" // WebDAV und http(s)
)
