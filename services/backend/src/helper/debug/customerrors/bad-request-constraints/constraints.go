package custombadrequestconstraints

type BadRequestContraint struct {
	KeyName           string
	ConstraintMessage string
}

var (
	BadRequest_AuthenticationProviderInvalid  *BadRequestContraint = &BadRequestContraint{KeyName: "provider", ConstraintMessage: "The provider is a enum and has to be one of our supported list."}
	BadRequest_InformationMismatch            *BadRequestContraint = &BadRequestContraint{KeyName: "information_missmatch", ConstraintMessage: "The given information do not match the given data"}
	BadRequest_FileInfected                   *BadRequestContraint = &BadRequestContraint{KeyName: "file_infected", ConstraintMessage: "The given file contains a known virus substring"}
	BadRequest_ApplicationStatusInvalid       *BadRequestContraint = &BadRequestContraint{KeyName: "application_status", ConstraintMessage: "The status could not be identified"}
	BadRequest_ApplicationStatusChangeInvalid *BadRequestContraint = &BadRequestContraint{KeyName: "application_status_change", ConstraintMessage: "The status change violates constraints"}
	BadRequest_IdNotUUID                      *BadRequestContraint = &BadRequestContraint{KeyName: "id", ConstraintMessage: "The provided id is not a valid UUID"}
)
