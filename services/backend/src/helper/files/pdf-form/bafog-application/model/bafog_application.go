package model

import (
	"fmt"
	"time"

	de_translations "github.com/dcssoftware/bafoeg-manager/src/helper/validator-i18n/de"
	de "github.com/go-playground/locales/de_DE"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type BescheidÜbermittelnAnEnum string
type VersichertEnum string
type VerhältnisElternteileEnum string
type SorgeberechtigterEnum string

const (
	BescheidÜbermittelnAnAntragstellerStändigerWohnsitz      BescheidÜbermittelnAnEnum = "mich (ständiger Wohnsitz)"
	BescheidÜbermittelnAnAntragstellerAusbildungsortWohnsitz BescheidÜbermittelnAnEnum = "mich (Wohnsitz am Ausbildungsort)"
	BescheidÜbermittelnAnElternteil01                        BescheidÜbermittelnAnEnum = "meinen ersten Elternteil"
	BescheidÜbermittelnAnElternteil02                        BescheidÜbermittelnAnEnum = "meinen zweiten Elternteil"
	BescheidÜbermittelnAnSorgeberechtiger                    BescheidÜbermittelnAnEnum = "meine/-n Sorgeberechtigte/-n"
	BescheidÜbermittelnAnBevollmächtigtePerson               BescheidÜbermittelnAnEnum = "die von mir bevollmächtigte Person"
	BescheidÜbermittelnAnUnbekannt                           BescheidÜbermittelnAnEnum = "Unbekannt"

	VersichertGesetzlichFamilienversichert   VersichertEnum = "gesetzlich familienversichert"
	VersichertStudentischFamilienversichert  VersichertEnum = "studentisch familienversichert"
	VersichertPrivatVersichert               VersichertEnum = "privat versichert"
	VersichertFreiwilligGesetzlichVersichert VersichertEnum = "freiwillig gesetzlich versichert"
	VersichertAndersVersichert               VersichertEnum = "anders versichert"

	VerhältnisElternteileLebenZusammen                      VerhältnisElternteileEnum = "ja"
	VerhältnisElternteileLebenZusammenAberDauerhaftGetrennt VerhältnisElternteileEnum = "ja, aber dauernd getrennt lebend"
	VerhältnisElternteileNichtZusammenlebend                VerhältnisElternteileEnum = "nein"

	SorgeberechtigterJa   SorgeberechtigterEnum = "ja, und zwar"
	SorgeberechtigterNein SorgeberechtigterEnum = "nein"
)

type RawBAföGApplicationBaseModel01 struct {

	/*
	*			############################################################
	*			#                                                          #
	*			#												PAGE 01                            #
	*			#                                                          #
	*			############################################################
	*
	 */

	Förderungsnummer                       string     `pdfpage:"1" pdf:"Förderungsnummer (falls vorhanden)" validate:"omitempty,alphanum"`
	Ausbildungsstätte                      string     `pdfpage:"1" pdf:"Ausbildungsstätte und Ausbildungsort" validate:"required,alphanumspace"`
	KlasseFachrichtung                     string     `pdfpage:"1" pdf:"Klasse/Fachrichtung" validate:"required,alphanumspace"`
	AngestrebterAbschluss                  string     `pdfpage:"1" pdf:"angestrebter Abschluss" validate:"required,alphanumspace"`
	Vollzeitausbildung                     bool       `pdfpage:"1" pdf:"Vollzeitausbildung" validate:"required,boolean"`
	FrühererBAföGAntrag                    bool       `pdfpage:"1" pdf:"früherer BAföG-Antrag" validate:"required,boolean"`
	FrühererBAföGAntragZuständigesAmt      string     `pdfpage:"1" pdf:"bisheriges Amt für Ausbildungsförderung" validate:"omitempty,alphanumspace"`
	FrühererBAföGAntragFörderungsnummer    string     `pdfpage:"1" pdf:"bisherige Förderungsnummer"`
	Nachname                               string     `pdfpage:"1" pdf:"Name" validate:"required,alphaspace"`
	Vorname                                string     `pdfpage:"1" pdf:"Vorname" validate:"required,alphaspace"`
	NachnameGeändertSeit                   *time.Time `pdfpage:"1" pdf:"bei Änderung seit"`
	Familienstand                          string     `pdfpage:"1" pdf:"Familienstand"`
	Geschlecht                             string     `pdfpage:"1" pdf:"Geschlecht"`
	Geburtsname                            string     `pdfpage:"1" pdf:"Geburtsname"`
	Geburtsdatum                           *time.Time `pdfpage:"1" pdf:"Geburtsdatum"`
	Geburtsort                             string     `pdfpage:"1" pdf:"Geburtsort"`
	EigeneStaatsangehörigkeit              string     `pdfpage:"1" pdf:"eigene Staatsangehörigkeit"`
	HatEigeneKinder                        bool       `pdfpage:"1" pdf:"eigene Kinder"`
	AnschriftHauptwohnsitzStraße           string     `pdfpage:"1" pdf:"Anschrift Straße"`
	AnschriftHauptwohnsitzHausnummer       string     `pdfpage:"1" pdf:"Anschrift Hausnummer"`
	AnschriftHauptwohnsitzAdresszusatz     string     `pdfpage:"1" pdf:"Anschrift Adresszusatz"`
	AnschriftHauptwohnsitzLand             string     `pdfpage:"1" pdf:"Anschrift Land"`
	AnschriftHauptwohnsitzPostleitzahl     string     `pdfpage:"1" pdf:"Anschrift Postleitzahl"`
	AnschriftHauptwohnsitzOrt              string     `pdfpage:"1" pdf:"Anschrift Ort"`
	AnschriftGemeinschaftMitEltern         bool       `pdfpage:"1" pdf:"häusliche Gemeinschaft mit Eltern"`
	AnschriftEigentumMieteigentumDerEltern bool       `pdfpage:"1" pdf:"Eigentum/Miteigentum der Eltern"`
	AnschriftAusbildungStraße              string     `pdfpage:"1" pdf:"Anschrift Ausbildung Straße"`
	AnschriftAusbildungHausnummer          string     `pdfpage:"1" pdf:"Anschrift Ausbildung Hausnummer"`
	AnschriftAusbildungAdresszusatz        string     `pdfpage:"1" pdf:"Anschrift Ausbildung Adresszusatz"`
	AnschriftAusbildungLand                string     `pdfpage:"1" pdf:"Anschrift Ausbildung Land"`
	AnschriftAusbildungPostleitzahl        string     `pdfpage:"1" pdf:"Anschrift Ausbildung Postleitzahl"`
	AnschriftAusbildungOrt                 string     `pdfpage:"1" pdf:"Anschrift Ausbildung Ort"`

	/*
	*			############################################################
	*			#                                                          #
	*			#												PAGE 02                            #
	*			#                                                          #
	*			############################################################
	*
	 */

	KontaktdatenEMail               string                     `pdfpage:"2" pdf:"Kontaktdaten E-Mail" validate:"required,email"`
	KontaktdatenTelefon             string                     `pdfpage:"2" pdf:"Kontaktdaten Telefon"`
	BescheidÜbermittelnAn           *BescheidÜbermittelnAnEnum `pdfpage:"2" pdf:"Der Bescheid soll übermittelt werden an"`
	IBAN1                           string                     `pdfpage:"2" pdf:"IBAN 1"`
	IBAN2                           string                     `pdfpage:"2" pdf:"IBAN 2"`
	IBAN3                           string                     `pdfpage:"2" pdf:"IBAN 3"`
	IBAN4                           string                     `pdfpage:"2" pdf:"IBAN 4"`
	IBAN5                           string                     `pdfpage:"2" pdf:"IBAN 5"`
	IBAN6                           string                     `pdfpage:"2" pdf:"IBAN 6"`
	GeldistitutName                 string                     `pdfpage:"2" pdf:"Name Geldistitut"`
	KontoinhaberName                string                     `pdfpage:"2" pdf:"Name, Vorname Kontoinhaber"`
	Krankenversicherung             *VersichertEnum            `pdfpage:"2" pdf:"Krankenversicherung"`
	Pflegeversicherung              bool                       `pdfpage:"2" pdf:"Pflegeversicherung"`
	Steueridentifikationsnummer     string                     `pdfpage:"2" pdf:"Steueridentifikationsnummer"`
	VerhältnisElternteile           *VerhältnisElternteileEnum `pdfpage:"2" pdf:"Verhältnis Elternteile"`
	Elternteil01Name                string                     `pdfpage:"2" pdf:"Name 1. Elternteil"`
	Elternteil01Vorname             string                     `pdfpage:"2" pdf:"Vorname Name 1. Elternteil"`
	Elternteil01Geschlecht          string                     `pdfpage:"2" pdf:"Geschlecht 1. Elternteil"`
	Elternteil01Geburtsdatum        *time.Time                 `pdfpage:"2" pdf:"Geburtsdatum 1. Elternteil"`
	Elternteil01Sterbedatum         *time.Time                 `pdfpage:"2" pdf:"Sterbedatum 1. Elternteil"`
	Elternteil01Staatsangehörigkeit string                     `pdfpage:"2" pdf:"Staatsangehörigkeit 1. Elternteil"`
	Elternteil01Straße              string                     `pdfpage:"2" pdf:"Straße 1. Elternteil"`
	Elternteil01Hausnummer          string                     `pdfpage:"2" pdf:"Hausnummer 1. Elternteil"`
	Elternteil01Adresszusatz        string                     `pdfpage:"2" pdf:"Adresszusatz 1. Elternteil"`
	Elternteil01Land                string                     `pdfpage:"2" pdf:"Land 1. Elternteil"`
	Elternteil01Postleitzahl        string                     `pdfpage:"2" pdf:"Postleitzahl 1. Elternteil"`
	Elternteil01Ort                 string                     `pdfpage:"2" pdf:"Ort 1. Elternteil"`
	Elternteil02Name                string                     `pdfpage:"2" pdf:"Name 2. Elternteil"`
	Elternteil02Vorname             string                     `pdfpage:"2" pdf:"Vorname Name 2. Elternteil"`
	Elternteil02Geschlecht          string                     `pdfpage:"2" pdf:"Geschlecht 2. Elternteil"`
	Elternteil02Geburtsdatum        *time.Time                 `pdfpage:"2" pdf:"Geburtsdatum 2. Elternteil"`
	Elternteil02Sterbedatum         *time.Time                 `pdfpage:"2" pdf:"Sterbedatum  2. Elternteil"`
	Elternteil02Staatsangehörigkeit string                     `pdfpage:"2" pdf:"Staatsangehörigkeit  2. Elternteil"`
	Elternteil02Straße              string                     `pdfpage:"2" pdf:"Straße  2. Elternteil"`
	Elternteil02Hausnummer          string                     `pdfpage:"2" pdf:"Hausnummer  2. Elternteil"`
	Elternteil02Adresszusatz        string                     `pdfpage:"2" pdf:"Adresszusatz  2. Elternteil"`
	Elternteil02Land                string                     `pdfpage:"2" pdf:"Land  2. Elternteil"`
	Elternteil02Postleitzahl        string                     `pdfpage:"2" pdf:"Postleitzahl  2. Elternteil"`
	Elternteil02Ort                 string                     `pdfpage:"2" pdf:"Ort  2. Elternteil"`

	/*
	*			############################################################
	*			#                                                          #
	*			#												PAGE 03                            #
	*			#                                                          #
	*			############################################################
	*
	 */

	AnwärterbezügeAusÖffentlichenMitteln        bool       `pdfpage:"3" pdf:"Anwärterbezüge oder ähnliche Leistungen aus öffentlichen Mitteln"`
	LeistungenSGB2SGB3                          bool       `pdfpage:"3" pdf:"Leistungen SGB II oder SGB III"`
	LeistungenBegabtenförderungswerk            bool       `pdfpage:"3" pdf:"Leistungen Begabtenförderungswerk"`
	KeineDerVorstehendenLeistungen              bool       `pdfpage:"3" pdf:"nein, keine der vorstehenden Leistungen"`
	BewilligungszeitraumVon                     *time.Time `pdfpage:"3" pdf:"Bewilligungszeitraum von" dateformat:"012006"`
	BewilligungszeitraumBis                     *time.Time `pdfpage:"3" pdf:"Bewilligungszeitraum bis" dateformat:"012006"`
	Sozialleistungen1                           string     `pdfpage:"3" pdf:"Sozialleistungen 1"`
	Sozialleistungen2                           string     `pdfpage:"3" pdf:"Sozialleistungen 2"`
	RiesterEuro                                 float64    `pdfpage:"3" pdf:"Riester Euro"`
	VoraussichtlicheEinnahmen                   bool       `pdfpage:"3" pdf:"voraussichtliche Einnahmen"`
	BruttoeinnahmenEuro                         float64    `pdfpage:"3" pdf:"Bruttoeinnahmen Euro"`
	AbeitgeberanteilEnthaltenEuro               bool       `pdfpage:"3" pdf:"Abeitgeberanteil enthalten Euro"`
	AusbildungsUndPraktikumsvergütungBruttoEuro float64    `pdfpage:"3" pdf:"Ausbildungs- und Praktikumsvergütung brutto Euro"`
	EinkünfteEuro                               float64    `pdfpage:"3" pdf:"Einkünfte Euro"`
	KapitalvermögenEuro                         float64    `pdfpage:"3" pdf:"Kapitalvermögen  Euro"`
	WaisenrenteWaisengeldRentenEuro             float64    `pdfpage:"3" pdf:"Waisenrente / Waisengeld / Renten  Euro"`
	AusbildungsbeihilfenEuro                    float64    `pdfpage:"3" pdf:"Ausbildungsbeihilfen Euro"`
	UnterhaltsleistungenEuro                    float64    `pdfpage:"3" pdf:"Unterhaltsleistungen Euro"`
	UnterhaltsvorschussleistungenEuro           float64    `pdfpage:"3" pdf:"Unterhaltsvorschussleistungen Euro"`
	WeitereEinnahmenEuro                        float64    `pdfpage:"3" pdf:"Weitere Einnahmen Euro"`

	/*
	*			############################################################
	*			#                                                          #
	*			#												PAGE 04                            #
	*			#                                                          #
	*			############################################################
	*
	 */

	BarvermögenEuro                   float64 `pdfpage:"4" pdf:"Barvermögen Euro"`
	BarvermögenNein                   bool    `pdfpage:"4" pdf:"Barvermögen nein" emptybooltrue:"1"`
	BankUndSparguthabenEuro           float64 `pdfpage:"4" pdf:"Bank- und Sparguthaben Euro"`
	BankUndSparguthabenNein           bool    `pdfpage:"4" pdf:"Bank- und Sparguthaben nein" emptybooltrue:"1"`
	BausparUndPrämienguthabenEuro     float64 `pdfpage:"4" pdf:"Wertpapiere Euro"`
	BausparUndPrämienguthabenNein     bool    `pdfpage:"4" pdf:"Wertpapiere nein" emptybooltrue:"1"`
	WertpapiereEuro                   float64 `pdfpage:"4" pdf:"Wertpapiere Euro"`
	WertpapiereNein                   bool    `pdfpage:"4" pdf:"Wertpapiere nein" emptybooltrue:"1"`
	KraftfahrzeugeEuro                float64 `pdfpage:"4" pdf:"Kraftfahrzeuge Euro"`
	KraftfahrzeugeNein                bool    `pdfpage:"4" pdf:"Kraftfahrzeuge nein" emptybooltrue:"1"`
	LebensversicherungenEuro          float64 `pdfpage:"4" pdf:"Lebensversicherungen Euro"`
	LebensversicherungenNein          bool    `pdfpage:"4" pdf:"Lebensversicherungen nein" emptybooltrue:"1"`
	AltersvorsorgevermögenEuro        float64 `pdfpage:"4" pdf:"Altersvorsorgevermögen Euro"`
	AltersvorsorgevermögenNein        bool    `pdfpage:"4" pdf:"Altersvorsorgevermögen nein" emptybooltrue:"1"`
	GrundstückeWohnungenEuro          float64 `pdfpage:"4" pdf:"Grundstücke, Wohnungen Euro"`
	GrundstückeWohnungenNein          bool    `pdfpage:"4" pdf:"Grundstücke, Wohnungen nein" emptybooltrue:"1"`
	BetriebsvermögenEuro              float64 `pdfpage:"4" pdf:"Betriebsvermögen Euro"`
	BetriebsvermögenNein              bool    `pdfpage:"4" pdf:"Betriebsvermögen nein" emptybooltrue:"1"`
	GeldforderungenEuro               float64 `pdfpage:"4" pdf:"Geldforderungen Euro"`
	GeldforderungenNein               bool    `pdfpage:"4" pdf:"Geldforderungen nein" emptybooltrue:"1"`
	VermögensgegenständeEuro          float64 `pdfpage:"4" pdf:"Vermögensgegenstände Euro"`
	VermögensgegenständeNein          bool    `pdfpage:"4" pdf:"Vermögensgegenstände nein" emptybooltrue:"1"`
	VermögenswerteAnrechnungsfreiEuro float64 `pdfpage:"4" pdf:"Vermögenswerte anrechnungsfrei Euro"`
	ÜbergangsbeihilfenEuro            float64 `pdfpage:"4" pdf:"Übergangsbeihilfen Euro"`
	SchuldenEuro                      float64 `pdfpage:"4" pdf:"Schulden Euro"`
	LastenEuro                        float64 `pdfpage:"4" pdf:"Lasten Euro"`

	/*
	*			############################################################
	*			#                                                          #
	*			#												PAGE 05                            #
	*			#                                                          #
	*			############################################################
	*
	 */

	Eintrag01Von                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr von 1"`
	Eintrag01Bis                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr bis 1"`
	Eintrag01Name                        string     `pdfpage:"5" pdf:"Name und Ort 1"`
	Eintrag01SchulartOderArtDerTätigkeit string     `pdfpage:"5" pdf:"Schulart/Fachrichtung/Tätigkeit 1"`
	Eintrag01AbschlussOderBruttolohn     string     `pdfpage:"5" pdf:"Abschluss / Bruttolohn /Leistung 1"`

	Eintrag02Von                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr von 2"`
	Eintrag02Bis                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr bis 2"`
	Eintrag02Name                        string     `pdfpage:"5" pdf:"Name und Ort 2"`
	Eintrag02SchulartOderArtDerTätigkeit string     `pdfpage:"5" pdf:"Schulart/Fachrichtung/Tätigkeit 2"`
	Eintrag02AbschlussOderBruttolohn     string     `pdfpage:"5" pdf:"Abschluss / Bruttolohn /Leistung 2"`

	Eintrag03Von                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr von 3"`
	Eintrag03Bis                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr bis 3"`
	Eintrag03Name                        string     `pdfpage:"5" pdf:"Name und Ort 3"`
	Eintrag03SchulartOderArtDerTätigkeit string     `pdfpage:"5" pdf:"Schulart/Fachrichtung/Tätigkeit 3"`
	Eintrag03AbschlussOderBruttolohn     string     `pdfpage:"5" pdf:"Abschluss / Bruttolohn /Leistung 3"`

	Eintrag04Von                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr von 4"`
	Eintrag04Bis                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr bis 4"`
	Eintrag04Name                        string     `pdfpage:"5" pdf:"Name und Ort 4"`
	Eintrag04SchulartOderArtDerTätigkeit string     `pdfpage:"5" pdf:"Schulart/Fachrichtung/Tätigkeit 4"`
	Eintrag04AbschlussOderBruttolohn     string     `pdfpage:"5" pdf:"Abschluss / Bruttolohn /Leistung 4"`

	Eintrag05Von                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr von 5"`
	Eintrag05Bis                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr bis 5"`
	Eintrag05Name                        string     `pdfpage:"5" pdf:"Name und Ort 5"`
	Eintrag05SchulartOderArtDerTätigkeit string     `pdfpage:"5" pdf:"Schulart/Fachrichtung/Tätigkeit 5"`
	Eintrag05AbschlussOderBruttolohn     string     `pdfpage:"5" pdf:"Abschluss / Bruttolohn /Leistung 5"`

	Eintrag06Von                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr von 6"`
	Eintrag06Bis                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr bis 6"`
	Eintrag06Name                        string     `pdfpage:"5" pdf:"Name und Ort 6"`
	Eintrag06SchulartOderArtDerTätigkeit string     `pdfpage:"5" pdf:"Schulart/Fachrichtung/Tätigkeit 6"`
	Eintrag06AbschlussOderBruttolohn     string     `pdfpage:"5" pdf:"Abschluss / Bruttolohn /Leistung 6"`

	Eintrag07Von                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr von 7"`
	Eintrag07Bis                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr bis 7"`
	Eintrag07Name                        string     `pdfpage:"5" pdf:"Name und Ort 7"`
	Eintrag07SchulartOderArtDerTätigkeit string     `pdfpage:"5" pdf:"Schulart/Fachrichtung/Tätigkeit 7"`
	Eintrag07AbschlussOderBruttolohn     string     `pdfpage:"5" pdf:"Abschluss / Bruttolohn /Leistung 7"`

	Eintrag08Von                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr von 8"`
	Eintrag08Bis                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr bis 8"`
	Eintrag08Name                        string     `pdfpage:"5" pdf:"Name und Ort 8"`
	Eintrag08SchulartOderArtDerTätigkeit string     `pdfpage:"5" pdf:"Schulart/Fachrichtung/Tätigkeit 8"`
	Eintrag08AbschlussOderBruttolohn     string     `pdfpage:"5" pdf:"Abschluss / Bruttolohn /Leistung 8"`

	Eintrag09Von                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr von 9"`
	Eintrag09Bis                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr bis 9"`
	Eintrag09Name                        string     `pdfpage:"5" pdf:"Name und Ort 9"`
	Eintrag09SchulartOderArtDerTätigkeit string     `pdfpage:"5" pdf:"Schulart/Fachrichtung/Tätigkeit 9"`
	Eintrag09AbschlussOderBruttolohn     string     `pdfpage:"5" pdf:"Abschluss / Bruttolohn /Leistung 9"`

	Eintrag10Von                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr von 10"`
	Eintrag10Bis                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr bis 10"`
	Eintrag10Name                        string     `pdfpage:"5" pdf:"Name und Ort 10"`
	Eintrag10SchulartOderArtDerTätigkeit string     `pdfpage:"5" pdf:"Schulart/Fachrichtung/Tätigkeit 10"`
	Eintrag10AbschlussOderBruttolohn     string     `pdfpage:"5" pdf:"Abschluss / Bruttolohn /Leistung 10"`

	Eintrag11Von                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr von 11"`
	Eintrag11Bis                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr bis 11"`
	Eintrag11Name                        string     `pdfpage:"5" pdf:"Name und Ort 11"`
	Eintrag11SchulartOderArtDerTätigkeit string     `pdfpage:"5" pdf:"Schulart/Fachrichtung/Tätigkeit 11"`
	Eintrag11AbschlussOderBruttolohn     string     `pdfpage:"5" pdf:"Abschluss / Bruttolohn /Leistung 11"`

	Eintrag12Von                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr von 12"`
	Eintrag12Bis                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr bis 12"`
	Eintrag12Name                        string     `pdfpage:"5" pdf:"Name und Ort 12"`
	Eintrag12SchulartOderArtDerTätigkeit string     `pdfpage:"5" pdf:"Schulart/Fachrichtung/Tätigkeit 12"`
	Eintrag12AbschlussOderBruttolohn     string     `pdfpage:"5" pdf:"Abschluss / Bruttolohn /Leistung 12"`

	Eintrag13Von                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr von 13"`
	Eintrag13Bis                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr bis 13"`
	Eintrag13Name                        string     `pdfpage:"5" pdf:"Name und Ort 13"`
	Eintrag13SchulartOderArtDerTätigkeit string     `pdfpage:"5" pdf:"Schulart/Fachrichtung/Tätigkeit 13"`
	Eintrag13AbschlussOderBruttolohn     string     `pdfpage:"5" pdf:"Abschluss / Bruttolohn /Leistung 13"`

	Eintrag14Von                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr von 14"`
	Eintrag14Bis                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr bis 14"`
	Eintrag14Name                        string     `pdfpage:"5" pdf:"Name und Ort 14"`
	Eintrag14SchulartOderArtDerTätigkeit string     `pdfpage:"5" pdf:"Schulart/Fachrichtung/Tätigkeit 14"`
	Eintrag14AbschlussOderBruttolohn     string     `pdfpage:"5" pdf:"Abschluss / Bruttolohn /Leistung 14"`

	Eintrag15Von                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr von 15"`
	Eintrag15Bis                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr bis 15"`
	Eintrag15Name                        string     `pdfpage:"5" pdf:"Name und Ort 15"`
	Eintrag15SchulartOderArtDerTätigkeit string     `pdfpage:"5" pdf:"Schulart/Fachrichtung/Tätigkeit 15"`
	Eintrag15AbschlussOderBruttolohn     string     `pdfpage:"5" pdf:"Abschluss / Bruttolohn /Leistung 15"`

	Eintrag16Von                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr von 16"`
	Eintrag16Bis                         *time.Time `pdfpage:"5" pdf:"Monat / Jahr bis 16"`
	Eintrag16Name                        string     `pdfpage:"5" pdf:"Name und Ort 16"`
	Eintrag16SchulartOderArtDerTätigkeit string     `pdfpage:"5" pdf:"Schulart/Fachrichtung/Tätigkeit 16"`
	Eintrag16AbschlussOderBruttolohn     string     `pdfpage:"5" pdf:"Abschluss / Bruttolohn /Leistung 16"`

	/*
	*			############################################################
	*			#                                                          #
	*			#												PAGE 06                            #
	*			#                                                          #
	*			############################################################
	*
	 */

	Sorgeberechtigter                 *SorgeberechtigterEnum `pdfpage:"6" pdf:"Sorgeberechtigte/r"`
	SorgeberechtigterName             string                 `pdfpage:"6" pdf:"Name, Vorname Sorgeberechtigte/r"`
	SorgeberechtigterStraße           string                 `pdfpage:"6" pdf:"Straße Sorgeberechtigte/r"`
	SorgeberechtigterHausnummer       string                 `pdfpage:"6" pdf:"Hausnummer Sorgeberechtigte/r"`
	SorgeberechtigterAddresszusatz    string                 `pdfpage:"6" pdf:"Adresszusatz Sorgeberechtigte/r"`
	SorgeberechtigterLand             string                 `pdfpage:"6" pdf:"Land Sorgeberechtigte/r"`
	SorgeberechtigterPostleitzahl     string                 `pdfpage:"6" pdf:"Postleitzahl Sorgeberechtigte/r"`
	SorgeberechtigterOrt              string                 `pdfpage:"6" pdf:"Ort Sorgeberechtigte/r"`
	UnterschriftAuszubildendePerson   string                 `pdfpage:"6" pdf:"Datum UnterschriftNamensangabe durch die auszubildende Person"`
	UnterschriftGesetzlicherVertreter string                 `pdfpage:"6" pdf:"Datum UnterschriftNamensangabe derdes gesetzlichen VertreterinVertreters bei Minderjährigen"`
}

type RawBAföGApplicationBaseModel01CV struct {
	Von                         *time.Time
	Bis                         *time.Time
	Name                        string
	SchulartOderArtDerTätigkeit string
	AbschlussOderBruttolohn     string
}

func (model RawBAföGApplicationBaseModel01) Validate() ([]string, error) {

	// Set up German translator for user‑friendly messages
	langDe := de.New()
	universalTranslator := ut.New(langDe, langDe)
	translator, languageFound := universalTranslator.GetTranslator("de_DE")
	if !languageFound {
		return []string{}, fmt.Errorf("could not find translator for de_DE")
	}

	// Initialise validator and register German translations
	validate := validator.New()
	if err := de_translations.RegisterDefaultTranslations(validate, translator); err != nil {
		return []string{}, err
	}

	// Perform validation on the struct
	if validationErr := validate.Struct(model); validationErr != nil {
		if ve, ok := validationErr.(validator.ValidationErrors); ok {
			var msgs []string
			for _, fe := range ve {
				msgs = append(msgs, fe.Translate(translator))
			}
			return msgs, nil
		}
		// Return the error string as a single element slice
		return []string{validationErr.Error()}, nil
	}

	// No validation errors
	return []string{}, nil
}

func (model RawBAföGApplicationBaseModel01) GetIBAN() string {
	return fmt.Sprintf(
		"%s %s %s %s %s %s",
		model.IBAN1,
		model.IBAN2,
		model.IBAN3,
		model.IBAN4,
		model.IBAN5,
		model.IBAN6,
	)
}

func (model BescheidÜbermittelnAnEnum) String() string {
	return string(model)
}
func (model VersichertEnum) String() string {
	return string(model)
}
func (model VerhältnisElternteileEnum) String() string {
	return string(model)
}

func (model SorgeberechtigterEnum) String() string {
	return string(model)
}

func (model RawBAföGApplicationBaseModel01) GetCV() []RawBAföGApplicationBaseModel01CV {
	return []RawBAföGApplicationBaseModel01CV{
		{
			Von:                         model.Eintrag01Von,
			Bis:                         model.Eintrag01Bis,
			Name:                        model.Eintrag01Name,
			SchulartOderArtDerTätigkeit: model.Eintrag01SchulartOderArtDerTätigkeit,
			AbschlussOderBruttolohn:     model.Eintrag01AbschlussOderBruttolohn,
		},
		{
			Von:                         model.Eintrag02Von,
			Bis:                         model.Eintrag02Bis,
			Name:                        model.Eintrag02Name,
			SchulartOderArtDerTätigkeit: model.Eintrag02SchulartOderArtDerTätigkeit,
			AbschlussOderBruttolohn:     model.Eintrag02AbschlussOderBruttolohn,
		},
		{
			Von:                         model.Eintrag03Von,
			Bis:                         model.Eintrag03Bis,
			Name:                        model.Eintrag03Name,
			SchulartOderArtDerTätigkeit: model.Eintrag03SchulartOderArtDerTätigkeit,
			AbschlussOderBruttolohn:     model.Eintrag03AbschlussOderBruttolohn,
		},
		{
			Von:                         model.Eintrag04Von,
			Bis:                         model.Eintrag04Bis,
			Name:                        model.Eintrag04Name,
			SchulartOderArtDerTätigkeit: model.Eintrag04SchulartOderArtDerTätigkeit,
			AbschlussOderBruttolohn:     model.Eintrag04AbschlussOderBruttolohn,
		},
		{
			Von:                         model.Eintrag05Von,
			Bis:                         model.Eintrag05Bis,
			Name:                        model.Eintrag05Name,
			SchulartOderArtDerTätigkeit: model.Eintrag05SchulartOderArtDerTätigkeit,
			AbschlussOderBruttolohn:     model.Eintrag05AbschlussOderBruttolohn,
		},
		{
			Von:                         model.Eintrag06Von,
			Bis:                         model.Eintrag06Bis,
			Name:                        model.Eintrag06Name,
			SchulartOderArtDerTätigkeit: model.Eintrag06SchulartOderArtDerTätigkeit,
			AbschlussOderBruttolohn:     model.Eintrag06AbschlussOderBruttolohn,
		},
		{
			Von:                         model.Eintrag07Von,
			Bis:                         model.Eintrag07Bis,
			Name:                        model.Eintrag07Name,
			SchulartOderArtDerTätigkeit: model.Eintrag07SchulartOderArtDerTätigkeit,
			AbschlussOderBruttolohn:     model.Eintrag07AbschlussOderBruttolohn,
		},
		{
			Von:                         model.Eintrag08Von,
			Bis:                         model.Eintrag08Bis,
			Name:                        model.Eintrag08Name,
			SchulartOderArtDerTätigkeit: model.Eintrag08SchulartOderArtDerTätigkeit,
			AbschlussOderBruttolohn:     model.Eintrag08AbschlussOderBruttolohn,
		},
		{
			Von:                         model.Eintrag09Von,
			Bis:                         model.Eintrag09Bis,
			Name:                        model.Eintrag09Name,
			SchulartOderArtDerTätigkeit: model.Eintrag09SchulartOderArtDerTätigkeit,
			AbschlussOderBruttolohn:     model.Eintrag09AbschlussOderBruttolohn,
		},
		{
			Von:                         model.Eintrag10Von,
			Bis:                         model.Eintrag10Bis,
			Name:                        model.Eintrag10Name,
			SchulartOderArtDerTätigkeit: model.Eintrag10SchulartOderArtDerTätigkeit,
			AbschlussOderBruttolohn:     model.Eintrag10AbschlussOderBruttolohn,
		},
		{
			Von:                         model.Eintrag11Von,
			Bis:                         model.Eintrag11Bis,
			Name:                        model.Eintrag11Name,
			SchulartOderArtDerTätigkeit: model.Eintrag11SchulartOderArtDerTätigkeit,
			AbschlussOderBruttolohn:     model.Eintrag11AbschlussOderBruttolohn,
		},
		{
			Von:                         model.Eintrag12Von,
			Bis:                         model.Eintrag12Bis,
			Name:                        model.Eintrag12Name,
			SchulartOderArtDerTätigkeit: model.Eintrag12SchulartOderArtDerTätigkeit,
			AbschlussOderBruttolohn:     model.Eintrag12AbschlussOderBruttolohn,
		},
		{
			Von:                         model.Eintrag13Von,
			Bis:                         model.Eintrag13Bis,
			Name:                        model.Eintrag13Name,
			SchulartOderArtDerTätigkeit: model.Eintrag13SchulartOderArtDerTätigkeit,
			AbschlussOderBruttolohn:     model.Eintrag13AbschlussOderBruttolohn,
		},
		{
			Von:                         model.Eintrag14Von,
			Bis:                         model.Eintrag14Bis,
			Name:                        model.Eintrag14Name,
			SchulartOderArtDerTätigkeit: model.Eintrag14SchulartOderArtDerTätigkeit,
			AbschlussOderBruttolohn:     model.Eintrag14AbschlussOderBruttolohn,
		},
		{
			Von:                         model.Eintrag15Von,
			Bis:                         model.Eintrag15Bis,
			Name:                        model.Eintrag15Name,
			SchulartOderArtDerTätigkeit: model.Eintrag15SchulartOderArtDerTätigkeit,
			AbschlussOderBruttolohn:     model.Eintrag15AbschlussOderBruttolohn,
		},
		{
			Von:                         model.Eintrag16Von,
			Bis:                         model.Eintrag16Bis,
			Name:                        model.Eintrag16Name,
			SchulartOderArtDerTätigkeit: model.Eintrag16SchulartOderArtDerTätigkeit,
			AbschlussOderBruttolohn:     model.Eintrag16AbschlussOderBruttolohn,
		},
	}
}
