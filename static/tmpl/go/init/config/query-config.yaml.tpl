envType: dev
dev:
  app:
    id: duxm-fundflow-queryservice
    httpPort: 9020
    rootUrl: /api/v1.0
  dapr:
    host: localhost
    httpPort: 9021
    grpcPort: 9022
    pubsubs:
      - "pubsub"
  log:
    level: debug
  mongo:
    host: 192.168.64.8:27018, 192.168.64.8:27019, 192.168.64.8:27020
    replicaSet: mongors
    dbname: duxm-fundflow-query-dev
    user: fundflow
    pwd: 123456
    maxPoolSize: 20
test :
  app:
    id: duxm-fundflow-queryservice
    httpPort: 8080
    rootUrl: /api/v1.0
  dapr:
    host: localhost
    pubsubs:
      - "pubsub"
  log:
    level: debug
  mongo:
    host: 192.168.64.8:27018, 192.168.64.8:27019, 192.168.64.8:27020
    replicaSet: mongors
    dbname: duxm-fundflow-query-test
    user: fundflow
    pwd: 123456
    maxPoolSize: 20
prod:
  app:
    id: duxm-fundflow-queryservice
    httpPort: 8080
    rootUrl: /api/v1.0
  dapr:
    host: localhost
    pubsubs:
      - "pubsub"
  log:
    level: debug
  mongo:
    host: 192.168.64.8:27018, 192.168.64.8:27019, 192.168.64.8:27020
    replicaSet: mongors
    dbname: duxm-fundflow-query-prod
    user: fundflow
    pwd: 123456
    maxPoolSize: 20