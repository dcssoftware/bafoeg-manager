swag fmt -g src/web-app/app.go 
swag init -g src/web-app/app.go -o swagger-docs

gotestsum --format testname ./src/...
