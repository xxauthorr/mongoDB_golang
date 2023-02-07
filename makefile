	
APP_BINARY = myApp

## building app	
build_app:
	@echo Building app binary...
	chdir .\ && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${APP_BINARY} ./cmd/api
	@echo Done!

## docker-compose up
up:
	@echo Running containers ...
	docker-compose up -d 
	@echo done !!

down:
	@echo stoping containers..
	docker-compose down --remove-orphans
	@echo done !!