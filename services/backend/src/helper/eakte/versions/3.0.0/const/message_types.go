package xdomeaconsts

type MessageType string

const (
	MessageType_Information_Information        MessageType = "0101"
	MessageType_Information_EmpfangBestaetigen MessageType = "0102"

	MessageType_Geschaeftsgang_Geschaeftsgang              MessageType = "0201"
	MessageType_Geschaeftsgang_EmpfangBestaetigen          MessageType = "0202"
	MessageType_Geschaeftsgang_GeaendertenLaufwegMitteilen MessageType = "0203"

	MessageType_Aktenplan_Aktenplan          MessageType = "0301"
	MessageType_Aktenplan_EmpfangBestaetigen MessageType = "0302"

	MessageType_Abgabe_Abgabe             MessageType = "0401"
	MessageType_Abgabe_ImportBestaetigen  MessageType = "0402"
	MessageType_Abgabe_EmpfangBestaetigen MessageType = "0403"

	MessageType_Aussonderung_Anbieteverzeichnis                      MessageType = "0501"
	MessageType_Aussonderung_Bewertungsverzeichnis                   MessageType = "0502"
	MessageType_Aussonderung_Aussonderung                            MessageType = "0503"
	MessageType_Aussonderung_AnbietungEmpfangBestaetigen             MessageType = "0504"
	MessageType_Aussonderung_BewertungEmpfangBestaetigen             MessageType = "0505"
	MessageType_Aussonderung_AussonderungImportBestaetigen           MessageType = "0506"
	MessageType_Aussonderung_AussonderungEmpfangBestaetigen          MessageType = "0507"
	MessageType_Aussonderung_AktenplanZurBewertung                   MessageType = "0511"
	MessageType_Aussonderung_AktenplanZurBewertungEmpfangBestaetigen MessageType = "0512"
	MessageType_Aussonderung_Bewertungskatalog                       MessageType = "0513"
	MessageType_Aussonderung_BewertungskatalogImportBestaetigen      MessageType = "0514"
	MessageType_Aussonderung_BewertungskatalogEmpfangBestaetigen     MessageType = "0515"

	MessageType_FVDaten_DokumentAktualisieren                MessageType = "0601"
	MessageType_FVDaten_SGOAnsehen                           MessageType = "0602"
	MessageType_FVDaten_SGOBearbeiten                        MessageType = "0603"
	MessageType_FVDaten_SGOErstellen                         MessageType = "0604"
	MessageType_FVDaten_SGOAblegen                           MessageType = "0605"
	MessageType_FVDaten_SGODrucken                           MessageType = "0606"
	MessageType_FVDaten_ProtokolleintragErstellen            MessageType = "0607"
	MessageType_FVDaten_SGOSuchen                            MessageType = "0608"
	MessageType_FVDaten_MetadatenAnlegen                     MessageType = "0609"
	MessageType_FVDaten_MetadatenAktualisieren               MessageType = "0610"
	MessageType_FVDaten_MetadatenAnsehen                     MessageType = "0611"
	MessageType_FVDaten_MetadatenLoeschen                    MessageType = "0612"
	MessageType_FVDaten_SGOLoeschenMarkieren                 MessageType = "0613"
	MessageType_FVDaten_SGOLoeschmarkierungAufheben          MessageType = "0614"
	MessageType_FVDaten_SGOEndgueltigLoeschen                MessageType = "0615"
	MessageType_FVDaten_SGOLoeschstatusAbfragen              MessageType = "0616"
	MessageType_FVDaten_SGOUngueltigKennzeichnen             MessageType = "0617"
	MessageType_FVDaten_DatensatzLoeschen                    MessageType = "0618"
	MessageType_FVDaten_BenachrichtigungAbrufen              MessageType = "0619"
	MessageType_FVDaten_VertretungAktivierenOderDeaktivieren MessageType = "0620"
	MessageType_FVDaten_VertretungsstatusAbfragen            MessageType = "0621"
	MessageType_FVDaten_ZustaendigkeitAendern                MessageType = "0622"
	MessageType_FVDaten_GesamtprotokollAblegen               MessageType = "0623"
	MessageType_FVDaten_SGOZDAVerfuegen                      MessageType = "0624"
	MessageType_FVDaten_SystemstatusAbfragen                 MessageType = "0625"
	MessageType_FVDaten_KonfigurationsparameterErstellen     MessageType = "0626"
	MessageType_FVDaten_KonfigurationsparameterAktualisieren MessageType = "0627"
	MessageType_FVDaten_KonfigurationsparameterAbrufen       MessageType = "0628"
	MessageType_FVDaten_EmpfangBestaetigen                   MessageType = "0629"
	MessageType_FVDaten_ImportBestaetigen                    MessageType = "0630"
	MessageType_FVDaten_PrimaerdokumentExportieren           MessageType = "0631"
	MessageType_FVDaten_SGOZDAAufheben                       MessageType = "0632"

	MessageType_Zwischenarchivierung_Auslagerung                         MessageType = "0701"
	MessageType_Zwischenarchivierung_AuslagerungEmpfangBestaetigen       MessageType = "0702"
	MessageType_Zwischenarchivierung_AuslagerungImportBestaetigen        MessageType = "0703"
	MessageType_Zwischenarchivierung_RueckleiheAnforderung               MessageType = "0711"
	MessageType_Zwischenarchivierung_RueckleiheUebergabe                 MessageType = "0712"
	MessageType_Zwischenarchivierung_RueckleiheEmpfangBestaetigen        MessageType = "0713"
	MessageType_Zwischenarchivierung_RueckuebertragungAnforderung        MessageType = "0721"
	MessageType_Zwischenarchivierung_RueckuebertragungUebergabe          MessageType = "0722"
	MessageType_Zwischenarchivierung_RueckuebertragungImportBestaetigen  MessageType = "0723"
	MessageType_Zwischenarchivierung_RueckuebertragungEmpfangBestaetigen MessageType = "0724"
)

func (mtype MessageType) GetCode() string {
	return string(mtype)
}
