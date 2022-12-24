# File Sharing Server

How to run:
```bash
# define the path you want to share
SHARE_PATH="$(pwd)"
# define the port you want to expose
PORT="19090"
docker run -d -p $PORT:19090 -v $SHARE_PATH:/file-sharing z1cheng/file-sharing-server:latest
```

Docker Repository: https://hub.docker.com/r/z1cheng/file-sharing-server
