package webapp

import (
	"errors"
	"fmt"
	"log"

	"github.com/go-sqlx/sqlx"
	"github.com/gofiber/fiber/v3"
	"github.com/minio/minio-go/v7"

	_ "github.com/dcssoftware/bafoeg-manager/swagger-docs"
	swagger "github.com/gofiber/contrib/v3/swaggo"

	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	dbHelper "github.com/dcssoftware/bafoeg-manager/src/helper/database"
	"github.com/dcssoftware/bafoeg-manager/src/helper/debug/logger"
	s3Helper "github.com/dcssoftware/bafoeg-manager/src/helper/s3-bucket"
	swaggercss "github.com/dcssoftware/bafoeg-manager/src/static/assets/swagger-css"

	customhttphandler "github.com/dcssoftware/bafoeg-manager/src/web-app/custom-http-handler"
	middlewareHttp "github.com/dcssoftware/bafoeg-manager/src/web-app/middlewares/http"
	middlewareService "github.com/dcssoftware/bafoeg-manager/src/web-app/middlewares/service"

	applicationsHttp "github.com/dcssoftware/bafoeg-manager/src/resources/applications/http"
	applicationsService "github.com/dcssoftware/bafoeg-manager/src/resources/applications/service"
	applicationsStorage "github.com/dcssoftware/bafoeg-manager/src/resources/applications/storage"

	applicationLabelsHttp "github.com/dcssoftware/bafoeg-manager/src/resources/application-labels/http"
	applicationLabelsService "github.com/dcssoftware/bafoeg-manager/src/resources/application-labels/service"
	applicationLabelsStorage "github.com/dcssoftware/bafoeg-manager/src/resources/application-labels/storage"

	authHttp "github.com/dcssoftware/bafoeg-manager/src/resources/auth/http"
	authService "github.com/dcssoftware/bafoeg-manager/src/resources/auth/service"
	authStorage "github.com/dcssoftware/bafoeg-manager/src/resources/auth/storage"

	eakteHttp "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/http"
	eakteService "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/service"
	eakteStorage "github.com/dcssoftware/bafoeg-manager/src/resources/eakte/storage"

	applicantHttp "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/http"
	applicantService "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/service"
	applicantStorage "github.com/dcssoftware/bafoeg-manager/src/resources/applicants/storage"

	filesService "github.com/dcssoftware/bafoeg-manager/src/resources/files/service"
	filesStorage "github.com/dcssoftware/bafoeg-manager/src/resources/files/storage"
	filesStorageS3 "github.com/dcssoftware/bafoeg-manager/src/resources/files/storage-s3"

	generalHttp "github.com/dcssoftware/bafoeg-manager/src/resources/general/http"
	generalService "github.com/dcssoftware/bafoeg-manager/src/resources/general/service"
	generalStorage "github.com/dcssoftware/bafoeg-manager/src/resources/general/storage"

	organizationHttp "github.com/dcssoftware/bafoeg-manager/src/resources/organization/http"
	organizationService "github.com/dcssoftware/bafoeg-manager/src/resources/organization/service"
	organizationStorage "github.com/dcssoftware/bafoeg-manager/src/resources/organization/storage"

	userHttp "github.com/dcssoftware/bafoeg-manager/src/resources/user/http"
	userService "github.com/dcssoftware/bafoeg-manager/src/resources/user/service"
	userStorage "github.com/dcssoftware/bafoeg-manager/src/resources/user/storage"
	userStorageS3 "github.com/dcssoftware/bafoeg-manager/src/resources/user/storage-s3"

	schoolHttp "github.com/dcssoftware/bafoeg-manager/src/resources/schools/http"
	schoolService "github.com/dcssoftware/bafoeg-manager/src/resources/schools/service"
	schoolStorage "github.com/dcssoftware/bafoeg-manager/src/resources/schools/storage"

	paymentsHttp "github.com/dcssoftware/bafoeg-manager/src/resources/payments/http"
	paymentsService "github.com/dcssoftware/bafoeg-manager/src/resources/payments/service"
	paymentsStorage "github.com/dcssoftware/bafoeg-manager/src/resources/payments/storage"

	ragHttp "github.com/dcssoftware/bafoeg-manager/src/resources/rag/http"
	ragService "github.com/dcssoftware/bafoeg-manager/src/resources/rag/service"
	ragStorage "github.com/dcssoftware/bafoeg-manager/src/resources/rag/storage"
	ragStorageS3 "github.com/dcssoftware/bafoeg-manager/src/resources/rag/storage-s3"

	cronjobs "github.com/dcssoftware/bafoeg-manager/src/cronjobs"
)

var (
	ErrCannotConnectToDatabase              = errors.New("could not connect to database")
	ErrCannotCreateDatabaseForUnknownReason = errors.New("cannot create the database for unknown reasons")

	ErrCannotConnectToS3Bucket = errors.New("could not connect to s3 bucket")
)

type App struct {
	StorageProviderPostgres *sqlx.DB
	StorageProviderMinio    *minio.Client

	MiddlewareHandler        MiddlewareHandler
	ApplicationsHandler      ApplicationsHandler
	AuthHandler              AuthHandler
	ApplicantHandler         ApplicantsHandler
	ApplicationLabelsHandler ApplicationLabelsHandler
	EakteHandler             EakteHandler
	GeneralHandler           GeneralHandler
	SchoolHandler            SchoolHandler
	UserHandler              UserHandler
	PaymentsHandler          PaymentsHandler
	RAGHandler               RAGHandler
	OrganizationHandler      OrganizationHandler
}

func NewApp() *App {
	services, servicesErr := SetupServices()
	if servicesErr != nil {
		panic(servicesErr)
	}

	registerErr := services.Cronjobs.RegisterRAGVectorProcessor()
	if registerErr != nil {
		panic(registerErr)
	}

	go func() {
		services.Cronjobs.StartCronjobs()
	}()

	return convertServicesToApp(services)
}

func NewIntegrationTestApp(databaseName uint64) *AppServices {
	services, servicesErr := SetupServicesForIntegrationTest(databaseName)
	if servicesErr != nil {
		panic(servicesErr)
	}

	return services
}

func convertServicesToApp(services *AppServices) *App {
	return &App{
		MiddlewareHandler: services.MiddlewareHandler,

		StorageProviderPostgres: services.StorageProviderPostgres,
		StorageProviderMinio:    services.StorageProviderMinio,

		ApplicationsHandler:      services.ApplicationsHandler,
		AuthHandler:              services.AuthHandler,
		ApplicantHandler:         services.ApplicantHandler,
		ApplicationLabelsHandler: services.ApplicationLabelsHandler,
		EakteHandler:             services.EakteHandler,
		GeneralHandler:           services.GeneralHandler,
		SchoolHandler:            services.SchoolHandler,
		UserHandler:              services.UserHandler,
		PaymentsHandler:          services.PaymentHandler,
		RAGHandler:               services.RAGHandler,
		OrganizationHandler:      services.OrganizationHandler,
	}
}

type AppServices struct {
	StorageProviderPostgres *sqlx.DB
	StorageProviderMinio    *minio.Client

	ApplicationsHandler *applicationsHttp.ApplicationsHandler
	ApplicationsService *applicationsService.ApplicationsService
	ApplicationsStore   *applicationsStorage.ApplicationsStorage

	AuthHandler *authHttp.AuthHandler
	AuthService *authService.AuthService
	AuthStore   *authStorage.AuthStore

	ApplicantHandler *applicantHttp.ApplicantHandler
	ApplicantSvc     *applicantService.ApplicantService
	ApplicantStore   *applicantStorage.ApplicantStorage

	ApplicationLabelsHandler *applicationLabelsHttp.ApplicationLabelsHandler
	ApplicationLabelsSvc     *applicationLabelsService.ApplicationLabelsService
	ApplicationLabelsStore   *applicationLabelsStorage.ApplicationLabelsStorage

	EakteHandler *eakteHttp.EakteHandler
	EakteSvc     *eakteService.EakteService
	EakteStore   *eakteStorage.EakteStorage

	FileSvc     *filesService.FileService
	FileStore   *filesStorage.FileStorage
	FileStoreS3 *filesStorageS3.FilesStorageS3

	GeneralHandler *generalHttp.GeneralHandler
	GeneralSvc     *generalService.GeneralSvc
	GeneralStore   *generalStorage.GeneralStore

	MiddlewareHandler *middlewareHttp.MiddlewareHandler
	MiddlewareSvc     *middlewareService.MiddlewareService

	OrganizationHandler *organizationHttp.OrganizationHandler
	OrganizationSvc     *organizationService.OrganizationService
	OrganizationStore   *organizationStorage.OrganizationStorage

	SchoolHandler *schoolHttp.SchoolHandler
	SchoolService *schoolService.SchoolService
	SchoolStorage *schoolStorage.SchoolStorage

	PaymentHandler *paymentsHttp.PaymentHandler
	PaymentService *paymentsService.PaymentService
	PaymentStorage *paymentsStorage.PaymentStorage

	UserHandler *userHttp.UserHandler
	UserSvc     *userService.UserService
	UserStore   *userStorage.UserStore

	PaymentsHandler *paymentsHttp.PaymentHandler
	PaymentsSvc     *paymentsService.PaymentService
	PaymentsStore   *paymentsStorage.PaymentStorage

	RAGHandler *ragHttp.RAGHandler
	RAGSvc     *ragService.RAGService
	RAGStore   *ragStorage.RAGStorage
	RAGStoreS3 *ragStorageS3.RAGStorageS3

	Cronjobs *cronjobs.Cronjob
}

/*
*
*	Why is there a setup services function?
*
*	Because for the integration tests, they shouldn't access the service functions
*	but the reality is, some functions are needed, like generating a jwt session for the request
*	so use this opportunity as a cautious gift, but not as a privilege
*
 */
func SetupServices() (*AppServices, error) {
	dbConn, dbConnErr := dbHelper.CreateDatabaseConnectionByConfig()
	if dbConn == nil || dbConn.DB == nil || dbConnErr != nil {
		log.Println(ErrCannotConnectToDatabase)

		databaseErr := ErrCannotCreateDatabaseForUnknownReason
		return nil, databaseErr
	}

	s3Conn, s3ConnErr := s3Helper.CreateS3BucketConnection()
	if s3Conn == nil || s3ConnErr != nil {
		log.Println(ErrCannotConnectToS3Bucket)
		return nil, s3ConnErr
	}

	return setupServicesCreateAppServices(dbConn, s3Conn)
}

func SetupServicesForIntegrationTest(testHash uint64) (*AppServices, error) {
	dbConn, dbConnErr := dbHelper.CreateDatabaseConnectionByConfigWithIntegrationTestID(testHash)
	if dbConn == nil || dbConn.DB == nil || dbConnErr != nil {
		log.Println(ErrCannotConnectToDatabase)

		databaseErr := ErrCannotCreateDatabaseForUnknownReason
		return nil, databaseErr
	}

	s3Conn, s3ConnErr := s3Helper.CreateS3BucketConnection()
	if s3Conn == nil || s3ConnErr != nil {
		log.Println(ErrCannotConnectToS3Bucket)
		return nil, s3ConnErr
	}

	return setupServicesCreateAppServices(dbConn, s3Conn)
}

func setupServicesCreateAppServices(dbConn *sqlx.DB, s3Conn *minio.Client) (*AppServices, error) {
	logger.NewLogger()

	filesStore := filesStorage.NewFileStorage(dbConn)
	filesStoreS3 := filesStorageS3.NewFilesStorageS3(s3Conn)
	filesSvc := filesService.NewFileService(filesStore, filesStoreS3)

	paymentsStore := paymentsStorage.NewPaymentStorage(dbConn)
	paymentsService := paymentsService.NewPaymentService(paymentsStore)
	paymentsHandler := paymentsHttp.NewPaymentHandler(paymentsService)

	applicantStore := applicantStorage.NewApplicantStorage(dbConn)
	applicantSvc := applicantService.NewApplicantService(applicantStore, paymentsService)
	applicantHandler := applicantHttp.NewApplicantHandler(applicantSvc)

	schoolStore := schoolStorage.NewSchoolStorage(dbConn)
	schoolSvc := schoolService.NewSchoolService(schoolStore)
	schoolHandler := schoolHttp.NewSchoolHandler(schoolSvc)

	applicationsStore := applicationsStorage.NewApplicationsStorage(dbConn)
	applicationsSvc := applicationsService.NewApplicationsService(applicationsStore, applicantSvc, schoolSvc, filesSvc)
	applicationsHandler := applicationsHttp.NewApplicationsHandler(applicationsSvc)

	applicationLabelStore := applicationLabelsStorage.NewApplicationLabelsStorage(dbConn)
	applicationLabelsSvc := applicationLabelsService.NewApplicationLabelsService(applicationLabelStore)
	applicationLabelsHandler := applicationLabelsHttp.NewApplicationLabelsHandler(applicationLabelsSvc)

	eakteStore := eakteStorage.NewEakteStorage(dbConn)
	eakteSvc := eakteService.NewEakteService(eakteStore, filesSvc)
	eakteHandler := eakteHttp.NewEakteHandler(eakteSvc)

	generalStore := generalStorage.NewGeneralStore(dbConn)
	generalSvc := generalService.NewGeneralSvc(generalStore)
	generalHandler := generalHttp.NewGeneralHandler(generalSvc)

	organizationStore := organizationStorage.NewOrganizationStorage(dbConn)
	organizationSvc := organizationService.NewOrganizationService(organizationStore)
	organizationHandler := organizationHttp.NewOrganizationHandler(organizationSvc)

	userStore := userStorage.NewUserStore(dbConn, s3Conn)
	userStoreS3 := userStorageS3.NewUserStoreS3(s3Conn)
	userSvc := userService.NewUserService(userStore, userStoreS3, filesSvc)
	userHandler := userHttp.NewUserHandler(userSvc)

	authStore := authStorage.NewAuthStore(dbConn)
	authService := authService.NewAuthService(authStore, userSvc)
	authHandler := authHttp.NewAuthHandler(authService)

	middlewareSvc := middlewareService.NewMiddlewareService(userSvc)
	middlewareHandler := middlewareHttp.NewMiddlewareHandler(middlewareSvc)

	ragStore := ragStorage.NewRAGStorage(dbConn)
	ragStoreS3 := ragStorageS3.NewRAGStorageS3(s3Conn)
	ragSvc := ragService.NewRAGService(ragStore, ragStoreS3)
	ragHandler := ragHttp.NewRAGHandler(ragSvc)

	cronjobInstance, cronjobInstanceErr := cronjobs.NewCronjob(dbConn, s3Conn)
	if cronjobInstanceErr != nil {
		panic(cronjobInstanceErr)
	}

	return &AppServices{
		StorageProviderPostgres: dbConn,
		StorageProviderMinio:    s3Conn,

		ApplicationsHandler: applicationsHandler,
		ApplicationsService: applicationsSvc,
		ApplicationsStore:   applicationsStore,

		AuthHandler: authHandler,
		AuthService: authService,
		AuthStore:   authStore,

		ApplicantHandler: applicantHandler,
		ApplicantSvc:     applicantSvc,
		ApplicantStore:   applicantStore,

		ApplicationLabelsHandler: applicationLabelsHandler,
		ApplicationLabelsSvc:     applicationLabelsSvc,
		ApplicationLabelsStore:   applicationLabelStore,

		EakteHandler: eakteHandler,
		EakteSvc:     eakteSvc,
		EakteStore:   eakteStore,

		FileSvc:     filesSvc,
		FileStore:   filesStore,
		FileStoreS3: filesStoreS3,

		GeneralHandler: generalHandler,
		GeneralSvc:     generalSvc,
		GeneralStore:   generalStore,

		OrganizationHandler: organizationHandler,
		OrganizationSvc:     organizationSvc,
		OrganizationStore:   organizationStore,

		MiddlewareHandler: middlewareHandler,
		MiddlewareSvc:     middlewareSvc,

		SchoolHandler: schoolHandler,
		SchoolService: schoolSvc,
		SchoolStorage: schoolStore,

		PaymentHandler: paymentsHandler,
		PaymentService: paymentsService,
		PaymentStorage: paymentsStore,

		UserHandler: userHandler,
		UserSvc:     userSvc,
		UserStore:   userStore,

		RAGHandler: ragHandler,
		RAGSvc:     ragSvc,

		Cronjobs: cronjobInstance,
	}, nil
}

// @title		DCS / BAföG-App
// @version	0.1
// @description
// @description	<img alt="coffee drinking gopher" src="/api/asset/gopher-coffee" height="200px">
// @description
//
// @host		web.nextreleaseplease.com
// @BasePath	/
func (a *App) RunApp() {
	router := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			// for internal use you can print error to response
			// "Unexpected Error Occurred " + "\n" + err.Error()
			return c.Status(fiber.StatusInternalServerError).SendString("Unexpected Error Occurred")
		},
	})

	a.CreateRoutes(router)

	if configuration.Webserver.Display.Swagger {
		router.Get(
			"/swagger/*",
			a.MiddlewareHandler.Authentication(),
			// a.MiddlewareHandler.PermissionsCheck([]string{}),
			swagger.New(swagger.Config{
				CustomStyle: swaggercss.GetSwaggerCustomCSS(),
			}),
		)
	}

	router.Use(customhttphandler.NotFoundHandler)

	serverPort := fmt.Sprintf(":%d", configuration.Webserver.Port)
	err := router.Listen(serverPort, fiber.ListenConfig{DisableStartupMessage: true})

	// err = router.Listen(serverPort, tlsCert)
	if err != nil {
		panic(err)
	}
}

func (appService *AppServices) ReturnAppInE2EMode() (*fiber.App, *sqlx.DB, *minio.Client) {

	var db *sqlx.DB = appService.StorageProviderPostgres
	var s3 *minio.Client = appService.StorageProviderMinio

	router := fiber.New(fiber.Config{ErrorHandler: customhttphandler.UnexpectedErrorHandler})

	app := convertServicesToApp(appService)

	app.CreateRoutes(router)
	router.Use(customhttphandler.NotFoundHandler)

	return router, db, s3
}
