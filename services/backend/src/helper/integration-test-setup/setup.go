package integrationtestsetup

import (
	"hash/maphash"
	"testing"

	"github.com/dcssoftware/bafoeg-manager/src/helper/database"
	"github.com/dcssoftware/bafoeg-manager/src/helper/jwt"
	jwtmodels "github.com/dcssoftware/bafoeg-manager/src/helper/jwt/models"
	"github.com/dcssoftware/bafoeg-manager/src/resources/auth/service/provider-const"
	webapp "github.com/dcssoftware/bafoeg-manager/src/web-app"
	"github.com/go-sqlx/sqlx"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/assert"
)

type TestInstance struct {
	App         *fiber.App
	AppServices *webapp.AppServices
	AppJWT      string

	DB *sqlx.DB
	S3 *minio.Client

	Test *testing.T

	Cleanup func()
}

func SetupTest(t *testing.T) *TestInstance {
	t.Parallel()
	testHash := new(maphash.Hash).Sum64()
	database.PrepareDatabaseForIntegrationTest(t, testHash)

	appServices := webapp.NewIntegrationTestApp(testHash)
	fiberApp, db, s3 := appServices.ReturnAppInE2EMode()

	_, err := db.Exec(
		"INSERT INTO users (username, display_name, email, email_verified) VALUES ($1, $2, $3, $4) ON CONFLICT DO NOTHING",
		"integration_tester",
		"Integration Tester",
		"noreply@example.com",
		true,
	)
	assert.NoError(t, err, "could not insert test user")

	result := &TestInstance{
		App:         fiberApp,
		AppServices: appServices,
		DB:          db,
		S3:          s3,
		Test:        t,
		Cleanup: func() {
			DropIntegrationDatabase(db, testHash)
		},
	}

	result.AppJWT = result.AuthenticateAs("integration_tester")
	return result
}

func (i *TestInstance) AuthenticateAs(username string) string {
	type DatabaseID struct {
		ID uuid.UUID `db:"id"`
	}
	var responseUser DatabaseID
	var responseSession DatabaseID

	err := i.DB.QueryRowx(
		"SELECT id FROM users WHERE username = $1",
		username,
	).StructScan(&responseUser)

	assert.NoErrorf(i.Test, err, "cannot retrieve application user \"%s\" from database\n", username)

	err = i.DB.QueryRowx(
		`INSERT INTO user_sessions (user_id, useragent_hash, created_ip_addr, last_jwt_refresh_ip_addr) VALUES ($1, $2, $3, $4) RETURNING id`,
		responseUser.ID,
		nil,
		"",
		"",
	).StructScan(&responseSession)

	assert.NoErrorf(i.Test, err, "cannot create application user \"%s\" session from database\n", username)

	jwtToken, jwtTokenErr := jwt.CreateJWT(provider.Authentik, &jwtmodels.JwtDataModel{
		UserUUID:  responseUser.ID.String(),
		SessionID: responseSession.ID.String(),
		Scopes:    []string{""},
	})
	assert.NoErrorf(i.Test, jwtTokenErr, "cannot create jwt token for user \"%s\"\n", username)

	return jwtToken
}
