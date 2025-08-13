docker exec -it kafka1 kafka-topics \
  --create --topic test \
  --bootstrap-server kafka1:9092,kafka2:9092,kafka3:9092 \
  --replication-factor 3 --partitions 3


  docker exec -it kafka1 kafka-topics --list --bootstrap-server kafka1:9092


  KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://192.168.0.158:19092