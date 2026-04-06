package configuration

import (
	"time"

	"github.com/dcssoftware/bafoeg-manager/src/resources/auth/service/provider-const"
)

/*

	DON'T FORGET TO ADD THE GIVEN PRESET

	->> APP_ <<-

*/

var (
	Webserver = WebserverModel{
		Displayname:   GetEnvOrDefaultString("WEBSERVER_DISPLAYHOST", "http://web.nextreleaseplease.com/"),
		Host:          GetEnvOrDefaultString("WEBSERVER_HOST", "127.0.0.1"),
		Port:          GetEnvOrDefaultInt("WEBSERVER_PORT", 3000),
		IsE2ETestMode: GetEnvOrDefaultBool("WEBSERVER_E2E_TEST_MODE", false),
		Display: WebserverModelDisplay{
			Swagger:                GetEnvOrDefaultBool("WEBSERVER_SHOW_SWAGGER", true),
			MaxResponseEntityCount: GetEnvOrDefaultUint("APP_CONFIG_RESPOMSE_ENTITY_COUNT", 25),
		},
	}

	Database = DatabaseModel{
		Addr:     GetEnvOrDefaultString("DATABASE_ADDR", "postgres"),
		Port:     GetEnvOrDefaultInt("DATABASE_PORT", 5432),
		Ssl:      GetEnvOrDefaultBool("DATABASE_SSL", false),
		Username: GetEnvOrDefaultString("DATABASE_USERNAME", "postgres"),
		Password: GetEnvOrDefaultString("DATABASE_PASSWORD", "postgres"),
		Database: GetEnvOrDefaultString("DATABASE_DATABASE", "postgres"),
	}

	S3Bucket = S3BucketModel{
		Endpoint:        GetEnvOrDefaultString("S3_BUCKET_ENDPOINT", "minio:9010"),
		AccessKeyID:     GetEnvOrDefaultString("S3_BUCKET_ACCESS_KEY_ID", "minio"),
		AccessKeySecret: GetEnvOrDefaultString("S3_BUCKET_ACCESS_KEY_SECRET", "minioadmin"),
		UseSSL:          GetEnvOrDefaultBool("S3_BUCKET_USE_SSL", false),
	}

	S3StoragePaths = S3StoragePathsModel{
		ProfilePictureBucket:  GetEnvOrDefaultString("S3_BUCKETPATH_PROFILEPICTURE_BUCKETNAME", "test"),
		ApplicationDataBucket: GetEnvOrDefaultString("S3_BUCKETPATH_APPLICATIONDATA_BUCKETNAME", "test"),
		EAktenDataBucket:      GetEnvOrDefaultString("S3_BUCKETPATH_EAKTENDATA_BUCKETNAME", "test"),
		RAGDataBucket:         GetEnvOrDefaultString("S3_BUCKETPATH_RAGDATA_BUCKETNAME", "test"),
	}

	EmailServer = EmailServerModel{
		Addr:     GetEnvOrDefaultString("EMAIL_SERVER_ADDR", "imap-server"),
		Port:     GetEnvOrDefaultString("EMAIL_SERVER_PORT", "993"),
		Username: GetEnvOrDefaultString("EMAIL_SERVER_USERNAME", "address@example.org"),
		Password: GetEnvOrDefaultString("EMAIL_SERVER_PASSWORD", "pass"),
	}

	ClamAVWrapper = ClamAVWrapperModel{
		Enabled:  GetEnvOrDefaultBool("CLAMAV_SCANNER_ENABLED", true),
		Address:  GetEnvOrDefaultString("CLAMAV_SCANNER_ADDRESS", "clamav"),
		Port:     GetEnvOrDefaultUint("CLAMAV_SCANNER_PORT", 33779),
		IsSecure: GetEnvOrDefaultBool("CLAMAV_SCANNER_IS_SECURE", false),
	}

	OllamaAPI = OllamaAPIModel{
		Address: GetEnvOrDefaultString("OLLAMA_ADDRESS", "host.docker.internal"),
		// Address:             GetEnvOrDefaultString("OLLAMA_ADDRESS", "127.0.0.1"),
		Port:                GetEnvOrDefaultUint("OLLAMA_PORT", 11434),
		ApiKey:              GetEnvOrDefaultString("OLLAMA_API_KEY", "HOMERSIMPSON"),
		IsSecure:            GetEnvOrDefaultBool("OLLAMA_IS_SECURE", false),
		EmbeddingModelname:  GetEnvOrDefaultString("OLLAMA_EMBEDDING_MODELNAME", "embeddinggemma:latest"),  // all-minilm
		RequestingModelname: GetEnvOrDefaultString("OLLAMA_REQUESTING_MODELNAME", "ministral-3:14b-cloud"), // ministral-3:14b-cloud
		// RequestingModelname: GetEnvOrDefaultString("OLLAMA_REQUESTING_MODELNAME", "gpt-oss:120b-cloud"),

		DatabaseTablenameRAGSchueler:     GetEnvOrDefaultString("OLLAMA_VECTOR_RAG_TABLENAME_SCHUELER", "rag_schueler"),
		DatabaseTablenameRAGStudierenden: GetEnvOrDefaultString("OLLAMA_VECTOR_RAG_TABLENAME_STUDIERENDEN", "rag_studierenden"),

		TextSplitterChunkSize:    GetEnvOrDefaultInt("OLLAMA_VECTOR_RAG_TEXTSPLITTER_CHUNK_SIZE", 5000),
		TextSplitterChunkOverlap: GetEnvOrDefaultInt("OLLAMA_VECTOR_RAG_TEXTSPLITTER_CHUNK_OVERLAP", 300),
	}

	Logger = LoggerModel{
		Loki: LoggerModelLoki{
			Enabled: GetEnvOrDefaultBool("LOGGER_LOKI_ENABLED", true),
			URL:     GetEnvOrDefaultString("LOGGER_LOKI_URL", "http://loki:3100/loki/api/v1/push"),
		},
	}

	Conjobs = ConjobsModel{
		RagVectorProcessor: ConjobsRagVectorProcessor{
			ParallelJobs: GetEnvOrDefaultUint("CRONJOB_RAG_VECTOR_PROCESSOR_PARALLEL_JOBS", 2),
			Interval:     time.Duration(GetEnvOrDefaultInt("CRONJOB_RAG_VECTOR_PROCESSOR_INTERVAL", 10)) * time.Second,
			Timeout:      time.Duration(GetEnvOrDefaultInt("CRONJOB_RAG_VECTOR_PROCESSOR_TIMEOUT", 10)) * time.Minute,
		},
	}

	Security = SecurityModel{
		JwtTokenValidInMinutes:     GetEnvOrDefaultUint("JWT_TOKEN_VALIDITY_MINS", 15),
		RefreshTokenValidityInDays: GetEnvOrDefaultUint("REFRESH_TOKEM_VALIDITY_MINS", 30),

		JwtTokenSigningKey:     GetEnvOrDefaultString("JWT_SIGNING_KEY", "qWeRtZ12345"),
		RefreshTokenSigningKey: GetEnvOrDefaultString("JWT_SIGNING_KEY", "qWeRtZ12345"),
	}

	Authentication = AuthenticationModel{
		Oauth: []OauthConfigurationModel{
			{
				Enabled:          GetEnvOrDefaultBool("AUTHORIZATION_OAUTH_AUTHENTIK", true),
				Identifier:       GetEnvOrDefaultString("AUTHORIZATION_OAUTH_AUTHENTIK_IDENTIFIER", "authentik"),
				TypeIdentifier:   provider.Authentik,
				AccessKeyID:      GetEnvOrDefaultString("AUTHORIZATION_OAUTH_AUTHENTIK_CLIENT_KEY", ""),
				AccessSecretID:   GetEnvOrDefaultString("AUTHORIZATION_OAUTH_AUTHENTIK_CLIENT_SECRET", ""),
				CallbackURL:      GetEnvOrDefaultString("AUTHORIZATION_OAUTH_AUTHENTIK_REDIRECT_URL", "http://web.nextreleaseplease.com/api/v1/callback"),
				AuthorizationURL: GetEnvOrDefaultString("AUTHORIZATION_OAUTH_AUTHENTIK_AUTHORIZATION_URL", "http://localhost:9000/application/o/authorize/"),
				TokenURL:         GetEnvOrDefaultString("AUTHORIZATION_OAUTH_AUTHENTIK_TOKEN_URL", ""),
				UserinfoURL:      GetEnvOrDefaultString("AUTHORIZATION_OAUTH_AUTHENTIK_USERINFO_URL", ""),
				Scopes:           GetEnvOrDefaultStringArray("AUTHORIZATION_OAUTH_AUTHENTIK_SCOPES", []string{}),
				LogoutURL:        GetEnvOrDefaultString("AUTHORIZATION_OAUTH_AUTHENTIK_LOGOUT_URL", ""),
			},
		},
	}

	ApplicationConfiguration = ApplicationConfigurationModel{
		MaxLegalProcessingTime: GetEnvOrDefaultUint("APP_CONFIG_APPLICATION_MAX_LEGAL_PROCESSING_TIME", 90), // Maximum legal processing time of an application in days
		IsDevEnvironment:       GetEnvOrDefaultBool("APP_CONFIG_APPLICATION_IS_DEV_ENVIRONMENT", true),
	}
)
