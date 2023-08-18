#!/bin/bash
docker run --net=host --ipc=host --uts=host --pid=host -it --security-opt=seccomp=unconfined --privileged --rm -it -v /:/host alpine
apk add util-linux