package models

import (
	"encoding/xml"
)

// FVDatenDokumentAktualisieren0601 is Die Informationen zu dem Dokument, das im DMS aktualisiert werden soll.
type FVDatenDokumentAktualisieren0601 struct {
	XMLName                  xml.Name                        `xml:"FVDaten.DokumentAktualisieren.0601"`
	Kopf                     *NkKeineRueckmeldungFVDatenType `xml:"Kopf"`
	DokumentZumAktualisieren []*DokumentZumAktualisierenType `xml:"DokumentZumAktualisieren"`
}

// FVDatenSGOAnsehen0602 is Die Informationen zu dem Schriftgutobjekt, das angesehen werden soll.
type FVDatenSGOAnsehen0602 struct {
	XMLName                    xml.Name                                      `xml:"FVDaten.SGOAnsehen.0602"`
	Kopf                       *NkKeineRueckmeldungFVDatenType               `xml:"Kopf"`
	SchriftgutobjektZumAnsehen *SchriftgutobjektZumAnsehenOderBearbeitenType `xml:"SchriftgutobjektZumAnsehen"`
}

// FVDatenSGOBearbeiten0603 is Die Informationen zu dem Schriftgutobjekt, das bearbeitet werden soll.
type FVDatenSGOBearbeiten0603 struct {
	XMLName                       xml.Name                                      `xml:"FVDaten.SGOBearbeiten.0603"`
	Kopf                          *NkKeineRueckmeldungFVDatenType               `xml:"Kopf"`
	SchriftgutobjektZumBearbeiten *SchriftgutobjektZumAnsehenOderBearbeitenType `xml:"SchriftgutobjektZumBearbeiten"`
}

// FVDatenSGOErstellen0604 is Die Informationen zu dem neuen Schriftgutobjekt, das erstellt werden soll.
type FVDatenSGOErstellen0604 struct {
	XMLName                      xml.Name                            `xml:"FVDaten.SGOErstellen.0604"`
	Kopf                         *NkKeineRueckmeldungFVDatenType     `xml:"Kopf"`
	SchriftgutobjektZumErstellen []*SchriftgutobjektZumErstellenType `xml:"SchriftgutobjektZumErstellen"`
}

// FVDatenSGOAblegen0605 is Die Informationen zu dem Schriftgutobjekt, das abgelegt werden soll.
type FVDatenSGOAblegen0605 struct {
	XMLName                    xml.Name                          `xml:"FVDaten.SGOAblegen.0605"`
	Kopf                       *NkKeineRueckmeldungFVDatenType   `xml:"Kopf"`
	SchriftgutobjektZumAblegen []*SchriftgutobjektZumAblegenType `xml:"SchriftgutobjektZumAblegen"`
}

// FVDatenSGODrucken0606 is Die Informationen zum zu druckenden Schriftgutobjekt.
type FVDatenSGODrucken0606 struct {
	XMLName                    xml.Name                          `xml:"FVDaten.SGODrucken.0606"`
	Kopf                       *NkKeineRueckmeldungFVDatenType   `xml:"Kopf"`
	SchriftgutobjektZumDrucken []*SchriftgutobjektZumDruckenType `xml:"SchriftgutobjektZumDrucken"`
}

// FVDatenProtokolleintragErstellen0607 is Die Informationen zu dem zu erstellenden Protokolleintrag.
type FVDatenProtokolleintragErstellen0607 struct {
	XMLName                      xml.Name                            `xml:"FVDaten.ProtokolleintragErstellen.0607"`
	Kopf                         *NkKeineRueckmeldungFVDatenType     `xml:"Kopf"`
	ProtokolleintragZumErstellen []*ProtokolleintragZumErstellenType `xml:"ProtokolleintragZumErstellen"`
}

// FVDatenSGOSuchen0608 is Die Parameter zu einer Suche nach Schriftgutobjekten.
type FVDatenSGOSuchen0608 struct {
	XMLName                   xml.Name                        `xml:"FVDaten.SGOSuchen.0608"`
	Kopf                      *NkKeineRueckmeldungFVDatenType `xml:"Kopf"`
	SchriftgutobjektZumSuchen *SchriftgutobjektZumSuchenType  `xml:"SchriftgutobjektZumSuchen"`
}

// FVDatenMetadatenAnlegen0609 is Die Informationen zu dem neuen Metadatum.
type FVDatenMetadatenAnlegen0609 struct {
	XMLName             xml.Name                        `xml:"FVDaten.MetadatenAnlegen.0609"`
	Kopf                *NkKeineRueckmeldungFVDatenType `xml:"Kopf"`
	MetadatenZumAnlegen []*MetadatenZumBearbeitenType   `xml:"MetadatenZumAnlegen"`
}

// FVDatenMetadatenAktualisieren0610 is Die Informationen zu dem zu aktualisierenden Metadatum.
type FVDatenMetadatenAktualisieren0610 struct {
	XMLName                   xml.Name                        `xml:"FVDaten.MetadatenAktualisieren.0610"`
	Kopf                      *NkKeineRueckmeldungFVDatenType `xml:"Kopf"`
	MetadatenZumAktualisieren []*MetadatenZumBearbeitenType   `xml:"MetadatenZumAktualisieren"`
}

// FVDatenMetadatenAnsehen0611 is Die Informationen zu dem Metadatum, das angesehen werden soll.
type FVDatenMetadatenAnsehen0611 struct {
	XMLName             xml.Name                        `xml:"FVDaten.MetadatenAnsehen.0611"`
	Kopf                *NkKeineRueckmeldungFVDatenType `xml:"Kopf"`
	MetadatenZumAnsehen []*MetadatenZumBearbeitenType   `xml:"MetadatenZumAnsehen"`
}

// FVDatenMetadatenLoeschen0612 is Die Informationen zu dem zu löschenden Metadatum.
type FVDatenMetadatenLoeschen0612 struct {
	XMLName              xml.Name                        `xml:"FVDaten.MetadatenLoeschen.0612"`
	Kopf                 *NkKeineRueckmeldungFVDatenType `xml:"Kopf"`
	MetadatenZumLoeschen []*MetadatenZumBearbeitenType   `xml:"MetadatenZumLoeschen"`
}

// FVDatenSGOLoeschenMarkieren0613 is Die Informationen zu dem Schriftgutobjekt, das zum Löschen markiert werden soll.
type FVDatenSGOLoeschenMarkieren0613 struct {
	XMLName                              xml.Name                           `xml:"FVDaten.SGOLoeschenMarkieren.0613"`
	Kopf                                 *NkKeineRueckmeldungFVDatenType    `xml:"Kopf"`
	SchriftgutobjektZumLoeschenMarkieren []*SchriftgutobjektZumLoeschenType `xml:"SchriftgutobjektZumLoeschenMarkieren"`
}

// FVDatenSGOLoeschmarkierungAufheben0614 is Die Informationen zu dem Schriftgutobjekt, dessen Löschmarkierung aufgehoben werden soll.
type FVDatenSGOLoeschmarkierungAufheben0614 struct {
	XMLName                                     xml.Name                           `xml:"FVDaten.SGOLoeschmarkierungAufheben.0614"`
	Kopf                                        *NkKeineRueckmeldungFVDatenType    `xml:"Kopf"`
	SchriftgutobjektZumLoeschmarkierungAufheben []*SchriftgutobjektZumLoeschenType `xml:"SchriftgutobjektZumLoeschmarkierungAufheben"`
}

// FVDatenSGOEndgueltigLoeschen0615 is Die Informationen zu dem Schriftgutobjekt, das endgültig gelöscht werden soll.
type FVDatenSGOEndgueltigLoeschen0615 struct {
	XMLName                               xml.Name                           `xml:"FVDaten.SGOEndgueltigLoeschen.0615"`
	Kopf                                  *NkKeineRueckmeldungFVDatenType    `xml:"Kopf"`
	SchriftgutobjektZumEndgueltigLoeschen []*SchriftgutobjektZumLoeschenType `xml:"SchriftgutobjektZumEndgueltigLoeschen"`
}

// FVDatenSGOLoeschstatusAbfragen0616 is Die Informationen zu dem Schriftgutobjekt, dessen Löschstatus abgefragt werden soll.
type FVDatenSGOLoeschstatusAbfragen0616 struct {
	XMLName                                 xml.Name                                       `xml:"FVDaten.SGOLoeschstatusAbfragen.0616"`
	Kopf                                    *NkKeineRueckmeldungFVDatenType                `xml:"Kopf"`
	SchriftgutobjektZumLoeschstatusAbfragen []*SchriftgutobjektZumLoeschstatusAbfragenType `xml:"SchriftgutobjektZumLoeschstatusAbfragen"`
}

// FVDatenSGOUngueltigKennzeichnen0617 is Die Informationen zu dem Schriftgutobjekt, das als ungültig gekennzeichnet werden soll.
type FVDatenSGOUngueltigKennzeichnen0617 struct {
	XMLName                                  xml.Name                           `xml:"FVDaten.SGOUngueltigKennzeichnen.0617"`
	Kopf                                     *NkKeineRueckmeldungFVDatenType    `xml:"Kopf"`
	SchriftgutobjektZumUngueltigKennzeichnen []*SchriftgutobjektZumLoeschenType `xml:"SchriftgutobjektZumUngueltigKennzeichnen"`
}

// FVDatenDatensatzLoeschen0618 is Die Information zu dem im Fachverfahren gelöschten Datensatz, zu dem auch alle vorhandenen Objekte im DMS gelöscht werden sollen.
type FVDatenDatensatzLoeschen0618 struct {
	XMLName              xml.Name                        `xml:"FVDaten.DatensatzLoeschen.0618"`
	Kopf                 *NkKeineRueckmeldungFVDatenType `xml:"Kopf"`
	DatensatzZumLoeschen []*DatensatzZumLoeschenType     `xml:"DatensatzZumLoeschen"`
}

// FVDatenBenachrichtigungAbrufen0619 is Der Kopf der Nachricht enthält allgemeine Informationen für den Empfänger der Nachricht.
type FVDatenBenachrichtigungAbrufen0619 struct {
	XMLName xml.Name                        `xml:"FVDaten.BenachrichtigungAbrufen.0619"`
	Kopf    *NkKeineRueckmeldungFVDatenType `xml:"Kopf"`
}

// FVDatenVertretungAktivierenOderDeaktivieren0620 is Die Informationen zur Aktivierung oder zur Deaktivierung einer Vertretung.
type FVDatenVertretungAktivierenOderDeaktivieren0620 struct {
	XMLName                                 xml.Name                                       `xml:"FVDaten.VertretungAktivierenOderDeaktivieren.0620"`
	Kopf                                    *NkKeineRueckmeldungFVDatenType                `xml:"Kopf"`
	VertretungZumAktivierenOderDeaktivieren []*VertretungZumAktivierenOderDeaktivierenType `xml:"VertretungZumAktivierenOderDeaktivieren"`
}

// FVDatenVertretungsstatusAbfragen0621 is Die Informationen zu dem Benutzer oder der Rolle, für den/die die Statusabfrage durchgeführt werden soll.
type FVDatenVertretungsstatusAbfragen0621 struct {
	XMLName                      xml.Name                            `xml:"FVDaten.VertretungsstatusAbfragen.0621"`
	Kopf                         *NkKeineRueckmeldungFVDatenType     `xml:"Kopf"`
	VertretungsstatusZumAbfragen []*VertretungsstatusZumAbfragenType `xml:"VertretungsstatusZumAbfragen"`
}

// FVDatenZustaendigkeitAendern0622 is Die Informationen zu der bisherigen und neuen Zuständigkeit des Schriftgutobjekts.
type FVDatenZustaendigkeitAendern0622 struct {
	XMLName                  xml.Name                        `xml:"FVDaten.ZustaendigkeitAendern.0622"`
	Kopf                     *NkKeineRueckmeldungFVDatenType `xml:"Kopf"`
	ZustaendigkeitZumAendern []*ZustaendigkeitZumAendernType `xml:"ZustaendigkeitZumAendern"`
}

// FVDatenGesamtprotokollAblegen0623 is Die Informationen zum Gesamtprotokoll, das an das DMS übermittelt werden soll.
type FVDatenGesamtprotokollAblegen0623 struct {
	XMLName                   xml.Name                         `xml:"FVDaten.GesamtprotokollAblegen.0623"`
	Kopf                      *NkKeineRueckmeldungFVDatenType  `xml:"Kopf"`
	GesamtprotokollZumAblegen []*GesamtprotokollZumAblegenType `xml:"GesamtprotokollZumAblegen"`
}

// FVDatenSGOZDAVerfuegen0624 is Die Informationen zu dem Schriftgutobjekt, das zdA-verfügt werden soll.
type FVDatenSGOZDAVerfuegen0624 struct {
	XMLName                         xml.Name                               `xml:"FVDaten.SGOZDAVerfuegen.0624"`
	Kopf                            *NkKeineRueckmeldungFVDatenType        `xml:"Kopf"`
	SchriftgutobjektZumZDAVerfuegen []*SchriftgutobjektZumZDAVerfuegenType `xml:"SchriftgutobjektZumZDAVerfuegen"`
}

// FVDatenSystemstatusAbfragen0625 is Die angeforderte Information zum Systemstatus.
type FVDatenSystemstatusAbfragen0625 struct {
	XMLName                 xml.Name                        `xml:"FVDaten.SystemstatusAbfragen.0625"`
	Kopf                    *NkKeineRueckmeldungFVDatenType `xml:"Kopf"`
	SystemstatusZumAbfragen []*SystemstatusZumAbfragenType  `xml:"SystemstatusZumAbfragen"`
}

// FVDatenKonfigurationsparameterErstellen0626 is Die Informationen zu dem neuen Konfigurationsparameter sowie zu dessen Ersteller.
type FVDatenKonfigurationsparameterErstellen0626 struct {
	XMLName                             xml.Name                                 `xml:"FVDaten.KonfigurationsparameterErstellen.0626"`
	Kopf                                *NkKeineRueckmeldungFVDatenType          `xml:"Kopf"`
	KonfigurationsparameterZumErstellen *KonfigurationsparameterZumErstellenType `xml:"KonfigurationsparameterZumErstellen"`
}

// FVDatenKonfigurationsparameterAktualisieren0627 is Die Information zum zu aktualisierenden Konfigurationsparameter.
type FVDatenKonfigurationsparameterAktualisieren0627 struct {
	XMLName                                 xml.Name                        `xml:"FVDaten.KonfigurationsparameterAktualisieren.0627"`
	Kopf                                    *NkKeineRueckmeldungFVDatenType `xml:"Kopf"`
	KonfigurationsparameterZumAktualisieren []*FeldType                     `xml:"KonfigurationsparameterZumAktualisieren"`
}

// FVDatenKonfigurationsparameterAbrufen0628 is Die Informationen zu dem abzurufenden Konfigurationsparameter.
type FVDatenKonfigurationsparameterAbrufen0628 struct {
	XMLName                           xml.Name                                 `xml:"FVDaten.KonfigurationsparameterAbrufen.0628"`
	Kopf                              *NkKeineRueckmeldungFVDatenType          `xml:"Kopf"`
	KonfigurationsparameterZumAbrufen []*KonfigurationsparameterZumAbrufenType `xml:"KonfigurationsparameterZumAbrufen"`
}

// FVDatenEmpfangBestaetigen0629 is Der Kopf der Nachricht enthält allgemeine Informationen für den Empfänger der Empfangsbestätigung.
type FVDatenEmpfangBestaetigen0629 struct {
	XMLName xml.Name       `xml:"FVDaten.EmpfangBestaetigen.0629"`
	Kopf    *NkFVDatenType `xml:"Kopf"`
}

// FVDatenImportBestaetigen0630 is Die Information über den erfolgreichen oder nicht erfolgreichen Import von Daten.
type FVDatenImportBestaetigen0630 struct {
	XMLName              xml.Name                         `xml:"FVDaten.ImportBestaetigen.0630"`
	Kopf                 *NkRueckmeldungFVDatenImportType `xml:"Kopf"`
	ErfolgOderMisserfolg *ErfolgOderMisserfolgDMSType     `xml:"ErfolgOderMisserfolg"`
}

// FVDatenPrimaerdokumentExportieren0631 is Die Informationen zu dem Dokument, dessen Primärdokumente exportiert werden sollen.
type FVDatenPrimaerdokumentExportieren0631 struct {
	XMLName                xml.Name                        `xml:"FVDaten.PrimaerdokumentExportieren.0631"`
	Kopf                   *NkKeineRueckmeldungFVDatenType `xml:"Kopf"`
	DokumentZumExportieren *DokumentZumExportierenType     `xml:"DokumentZumExportieren"`
}

// FVDatenSGOZDAAufheben0632 is Die Informationen zu dem Schriftgutobjekt, dessen zdA-Verfügung im DMS aufgehoben werden soll.
type FVDatenSGOZDAAufheben0632 struct {
	XMLName           xml.Name                              `xml:"FVDaten.SGOZDAAufheben.0632"`
	Kopf              *NkKeineRueckmeldungFVDatenType       `xml:"Kopf"`
	SGOZumZDAAufheben []*SchriftgutobjektZumZDAAufhebenType `xml:"SGOZumZDAAufheben"`
}
