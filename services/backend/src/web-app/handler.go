package webapp

import "github.com/gofiber/fiber/v3"

type MiddlewareHandler interface {
	Authentication() func(c fiber.Ctx) error
	AuthenticationRefreshToken() func(c fiber.Ctx) error
	PermissionsCheck(permissions []string) func(c fiber.Ctx) error
	Logger() func(c fiber.Ctx) error
}

type ApplicationsHandler interface {
	GetApplicationByID(c fiber.Ctx) error
	GetApplications(c fiber.Ctx) error
	GetApplicationsByApplicantID(c fiber.Ctx) error
	GetApplicationRevisionsByApplicationID(c fiber.Ctx) error
	GetApplicationsMetrics(c fiber.Ctx) error
	GetApplicationFilesByApplicationID(c fiber.Ctx) error
	GetApplicationFileByFileID(c fiber.Ctx) error
	UploadApplicationFile(c fiber.Ctx) error
	UpdateApplicationStatus(c fiber.Ctx) error
	UpdateApplicationAssignedUser(c fiber.Ctx) error
	DeleteApplicationFile(c fiber.Ctx) error
	UpdateApplicationAssignedSchoolDegree(c fiber.Ctx) error
	IsApplicationUpdatableByApplicationID(c fiber.Ctx) error
	UpdateApplicationAssignedApplicant(c fiber.Ctx) error
	CreateApplicationFromEakte(c fiber.Ctx) error
}

type ApplicationLabelsHandler interface {
	GetApplicationLabels(c fiber.Ctx) error
}

type AuthHandler interface {
	CreateRedirect(c fiber.Ctx) error
	CallbackHandler(c fiber.Ctx) error
	Logout(c fiber.Ctx) error
	Refresh(c fiber.Ctx) error
	E2ETestToken(c fiber.Ctx) error
}

type ApplicantsHandler interface {
	GetApplicant(c fiber.Ctx) error
	GetApplicants(c fiber.Ctx) error
	GetApplicantsBySchoolID(c fiber.Ctx) error
	CreateApplicant(c fiber.Ctx) error
}

type EakteHandler interface {
	GetEakten(c fiber.Ctx) error
	GetEakteByID(c fiber.Ctx) error
	GetVorgänge(c fiber.Ctx) error
	GetDocumentFilesByEakteID(c fiber.Ctx) error
	GetDocumentFilesByVorgangsID(c fiber.Ctx) error
	GetEaktenApplicationByEakteID(c fiber.Ctx) error
	GetDocumentFileByID(c fiber.Ctx) error
	UploadEakte(c fiber.Ctx) error
}

type GeneralHandler interface {
	Ping(c fiber.Ctx) error
	SystemHealthCheck(c fiber.Ctx) error
	AssetGopherCoffee(c fiber.Ctx) error
	AssetLogo(c fiber.Ctx) error
	AssetLogoBranding(c fiber.Ctx) error
	AssetLicenses(c fiber.Ctx) error
}

type OrganizationHandler interface {
	GetAbteilungen(c fiber.Ctx) error
	GetBehörden(c fiber.Ctx) error
	GetRegions(c fiber.Ctx) error
}

type RAGHandler interface {
	GetRagInformation(c fiber.Ctx) error
	GetRAGrequestSchüler(c fiber.Ctx) error
	GetRAGrequestStudierende(c fiber.Ctx) error
	GetDocumentsSchüler(c fiber.Ctx) error
	GetDocumentsSchülerByID(c fiber.Ctx) error
	GetDocumentsStudierenden(c fiber.Ctx) error
	GetDocumentsStudierendenByID(c fiber.Ctx) error

	UploadRAGrelevantDocumentsSchüler(c fiber.Ctx) error
	UploadRAGrelevantDocumentsStudierende(c fiber.Ctx) error

	InsertRagConversationSchueler(c fiber.Ctx) error
	InsertRagConversationStudierende(c fiber.Ctx) error

	DeleteRagFilesSchüler(c fiber.Ctx) error
	DeleteRagFilesStudierenden(c fiber.Ctx) error
}

type SchoolHandler interface {
	GetSchools(c fiber.Ctx) error
	GetSchoolByID(c fiber.Ctx) error

	GetSchoolDegreesBySchoolID(c fiber.Ctx) error
}

type UserHandler interface {
	GetSelfPermissions(c fiber.Ctx) error
	GetSelfInformation(c fiber.Ctx) error
	GetUsers(c fiber.Ctx) error
	GetUserSelection(c fiber.Ctx) error
	GetSelfProfilePicture(c fiber.Ctx) error
	GetProfilePictureByUserID(c fiber.Ctx) error
}

type PaymentsHandler interface {
	GetPaymentHistoryByApplicantID(c fiber.Ctx) error
}
