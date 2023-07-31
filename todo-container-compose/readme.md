## Todo and Redis

This folder covers how to orchestrate the todo-api with redis using docker compose.  There are several different examples:

* `1-basic`: This example starts both the redis and todo-api container and allows them to work together.  Redis and the API are fully exposed to the outside host
* `2-network`: This example starts both the redis and todo-api container and allows them to work together.  The `todo-api` is exposed externally, but the network running `redis` is not.  Thus we are simulating `redis` running as a backend service and the `todo-api` running as a front end service.
* `3-volume`: This example builds on the `2-network` example and adds persistance to the redis container.
* `4-init-redis`:  This example builds on the `3-volume` example.  It not only has persistance it demonstrates the concept of an `init` container that runs to initialize something and then goes away.  Specifically, the `cache-init` container runs, executes the `\data\load-redis.sh` script to connect to the `cache` container and then load some sample data. 


#### Changes to the ToDo API

Note the `/api` directory, this API adds a `/kill` endpoint to show how we can use the restart capabilities of docker compose to add some resiliency.  You need to build this container for this demonstration.  There is a build-docker script in the api directory.  Note that this will create the container named `todo-api-basic:v3`.  Thus all of the demos here will use `v3` of our todo playground container. 