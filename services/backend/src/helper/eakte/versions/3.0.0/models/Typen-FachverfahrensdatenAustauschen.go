package models

import (
	"time"
)

// BearbeitungFVType is Die Anmerkung, die Aufschluss über die durchgeführte(n) Aktion(en) während einer Bearbeitung gibt.
type BearbeitungFVType struct {
	Bearbeiter string    `xml:"Bearbeiter"`
	Datum      string    `xml:"Datum"`
	Uhrzeit    time.Time `xml:"Uhrzeit,omitempty"`
	Notiz      string    `xml:"Notiz"`
}

// DatensatzZumLoeschenType is Die ID des Datensatzes im Fachverfahren.
type DatensatzZumLoeschenType struct {
	NummerImStapel           uint32 `xml:"NummerImStapel,omitempty"`
	FachverfahrenDatensatzID string `xml:"FachverfahrenDatensatzID"`
}

// DokumentZumAktualisierenType is Das Dokument, das aktualisiert werden soll.
type DokumentZumAktualisierenType struct {
	NummerImStapel       uint32        `xml:"NummerImStapel,omitempty"`
	IDDokument           string        `xml:"IDDokument"`
	Ablageort            string        `xml:"Ablageort"`
	NeueVersionErstellen bool          `xml:"NeueVersionErstellen,omitempty"`
	Dokument             *DokumentType `xml:"Dokument"`
}

// DokumentZumExportierenType is Die Versionsnummer des Dokuments.
type DokumentZumExportierenType struct {
	IDSGO          string `xml:"IDSGO"`
	Versionsnummer string `xml:"Versionsnummer,omitempty"`
}

// Rueckmeldung is Eine sonstige, nicht näher in xdomea definierte Fehlermeldung, die als Code angegeben werden kann.
type Rueckmeldung struct {
	AllgemeineRueckmeldung       *AllgemeineRueckmeldungCodeType       `xml:"AllgemeineRueckmeldung"`
	SpezifischeFehlermeldung0601 *SpezifischeFehlermeldung0601CodeType `xml:"SpezifischeFehlermeldung0601"`
	SpezifischeFehlermeldung0602 *SpezifischeFehlermeldung0602CodeType `xml:"SpezifischeFehlermeldung0602"`
	SpezifischeFehlermeldung0603 *SpezifischeFehlermeldung0603CodeType `xml:"SpezifischeFehlermeldung0603"`
	SpezifischeFehlermeldung0604 *SpezifischeFehlermeldung0604CodeType `xml:"SpezifischeFehlermeldung0604"`
	SpezifischeFehlermeldung0605 *SpezifischeFehlermeldung0605CodeType `xml:"SpezifischeFehlermeldung0605"`
	SpezifischeFehlermeldung0606 *SpezifischeFehlermeldung0606CodeType `xml:"SpezifischeFehlermeldung0606"`
	SpezifischeFehlermeldung0607 *SpezifischeFehlermeldung0607CodeType `xml:"SpezifischeFehlermeldung0607"`
	SpezifischeFehlermeldung0608 *SpezifischeFehlermeldung0608CodeType `xml:"SpezifischeFehlermeldung0608"`
	SpezifischeFehlermeldung0609 *SpezifischeFehlermeldung0609CodeType `xml:"SpezifischeFehlermeldung0609"`
	SpezifischeFehlermeldung0610 *SpezifischeFehlermeldung0610CodeType `xml:"SpezifischeFehlermeldung0610"`
	SpezifischeFehlermeldung0611 *SpezifischeFehlermeldung0611CodeType `xml:"SpezifischeFehlermeldung0611"`
	SpezifischeFehlermeldung0612 *SpezifischeFehlermeldung0612CodeType `xml:"SpezifischeFehlermeldung0612"`
	SpezifischeFehlermeldung0613 *SpezifischeFehlermeldung0613CodeType `xml:"SpezifischeFehlermeldung0613"`
	SpezifischeFehlermeldung0614 *SpezifischeFehlermeldung0614CodeType `xml:"SpezifischeFehlermeldung0614"`
	SpezifischeFehlermeldung0615 *SpezifischeFehlermeldung0615CodeType `xml:"SpezifischeFehlermeldung0615"`
	SpezifischeFehlermeldung0616 *SpezifischeFehlermeldung0616CodeType `xml:"SpezifischeFehlermeldung0616"`
	SpezifischeFehlermeldung0617 *SpezifischeFehlermeldung0617CodeType `xml:"SpezifischeFehlermeldung0617"`
	SpezifischeFehlermeldung0618 *SpezifischeFehlermeldung0618CodeType `xml:"SpezifischeFehlermeldung0618"`
	SpezifischeFehlermeldung0619 *SpezifischeFehlermeldung0619CodeType `xml:"SpezifischeFehlermeldung0619"`
	SpezifischeFehlermeldung0620 *SpezifischeFehlermeldung0620CodeType `xml:"SpezifischeFehlermeldung0620"`
	SpezifischeFehlermeldung0621 *SpezifischeFehlermeldung0621CodeType `xml:"SpezifischeFehlermeldung0621"`
	SpezifischeFehlermeldung0622 *SpezifischeFehlermeldung0622CodeType `xml:"SpezifischeFehlermeldung0622"`
	SpezifischeFehlermeldung0623 *SpezifischeFehlermeldung0623CodeType `xml:"SpezifischeFehlermeldung0623"`
	SpezifischeFehlermeldung0624 *SpezifischeFehlermeldung0624CodeType `xml:"SpezifischeFehlermeldung0624"`
	SpezifischeFehlermeldung0625 *SpezifischeFehlermeldung0625CodeType `xml:"SpezifischeFehlermeldung0625"`
	SpezifischeFehlermeldung0626 *SpezifischeFehlermeldung0626CodeType `xml:"SpezifischeFehlermeldung0626"`
	SpezifischeFehlermeldung0627 *SpezifischeFehlermeldung0627CodeType `xml:"SpezifischeFehlermeldung0627"`
	SpezifischeFehlermeldung0628 *SpezifischeFehlermeldung0628CodeType `xml:"SpezifischeFehlermeldung0628"`
	SpezifischeFehlermeldung0631 *SpezifischeFehlermeldung0631CodeType `xml:"SpezifischeFehlermeldung0631"`
	SpezifischeFehlermeldung0632 *SpezifischeFehlermeldung0632CodeType `xml:"SpezifischeFehlermeldung0632"`
	SonstigeFehlermeldung        string                                `xml:"SonstigeFehlermeldung"`
	SonstigeFehlermeldungCode    *SonstigeFehlermeldungCodeType        `xml:"SonstigeFehlermeldungCode"`
}

// ErfolgOderMisserfolgDMSType is Die Information zum erfolgreichen oder nicht erfolgreichen Import von Daten.
type ErfolgOderMisserfolgDMSType struct {
	Rueckmeldung       []*Rueckmeldung                  `xml:"Rueckmeldung"`
	Rueckgabeparameter []*RueckgabeparameterFVDatenType `xml:"Rueckgabeparameter,omitempty"`
}

// GesamtprotokollZumAblegenType is Die Bearbeitungs- und Protokollinformationen eines Schriftgutobjekts.
type GesamtprotokollZumAblegenType struct {
	NummerImStapel                uint32                               `xml:"NummerImStapel,omitempty"`
	IDSGO                         string                               `xml:"IDSGO"`
	Ablageort                     string                               `xml:"Ablageort"`
	Dokument                      *DokumentType                        `xml:"Dokument,omitempty"`
	Geschaeftsgang                *GeschaeftsgangType                  `xml:"Geschaeftsgang,omitempty"`
	HistorienProtokollInformation []*HistorienProtokollInformationType `xml:"HistorienProtokollInformation,omitempty"`
}

// Wert is Es handelt sich bei dem Wert des Konfigurationsparameters um Angaben zur Standardablage.
type Wert struct {
	WertOffen                   string                         `xml:"WertOffen"`
	Transportweg                *TransportwegCodeType          `xml:"Transportweg"`
	Datenaustauschart           *DatenaustauschartCodeType     `xml:"Datenaustauschart"`
	KomprimierterDatenaustausch bool                           `xml:"KomprimierterDatenaustausch"`
	Kompressionsverfahren       *KompressionsverfahrenCodeType `xml:"Kompressionsverfahren"`
	Standardablage              *StandardablageType            `xml:"Standardablage"`
}

// KonfigurationsparameterType is Der Datentyp des Konfigurationsparameters.
type KonfigurationsparameterType struct {
	Parameter *KonfigurationsparameterCodeType       `xml:"Parameter"`
	Datentyp  *DatentypCodeType                      `xml:"Datentyp,omitempty"`
	Wert      *Wert                                  `xml:"Wert,omitempty"`
	Gruppe    *KonfigurationsparameterGruppeCodeType `xml:"Gruppe"`
}

// KonfigurationsparameterZumAbrufenType is Der Name des Feldes.
type KonfigurationsparameterZumAbrufenType struct {
	Name string `xml:"Name"`
}

// KonfigurationsparameterZumErstellenType is Die Informationen zu dem (den) zu erstellenden Konfigurationsparameter(n).
type KonfigurationsparameterZumErstellenType struct {
	BenutzerOderRolle                  string                         `xml:"BenutzerOderRolle,omitempty"`
	DefinierterKonfigurationsparameter []*KonfigurationsparameterType `xml:"DefinierterKonfigurationsparameter,omitempty"`
	Konfigurationsparameter            []*FeldType                    `xml:"Konfigurationsparameter,omitempty"`
}

// MetadatenZumBearbeitenType is Die Informationen zu einem zu bearbeitenden Metadatum, das nicht in xdomea vordefiniert ist.
type MetadatenZumBearbeitenType struct {
	NummerImStapel uint32      `xml:"NummerImStapel,omitempty"`
	IDSGO          string      `xml:"IDSGO"`
	Metadatum      []*FeldType `xml:"Metadatum"`
}

// NkKeineRueckmeldungFVDatenType is Die DMS-Session-ID wird verwendet, um den Benutzer bzw. die Rolle gegenüber dem DMS zu authentifizieren. Sie wird bei der Anmeldung im DMS erzeugt und wird nach einer konfigurierbaren Zeit der Inaktivität ungültig.
type NkKeineRueckmeldungFVDatenType struct {
	WeiteresEmpfangendesSystem []*SystemType `xml:"WeiteresEmpfangendesSystem,omitempty"`
	Empfangsbestaetigung       bool          `xml:"Empfangsbestaetigung"`
	Importbestaetigung         bool          `xml:"Importbestaetigung"`
	Erstellerkennung           string        `xml:"Erstellerkennung"`
	ErstellerRolle             string        `xml:"ErstellerRolle,omitempty"`
	Stapel                     bool          `xml:"Stapel"`
	Stapellaenge               uint32        `xml:"Stapellaenge,omitempty"`
	DMSSessionID               string        `xml:"DMSSessionID,omitempty"`
	*NkFVDatenType
}

// NkRueckmeldungFVDatenImportType is Die FV-Session-ID wird verwendet, um den Benutzer bzw. die Rolle aus dem DMS gegenüber dem Fachverfahren zu authentifizieren. Sie wird bei der Anmeldung im Fachverfahren erzeugt und wird nach einer konfigurierbaren Zeit der Inaktivität ungültig.
type NkRueckmeldungFVDatenImportType struct {
	FVSessionID string `xml:"FVSessionID"`
	*NkFVDatenType
}

// ProtokolleintragZumErstellenType is Der Protokolleintrag zu dem Schriftgutobjekt.
type ProtokolleintragZumErstellenType struct {
	NummerImStapel   uint32             `xml:"NummerImStapel,omitempty"`
	IDSGO            string             `xml:"IDSGO"`
	Protokolleintrag *BearbeitungFVType `xml:"Protokolleintrag"`
}

// Parameter is Der Ablageort des Schriftgutobjekts
type Parameter struct {
	AllgemeinerRueckgabeparameter *FeldType                                  `xml:"AllgemeinerRueckgabeparameter"`
	ZumLoeschenMarkiert           bool                                       `xml:"ZumLoeschenMarkiert"`
	Vertretungsstatus             []*RueckgabeparameterVertretungsstatusType `xml:"Vertretungsstatus"`
	Ablageort                     string                                     `xml:"Ablageort"`
}

// RueckgabeparameterFVDatenType is Die ID des Schriftgutobjekts im DMS, zu dem die Rückmeldung erfolgt.
type RueckgabeparameterFVDatenType struct {
	NummerImStapel uint32          `xml:"NummerImStapel,omitempty"`
	IDSGO          string          `xml:"IDSGO,omitempty"`
	Parameter      *Parameter      `xml:"Parameter"`
	Erfolgreich    bool            `xml:"Erfolgreich"`
	Rueckmeldung   []*Rueckmeldung `xml:"Rueckmeldung,omitempty"`
}

// RueckgabeparameterVertretungsstatusType is Der Status zur Vertretung, der abgefragt wird. Für aktive Vertretungen ist der Wert 1 anzugeben, für inaktive der Wert 0.
type RueckgabeparameterVertretungsstatusType struct {
	BenutzerOderRolle string `xml:"BenutzerOderRolle"`
	VertretungAktiv   bool   `xml:"VertretungAktiv"`
}

// SchriftgutobjektZumAblegenOderErstellenType is Die ID des übergeordneten Schriftgutobjekts, in dem das Schriftgutobjekt erstellt werden soll.
type SchriftgutobjektZumAblegenOderErstellenType struct {
	NummerImStapel        uint32            `xml:"NummerImStapel,omitempty"`
	FachverfahrenObjektID string            `xml:"FachverfahrenObjektID"`
	IDContainer           string            `xml:"IDContainer,omitempty"`
	Schriftgutobjekt      *Schriftgutobjekt `xml:"Schriftgutobjekt"`
}

// SchriftgutobjektZumAblegenType is Der Ablageort des Schriftgutobjekts, das im DMS abgelegt werden soll. Die Primärdokumente zu Dokumenten finden sich in Dateien, die durch Primaerdokument.Dateiname spezifiziert werden. In Abhängigkeit des Konfigurationsparameters Datenaustauschart kann dieser Parameter unterschiedliche Werte annehmen: Übergabe über ein freigegebenes Verzeichnis: Name des Verzeichnisses. Übergabe per WebDAV und http(s): URL. Übergabe über eine Datenbank: der Parameter bleibt leer (bzw. wird ignoriert).
type SchriftgutobjektZumAblegenType struct {
	Ablageort string `xml:"Ablageort,omitempty"`
	*SchriftgutobjektZumAblegenOderErstellenType
}

// SchriftgutobjektZumAnsehenOderBearbeitenType is Der Client, auf dem das Schriftgutobjekt zu öffnen ist.
type SchriftgutobjektZumAnsehenOderBearbeitenType struct {
	IDSGO          string `xml:"IDSGO"`
	Versionsnummer string `xml:"Versionsnummer,omitempty"`
	Client         string `xml:"Client,omitempty"`
}

// SchriftgutobjektZumDruckenType is Der Client, von dem aus das Schriftgutobjekt gedruckt wird.
type SchriftgutobjektZumDruckenType struct {
	NummerImStapel         uint32   `xml:"NummerImStapel,omitempty"`
	IDSGO                  string   `xml:"IDSGO"`
	Drucker                string   `xml:"Drucker"`
	MetadatenDrucken       bool     `xml:"MetadatenDrucken"`
	PrimaerdokumentDrucken bool     `xml:"PrimaerdokumentDrucken"`
	UnterSGODrucken        bool     `xml:"UnterSGODrucken"`
	UnterSGOTyp            []string `xml:"UnterSGOTyp,omitempty"`
	Versionsnummer         string   `xml:"Versionsnummer,omitempty"`
	Client                 string   `xml:"Client,omitempty"`
}

// SchriftgutobjektZumErstellenType is Der Client, auf dem das Schriftgutobjekt zu öffnen ist.
type SchriftgutobjektZumErstellenType struct {
	SGOOeffnen bool   `xml:"SGOOeffnen,omitempty"`
	Client     string `xml:"Client,omitempty"`
	*SchriftgutobjektZumAblegenOderErstellenType
}

// SchriftgutobjektZumLoeschenType is Die ID des Schriftgutobjekts im DMS.
type SchriftgutobjektZumLoeschenType struct {
	NummerImStapel uint32 `xml:"NummerImStapel,omitempty"`
	IDSGO          string `xml:"IDSGO"`
}

// SchriftgutobjektZumLoeschstatusAbfragenType is Die ID des Schriftgutobjekts im DMS, dessen Löschtstatus abgefragt werden soll.
type SchriftgutobjektZumLoeschstatusAbfragenType struct {
	NummerImStapel uint32 `xml:"NummerImStapel,omitempty"`
	IDSGO          string `xml:"IDSGO"`
}

// SchriftgutobjektZumSuchenType is Metadaten legt fest, ob die Suche auch in den Metadaten der Schriftgutobjekte durchgeführt werden soll. Soll eine Suche auch in den Metadaten der Schriftgutobjekte stattfinden, so ist der Wert 1 anzugeben. Soll keine Suche in den Metadaten der Schriftgutobjekte stattfinden, so ist der Wert 0 anzugeben.
type SchriftgutobjektZumSuchenType struct {
	Suchbegriffe string `xml:"Suchbegriffe"`
	Volltext     bool   `xml:"Volltext"`
	Metadaten    bool   `xml:"Metadaten"`
}

// SchriftgutobjektZumZDAAufhebenType is Die ID des Schriftgutobjektes, dessen zdA-Verfügung aufgehoben werden soll.
type SchriftgutobjektZumZDAAufhebenType struct {
	NummerImStapel uint32 `xml:"NummerImStapel,omitempty"`
	IDSGO          string `xml:"IDSGO"`
}

// SchriftgutobjektZumZDAVerfuegenType is Die ID des Schriftgutobjekts im DMS, das zdA-verfügt werden soll.
type SchriftgutobjektZumZDAVerfuegenType struct {
	NummerImStapel uint32 `xml:"NummerImStapel,omitempty"`
	IDSGO          string `xml:"IDSGO"`
}

// StandardablageType is Der Benutzer oder die Rolle, dem/der der Ablageort zugewiesen wird.
type StandardablageType struct {
	Ablageort         string `xml:"Ablageort"`
	BenutzerOderRolle string `xml:"BenutzerOderRolle,omitempty"`
}

// SystemstatusZumAbfragenType is Die ID der Systemstatus-Information, die abgefragt werden soll.
type SystemstatusZumAbfragenType struct {
	SystemstatusID int `xml:"SystemstatusID"`
}

// VertretungZumAktivierenOderDeaktivierenType is Aktivieren bestimmt, ob eine Vertretung aktiviert oder deaktiviert werden soll. Hat Aktivieren den Wert 1, so wird die Vertretung aktiviert. Hat Aktivieren den Wert 0, so wird die Vertretung deaktiviert.
type VertretungZumAktivierenOderDeaktivierenType struct {
	ZuVertretender string `xml:"ZuVertretender"`
	Vertreter      string `xml:"Vertreter"`
	Aktivieren     bool   `xml:"Aktivieren"`
}

// VertretungsstatusZumAbfragenType is Der Benutzer oder die Rolle, für den/die der Vertretungsstatus abgefragt wird.
type VertretungsstatusZumAbfragenType struct {
	BenutzerOderRolle string `xml:"BenutzerOderRolle"`
}

// ZustaendigkeitZumAendernType is Ein im Rahmen der Zuständigkeitsänderung angepasstes Metadatum zu dem Schriftgutobjekt.
type ZustaendigkeitZumAendernType struct {
	NummerImStapel                        uint32      `xml:"NummerImStapel,omitempty"`
	IDSGO                                 string      `xml:"IDSGO"`
	JetztZustaendig                       string      `xml:"JetztZustaendig"`
	BisherZustaendig                      string      `xml:"BisherZustaendig,omitempty"`
	AufnehmendAktenplaneinheitKennzeichen string      `xml:"AufnehmendAktenplaneinheitKennzeichen,omitempty"`
	AufnehmendID                          string      `xml:"AufnehmendID,omitempty"`
	AbgebendAktenplaneinheitKennzeichen   string      `xml:"AbgebendAktenplaneinheitKennzeichen,omitempty"`
	AbgebendID                            string      `xml:"AbgebendID,omitempty"`
	Metadatum                             []*FeldType `xml:"Metadatum,omitempty"`
}
