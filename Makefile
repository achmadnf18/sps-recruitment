.PHONY= setup compile build
LIB=sps-app
APP_EXEC="./$(LIB)"

setup:
	sudo apt-get install httpie
	go get -u github.com/urfave/cli
	go get -u github.com/lib/pq
	go get -u github.com/jmoiron/sqlx
	go get -u github.com/labstack/echo
compile:
	go build -o $(APP_EXEC)
build: compile

