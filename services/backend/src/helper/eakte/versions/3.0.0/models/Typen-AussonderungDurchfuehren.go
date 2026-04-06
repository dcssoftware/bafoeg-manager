package models

// AktenplanBewertungskatalogType is Die Datei zum Aktenplan als Primärdokument.
type AktenplanBewertungskatalogType struct {
	ID             string                                            `xml:"ID"`
	Bezeichnung    string                                            `xml:"Bezeichnung"`
	Typ            string                                            `xml:"Typ,omitempty"`
	Version        string                                            `xml:"Version"`
	Einheit        []*AktenplaneinheitAktenplanBewertungskatalogType `xml:"Einheit,omitempty"`
	Gueltigkeit    *ZeitraumType                                     `xml:"Gueltigkeit,omitempty"`
	Aktenplandatei []*FormatType                                     `xml:"Aktenplandatei,omitempty"`
}

// AktenplaneinheitAktenplanBewertungskatalogType is Die Aussonderungsart gibt das Ergebnis der archivischen Bewertung an. Die Aussonderungsart wird vom Aktenplan (zweistufiges Aussonderungsverfahren) auf zugehörige Akten und Vorgänge vererbt.
type AktenplaneinheitAktenplanBewertungskatalogType struct {
	ID               string                `xml:"ID"`
	Aussonderungsart *AussonderungsartType `xml:"Aussonderungsart"`
}

// ErfolgOderMisserfolgAussonderungType is Die ID, unter der das ausgesonderte Schriftgutobjekt im Archiv verwahrt wird. Für den Fall, dass "Erfolgreich" positiv belegt ist und von der aussondernden Behörde die Rückgabe der Archivkennungen gewünscht ist ("RueckmeldungArchivkennung" ist positiv belegt), können durch das archivierende System die Archivkennungen zu den einzelnen Schriftgutobjekt übergeben werden.
type ErfolgOderMisserfolgAussonderungType struct {
	IDSGO             string                           `xml:"IDSGO"`
	Erfolgreich       bool                             `xml:"Erfolgreich"`
	Fehlermeldung     []string                         `xml:"Fehlermeldung,omitempty"`
	FehlermeldungCode []*SonstigeFehlermeldungCodeType `xml:"FehlermeldungCode,omitempty"`
	Archivkennung     string                           `xml:"Archivkennung,omitempty"`
}

// ErfolgOderMisserfolgBewertungskatalogType is Erläuterung des Grundes als Code für den nicht erfolgreichen Import einer Bewertung zu einer Aktenplaneinheit.
type ErfolgOderMisserfolgBewertungskatalogType struct {
	ID                string                           `xml:"ID"`
	Erfolgreich       bool                             `xml:"Erfolgreich"`
	Fehlermeldung     []string                         `xml:"Fehlermeldung,omitempty"`
	FehlermeldungCode []*SonstigeFehlermeldungCodeType `xml:"FehlermeldungCode,omitempty"`
}

// NkAussonderungType is Die Angabe, ob vom Empfänger der Nachricht eine Empfangsbestätigung erwünscht wird. Eine erwünschte Empfangsbestätigung wird mit 1 gekennzeichnet. Wird keine Empfangsbestätigung gewünscht, so wird dies mit 0 gekennzeichnet.
type NkAussonderungType struct {
	Importbestaetigung        bool `xml:"Importbestaetigung"`
	RueckmeldungArchivkennung bool `xml:"RueckmeldungArchivkennung"`
	Empfangsbestaetigung      bool `xml:"Empfangsbestaetigung"`
	*NkBasisType
}

// RueckgabeparameterAnbietungType is Die Beschreibung des Bewertungsergebnisses für ein angebotenes Schriftgutobjekt.
type RueckgabeparameterAnbietungType struct {
	ID               string                `xml:"ID"`
	Aussonderungsart *AussonderungsartType `xml:"Aussonderungsart"`
}
