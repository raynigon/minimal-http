# minimal-http
Minimal HTTP Server with small integrated CLI to modify files

## Container Usage
Run the container with `docker run -p 8080:8080 -d --rm raynigon/minimal-http` and call the cli with `docker exec -ti <CONTAINER_ID> /cli copyTo /data/index.html`.

## CLI Usage
The CLI currently only supports two commands.
The first command allows the user to create a new file and copy content from stdin to this file.
This can be used to create files in the docker container after it was started without the need of a mounted volume.
The second command can be used to create directories.