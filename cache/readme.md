## Cache Demo

This folder contains helpers to run redis locally in a container and to interact with it. There is no code required, although you must have docker (preferrably docker desktop) installed and running on your machine.

#### Scripts
The following describes each of the scripts:

1. `start-redis.sh`: This command starts redis on your machine as a container.  Note the first time you run this it might take a little time as the container must be downloaded.  After you run this `redis` will be available on your local machine over the standard redis port of `6379`.  **NOTE: THIS COMMAND WILL NOT SAVE YOUR DATA LONG TERM - EVERY TIME YOU STOP REDIS AND RESTART IT WITH THIS COMMAND YOU WILL START WITH A FRESH REDIS CACHE CONTAINING NO DATA.**
   
2. `start-redis-volume.sh`.  Just like the `start-redis.sh` command, this will start `redis` in a container on your local machine.  If you look at this command it includes a `-v ./cache-data/:/data` flag.  This instructs the container to _map_ the local subdirectory `./cache-data` to the `\data` directory in the container.  After running this you should see a file named `appendonly.aof` appear in your local `.\cache-data\` directory.  This will save your data from one run of this container to the next, effectively persisting your data long term.

3. `stop-redis.sh`:  This script stops the redis container on your machine.

4. `start-redis-cli.sh`: Redis has a helpful command line interface program called `redis-cli` that enables you to directly interact with redis.  This command will connect to your redis container (which must be running) and will execute the `redis-cli` command for you within the container so that you can interact with it.  Information on what you can do with `redis-cli` can be found here: https://redis.io/docs/ui/cli/

5. `load-data-from-container.sh`:  This command executes the script `\data\load-redis.sh` from within the container to pre-populate the cache with some sample data.  If you want to change the sample data (content or size), you can modify the `./cache-data/redis-load.redis` file.  I am just loading 4 sample records.  Note the first command executed is `flushdb` so each time this runs it will start by purging all of the data in the cache and then inserting the sample records that are shown in the file.

#### Other Noteworthy Items

1. Note that I am providing these files as shell scripts, which assume you are running linux, macos, or an equivalent _bash_ environment on windows. You can download a `bash` emulator for windows, and I think powershell will execute these files. You can also convert them to windows `.bat` files.  If somebody has the initiative to do this (and test it please), please feel free to send me a PR for **EXTRA CREDIT**.  I dont have a windows machine to test this, or I would have done this.

2. I decided to use the expanded `redis-stack` distribution, which has a number of nice features.  If you want to interact with `redis` via nice gui head to `http://localhost:8001` and you can interact with it via a web browser.