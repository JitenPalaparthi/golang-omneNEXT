podman exec -it redis-7000 redis-cli --cluster create redis-7000:7000 redis-7001:7001 redis-7002:7002 --cluster-yes

 podman  exec -it redis-7000 redis-cli -p 7000
podman volume create redis-data
podman run -d --name redis \
  -p 6379:6379 \
  -v redis-data:/data:Z \
  redis:7.2-alpine \
  redis-server --appendonly yes
