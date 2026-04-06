package models

import (
	"encoding/xml"
)

// ZwischenarchivierungAuslagerung0701 is Ein Dokument, das dem Empfänger zur Erläuterung der Auslagerung mitgegeben wird.
type ZwischenarchivierungAuslagerung0701 struct {
	XMLName          xml.Name                    `xml:"Zwischenarchivierung.Auslagerung.0701"`
	Kopf             *NkZwischenarchivierungType `xml:"Kopf"`
	Anschreiben      *DokumentType               `xml:"Anschreiben,omitempty"`
	Schriftgutobjekt []*Schriftgutobjekt         `xml:"Schriftgutobjekt"`
	Aktenplan        *AktenplanType              `xml:"Aktenplan,omitempty"`
}

// ZwischenarchivierungAuslagerungEmpfangBestaetigen0702 is Der Kopf der Nachricht enthält allgemeine Informationen für den Empfänger der Empfangsbestätigung.
type ZwischenarchivierungAuslagerungEmpfangBestaetigen0702 struct {
	XMLName xml.Name     `xml:"Zwischenarchivierung.AuslagerungEmpfangBestaetigen.0702"`
	Kopf    *NkBasisType `xml:"Kopf"`
}

// ZwischenarchivierungAuslagerungImportBestaetigen0703 is Die Information zum erfolgreichen oder nicht erfolgreichen Import eines Schriftgutobjekts zur Auslagerung.
type ZwischenarchivierungAuslagerungImportBestaetigen0703 struct {
	XMLName             xml.Name                                        `xml:"Zwischenarchivierung.AuslagerungImportBestaetigen.0703"`
	Kopf                *NkBasisType                                    `xml:"Kopf"`
	AusgelagertesObjekt []*ErfolgOderMisserfolgZwischenarchivierungType `xml:"AusgelagertesObjekt"`
}

// ZwischenarchivierungRueckleiheAnforderung0711 is Das Identifikationsmerkmal zu einem Schriftgutobjekt, das zurückgeliehen werden soll.
type ZwischenarchivierungRueckleiheAnforderung0711 struct {
	XMLName                xml.Name     `xml:"Zwischenarchivierung.RueckleiheAnforderung.0711"`
	Kopf                   *NkBasisType `xml:"Kopf"`
	RueckzuleihendesObjekt []string     `xml:"RueckzuleihendesObjekt"`
}

// ZwischenarchivierungRueckleiheUebergabe0712 is Ein Dokument, das dem Empfänger zur Erläuterung der Rückleihe mitgegeben wird.
type ZwischenarchivierungRueckleiheUebergabe0712 struct {
	XMLName          xml.Name                    `xml:"Zwischenarchivierung.RueckleiheUebergabe.0712"`
	Kopf             *NkZwischenarchivierungType `xml:"Kopf"`
	Anschreiben      *DokumentType               `xml:"Anschreiben,omitempty"`
	Schriftgutobjekt []*Schriftgutobjekt         `xml:"Schriftgutobjekt"`
	Aktenplan        *AktenplanType              `xml:"Aktenplan,omitempty"`
}

// ZwischenarchivierungRueckleiheEmpfangBestaetigen0713 is Der Kopf der Nachricht enthält allgemeine Informationen für den Empfänger der Empfangsbestätigung.
type ZwischenarchivierungRueckleiheEmpfangBestaetigen0713 struct {
	XMLName xml.Name     `xml:"Zwischenarchivierung.RueckleiheEmpfangBestaetigen.0713"`
	Kopf    *NkBasisType `xml:"Kopf"`
}

// ZwischenarchivierungRueckuebertragungAnforderung0721 is Das Identifikationsmerkmal zu einem Schriftgutobjekt, das zurückübertragen werden soll.
type ZwischenarchivierungRueckuebertragungAnforderung0721 struct {
	XMLName                     xml.Name     `xml:"Zwischenarchivierung.RueckuebertragungAnforderung.0721"`
	Kopf                        *NkBasisType `xml:"Kopf"`
	RueckzuuebertragendesObjekt []string     `xml:"RueckzuuebertragendesObjekt"`
}

// ZwischenarchivierungRueckuebertragungUebergabe0722 is Ein Dokument, das dem Empfänger zur Erläuterung der Rückübertragung mitgegeben wird.
type ZwischenarchivierungRueckuebertragungUebergabe0722 struct {
	XMLName          xml.Name                    `xml:"Zwischenarchivierung.RueckuebertragungUebergabe.0722"`
	Kopf             *NkZwischenarchivierungType `xml:"Kopf"`
	Anschreiben      *DokumentType               `xml:"Anschreiben,omitempty"`
	Schriftgutobjekt []*Schriftgutobjekt         `xml:"Schriftgutobjekt"`
	Aktenplan        *AktenplanType              `xml:"Aktenplan,omitempty"`
}

// ZwischenarchivierungRueckuebertragungImportBestaetigen0723 is Die Information zum erfolgreichen oder nicht erfolgreichen Import eines Schriftgutobjekts zur Rueckuebertragung.
type ZwischenarchivierungRueckuebertragungImportBestaetigen0723 struct {
	XMLName                  xml.Name                                        `xml:"Zwischenarchivierung.RueckuebertragungImportBestaetigen.0723"`
	Kopf                     *NkBasisType                                    `xml:"Kopf"`
	RueckuebertragenesObjekt []*ErfolgOderMisserfolgZwischenarchivierungType `xml:"RueckuebertragenesObjekt"`
}

// ZwischenarchivierungRueckuebertragungEmpfangBestaetigen0724 is Der Kopf der Nachricht enthält allgemeine Informationen für den Empfänger der Empfangsbestätigung.
type ZwischenarchivierungRueckuebertragungEmpfangBestaetigen0724 struct {
	XMLName xml.Name     `xml:"Zwischenarchivierung.RueckuebertragungEmpfangBestaetigen.0724"`
	Kopf    *NkBasisType `xml:"Kopf"`
}
