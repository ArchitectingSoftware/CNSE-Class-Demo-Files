package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"architectingsoftware.com/reading-list-api/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	hostFlag  string
	portFlag  uint
	cacheURL  string
	pubAPIURL string
)

func processCmdLineFlags() {
	flag.StringVar(&hostFlag, "h", "0.0.0.0", "Listen on all interfaces")
	flag.StringVar(&pubAPIURL, "pubapi", "http://localhost:2080", "Default endpoint for publication API")
	flag.StringVar(&cacheURL, "c", "0.0.0.0:6379", "Default cache location")
	flag.UintVar(&portFlag, "p", 3080, "Default Port")

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
	cacheURL = envVarOrDefault("RLAPI_CACHE_URL", cacheURL)
	pubAPIURL = envVarOrDefault("RLAPI_PUB_API_URL", pubAPIURL)
	hostFlag = envVarOrDefault("RLAPI_HOST", hostFlag)

	pfNew, err := strconv.Atoi(envVarOrDefault("RLAPI_PORT", fmt.Sprintf("%d", portFlag)))
	//only update the port if we were able to convert the env var to an int, else
	//we will use the default we got from the command line, or command line defaults
	if err == nil {
		portFlag = uint(pfNew)
	}

}

func main() {
	//this will allow the user to override key parameters and also setup defaults
	setupParms()
	log.Println("Init/cacheURL: " + cacheURL)
	log.Println("Init/pubAPIURL: " + pubAPIURL)
	log.Println("Init/hostFlag: " + hostFlag)
	log.Printf("Init/portFlag: %d", portFlag)

	apiHandler, err := api.NewReadingListAPI(cacheURL, pubAPIURL)

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/publists", apiHandler.GetReadingLists)
	r.GET("/publists/:id", apiHandler.GetReadingList)
	r.GET("/publists/:id/:idx", apiHandler.GetPubFromReadingList)
	r.GET("/publists/:id/:idx/paper", apiHandler.RedirectWithPublication)

	//For now we will just support gets
	serverPath := fmt.Sprintf("%s:%d", hostFlag, portFlag)
	r.Run(serverPath)

}
