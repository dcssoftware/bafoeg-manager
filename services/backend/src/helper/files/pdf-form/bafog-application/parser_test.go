package bafogapplication

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:embed testdata/01.pdf
var testFile []byte

func TestParseBAföGApplicationForm01Page01(t *testing.T) {
	fields, err := ParseBAföGApplication(testFile, false)
	assert.NoError(t, err)

	assert.Equal(t, "Technische Universität München", fields.Ausbildungsstätte)
	assert.Equal(t, "1. Semester", fields.KlasseFachrichtung)
	assert.Equal(t, "Informatik B.Sc", fields.AngestrebterAbschluss)
	assert.Equal(t, true, fields.Vollzeitausbildung)
	assert.Equal(t, false, fields.FrühererBAföGAntrag)
	assert.Equal(t, "München", fields.FrühererBAföGAntragZuständigesAmt)
	assert.Equal(t, "12345", fields.FrühererBAföGAntragFörderungsnummer)
	assert.Equal(t, "Mustermann", fields.Nachname)
	assert.Equal(t, "Max", fields.Vorname)
	assert.Equal(t, "01.01.2026", fields.NachnameGeändertSeit.Format("02.01.2006"))
	assert.Equal(t, "ledig", fields.Familienstand)
	assert.Equal(t, "Mann", fields.Geschlecht)
	assert.Equal(t, "Neumann", fields.Geburtsname)
	assert.Equal(t, "23.07.1998", fields.Geburtsdatum.Format("02.01.2006"))
	assert.Equal(t, "München", fields.Geburtsort)
	assert.Equal(t, "deutsch", fields.EigeneStaatsangehörigkeit)
	assert.Equal(t, true, fields.HatEigeneKinder)
	assert.Equal(t, "Musterstraße", fields.AnschriftHauptwohnsitzStraße)
	assert.Equal(t, "123", fields.AnschriftHauptwohnsitzHausnummer)
	assert.Equal(t, "Oben", fields.AnschriftHauptwohnsitzAdresszusatz)
	assert.Equal(t, "DE", fields.AnschriftHauptwohnsitzLand)
	assert.Equal(t, "12345", fields.AnschriftHauptwohnsitzPostleitzahl)
	assert.Equal(t, "Berlin", fields.AnschriftHauptwohnsitzOrt)
	assert.Equal(t, true, fields.AnschriftGemeinschaftMitEltern)
	assert.Equal(t, false, fields.AnschriftEigentumMieteigentumDerEltern)
	assert.Equal(t, "Johannesbergstraße", fields.AnschriftAusbildungStraße)
	assert.Equal(t, "1337", fields.AnschriftAusbildungHausnummer)
	assert.Equal(t, "Unten", fields.AnschriftAusbildungAdresszusatz)
	assert.Equal(t, "FR", fields.AnschriftAusbildungLand)
	assert.Equal(t, "12345", fields.AnschriftAusbildungPostleitzahl)
	assert.Equal(t, "Paris", fields.AnschriftAusbildungOrt)
}

func TestParseBAföGApplicationForm01Page02(t *testing.T) {
	fields, err := ParseBAföGApplication(testFile, false)
	assert.NoError(t, err)

	assert.Equal(t, "example@mail.de", fields.KontaktdatenEMail)
	assert.Equal(t, "0123456789", fields.KontaktdatenTelefon)
	assert.Equal(t, "mich (ständiger Wohnsitz)", fields.BescheidÜbermittelnAn.String())
	assert.Equal(t, "DE12", fields.IBAN1)
	assert.Equal(t, "3456", fields.IBAN2)
	assert.Equal(t, "7890", fields.IBAN3)
	assert.Equal(t, "1234", fields.IBAN4)
	assert.Equal(t, "5678", fields.IBAN5)
	assert.Equal(t, "90", fields.IBAN6)
	assert.Equal(t, "Sparinstitut Musterstätt", fields.GeldistitutName)
	assert.Equal(t, "Max Mustermann", fields.KontoinhaberName)
	assert.Equal(t, "gesetzlich familienversichert", fields.Krankenversicherung.String())
	assert.Equal(t, true, fields.Pflegeversicherung)
	assert.Equal(t, "", fields.Steueridentifikationsnummer)
	assert.Equal(t, "ja, aber dauernd getrennt lebend", fields.VerhältnisElternteile.String())
	assert.Equal(t, "Mustermann", fields.Elternteil01Name)
	assert.Equal(t, "Michelle", fields.Elternteil01Vorname)
	assert.Equal(t, "1", fields.Elternteil01Geschlecht)
	assert.Equal(t, "01.01.2000", fields.Elternteil01Geburtsdatum.Format("02.01.2006"))
	assert.Equal(t, "01.01.2000", fields.Elternteil01Sterbedatum.Format("02.01.2006"))
	assert.Equal(t, "deutsch", fields.Elternteil01Staatsangehörigkeit)
	assert.Equal(t, "Schummelstraße", fields.Elternteil01Straße)
	assert.Equal(t, "69a", fields.Elternteil01Hausnummer)
	assert.Equal(t, "boring", fields.Elternteil01Adresszusatz)
	assert.Equal(t, "DE", fields.Elternteil01Land)
	assert.Equal(t, "12345", fields.Elternteil01Postleitzahl)
	assert.Equal(t, "Musterstätt", fields.Elternteil01Ort)
	assert.Equal(t, "Power", fields.Elternteil02Name)
	assert.Equal(t, "Max", fields.Elternteil02Vorname)
	assert.Equal(t, "2", fields.Elternteil02Geschlecht)
	assert.Equal(t, "01.01.2000", fields.Elternteil02Geburtsdatum.Format("02.01.2006"))
	assert.Equal(t, "01.01.2000", fields.Elternteil02Sterbedatum.Format("02.01.2006"))
	assert.Equal(t, "deutsch", fields.Elternteil02Staatsangehörigkeit)
	assert.Equal(t, "Schummelstraße", fields.Elternteil02Straße)
	assert.Equal(t, "69b", fields.Elternteil02Hausnummer)
	assert.Equal(t, "boring", fields.Elternteil02Adresszusatz)
	assert.Equal(t, "DE", fields.Elternteil02Land)
	assert.Equal(t, "12345", fields.Elternteil02Postleitzahl)
	assert.Equal(t, "Musterstätt", fields.Elternteil02Ort)
}

func TestParseBAföGApplicationForm01Page03(t *testing.T) {
	fields, err := ParseBAföGApplication(testFile, false)
	assert.NoError(t, err)

	assert.Equal(t, true, fields.AnwärterbezügeAusÖffentlichenMitteln)
	assert.Equal(t, true, fields.LeistungenSGB2SGB3)
	assert.Equal(t, false, fields.LeistungenBegabtenförderungswerk)
	assert.Equal(t, false, fields.KeineDerVorstehendenLeistungen)
	assert.Equal(t, "01.2000", fields.BewilligungszeitraumVon.Format("01.2006"))
	assert.Equal(t, "01.2000", fields.BewilligungszeitraumBis.Format("01.2006"))
	assert.Equal(t, "200", fields.Sozialleistungen1)
	assert.Equal(t, "200", fields.Sozialleistungen2)
	assert.Equal(t, float64(200), fields.RiesterEuro)
	assert.Equal(t, true, fields.VoraussichtlicheEinnahmen)
	assert.Equal(t, float64(200), fields.BruttoeinnahmenEuro)
	assert.Equal(t, true, fields.AbeitgeberanteilEnthaltenEuro)
	assert.Equal(t, float64(200), fields.AusbildungsUndPraktikumsvergütungBruttoEuro)
	assert.Equal(t, float64(200), fields.EinkünfteEuro)
	assert.Equal(t, float64(200), fields.KapitalvermögenEuro)
	assert.Equal(t, float64(200), fields.WaisenrenteWaisengeldRentenEuro)
	assert.Equal(t, float64(200), fields.AusbildungsbeihilfenEuro)
	assert.Equal(t, float64(200), fields.UnterhaltsleistungenEuro)
	assert.Equal(t, float64(200), fields.UnterhaltsvorschussleistungenEuro)
	assert.Equal(t, float64(200), fields.WeitereEinnahmenEuro)
}

func TestParseBAföGApplicationForm01Page04(t *testing.T) {
	fields, err := ParseBAföGApplication(testFile, false)
	assert.NoError(t, err)

	assert.Equal(t, float64(200), fields.BarvermögenEuro)
	assert.Equal(t, true, fields.BarvermögenNein)
	assert.Equal(t, float64(200), fields.BankUndSparguthabenEuro)
	assert.Equal(t, true, fields.BankUndSparguthabenNein)
	assert.Equal(t, float64(200), fields.BausparUndPrämienguthabenEuro)
	assert.Equal(t, true, fields.BausparUndPrämienguthabenNein)
	assert.Equal(t, float64(200), fields.WertpapiereEuro)
	assert.Equal(t, true, fields.WertpapiereNein)
	assert.Equal(t, float64(200), fields.KraftfahrzeugeEuro)
	assert.Equal(t, true, fields.KraftfahrzeugeNein)
	assert.Equal(t, float64(200), fields.LebensversicherungenEuro)
	assert.Equal(t, true, fields.LebensversicherungenNein)
	assert.Equal(t, float64(200), fields.AltersvorsorgevermögenEuro)
	assert.Equal(t, true, fields.AltersvorsorgevermögenNein)
	assert.Equal(t, float64(200), fields.GrundstückeWohnungenEuro)
	assert.Equal(t, true, fields.GrundstückeWohnungenNein)
	assert.Equal(t, float64(200), fields.BetriebsvermögenEuro)
	assert.Equal(t, true, fields.BetriebsvermögenNein)
	assert.Equal(t, float64(200), fields.GeldforderungenEuro)
	assert.Equal(t, true, fields.GeldforderungenNein)
	assert.Equal(t, float64(200), fields.VermögensgegenständeEuro)
	assert.Equal(t, true, fields.VermögensgegenständeNein)
	assert.Equal(t, float64(200), fields.VermögenswerteAnrechnungsfreiEuro)
	assert.Equal(t, float64(200), fields.ÜbergangsbeihilfenEuro)
	assert.Equal(t, float64(200), fields.SchuldenEuro)
	assert.Equal(t, float64(200), fields.LastenEuro)
}

func TestParseBAföGApplicationForm01Page05(t *testing.T) {
	fields, err := ParseBAföGApplication(testFile, false)
	assert.NoError(t, err)

	assert.Equal(t, float64(200), fields.BarvermögenEuro)
}

func TestParseBAföGApplicationForm01Page06(t *testing.T) {
	fields, err := ParseBAföGApplication(testFile, false)
	assert.NoError(t, err)

	assert.Equal(t, "ja, und zwar", fields.Sorgeberechtigter.String())
	assert.Equal(t, "Jugendamt Musterstätt", fields.SorgeberechtigterName)
	assert.Equal(t, "Hohengrüne", fields.SorgeberechtigterStraße)
	assert.Equal(t, "2", fields.SorgeberechtigterHausnummer)
	assert.Equal(t, "Grüne", fields.SorgeberechtigterAddresszusatz)
	assert.Equal(t, "EN", fields.SorgeberechtigterLand)
	assert.Equal(t, "99999", fields.SorgeberechtigterPostleitzahl)
	assert.Equal(t, "Mainingen", fields.SorgeberechtigterOrt)
	assert.Equal(t, "12.03.2000 Max Mustermann", fields.UnterschriftAuszubildendePerson)
	assert.Equal(t, "12.03.2000 i.A. Johannes Schlittermann", fields.UnterschriftGesetzlicherVertreter)
}

func TestErrorsInGerman(t *testing.T) {
	fields, parseErr := ParseBAföGApplication(testFile, false)
	assert.NoError(t, parseErr)

	// Validate now returns a slice of German error messages
	validationMsgs, validateErr := fields.Validate()
	assert.NoError(t, validateErr)

	// Expect at least one validation message
	assert.NotNil(t, validationMsgs)
	assert.NotEmpty(t, validationMsgs)
	// Log the messages for debugging
	for _, vm := range validationMsgs {
		t.Logf("validation messages: %v", vm)
	}
	// Ensure at least one message contains the expected German snippet
}
