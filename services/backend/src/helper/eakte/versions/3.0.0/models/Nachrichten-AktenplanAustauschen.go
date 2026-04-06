package models

import (
	"encoding/xml"
)

// AktenplanAktenplan0301 is Der zu übergebende Aktenplan enthält ein aufgabenbezogenes mehrstufiges Ordnungssystem mit hierarchischer Gliederung für das Bilden und Kennzeichnen von Akten und Vorgängen sowie das Zuordnen von Dokumenten.
type AktenplanAktenplan0301 struct {
	XMLName   xml.Name                                                    `xml:"Aktenplan.Aktenplan.0301"`
	Kopf      *NkNichtFVDatenWeitereEmpfaengerMitEmpfangsbestaetigungType `xml:"Kopf"`
	Aktenplan *AktenplanType                                              `xml:"Aktenplan"`
}

// AktenplanEmpfangBestaetigen0302 is Der Kopf der Nachricht enthält allgemeine Informationen für den Empfänger der Empfangsbestätigung.
type AktenplanEmpfangBestaetigen0302 struct {
	XMLName xml.Name     `xml:"Aktenplan.EmpfangBestaetigen.0302"`
	Kopf    *NkBasisType `xml:"Kopf"`
}
