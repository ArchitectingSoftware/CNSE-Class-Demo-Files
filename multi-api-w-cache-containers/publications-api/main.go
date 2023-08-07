package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"architectingsoftware.com/pub-api/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	hostFlag string
	portFlag uint
	cacheURL string
)

func processCmdLineFlags() {
	flag.StringVar(&hostFlag, "h", "0.0.0.0", "Listen on all interfaces")
	flag.StringVar(&cacheURL, "c", "0.0.0.0:6379", "Default cache location")
	flag.UintVar(&portFlag, "p", 2080, "Default Port")

	flag.Parse()
}

func envVarOrDefault(envVar string, defaultVal string) string {
	envVal := os.Getenv(envVar)
	if envVal != "" {
		return envVal
	}
	return defaultVal
}

func setupParms() {
	//first process any command line flags
	processCmdLineFlags()

	//now process any environment variables
	cacheURL = envVarOrDefault("PUBAPI_CACHE_URL", cacheURL)
	hostFlag = envVarOrDefault("PUBAPI_HOST", hostFlag)
	pfNew, err := strconv.Atoi(envVarOrDefault("PUBAPI_PORT", fmt.Sprintf("%d", portFlag)))
	//only update the port if we were able to convert the env var to an int, else
	//we will use the default we got from the command line, or command line defaults
	if err == nil {
		portFlag = uint(pfNew)
	}

}

func main() {
	//this will allow the user to override key parameters and also setup defaults
	setupParms()

	apiHandler, err := api.NewPubAPI(cacheURL)

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/pubs", apiHandler.GetPublications)
	r.GET("/pubs/:id", apiHandler.GetPublication)

	//For now we will just support gets
	serverPath := fmt.Sprintf("%s:%d", hostFlag, portFlag)
	r.Run(serverPath)

}
