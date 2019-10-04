## Backend configuration

To configure the backend, specify the following environment variables in an `.env` file located in `tool-inventory\pkg\backend\.env`:  
`PORT`  
`PROD_DATABASE_URL`  
`TEST_DATABASE_URL`  

The database URLs should point to a postgres databases.  

While in `tool-inventory\pkg\backend`, you can run `go run main.go` to serve up the backend, or `go build main.go` and then just `main`.  

## Frontend configuration  

In `tool-inventory\pkg\frontend`, create an `.env` file and put this environment variable in it:  
`VUE_APP_HOST`  

This should point to the backend you are serving up.  The value will most likely be something like:
`VUE_APP_HOST=http://localhost:3000/api/v1`  

Remember to replace `3000` with whatever port you specified the backend .env file.  

Then run `yarn install` and `yarn serve`.  The app should be accessible in the browser at `localhost:8080` and will initially not have any data in it.    


