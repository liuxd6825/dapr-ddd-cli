env: dev
envs:
  dev:
    app:
      id: {{.CommandServiceName}}
      httpPort: {{.Metadata.CommandService.Dev.AppHttpPort}}
      rootUrl: /api/{{.ApiVersion}}
    dapr:
      host: localhost
      httpPort: {{.Metadata.CommandService.Dev.DaprHttpPort}}
      grpcPort: {{.Metadata.CommandService.Dev.DaprGrpcPort}}
      pubsubs:
        - "pubsub"
    log:
      level: debug
  test:
    app:
      id: {{.CommandServiceName}}
      httpPort: 8080
      rootUrl: /api/{{.ApiVersion}}
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
      rootUrl: /api/{{.ApiVersion}}
    dapr:
      host: localhost
      #httpPort: 3500
      #grpc-port: 50001
      pubsubs:
        - "pubsub"
    log:
      level: debug
