.PHONY= compile build
LIB=sps-app
APP_EXEC="./$(LIB)"

compile:
	go build -o $(APP_EXEC)
build: compile

