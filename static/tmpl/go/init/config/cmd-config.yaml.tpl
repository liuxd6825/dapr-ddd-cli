envType: dev
dev:
  app:
    id: {{.CommandServiceName}} 
    httpPort: 9010
    rootUrl: /api/v1.0
  dapr:
    host: localhost
    httpPort: 9011
    grpcPort: 9012
    pubsubs:
      - "pubsub"
  log:
    level: debug
test:
  app:
    id: {{.CommandServiceName}}
    httpPort: 8080
    rootUrl: /api/v1.0
  dapr:
    host: localhost
    #httpPort: 3500
    #grpc-port: 50001
    pubsubs:
      - "pubsub"
  log:
    level: debug
prod:
  app:
    id: {{.CommandServiceName}}
    httpPort: 8080
    rootUrl: /api/v1.0
  dapr:
    host: localhost
    #httpPort: 3500
    #grpc-port: 50001
    pubsubs:
      - "pubsub"
  log:
    level: debug

