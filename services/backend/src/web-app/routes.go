package webapp

import (
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	customhttphandler "github.com/dcssoftware/bafoeg-manager/src/web-app/custom-http-handler"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (a *App) CreateRoutes(router *fiber.App) {

	api := router.Group("/api")
	api.Use(a.MiddlewareHandler.Logger())

	a.setupGeneralRoutes(api)

	apiV1 := api.Group("v1")

	// authentication
	a.setupAuthentication(apiV1)

	// from here only authenticated
	apiV1.Use(a.MiddlewareHandler.Authentication())

	a.setupUserRoutes(apiV1)
	a.setupUserManagementRoutes(apiV1)

	a.setupEaktenRoutes(apiV1)

	a.setupSchoolManagementRoutes(apiV1)
	a.setupRagManagementRoutes(apiV1)

	a.setupOrganizationRoutes(apiV1)

	a.setupApplicantRoutes(apiV1)
	a.setupApplicationRevisionsRoutes(apiV1)

	a.setupApplicantPaymentRoutes(apiV1)
	a.setupApplicationLabelsRoutes(apiV1)

	a.setupApplicationUserRoutes(apiV1)

	a.setupApplicationFilesRoutes(apiV1)
	a.setupApplicationRoutes(apiV1)

	// apiV1Admin.Use(Handle404)
	apiV1.Use(customhttphandler.NotFoundHandler)
	api.Use(customhttphandler.NotFoundHandler)
}

func (a *App) setupGeneralRoutes(api fiber.Router) {
	api.All("/", func(c fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	api.Get("/ping", a.GeneralHandler.Ping)
	api.Get("/system-healthcheck", a.GeneralHandler.SystemHealthCheck)
	api.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	api.Get("/asset/gopher-coffee", a.GeneralHandler.AssetGopherCoffee)
	api.Get("/asset/logo", a.GeneralHandler.AssetLogo)
	api.Get("/asset/logo-branding", a.GeneralHandler.AssetLogoBranding)
	api.Get("/asset/licenses", a.GeneralHandler.AssetLicenses)
}

func (a *App) setupAuthentication(apiV1 fiber.Router) {
	apiV1.Get("/oauth/redirect", a.AuthHandler.CreateRedirect)
	apiV1.Get("/oauth/callback", a.AuthHandler.CallbackHandler)
	apiV1.Get("/auth/logout", a.AuthHandler.Logout)
	apiV1.Get("/auth/refresh", a.MiddlewareHandler.AuthenticationRefreshToken(), a.AuthHandler.Refresh)

	if configuration.Webserver.IsE2ETestMode {
		apiV1.Get("/auth/e2e-test-token", a.AuthHandler.E2ETestToken)
	}
}

func (a *App) setupUserRoutes(apiV1 fiber.Router) {
	apiV1.Get("/self", a.UserHandler.GetSelfInformation)
	apiV1.Get("/self/permissions", a.UserHandler.GetSelfPermissions)
	apiV1.Get("/self/picture", a.UserHandler.GetSelfProfilePicture)
	apiV1.Get("/users/profilpictures/:userid", a.UserHandler.GetProfilePictureByUserID)
}

func (a *App) setupUserManagementRoutes(apiV1 fiber.Router) {
	apiV1.Get("/user-management", a.MiddlewareHandler.PermissionsCheck([]string{"read:user-management"}), a.UserHandler.GetUsers)
}

func (a *App) setupSchoolManagementRoutes(apiV1 fiber.Router) {
	apiV1.Get("/schools/", a.SchoolHandler.GetSchools)
	apiV1.Get("/schools/:schoolID", a.SchoolHandler.GetSchoolByID)

	apiV1.Get("/schools/:schoolID/degrees", a.SchoolHandler.GetSchoolDegreesBySchoolID)
}

func (a *App) setupRagManagementRoutes(apiV1 fiber.Router) {

	viewPermissionSchüler := a.MiddlewareHandler.PermissionsCheck([]string{"read:rag-management-schueler-files"})
	requestPermissionSchüler := a.MiddlewareHandler.PermissionsCheck([]string{"read:rag-management-schueler-files"})
	uploadPermissionSchüler := a.MiddlewareHandler.PermissionsCheck([]string{"read:rag-management-schueler-files", "upload:rag-management-schueler-files"})
	deletePermissionSchüler := a.MiddlewareHandler.PermissionsCheck([]string{"read:rag-management-schueler-files", "delete:rag-management-schueler-files"})

	viewPermissionStudierenden := a.MiddlewareHandler.PermissionsCheck([]string{"read:rag-management-studierenden-files"})
	requestPermissionStudierenden := a.MiddlewareHandler.PermissionsCheck([]string{"read:rag-management-studierenden-files"})
	uploadPermissionStudierenden := a.MiddlewareHandler.PermissionsCheck([]string{"read:rag-management-studierenden-files", "upload:rag-management-studierenden-files"})
	deletePermissionStudierenden := a.MiddlewareHandler.PermissionsCheck([]string{"read:rag-management-studierenden-files", "delete:rag-management-studierenden-files"})

	// get general data from rag configuration
	apiV1.Get("/rag/information", a.RAGHandler.GetRagInformation)

	// interact with the rag (general requests)
	apiV1.Get("/rag/bafoeg/schueler/request", requestPermissionSchüler, a.RAGHandler.GetRAGrequestSchüler)
	apiV1.Get("/rag/bafoeg/studierenden/request", requestPermissionStudierenden, a.RAGHandler.GetRAGrequestStudierende)

	// view file by file id
	apiV1.Get("/rag/bafoeg/schueler/:fileID", viewPermissionSchüler, a.RAGHandler.GetDocumentsSchülerByID)
	apiV1.Get("/rag/bafoeg/studierenden/:fileID", viewPermissionStudierenden, a.RAGHandler.GetDocumentsStudierendenByID)

	apiV1.Delete("/rag/bafoeg/schueler/:fileID", deletePermissionSchüler, a.RAGHandler.DeleteRagFilesSchüler)
	apiV1.Delete("/rag/bafoeg/studierenden/:fileID", deletePermissionStudierenden, a.RAGHandler.DeleteRagFilesStudierenden)

	// insert new conversation session
	apiV1.Put("/rag/bafoeg/schueler", a.RAGHandler.InsertRagConversationSchueler)
	apiV1.Put("/rag/bafoeg/studierenden", a.RAGHandler.InsertRagConversationStudierende)

	// view all available files inside the rag database
	apiV1.Get("/rag/bafoeg/schueler", viewPermissionSchüler, a.RAGHandler.GetDocumentsSchüler)
	apiV1.Get("/rag/bafoeg/studierenden", viewPermissionStudierenden, a.RAGHandler.GetDocumentsStudierenden)

	// upload files
	apiV1.Post("/rag/bafoeg/schueler", uploadPermissionSchüler, a.RAGHandler.UploadRAGrelevantDocumentsSchüler)
	apiV1.Post("/rag/bafoeg/studierenden", uploadPermissionStudierenden, a.RAGHandler.UploadRAGrelevantDocumentsStudierende)
}

func (a *App) setupApplicantRoutes(apiV1 fiber.Router) {
	// applicants
	apiV1.Get("/applications/applicants/", a.ApplicantHandler.GetApplicants)
	apiV1.Put("/applications/applicants", a.ApplicantHandler.CreateApplicant)
	apiV1.Get("/applications/applicants/:id", a.ApplicantHandler.GetApplicant)
	apiV1.Get("/applications/applicants/by-school/:schoolID", a.ApplicantHandler.GetApplicantsBySchoolID)

}

func (a *App) setupApplicantPaymentRoutes(apiV1 fiber.Router) {
	apiV1.Get("/applications/applicants/:id/payments", a.PaymentsHandler.GetPaymentHistoryByApplicantID)
}

func (a *App) setupApplicationRevisionsRoutes(apiV1 fiber.Router) {
	// revision
	apiV1.Get("/applications/:applicationID/revision", a.ApplicationsHandler.GetApplicationRevisionsByApplicationID)
}

func (a *App) setupApplicationFilesRoutes(apiV1 fiber.Router) {
	apiV1.Get("/applications/:applicationID/files", a.ApplicationsHandler.GetApplicationFilesByApplicationID)
	apiV1.Get("/applications/:applicationID/files/:fileID", a.ApplicationsHandler.GetApplicationFileByFileID)
	// apiV1.Delete("/applications/:applicationID/files/:fileID", a.ApplicationsHandler.DeleteApplicationFile)
	apiV1.Post("/applications/:applicationID/uploadFiles", a.ApplicationsHandler.UploadApplicationFile)
}

func (a *App) setupApplicationRoutes(apiV1 fiber.Router) {
	apiV1.Put("/applications/eakte-new-application", a.ApplicationsHandler.CreateApplicationFromEakte)
	apiV1.Get("/applications/by-applicant-id/:applicantID", a.ApplicationsHandler.GetApplicationsByApplicantID)
	apiV1.Get("/applications/metrics", a.ApplicationsHandler.GetApplicationsMetrics)
	apiV1.Get("/applications/", a.ApplicationsHandler.GetApplications)
	apiV1.Patch("/applications/:applicationID/change-assigned-user/:userID", a.ApplicationsHandler.UpdateApplicationAssignedUser)
	apiV1.Patch("/applications/:applicationID/change-status/:status", a.ApplicationsHandler.UpdateApplicationStatus)
	apiV1.Patch("/applications/:applicationID/change-assigned-school-degree/:schoolDegreeID", a.ApplicationsHandler.UpdateApplicationAssignedSchoolDegree)
	apiV1.Patch("/applications/:applicationID/change-assigned-applicant/:applicantID", a.ApplicationsHandler.UpdateApplicationAssignedApplicant)
	apiV1.Get("/applications/:applicationID/is-updatable", a.ApplicationsHandler.IsApplicationUpdatableByApplicationID)
	apiV1.Get("/applications/:applicationID", a.ApplicationsHandler.GetApplicationByID)
}

func (a *App) setupApplicationLabelsRoutes(apiV1 fiber.Router) {
	apiV1.Get("/applications/labels", a.ApplicationLabelsHandler.GetApplicationLabels)
}

func (a *App) setupApplicationUserRoutes(apiV1 fiber.Router) {
	apiV1.Get("/applications/assignable-users/", a.UserHandler.GetUserSelection)
}

func (a *App) setupEaktenRoutes(apiV1 fiber.Router) {
	apiV1.Post("/eakten/documents/upload", a.EakteHandler.UploadEakte)
	apiV1.Get("/eakten/documents/:documentID", a.EakteHandler.GetDocumentFileByID)
	apiV1.Get("/eakten/", a.EakteHandler.GetEakten)
	apiV1.Get("/eakten/:eakteID", a.EakteHandler.GetEakteByID)
	apiV1.Get("/eakten/:eakteID/get-application-mapping", a.EakteHandler.GetEaktenApplicationByEakteID)
	apiV1.Get("/eakten/:eakteID/vorgang", a.EakteHandler.GetVorgänge)
	apiV1.Get("/eakten/:eakteID/documents", a.EakteHandler.GetDocumentFilesByEakteID)
	apiV1.Get("/eakten/:eakteID/vorgang/:vorgangsID/documents", a.EakteHandler.GetDocumentFilesByVorgangsID)
}

func (a *App) setupOrganizationRoutes(apiV1 fiber.Router) {
	apiV1.Get("/organization/regions", a.OrganizationHandler.GetRegions)
	apiV1.Get("/organization/behoerden/:regionID", a.OrganizationHandler.GetBehörden)
	apiV1.Get("/organization/abteilungen/:behoerdeID", a.OrganizationHandler.GetAbteilungen)
}
