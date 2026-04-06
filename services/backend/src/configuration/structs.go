package configuration

import (
	"time"

	"github.com/dcssoftware/bafoeg-manager/src/resources/auth/service/provider-const"
)

type WebserverModel struct {
	Displayname   string
	Host          string
	Port          int
	IsE2ETestMode bool
	Display       WebserverModelDisplay
}

type WebserverModelDisplay struct {
	Swagger                bool
	MaxResponseEntityCount uint
}

type ConjobsModel struct {
	RagVectorProcessor ConjobsRagVectorProcessor
}

type ConjobsRagVectorProcessor struct {
	ParallelJobs uint
	Interval     time.Duration
	Timeout      time.Duration
}

type DatabaseModel struct {
	Addr     string
	Port     int
	Ssl      bool
	Username string
	Password string
	Database string
}

type S3BucketModel struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	UseSSL          bool
}

type S3StoragePathsModel struct {
	ProfilePictureBucket  string
	ApplicationDataBucket string
	EAktenDataBucket      string
	RAGDataBucket         string
}

type EmailServerModel struct {
	Addr           string
	Port           string
	Authentication string
	Username       string
	Password       string
}

type ClamAVWrapperModel struct {
	Enabled  bool
	Address  string
	Port     uint
	IsSecure bool
}

type LoggerModel struct {
	Loki LoggerModelLoki
}

type LoggerModelLoki struct {
	Enabled bool
	URL     string
}

type AuthenticationModel struct {
	Oauth []OauthConfigurationModel
}

type OauthConfigurationModel struct {
	Enabled          bool
	Identifier       string
	TypeIdentifier   provider.AuthProvider
	AccessKeyID      string
	AccessSecretID   string
	CallbackURL      string
	AuthorizationURL string
	TokenURL         string
	UserinfoURL      string
	Scopes           []string
	LogoutURL        string
}

type OllamaAPIModel struct {
	Address             string
	Port                uint
	ApiKey              string
	IsSecure            bool
	EmbeddingModelname  string
	RequestingModelname string

	DatabaseTablenameRAGSchueler     string
	DatabaseTablenameRAGStudierenden string

	TextSplitterChunkSize    int
	TextSplitterChunkOverlap int
}

type ApplicationConfigurationModel struct {
	MaxLegalProcessingTime uint
	IsDevEnvironment       bool
}

type SecurityModel struct {
	JwtTokenValidInMinutes     uint
	RefreshTokenValidityInDays uint
	JwtTokenSigningKey         string
	RefreshTokenSigningKey     string
}
