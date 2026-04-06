package xdomeaconsts

type SchriftgutObjektTyp string

const (
	SchriftgutObjektTyp_Akte     SchriftgutObjektTyp = "001"
	SchriftgutObjektTyp_Dokument SchriftgutObjektTyp = "002"
	SchriftgutObjektTyp_Vorgang  SchriftgutObjektTyp = "003"

	// ?? in der V4 Doku erwähnt, aber noch nicht ausreichend geklärt
	// SchriftgutObjektTyp_Geschäftsgang SchriftgutObjektTyp = ""
	// SchriftgutObjektTyp_Schriftstück  SchriftgutObjektTyp = ""
)
