podman exec -it redis-7000 redis-cli --cluster create redis-7000:7000 redis-7001:7001 redis-7002:7002 --cluster-yes

 podman  exec -it redis-7000 redis-cli -p 7000