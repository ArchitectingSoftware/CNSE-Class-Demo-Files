## ToDo API Demo - With Cache

This is an extension of the todo API that persists todo data in a Redis cache.  

### Application Objectives

1. Extend our reference API to integrate with a popular cloud-native cache called redis.
2. Gain experience using a service, redis in this case, that is packaged and deployed in a docker container.
3. Introduce packaging this API into a docker container.

Installing Redis GoLang Client `go get github.com/redis/go-redis/v8`
Installing Redis JSON Extension for GoLang `go get github.com/nitishm/go-rejson/v4`

**IMPORTANT:  REDIS MUST BE RUNNING AND AVAILABLE ON ITS STANDARD PORT 6973 FOR THIS API TO WORK PROPERLY.  DIRECTIONS FOR HOW TO INSTALL AND RUN REDIS LOCALLY VIA A CONTAINER ARE AVAILABLE VIA THE [cache](/cache) DIRECTORY**

### Docker Objectives

This will be our first introduction to creating our own docker containers.  Note that I will be showing building the container 2 different ways.  The first way is highlighted in the `dockerfile.basic` file, the other way is highlighted in the `dockerfile.better` file.

1. **BASIC**: I have provided 2 scripts to create a basic docker container on your machine.  The script `build-basic-docker.sh` is used to create the container. and the `run-basic-docker.sh` is used to run the container.

2. **BETTER**:  I have also provided 2 scripts for a better way to create containers that I will be explaining.  The script `build-better-docker.sh` is used to create the container. and the `run-better-docker.sh` is used to run the container.

__**What Is The Difference?**__
The _basic_ version simply copies your code into a base container that already has the go tooling installed.  From there the code is compiled into a binary, and then set to execute.  The _better_ version does the same as the _basic_ version for the first step.  Inside we create a _build_ container that is based off of a standard container that has the go tooling previously installed.  This _build_ container does the same as the _basic_ container in that it builds the binary in a suitable linux format.  The difference is that it then creates the final container that copies newly created binary from the build container into the final container.  This way the final container does not have any of the go tooling, and is based off of the small alpine linux base image.  Lets get into some other things you should know:

1. Make sure you look at both dockerfiles and understand them.

2. Both dockerfiles build the go program using the command `CGO_ENABLED=0 GOOS=linux go build -o /todo-api`.  We have seen `go build` before but not some of the other flags.  The `-o /todo-api` flag simply states build an executable and name it `todo-api`.  The more interesting things are before the `go` command:

   * The `GOOS=linux` environment variable setup instructs the `go` compiler to build a linux binary. Since the ultimate runnable containers are based on linux this enures the binary will work.
   * The `CGO_ENABLED=0` is a little more involved.  The `go` compiler will make some assumptions about dynamically loadable components that it requires at runtime based on the operating system architecture, linux in this case. Thus the resultant binary will dynamically load these components as needed.  By doing `CGO_ENABLED=0` you are instructing the `go` compiler to statically link all of its dynamic dependencies.  This will result in a larger binary file, however, this file has everything that it requires inside, this making it a bit more portable across linux container runtimes.

3. The next thing I want to call out is inside of both dockerfiles you will see the command `ENV REDIS_URL=host.docker.internal:6379`.  This sets up an environment variable that is used by our API to locate the redis cache.  For now we are executing our API in one container, and the redis cache in another container.  Down the road we will look at container orchestration. Doing things this way demonstrates some best practices:
   * In many cases its preferred that docker containers obtain config and runtime information via environment variables.  Since they are ephemeral components, the runtime aspects may change every time they start, so injecting proper information at startup time via environment variables is a good practice.
   * In our go code, specifically the `todo.go` file we specifiy the _DEFAULT_ location for where this container expects to find redis - `RedisDefaultLocation = "0.0.0.0:6379"`.  Thus by default, its expected to be running locally over poert `6379`.  This is a good default for running this API in development without docker.  If you scroll down a little in the `New()` function you will see:

   ```go
   redisUrl := os.Getenv("REDIS_URL")
	 if redisUrl == "" {
		 redisUrl = RedisDefaultLocation
	 }
   return NewWithCacheInstance(redisUrl)
  

  What this code does is that it first checks to see if the `REDIS_URL` environment varaible is set, if so it sets a local variable `redisUrl` to this value.  The `if` statement handles the case where its not set and then sets the `redisUrl` value to the default discussed above.  The actual connection to redis is handled in the `NewWithCachInstance(redisUrl)` function. This function requires the URL of where redis is actually running. 

