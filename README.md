

## Install

    go get -u github.com/achmadnf18/sps-recruitment
    
## Restore Database
    sudo su - postgres
    createuser spacestock
    createdb spacestock
    psql -c "GRANT ALL PRIVILEGES ON DATABASE spacestock TO spacestock;"
    exit
    cd $GOPATH/PATH_TO_PROJECT/rest/ 
    psql -U spacestock spacestock < models/migrations/spacestock_bak.sql
    
## Run Makefile
    cd $GOPATH/PATH_TO_PROJECT/
    make setup
  
## How to Run Assignment
  
  #### Function Sum
    1. ./sps-app sum
  
  #### Function Contains
    2. ./sps-app contains
  
  #### Function Get Person Struct
    3. ./sps-app person
  
  #### Run Unit Test
    4. cd $GOPATH/PATH_TO_PROJECT/application/test && go test -v
  
  #### Function Rearrange
    5. ./sps-app rearrange
    
  #### REST
    a. cd rest/
    b. go run app.go
    c. Open new terminal
    d. fajri@fajri: http GET localhost:8000/apartment
    e. fajri@fajri: http -f POST localhost:8000/apartment name='Apartement SEA'
    f. fajri@fajri: http -f PUT localhost:8000/apartment id=1 name='Apartement Cikarang'
    g. fajri@fajri: http DELETE localhost:8000/apartment/1
