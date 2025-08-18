## Create an image and publish to docker hub

1. Dockerfile 

2. podman build . -f Dockerfile -t docker.io/jpalaparthi/user-service1:v1

3. podman login docker.io

4. podman push docker.io/jpalaparthi/user-service1:v1

5. podman rmi docker.io/jpalaparthi/user-service1:v1

6. podman pull docker.io/jpalaparthi/user-service1:v1

## Create containers and run 

1. create podman network 

2. create postgres, adminer and also app container on the same network.

3. If using nginx , then create nginx container in the same network.

4. User similar config file of nginx that is used for the demo

5. Test the application

- Run database and ui

```bash
podman run -d --name pg -p 5432:5432 --network demo-network -e POSTGRES_USER=app -e POSTGRES_PASSWORD=app123 -e POSTGRES_DB=usersdb postgres:latest

podman run -d --name dbui -p 28080:8080 --network demo-network adminer

```
- Run kafka

```bash
podman compose . -f ../../kafka/docker-compose-kafka-multinode.yaml up -d
```
-- run prometheus with config file
```bash

 podman run -d --name prometheus --network demo-network -v C:\Users\PalaparthiJitendrana\workspace\go-demos\http\prometheus\prometheus.yml:/opt/bitnami/prometheus/conf/prometheus.yml:ro -p 9090:9090 bitnami/prometheus:latest

```
- Run the built application 
- IF the application is built 

```bash
  podman run -d --name user-service -p 8089:8089 --network demo-network -e PORT=8089 docker.io/jpalaparthi/user-service-prom:v02   
```                                                                      

## If your Kafka is reachable on the HOST as localhost:19092,29092,39092
podman run -d --name kafka-ui -p 8080:8080 `
  -e KAFKA_CLUSTERS_0_NAME=local `
  -e KAFKA_CLUSTERS_0_BOOTSTRAPSERVERS=host.containers.internal:19092,host.containers.internal:29092,host.containers.internal:39092 `
  provectuslabs/kafka-ui:latest

### Postgres Exporter

```
podman run -d --name pg-exporter --network demo-network 
  -e DATA_SOURCE_NAME="postgresql://app:app123@pg:5432/userdb?sslmode=disable"
  -p 9187:9187 `
  wrouesnel/postgres_exporter
  ```