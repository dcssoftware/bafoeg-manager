ALLOWED_LICENSES_GOLANG_BACKEND ?= "MIT,BSD-2-Clause,BSD-3-Clause,Apache-2.0,MPL-2.0,ISC"
ALLOWED_LICENSES_NPM_FRONTEND ?= "MIT;BSD-3-Clause;BSD-2-Clause;BSD;BSD*;Apache-2.0;ISC;OFL-1.1;CC0-1.0;0BSD;UNLICENSED;Python-2.0;BlueOak-1.0.0;MPL-2.0"
IGNORE ?= "github.com/dcssoftware/bafoeg-manager,golang.org/x,"

GIT_COMPARE_SINCE_TAG = $(git rev-parse main)

HELP_COMMAND_PADDING = 50

.install-deps: ## install all dependencies
	@bash  ./devops/scripts/utils/install-dependencies.sh

dev: .install-deps ## start debugging in docker compose microservices (auto reload)
	@docker compose -f docker-compose.dev.yaml up --build --abort-on-container-failure backend frontend

deadcode-lint:
	@(cd services/backend ; deadcode -f='{{println .Path}}{{range .Funcs}}{{printf "\t%s\n" .Name}}{{end}}{{println}}' -test ./src ./integration-tests/...)

prepare-integration-test: ## prepare golang integration tests
	@docker compose -f docker-compose.dev.yaml up -d --wait integration-tests-postgres clamav
	@docker compose -f docker-compose.dev.yaml exec integration-tests-postgres bash -c "psql -U postgres -d postgres -c \"SELECT 1 FROM pg_database WHERE datname = 'unittest_template'\" | grep -q 1 || psql -U postgres -d postgres -c 'CREATE DATABASE unittest_template;'"
	@docker compose -f docker-compose.dev.yaml exec integration-tests-postgres bash -c "psql -U postgres -d postgres -c 'GRANT ALL PRIVILEGES ON DATABASE unittest_template TO postgres;'"
	@docker compose -f docker-compose.dev.yaml exec integration-tests-postgres bash -c "psql -U postgres -d postgres -c 'UPDATE pg_database SET datistemplate=true, datallowconn=true WHERE datname='\''unittest_template'\'''"
	@(cd services/backend ; APP_DATABASE_ADDR="127.0.0.1" APP_DATABASE_PORT=5400 APP_DATABASE_DATABASE=unittest_template go run ./src/main.go migrate-db)
	@docker compose -f docker-compose.dev.yaml exec integration-tests-postgres bash -c "psql -U postgres -d postgres -c 'UPDATE pg_database SET datistemplate=true, datallowconn=false WHERE datname='\''unittest_template'\'''"

test-be: prepare-integration-test ## run golang unit tests
	@(cd services/backend ;  APP_CLAMAV_SCANNER_ADDRESS="127.0.0.1" APP_S3_BUCKET_PROFILEPICTURE_BUCKETNAME="integration-tests" APP_LOGGER_LOKI_URL="http://127.0.0.1:3100/loki/api/v1/push" APP_S3_BUCKET_ENDPOINT="127.0.0.1:9010" APP_DATABASE_ADDR="127.0.0.1" APP_DATABASE_PORT=5400 gotestsum --format testname ./...)

test-be-ci: prepare-integration-test ## run golang unit tests (if gotestsum is not installed)
	@(cd services/backend ; APP_CLAMAV_SCANNER_ADDRESS="127.0.0.1" APP_S3_BUCKET_PROFILEPICTURE_BUCKETNAME="integration-tests" APP_LOGGER_LOKI_URL="http://127.0.0.1:3100/loki/api/v1/push" APP_S3_BUCKET_ENDPOINT="127.0.0.1:9010" APP_DATABASE_ADDR="127.0.0.1" APP_DATABASE_PORT=5400 go run gotest.tools/gotestsum@latest --format testname ./...)

test-fe: ## run sveltekit unit tests
	@docker compose -f docker-compose.dev.yaml up --abort-on-container-exit frontend-tests

deadcode-be: ## run deadcode for backend
	@(cd services/backend ; deadcode ./...)

deadcode-fe: ## run deadcode for frontend
	@(cd services/frontend ; npm run knip)

git-has-change-be: ## check if there are changes in the backend since (latest) commit
	@echo "Compare to Git Tag: $(GIT_COMPARE_SINCE_TAG)"
	@git diff --quiet HEAD $(GIT_COMPARE_SINCE_TAG) -- ./services/backend || echo changed

git-has-change-fe: ## check if there are changes in the frontend since (latest) commit
	@echo "Compare to Git Tag: $(GIT_COMPARE_SINCE_TAG)"
	@git diff --quiet HEAD $(GIT_COMPARE_SINCE_TAG) -- ./services/frontend || echo changed

git-has-change-actions: ## check if there are changes in the github / gitea actions since (latest) commit
	@echo "Compare to Git Tag: $(GIT_COMPARE_SINCE_TAG)"
	@git diff --quiet HEAD $(GIT_COMPARE_SINCE_TAG) -- ./github/workflows || echo changed

dependency-manual-updater-fe: ## run npm-check-updates for frontend to update dependencies manually
	@(cd services/frontend ; ncu -u -i)

dependency-manual-updater-fe-migrate-storybook: ## migrate storybook to latest version (if needed)
	@(cd services/frontend ; npx storybook@latest upgrade)

dependency-manual-updater-be: ## run go-mod-upgrade for backend to update dependencies manually
	@(cd services/backend ; go-mod-upgrade)

build-dockerfile-frontend: ## build the frontend microservice
	@bash ./devops/scripts/build-container/frontend.sh

build-dockerfile-backend: ## build the backend microservice
	@bash ./devops/scripts/build-container/backend.sh

dev-storybook: .install-deps ## start debugging storybook ui lib service (auto reload)
	@(cd services/frontend ; npm run storybook)

storybook-build: .install-deps ## build the storybook ui lib
	@(cd services/frontend ; npm run build-storybook)

license-check-be: ## runs golang license check
	@(cd services/backend ; go-licenses check --allowed_licenses=$(ALLOWED_LICENSES_GOLANG_BACKEND) --ignore=$(IGNORE) --one_output ./...); \
	STATUS=$$?; \
	exit $$STATUS

license-check-fe: ## runs npm license check
	@(cd services/frontend ; license-checker --onlyAllow=$(ALLOWED_LICENSES_NPM_FRONTEND)); \
	STATUS=$$?; \
	exit $$STATUS

dev-generate-licenses-for-ui: ## generate license.json file for backend to show to user in frontend
	@(cd services/dev-utils/prepare-licenses ; go run ./main.go); \
	STATUS=$$?; \
	exit $$STATUS

renovate-config-validator: ## validate renovate.json configuration
	@npx --yes --package renovate -- renovate-config-validator

help: ## print our all commands to commandline
	@echo "\033[34m"
	@echo "		BAföG Manager Application"
	@echo "	   ------------------------------------"
	@echo "   		@dcssoftware + @uvulpos"
	@echo "\033[0m"
	@echo ""
	@echo "\033[33mDevelopment Commands:\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-$(HELP_COMMAND_PADDING)s\033[0m %s\n", $$1, $$2}' | grep -B0 -E 'dev'
	@echo ""
	@echo "\033[33mLicense Commands:\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-$(HELP_COMMAND_PADDING)s\033[0m %s\n", $$1, $$2}' | grep -B0 -E 'license-'
	@echo ""
	@echo "\033[33mStorybook Commands:\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-$(HELP_COMMAND_PADDING)s\033[0m %s\n", $$1, $$2}' | grep -B0 -E 'storybook-'
	@echo ""
	@echo "\033[33mDeadcode Commands:\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-$(HELP_COMMAND_PADDING)s\033[0m %s\n", $$1, $$2}' | grep -B0 -E 'deadcode-'
	@echo ""
	@echo "\033[33mGit Commands:\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-$(HELP_COMMAND_PADDING)s\033[0m %s\n", $$1, $$2}' | grep -B0 -E 'git-'
	@echo ""
	@echo "\033[33mDependency Updates Commands:\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-$(HELP_COMMAND_PADDING)s\033[0m %s\n", $$1, $$2}' | grep -B0 -E 'dependency-'
	@echo ""
	@echo "\033[33mDocker Commands:\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-$(HELP_COMMAND_PADDING)s\033[0m %s\n", $$1, $$2}' | grep -B0 -E 'build-dockerfile-'
	@echo ""
	@echo "\033[33mTest Commands:\033[0m"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-$(HELP_COMMAND_PADDING)s\033[0m %s\n", $$1, $$2}' | grep -B0 -E 'test-be|test-fe'
