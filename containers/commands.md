## Docker or Podman commands

- To run docker replace podman by docker and viceversa

### To pull the image

```bash
podman pull docker.io/library/postgres
```

### To list down images

```bash
podman images
```

### To run container

```bash
podman run -d --name pg -p 5432:5432 -e POSTGRES_USER=app -e POSTGRES_PASSWORD=app123 -e POSTGRES_DB=usersdb docker.io/library/postgres
```

### To list down running containers

```bash
podman ps
```
### To list down all conatiners(running + non running)

```bash
podman ps -a
```

### To remove a running container

```bash
 podman rm -f pg
 ```
### To start a stopped container or to stop a runnign container

```bash
podman start pg
podman stop pg
```

-- To run postgres db ui

```bash
 podman run -d --name pgui -p 8088:8080 adminer
 ```

 ### TO create a network

 ```bash
  podman network create demo-network
  ```