package models

import (
	"encoding/xml"
)

// AussonderungAnbieteverzeichnis0501 is Ein Dokument, in dem allgemeine Informationen zu den anzubietenden Schriftgutobjekten an das Archiv mitgegeben werden. Hier können auch Formblätter (z.B. Vorblatt zur Abgabe) mitgegeben werden.
type AussonderungAnbieteverzeichnis0501 struct {
	XMLName          xml.Name                                                `xml:"Aussonderung.Anbieteverzeichnis.0501"`
	Kopf             *NkNichtFVDatenEinEmpfaengerMitEmpfangsbestaetigungType `xml:"Kopf"`
	Anschreiben      []*DokumentType                                         `xml:"Anschreiben,omitempty"`
	Schriftgutobjekt []*Schriftgutobjekt                                     `xml:"Schriftgutobjekt"`
	Aktenplan        *AktenplanType                                          `xml:"Aktenplan,omitempty"`
}

// AussonderungBewertungsverzeichnis0502 is Die Information zu einem Schriftgutobjekt über die durch das Archiv erfolgte Bewertung im Anbieteverzeichnis.
type AussonderungBewertungsverzeichnis0502 struct {
	XMLName          xml.Name                                                `xml:"Aussonderung.Bewertungsverzeichnis.0502"`
	Kopf             *NkNichtFVDatenEinEmpfaengerMitEmpfangsbestaetigungType `xml:"Kopf"`
	BewertetesObjekt []*RueckgabeparameterAnbietungType                      `xml:"BewertetesObjekt"`
}

// AussonderungAussonderung0503 is Ein Dokument, in dem allgemeine Informationen zu den auszusondernden Schriftgutobjekten an das Archiv mitgegeben werden können. Hier können auch Formblätter (z.B. "Vorblatt zur Abgabe" gemäß Registraturrichtlinie) mitgegeben werden.
type AussonderungAussonderung0503 struct {
	XMLName          xml.Name            `xml:"Aussonderung.Aussonderung.0503"`
	Kopf             *NkAussonderungType `xml:"Kopf"`
	Anschreiben      []*DokumentType     `xml:"Anschreiben,omitempty"`
	Schriftgutobjekt []*Schriftgutobjekt `xml:"Schriftgutobjekt"`
	Aktenplan        *AktenplanType      `xml:"Aktenplan,omitempty"`
}

// AussonderungAnbietungEmpfangBestaetigen0504 is Der Kopf der Nachricht enthält allgemeine Informationen für den Empfänger der Empfangsbestätigung.
type AussonderungAnbietungEmpfangBestaetigen0504 struct {
	XMLName xml.Name     `xml:"Aussonderung.AnbietungEmpfangBestaetigen.0504"`
	Kopf    *NkBasisType `xml:"Kopf"`
}

// AussonderungBewertungEmpfangBestaetigen0505 is Der Kopf der Nachricht enthält allgemeine Informationen für den Empfänger der Empfangsbestätigung.
type AussonderungBewertungEmpfangBestaetigen0505 struct {
	XMLName xml.Name     `xml:"Aussonderung.BewertungEmpfangBestaetigen.0505"`
	Kopf    *NkBasisType `xml:"Kopf"`
}

// AussonderungAussonderungImportBestaetigen0506 is Die Information zum erfolgreichen oder nicht erfolgreichen Import eines Schriftgutobjekts zur Aussonderung.
type AussonderungAussonderungImportBestaetigen0506 struct {
	XMLName           xml.Name                                `xml:"Aussonderung.AussonderungImportBestaetigen.0506"`
	Kopf              *NkBasisType                            `xml:"Kopf"`
	AusgesondertesSGO []*ErfolgOderMisserfolgAussonderungType `xml:"AusgesondertesSGO"`
}

// AussonderungAussonderungEmpfangBestaetigen0507 is Der Kopf der Nachricht enthält allgemeine Informationen für den Empfänger der Empfangsbestätigung.
type AussonderungAussonderungEmpfangBestaetigen0507 struct {
	XMLName xml.Name     `xml:"Aussonderung.AussonderungEmpfangBestaetigen.0507"`
	Kopf    *NkBasisType `xml:"Kopf"`
}

// AussonderungAktenplanZurBewertung0511 is Der für die Bewertung zu übergebende Aktenplan enthält ein aufgabenbezogenes mehrstufiges Ordnungssystem mit hierarchischer Gliederung für das Bilden und Kennzeichnen von Akten und Vorgängen sowie das Zuordnen von Dokumenten.
type AussonderungAktenplanZurBewertung0511 struct {
	XMLName   xml.Name                                                    `xml:"Aussonderung.AktenplanZurBewertung.0511"`
	Kopf      *NkNichtFVDatenWeitereEmpfaengerMitEmpfangsbestaetigungType `xml:"Kopf"`
	Aktenplan *AktenplanBewertungskatalogType                             `xml:"Aktenplan"`
}

// AussonderungAktenplanZurBewertungEmpfangBestaetigen0512 is Der Kopf der Nachricht enthält allgemeine Informationen für den Empfänger der Empfangsbestätigung.
type AussonderungAktenplanZurBewertungEmpfangBestaetigen0512 struct {
	XMLName xml.Name     `xml:"Aussonderung.AktenplanZurBewertungEmpfangBestaetigen.0512"`
	Kopf    *NkBasisType `xml:"Kopf"`
}

// AussonderungBewertungskatalog0513 is Die bewertete Einheit eines Aktenplans. Eine Aktenplanheit entspricht z.B. einer Hauptgruppe, Obergruppe, Gruppe, Untergruppe oder Betreffseinheit.
type AussonderungBewertungskatalog0513 struct {
	XMLName                   xml.Name                                                    `xml:"Aussonderung.Bewertungskatalog.0513"`
	Kopf                      *NkNichtFVDatenWeitereEmpfaengerMitEmpfangsbestaetigungType `xml:"Kopf"`
	BewerteteAktenplaneinheit []*AktenplaneinheitAktenplanBewertungskatalogType           `xml:"BewerteteAktenplaneinheit"`
}

// AussonderungBewertungskatalogImportBestaetigen0514 is Die Information über den erfolgreichen oder nicht erfolgreichen Import des Wertes in "Aussonderungsart" oder "AussonderungsartKonfigurierbar" zu der jeweiligen Aktenplaneinheit.
type AussonderungBewertungskatalogImportBestaetigen0514 struct {
	XMLName                   xml.Name                                     `xml:"Aussonderung.BewertungskatalogImportBestaetigen.0514"`
	Kopf                      *NkBasisType                                 `xml:"Kopf"`
	BewerteteAktenplaneinheit []*ErfolgOderMisserfolgBewertungskatalogType `xml:"BewerteteAktenplaneinheit"`
}

// AussonderungBewerteterAktenplanEmpfangBestaetigen0515 is Der Kopf der Nachricht enthält allgemeine Informationen für den Empfänger der Empfangsbestätigung.
type AussonderungBewerteterAktenplanEmpfangBestaetigen0515 struct {
	XMLName xml.Name     `xml:"Aussonderung.BewerteterAktenplanEmpfangBestaetigen.0515"`
	Kopf    *NkBasisType `xml:"Kopf"`
}
