package models

import (
	"time"
)

// Akteninhalt is Über diese Eigenschaft wird eine Teilakte hierarchisch in eine Akte eingebunden bzw. einer Akte untergeordnet.
type Akteninhalt struct {
	Dokument []*DokumentType `xml:"Dokument,omitempty"`
	Vorgang  []*VorgangType  `xml:"Vorgang,omitempty"`
	Teilakte []*AkteType     `xml:"Teilakte,omitempty"`
}

// AkteType is Über InternerGeschaeftsgang wird der Akte zum Nachweis des internen Geschäftsgangs eine spezifische Ausprägung eines Geschäftsgangs zugeordnet.
type AkteType struct {
	Identifikation                      *IdentifikationObjektType                `xml:"Identifikation"`
	AllgemeineMetadaten                 *AllgemeineMetadatenType                 `xml:"AllgemeineMetadaten,omitempty"`
	ArchivspezifischeMetadaten          *MetadatenAussonderungType               `xml:"ArchivspezifischeMetadaten,omitempty"`
	Standort                            string                                   `xml:"Standort,omitempty"`
	Typ                                 string                                   `xml:"Typ,omitempty"`
	Laufzeit                            *ZeitraumType                            `xml:"Laufzeit,omitempty"`
	HistorienProtokollInformation       []*HistorienProtokollInformationType     `xml:"HistorienProtokollInformation,omitempty"`
	InternerGeschaeftsgang              *GeschaeftsgangType                      `xml:"InternerGeschaeftsgang,omitempty"`
	Akteninhalt                         *Akteninhalt                             `xml:"Akteninhalt,omitempty"`
	Verweis                             []*VerweisType                           `xml:"Verweis,omitempty"`
	Kontakt                             []*KontaktType                           `xml:"Kontakt,omitempty"`
	ZdA                                 bool                                     `xml:"ZdA,omitempty"`
	ZdADatum                            string                                   `xml:"ZdADatum,omitempty"`
	AnwendungsspezifischeErweiterung    *AnwendungsspezifischeErweiterungType    `xml:"AnwendungsspezifischeErweiterung,omitempty"`
	AnwendungsspezifischeErweiterungXML *AnwendungsspezifischeErweiterungXMLType `xml:"AnwendungsspezifischeErweiterungXML,omitempty"`
}

// AktenplanType is Die Datei zum Aktenplan als Primärdokument.
type AktenplanType struct {
	Bezeichnung    string                           `xml:"Bezeichnung"`
	Typ            string                           `xml:"Typ,omitempty"`
	Version        string                           `xml:"Version"`
	Einheit        []*AktenplaneinheitAktenplanType `xml:"Einheit,omitempty"`
	Gueltigkeit    *ZeitraumType                    `xml:"Gueltigkeit,omitempty"`
	Aktenplandatei []*FormatType                    `xml:"Aktenplandatei,omitempty"`
}

// AktenplaneinheitAktenplanType is AnwendungsspezifischeErweiterungXML darf nur dazu genutzt werden, weitere (z.B. fachspezifische) Metadaten zu spezifizieren, deren Übermittlung mit den bereits in xdomea spezifizierten Metadaten nicht möglich ist. Die AnwendungsspezifischeErweiterungXML bietet die Möglichkeit, mittels Einbindung externer XML-Schemata diese Metadaten zu spezifizieren. Es können beliebige XML-Schemata mit unterschiedlichen Namensräumen angegeben werden. Die XML-Schema-Validierung der weiterführenden Metadaten erfolgt innerhalb der xdomea-Nachricht selbst "lax".
type AktenplaneinheitAktenplanType struct {
	Aussonderungsart                    *AussonderungsartType                    `xml:"Aussonderungsart"`
	Gueltigkeit                         *ZeitraumType                            `xml:"Gueltigkeit,omitempty"`
	Stillgelegt                         bool                                     `xml:"Stillgelegt"`
	Aufbewahrungsdauer                  *AufbewahrungsdauerType                  `xml:"Aufbewahrungsdauer"`
	Einheit                             []*AktenplaneinheitAktenplanType         `xml:"Einheit,omitempty"`
	VerweisAktenplaneinheit             []string                                 `xml:"VerweisAktenplaneinheit,omitempty"`
	AnwendungsspezifischeErweiterung    *AnwendungsspezifischeErweiterungType    `xml:"AnwendungsspezifischeErweiterung,omitempty"`
	AnwendungsspezifischeErweiterungXML *AnwendungsspezifischeErweiterungXMLType `xml:"AnwendungsspezifischeErweiterungXML,omitempty"`
	*AktenplaneinheitType
}

// AktenplaneinheitType is Die kurze Beschreibung der Aufgabe, die durch die Aktenplaneinheit repräsentiert wird.
type AktenplaneinheitType struct {
	Kennzeichen   string `xml:"Kennzeichen,omitempty"`
	Inhaltsangabe string `xml:"Inhaltsangabe,omitempty"`
	BetreffKurz   string `xml:"BetreffKurz,omitempty"`
}

// AllgemeineMetadatenType is Die Aktenplaneinheit, dem das Schriftgutobjekt zugeordnet ist.
type AllgemeineMetadatenType struct {
	Betreff               string                         `xml:"Betreff,omitempty"`
	Kennzeichen           string                         `xml:"Kennzeichen,omitempty"`
	Federfuehrung         string                         `xml:"Federfuehrung,omitempty"`
	Aktenfuehrung         string                         `xml:"Aktenfuehrung,omitempty"`
	Vertraulichkeitsstufe *VertraulichkeitsstufeCodeType `xml:"Vertraulichkeitsstufe,omitempty"`
	Bemerkung             string                         `xml:"Bemerkung,omitempty"`
	Medium                *MediumCodeType                `xml:"Medium,omitempty"`
	Aktenplaneinheit      *AktenplaneinheitType          `xml:"Aktenplaneinheit,omitempty"`
}

// AllgemeinerNameType is Der Name ist der eigentliche Familien- oder Vorname als Zeichenkette. Nachnamen, z.B. mit Adelstiteln bzw. ausländische Nachnamen werden als ein Name übermittelt und nicht in verschiedene Bestandteile aufgeteilt.
type AllgemeinerNameType struct {
	Name string `xml:"Name,omitempty"`
}

// AnlageDokumentType is Die fortlaufende Nummer eines Anlagendokuments zu einem Dokument.
type AnlageDokumentType struct {
	Nummer string `xml:"Nummer,omitempty"`
	*DokumentType
}

// AnschriftType is Im Typ wird beschrieben, um welche Art der Anschrift es sich handelt. Mögliche Werte sind "Aktuelle Anschrift", "Hauptsitz" oder "Zweitsitz".
type AnschriftType struct {
	Staat        *StaatType             `xml:"Staat,omitempty"`
	Strasse      string                 `xml:"Strasse,omitempty"`
	Hausnummer   string                 `xml:"Hausnummer,omitempty"`
	Postfach     string                 `xml:"Postfach,omitempty"`
	Postleitzahl string                 `xml:"Postleitzahl,omitempty"`
	Ort          string                 `xml:"Ort,omitempty"`
	Zusatz       string                 `xml:"Zusatz,omitempty"`
	Typ          *AnschriftstypCodeType `xml:"Typ,omitempty"`
}

// AnwendungsspezifischeErweiterungType is Ein Feld, das der anwendungsspezifischen Erweiterung zugeordnet ist.
type AnwendungsspezifischeErweiterungType struct {
	Kennung        string            `xml:"Kennung"`
	Name           string            `xml:"Name"`
	Beschreibung   string            `xml:"Beschreibung,omitempty"`
	Versionsnummer string            `xml:"Versionsnummer,omitempty"`
	Versionsdatum  string            `xml:"Versionsdatum,omitempty"`
	Feldgruppe     []*FeldgruppeType `xml:"Feldgruppe,omitempty"`
	Feld           []*FeldType       `xml:"Feld,omitempty"`
}

// AnwendungsspezifischeErweiterungXMLType is AnwendungsspezifischeErweiterungXML darf nur dazu genutzt werden, weitere (z.B. fachspezifische) Metadaten zu spezifizieren, deren Übermittlung mit den bereits in xdomea spezifizierten Metadaten nicht möglich ist. Die AnwendungsspezifischeErweiterungXML bietet über ein xs:any-Element die Möglichkeit, mittels Einbindung externer XML-Schemata diese Metadaten zu spezifizieren. Es können beliebige XML-Schemata mit unterschiedlichen Namensräumen angegeben werden. Die XML-Schema-Validierung der weiterführenden Metadaten erfolgt innerhalb der xdomea-Nachricht selbst "lax".
type AnwendungsspezifischeErweiterungXMLType struct {
}

// AufbewahrungsdauerType is Eine Akte oder ein Vorgang ist nach der Verfügung zur Akte (zdA-Verfügung) innerhalb der aktenführenden Stelle unbefristet aufzubewahren.
type AufbewahrungsdauerType struct {
	AnzahlJahre uint16 `xml:"AnzahlJahre"`
	Unbefristet bool   `xml:"Unbefristet"`
}

// AussonderungsartType is Die Aussonderungsart als frei konfigurierbarer Wert.
type AussonderungsartType struct {
	Aussonderungsart               *AussonderungsartCodeType `xml:"Aussonderungsart"`
	AussonderungsartKonfigurierbar string                    `xml:"AussonderungsartKonfigurierbar"`
}

// BearbeitungType is Ein Dokument, das der Bearbeiter des Beteiligungsschrittes dem Beteiligungsschritt z.B. als Stellungnahme beifügt.
type BearbeitungType struct {
	Bearbeiter *KontaktType    `xml:"Bearbeiter"`
	Datum      string          `xml:"Datum"`
	Uhrzeit    time.Time       `xml:"Uhrzeit,omitempty"`
	Vermerk    string          `xml:"Vermerk,omitempty"`
	Notiz      string          `xml:"Notiz,omitempty"`
	Anlage     []*DokumentType `xml:"Anlage,omitempty"`
}

// BehoerdenkennungType is Der Präfix bezeichnet eine Klasse von Behördenkennungen. Beispiel: So werden u.a. alle Behördenkennungen der Behörden, die anhand des amtlichen Gemeindeschlüssels (AGS) identifiziert werden können, den Präfix ags: erhalten. Die Liste der Präfixe für Behördenkennungen werden im Zusammenhang mit dem DVDV durch das Bundesverwaltungsamt als koordinierende Stelle für das DVDV verwaltet.
type BehoerdenkennungType struct {
	Behoerdenschluessel string `xml:"Behoerdenschluessel,omitempty"`
	Praefix             string `xml:"Praefix,omitempty"`
}

// BeteiligungsschrittType is In der Bearbeitung sind die Informationen zum Bearbeitungsteil des Beteiligungsschrittes zusammengefasst.
type BeteiligungsschrittType struct {
	Nummer      string                      `xml:"Nummer"`
	Status      *BeteiligungsstatusCodeType `xml:"Status"`
	Verfuegung  *VerfuegungType             `xml:"Verfuegung"`
	Bearbeitung *BearbeitungType            `xml:"Bearbeitung,omitempty"`
}

// DokumentType is AnwendungsspezifischeErweiterungXML darf nur dazu genutzt werden, weitere (z.B. fachspezifische) Metadaten zu spezifizieren, deren Übermittlung mit den bereits in xdomea spezifizierten Metadaten nicht möglich ist. Die AnwendungsspezifischeErweiterungXML bietet die Möglichkeit, mittels Einbindung externer XML-Schemata diese Metadaten zu spezifizieren. Es können beliebige XML-Schemata mit unterschiedlichen Namensräumen angegeben werden. Die XML-Schema-Validierung der weiterführenden Metadaten erfolgt innerhalb der xdomea-Nachricht selbst "lax".
type DokumentType struct {
	Identifikation                      *IdentifikationObjektType                `xml:"Identifikation"`
	AllgemeineMetadaten                 *AllgemeineMetadatenType                 `xml:"AllgemeineMetadaten,omitempty"`
	FremdesGeschaeftszeichen            string                                   `xml:"FremdesGeschaeftszeichen,omitempty"`
	Posteingangsdatum                   string                                   `xml:"Posteingangsdatum,omitempty"`
	Postausgangsdatum                   string                                   `xml:"Postausgangsdatum,omitempty"`
	DatumDesSchreibens                  string                                   `xml:"DatumDesSchreibens,omitempty"`
	Bezug                               string                                   `xml:"Bezug,omitempty"`
	Hier                                string                                   `xml:"Hier,omitempty"`
	Bearbeiter                          string                                   `xml:"Bearbeiter,omitempty"`
	Typ                                 string                                   `xml:"Typ,omitempty"`
	HistorienProtokollInformation       []*HistorienProtokollInformationType     `xml:"HistorienProtokollInformation,omitempty"`
	InternerGeschaeftsgang              *GeschaeftsgangType                      `xml:"InternerGeschaeftsgang,omitempty"`
	Version                             []*VersionType                           `xml:"Version,omitempty"`
	Verweis                             []*VerweisType                           `xml:"Verweis,omitempty"`
	Anlage                              []*AnlageDokumentType                    `xml:"Anlage,omitempty"`
	Absender                            []*KontaktType                           `xml:"Absender,omitempty"`
	Empfaenger                          []*KontaktType                           `xml:"Empfaenger,omitempty"`
	WeitererKontakt                     []*KontaktType                           `xml:"WeitererKontakt,omitempty"`
	AnwendungsspezifischeErweiterung    []*AnwendungsspezifischeErweiterungType  `xml:"AnwendungsspezifischeErweiterung,omitempty"`
	AnwendungsspezifischeErweiterungXML *AnwendungsspezifischeErweiterungXMLType `xml:"AnwendungsspezifischeErweiterungXML,omitempty"`
}

// FeldType is Der Wert des Feldes.
type FeldType struct {
	Name         string            `xml:"Name"`
	Beschreibung string            `xml:"Beschreibung,omitempty"`
	Datentyp     *DatentypCodeType `xml:"Datentyp,omitempty"`
	Wert         string            `xml:"Wert,omitempty"`
}

// FeldgruppeType is Über Feld erfolgt die Unterteilung einer Feldgruppe in konkrete Felder.
type FeldgruppeType struct {
	Name            string            `xml:"Name"`
	Beschreibung    string            `xml:"Beschreibung"`
	Unterfeldgruppe []*FeldgruppeType `xml:"Unterfeldgruppe,omitempty"`
	Feld            []*FeldType       `xml:"Feld,omitempty"`
}

// FormatType is Über Primaerdokument werden Dateiangaben zum tatsächlich beschriebenen Primärdokument eines Formats angegeben.
type FormatType struct {
	Name            *DateiformatCodeType `xml:"Name"`
	SonstigerName   string               `xml:"SonstigerName,omitempty"`
	Version         string               `xml:"Version"`
	Primaerdokument *PrimaerdokumentType `xml:"Primaerdokument"`
}

// GeburtType is Das Datum der Geburt.
type GeburtType struct {
	Datum string `xml:"Datum"`
}

// GeschaeftsgangType is Ein dem Geschäftsgang zugeordneter Beteiligungsschritt, der im Zuge des Geschäftsgangs durchgeführt wird. Die Beteiligungsschritte sind fortlaufend nummeriert.
type GeschaeftsgangType struct {
	Identifikation      *IdentifikationObjektType  `xml:"Identifikation"`
	Beteiligungsschritt []*BeteiligungsschrittType `xml:"Beteiligungsschritt"`
}

// HistorienProtokollInformationType is Die Aktion, die die konkrete Änderung des Metadatums näher beschreibt, z.B. gelöscht, neu.
type HistorienProtokollInformationType struct {
	MetadatumName      string `xml:"MetadatumName,omitempty"`
	MetadatumAlterWert string `xml:"MetadatumAlterWert,omitempty"`
	MetadatumNeuerWert string `xml:"MetadatumNeuerWert,omitempty"`
	Akteur             string `xml:"Akteur"`
	DatumUhrzeit       string `xml:"DatumUhrzeit"`
	Bemerkung          string `xml:"Bemerkung,omitempty"`
	Aktion             string `xml:"Aktion"`
}

// IdentifikationObjektType is Die laufende Nummer des Schriftgutobjekts im übergeordneten Objekt (z.B. die Heftungsnummer eines Dokuments in einem Vorgang, die Nummer eines Bandes in einer Akte).
type IdentifikationObjektType struct {
	ID                               string `xml:"ID"`
	NummerImUebergeordnetenContainer uint32 `xml:"NummerImUebergeordnetenContainer,omitempty"`
}

// KommunikationType is Mit IstInstitution kann angegeben werden, ob es sich um Kommunikationsdaten einer Institution handelt oder nicht. Handelt es sich um eine Institution, so ist der Wert 1 anzugeben. Handelt es sich um keine Institution, so ist der Wert 0 anzugeben.
type KommunikationType struct {
	IstDienstlich  bool                       `xml:"IstDienstlich,omitempty"`
	Kanal          *KommunikationsartCodeType `xml:"Kanal,omitempty"`
	Kennung        string                     `xml:"Kennung,omitempty"`
	Zusatz         string                     `xml:"Zusatz,omitempty"`
	IstInstitution bool                       `xml:"IstInstitution,omitempty"`
}

// KontaktType is Die Geburtsangaben des Ansprechpartners zum Kontakt.
type KontaktType struct {
	Behoerdenkennung         *BehoerdenkennungType           `xml:"Behoerdenkennung,omitempty"`
	Institution              *NameOrganisationType           `xml:"Institution,omitempty"`
	Organisationseinheit     *OrganisationseinheitType       `xml:"Organisationseinheit,omitempty"`
	Name                     *NameNatuerlichePersonType      `xml:"Name,omitempty"`
	Taetigkeit               string                          `xml:"Taetigkeit,omitempty"`
	Zustaendigkeit           string                          `xml:"Zustaendigkeit,omitempty"`
	Anschrift                []*AnschriftType                `xml:"Anschrift,omitempty"`
	Kommunikation            []*KommunikationType            `xml:"Kommunikation,omitempty"`
	Rolle                    string                          `xml:"Rolle,omitempty"`
	UnstrukturierteAnschrift []*UnstrukturierteAnschriftType `xml:"UnstrukturierteAnschrift,omitempty"`
	Geburt                   *GeburtType                     `xml:"Geburt,omitempty"`
}

// MetadatenAussonderungType is Das Aufbewahrungsende gibt taggenau das Ende der Aufbewahrungsfrist an.
type MetadatenAussonderungType struct {
	Aufbewahrungsdauer  *AufbewahrungsdauerType      `xml:"Aufbewahrungsdauer,omitempty"`
	Aussonderungsart    *AussonderungsartType        `xml:"Aussonderungsart,omitempty"`
	Kennung             string                       `xml:"Kennung,omitempty"`
	Bewertungsvorschlag *BewertungsvorschlagCodeType `xml:"Bewertungsvorschlag,omitempty"`
	Aufbewahrungsende   string                       `xml:"Aufbewahrungsende,omitempty"`
}

// NameNatuerlichePersonType is Der Vorname ist der Name bzw. der Teil des Namens, der nicht die Zugehörigkeit zu einer Familie ausdrückt, sondern das Individuum innerhalb der Familie bezeichnet und dazu dient, es von anderen Familienmitgliedern zu unterscheiden.
type NameNatuerlichePersonType struct {
	Anrede       string               `xml:"Anrede,omitempty"`
	Titel        string               `xml:"Titel,omitempty"`
	Familienname *AllgemeinerNameType `xml:"Familienname,omitempty"`
	Vorname      *AllgemeinerNameType `xml:"Vorname,omitempty"`
}

// NameOrganisationType is Die Kurzbezeichnung des Namens einer Organisation.
type NameOrganisationType struct {
	Name            string `xml:"Name,omitempty"`
	Kurzbezeichnung string `xml:"Kurzbezeichnung,omitempty"`
}

// NkBasisType is AnwendungsspezifischeErweiterungXML darf nur dazu genutzt werden, weitere (z.B. fachspezifische) Metadaten zu spezifizieren, deren Übermittlung mit den bereits in xdomea spezifizierten Metadaten nicht möglich ist. Die AnwendungsspezifischeErweiterungXML bietet die Möglichkeit, mittels Einbindung externer XML-Schemata diese Metadaten zu spezifizieren. Es können beliebige XML-Schemata mit unterschiedlichen Namensräumen angegeben werden. Die XML-Schema-Validierung der weiterführenden Metadaten erfolgt innerhalb der xdomea-Nachricht selbst "lax".
type NkBasisType struct {
	ProzessID                           string                                   `xml:"ProzessID"`
	Nachrichtentyp                      *NachrichtentypCodeType                  `xml:"Nachrichtentyp"`
	Erstellungszeitpunkt                string                                   `xml:"Erstellungszeitpunkt"`
	Absender                            *KontaktType                             `xml:"Absender"`
	Empfaenger                          *KontaktType                             `xml:"Empfaenger"`
	SendendesSystem                     *SystemType                              `xml:"SendendesSystem,omitempty"`
	Hinweis                             string                                   `xml:"Hinweis,omitempty"`
	AnwendungsspezifischeErweiterung    *AnwendungsspezifischeErweiterungType    `xml:"AnwendungsspezifischeErweiterung,omitempty"`
	AnwendungsspezifischeErweiterungXML *AnwendungsspezifischeErweiterungXMLType `xml:"AnwendungsspezifischeErweiterungXML,omitempty"`
}

// NkFVDatenType is Die Informationen zu dem System, das die Daten erhalten soll.
type NkFVDatenType struct {
	EmpfangendesSystem []*SystemType `xml:"EmpfangendesSystem,omitempty"`
	*NkBasisType
}

// NkNichtFVDatenEinEmpfaengerMitEmpfangsbestaetigungType is Die Angabe, ob vom Empfänger der Nachricht eine Empfangsbestätigung erwünscht wird. Eine erwünschte Empfangsbestätigung wird mit 1 gekennzeichnet. Wird keine Empfangsbestätigung gewünscht, so wird dies mit 0 gekennzeichnet.
type NkNichtFVDatenEinEmpfaengerMitEmpfangsbestaetigungType struct {
	Empfangsbestaetigung bool `xml:"Empfangsbestaetigung"`
	*NkBasisType
}

// NkNichtFVDatenWeitereEmpfaengerMitEmpfangsbestaetigungType is Die Angabe, ob vom Empfänger der Nachricht eine Empfangsbestätigung erwünscht wird. Eine erwünschte Empfangsbestätigung wird mit 1 gekennzeichnet. Wird keine Empfangsbestätigung gewünscht, so wird dies mit 0 gekennzeichnet.
type NkNichtFVDatenWeitereEmpfaengerMitEmpfangsbestaetigungType struct {
	Empfangsbestaetigung bool `xml:"Empfangsbestaetigung"`
	*NkNichtFVDatenWeitereEmpfaengerType
}

// NkNichtFVDatenWeitereEmpfaengerType is Ein weiterer Empfänger der Nachricht.
type NkNichtFVDatenWeitereEmpfaengerType struct {
	WeitererEmpfaenger []*KontaktType `xml:"WeitererEmpfaenger,omitempty"`
	*NkBasisType
}

// OrganisationseinheitType is Die Bezeichnung der Organisationseinheit.
type OrganisationseinheitType struct {
	Name string `xml:"Name"`
}

// PrimaerdokumentType is Angaben zur elektronischen Signatur oder zum elektronischen Siegel, die zum Dokument gehören.
type PrimaerdokumentType struct {
	Dateiname         string              `xml:"Dateiname"`
	DateinameOriginal string              `xml:"DateinameOriginal,omitempty"`
	Ersteller         string              `xml:"Ersteller,omitempty"`
	DatumUhrzeit      string              `xml:"DatumUhrzeit,omitempty"`
	SignaturSiegel    *SignaturSiegelType `xml:"SignaturSiegel,omitempty"`
}

// SignaturSiegelType is Die elektronische Signatur oder das elektronische Siegel zum Dokument liegt in einer separaten Datei vor.
type SignaturSiegelType struct {
	SignaturSiegelEingebettet bool     `xml:"SignaturSiegelEingebettet"`
	SignaturSiegelDatei       []string `xml:"SignaturSiegelDatei"`
}

// StaatType is Staat enthält einen Schlüssel zur Identifikation eines Staates.
type StaatType struct {
	Staat *StaatCodeType `xml:"Staat"`
}

// SystemType is Die Version des Produktes.
type SystemType struct {
	InstanzID   string `xml:"InstanzID,omitempty"`
	Produktname string `xml:"Produktname,omitempty"`
	Version     string `xml:"Version,omitempty"`
}

// UnstrukturierteAnschriftType is Ein Anschriftenzusatz beinhaltet ggf. erforderliche weitere Präzisierungen zu einer Anschrift.
type UnstrukturierteAnschriftType struct {
	Typ    *AnschriftstypCodeType `xml:"Typ,omitempty"`
	Zeile1 string                 `xml:"Zeile1,omitempty"`
	Zeile2 string                 `xml:"Zeile2,omitempty"`
	Zeile3 string                 `xml:"Zeile3,omitempty"`
	Zeile4 string                 `xml:"Zeile4,omitempty"`
	Zeile5 string                 `xml:"Zeile5,omitempty"`
	Zeile6 string                 `xml:"Zeile6,omitempty"`
	Zusatz string                 `xml:"Zusatz,omitempty"`
}

// VerfuegungType is Die Hinweise oder Erläuterungen zu einer Verfügung.
type VerfuegungType struct {
	Ersteller          *KontaktType `xml:"Ersteller"`
	Adressat           *KontaktType `xml:"Adressat"`
	Erstellungsdatum   string       `xml:"Erstellungsdatum"`
	Erstellungsuhrzeit time.Time    `xml:"Erstellungsuhrzeit,omitempty"`
	Verfuegung         string       `xml:"Verfuegung,omitempty"`
	TerminDatum        string       `xml:"TerminDatum,omitempty"`
	TerminUhrzeit      time.Time    `xml:"TerminUhrzeit,omitempty"`
	Notiz              string       `xml:"Notiz,omitempty"`
}

// VersionType is Ein Format, das der Dokumentversion zugeordnet ist.
type VersionType struct {
	Nummer string        `xml:"Nummer"`
	Format []*FormatType `xml:"Format"`
}

// VerweisType is Die Hinweise und Bemerkungen zu diesem Verweis.
type VerweisType struct {
	ID             string                       `xml:"ID,omitempty"`
	SGOTyp         *SchriftgutobjekttypCodeType `xml:"SGOTyp"`
	SGOKennzeichen string                       `xml:"SGOKennzeichen,omitempty"`
	Bemerkung      string                       `xml:"Bemerkung,omitempty"`
}

// VorgangType is AnwendungsspezifischeErweiterungXML darf nur dazu genutzt werden, weitere (z.B. fachspezifische) Metadaten zu spezifizieren, deren Übermittlung mit den bereits in xdomea spezifizierten Metadaten nicht möglich ist. Die AnwendungsspezifischeErweiterungXML bietet die Möglichkeit, mittels Einbindung externer XML-Schemata diese Metadaten zu spezifizieren. Es können beliebige XML-Schemata mit unterschiedlichen Namensräumen angegeben werden. Die XML-Schema-Validierung der weiterführenden Metadaten erfolgt innerhalb der xdomea-Nachricht selbst "lax".
type VorgangType struct {
	Identifikation                      *IdentifikationObjektType                `xml:"Identifikation"`
	AllgemeineMetadaten                 *AllgemeineMetadatenType                 `xml:"AllgemeineMetadaten,omitempty"`
	ArchivspezifischeMetadaten          *MetadatenAussonderungType               `xml:"ArchivspezifischeMetadaten,omitempty"`
	Aktenbetreff                        string                                   `xml:"Aktenbetreff,omitempty"`
	Typ                                 string                                   `xml:"Typ,omitempty"`
	ZdA                                 bool                                     `xml:"ZdA,omitempty"`
	ZdADatum                            string                                   `xml:"ZdADatum,omitempty"`
	Laufzeit                            *ZeitraumType                            `xml:"Laufzeit,omitempty"`
	HistorienProtokollInformation       []*HistorienProtokollInformationType     `xml:"HistorienProtokollInformation,omitempty"`
	InternerGeschaeftsgang              *GeschaeftsgangType                      `xml:"InternerGeschaeftsgang,omitempty"`
	Dokument                            []*DokumentType                          `xml:"Dokument,omitempty"`
	Verweis                             []*VerweisType                           `xml:"Verweis,omitempty"`
	Kontakt                             []*KontaktType                           `xml:"Kontakt,omitempty"`
	Teilvorgang                         []*VorgangType                           `xml:"Teilvorgang,omitempty"`
	AnwendungsspezifischeErweiterung    *AnwendungsspezifischeErweiterungType    `xml:"AnwendungsspezifischeErweiterung,omitempty"`
	AnwendungsspezifischeErweiterungXML *AnwendungsspezifischeErweiterungXMLType `xml:"AnwendungsspezifischeErweiterungXML,omitempty"`
}

// ZeitraumType is Das Ende eines Zeitraumes beschreibt den Zeitpunkt, ab dem ein Sachverhalt endet bzw. nicht mehr rechtskräftig ist. Das Ende ist selbst Teil der Dauer des Zeitraumes. Beispiele sind: Fristdatum (Bau) Ablaufdatum (Finanz) Faelligkeitsdatum (Finanz) Wirksamkeitsdatum der Aufhebung/Scheidung der Ehe (Personenstand)
type ZeitraumType struct {
	Beginn string `xml:"Beginn,omitempty"`
	Ende   string `xml:"Ende,omitempty"`
}
