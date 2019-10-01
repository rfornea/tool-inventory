package utils

import (
	"errors"
	"log"

	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

/*ServerEnv represents our environment*/
type ServerEnv struct {
	// Database information
	ProdDatabaseURL string `env:"PROD_DATABASE_URL" envDefault:""`
	TestDatabaseURL string `env:"TEST_DATABASE_URL" envDefault:""`
	DatabaseURL     string `envDefault:""`

	// Go environment
	GoEnv string `env:"GO_ENV" envDefault:"GO_ENV not set!"`
}

/*Env is an instance of ServerEnv while the server is running*/
var Env ServerEnv

func initEnv(filenames ...string) {
	// Load ENV variables
	err := godotenv.Load(filenames...)
	if err != nil {
		lookupErr := tryLookUp()
		if lookupErr != nil {
			log.Fatal("Error loading environment variables: " + CollectErrors([]error{err, lookupErr}).Error())
		}
		return
	}

	serverEnv := ServerEnv{}
	env.Parse(&serverEnv)

	Env = serverEnv
}

/*SetLive sets the live environment*/
func SetLive() {
	initEnv()
	Env.GoEnv = "live"
	Env.DatabaseURL = Env.ProdDatabaseURL
}

/*SetTesting sets the testing environment*/
func SetTesting(filenames ...string) {
	initEnv(filenames...)
	Env.GoEnv = "test"
	Env.DatabaseURL = Env.TestDatabaseURL
}

/*IsTestEnv returns whether we are in the test environment*/
func IsTestEnv() bool {
	return Env.GoEnv == "test"
}

func tryLookUp() error {
	var collectedErrors []error
	prodDBUrl := AppendLookupErrors("PROD_DATABASE_URL", &collectedErrors)
	testDBUrl := AppendLookupErrors("TEST_DATABASE_URL", &collectedErrors)

	serverEnv := ServerEnv{
		ProdDatabaseURL: prodDBUrl,
		TestDatabaseURL: testDBUrl,
	}

	Env = serverEnv
	return CollectErrors(collectedErrors)
}

/*AppendLookupErrors will look up an environment variable, and if there was an
error, it will append it to the array of errors that are passed in*/
func AppendLookupErrors(property string, collectedErrors *[]error) string {
	value, exists := os.LookupEnv(property)
	if !exists || value == "" {
		AppendIfError(errors.New("in tryLookup, failed to load .env variable: "+property), collectedErrors)
	}
	return value
}
