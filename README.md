## Pre-requisites   
You have the following installed on your machine:  
* go  
* yarn (or you can probably just use npm)  
* docker-compose  

## Backend configuration

To configure the backend, specify the following environment variables in an `.env` file located in `tool-inventory\pkg\backend\.env`:  
`PORT`  
`PROD_DATABASE_URL`  
`TEST_DATABASE_URL`  

The database URLs should point to a sql database.  

This project includes a docker-compose file to serve up a sql database through docker.  This will be the simplest option for most people.  If you decide to do this, copy and paste the following into `tool-inventory\pkg\backend\.env`:   
```
PORT=3000

PROD_DATABASE_URL=root:secret@tcp(localhost:33060)/prod?parseTime=true&multiStatements=true&readTimeout=1s

TEST_DATABASE_URL=root:secret@tcp(localhost:33060)/test?parseTime=true&multiStatements=true&readTimeout=1s
```

While in `tool-inventory\pkg\backend`, you can run `go run main.go` to serve up the backend, or `go build main.go` and then just `main`.  

## Frontend configuration  

In `tool-inventory\pkg\frontend`, create an `.env` file and put this environment variable in it:  
`VUE_APP_HOST`  

This should point to the backend you are serving up.  The value will most likely be something like:
`VUE_APP_HOST=http://localhost:3000/api/v1`  

Remember to replace `3000` with whatever port you specified the backend `.env` file.  

Then run `yarn install` and `yarn serve`.  The app should be accessible in the browser at `localhost:8080` and will initially not have any data in it.    

## Database configuration

If you want to use the docker-compose file to spin up a sql database locally, you should be able to just run `docker-compose up` or `docker-compose up -d` if you don't care about seeing the output once the container is built.  If you use the docker option for your sql database, you'll want your backend's `.env` file to point to the URLs provided in the backend configuration section.



