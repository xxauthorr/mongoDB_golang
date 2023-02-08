	
APP_BINARY = myApp

## building app	
build_appBinary:
	@echo Building app binary...
	chdir .\ && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${APP_BINARY} ./cmd/api
	@echo Done!

build_image:
	@echo Building app image..
	docker build -t saavy .
	@echo done !
	
## docker-compose up
up:
	@echo Running containers ...
	docker-compose up --build -d 
	@echo done !!

## docker-compose up
down:
	@echo stoping containers..
	docker-compose down --remove-orphans
	@echo done !!

run_app:
	@echo Running binary..
	go run ./cmd/api