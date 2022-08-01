env: dev
envs:
  dev:
    app:
      id: {{.QueryServiceName}}
      httpPort: {{.Metadata.QueryService.Dev.AppHttpPort}}
      rootUrl: /api/{{.ApiVersion}}
    dapr:
      host: localhost
      httpPort: {{.Metadata.QueryService.Dev.DaprHttpPort}}
      grpcPort: {{.Metadata.QueryService.Dev.DaprGrpcPort}}
      pubsubs:
        - "pubsub"
    log:
      level: debug
    mongo:
      default:
        host: {{.Metadata.QueryService.Dev.Mongo.host}}
        replicaSet: {{.Metadata.QueryService.Dev.Mongo.replicaSet}}
        dbname: {{.Metadata.QueryService.Dev.Mongo.dbname}}
        user: {{.Metadata.QueryService.Dev.Mongo.user}}
        pwd: {{.Metadata.QueryService.Dev.Mongo.pwd}}
        maxPoolSize: {{.Metadata.QueryService.Dev.Mongo.maxPoolSize}}
    neo4j:
      default:
        host: {{.Metadata.QueryService.Dev.Neoj4.host}}
        port: {{.Metadata.QueryService.Dev.Neoj4.port}}
        user: {{.Metadata.QueryService.Dev.Neoj4.user}}
        pwd: {{.Metadata.QueryService.Dev.Neoj4.pwd}}
  test:
    app:
      id: {{.QueryServiceName}}
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
    mongo:
      default:
        host: {{.Metadata.QueryService.Test.Mongo.host}}
        replicaSet: {{.Metadata.QueryService.Test.Mongo.replicaSet}}
        dbname: {{.Metadata.QueryService.Test.Mongo.dbname}}
        user: {{.Metadata.QueryService.Test.Mongo.user}}
        pwd: {{.Metadata.QueryService.Test.Mongo.pwd}}
        maxPoolSize: {{.Metadata.QueryService.Test.Mongo.maxPoolSize}}
    neo4j:
      default:
        host: {{.Metadata.QueryService.Test.Neoj4.host}}
        port: {{.Metadata.QueryService.Test.Neoj4.port}}
        user: {{.Metadata.QueryService.Test.Neoj4.user}}
        pwd: {{.Metadata.QueryService.Test.Neoj4.pwd}}
  prod:
    app:
      id: {{.QueryServiceName}}
      httpPort: 8080
      rootUrl: /api/{{.ApiVersion}}
    dapr:
      host: localhost
      pubsubs:
      - "pubsub"
    log:
      level: debug
    mongo:
      default:
        host: {{.Metadata.QueryService.Prod.Mongo.host}}
        replicaSet: {{.Metadata.QueryService.Prod.Mongo.replicaSet}}
        dbname: {{.Metadata.QueryService.Prod.Mongo.dbname}}
        user: {{.Metadata.QueryService.Prod.Mongo.user}}
        pwd: {{.Metadata.QueryService.Prod.Mongo.pwd}}
        maxPoolSize: {{.Metadata.QueryService.Prod.Mongo.maxPoolSize}}
    neo4j:
      default:
        host: {{.Metadata.QueryService.Prod.Neoj4.host}}
        port: {{.Metadata.QueryService.Prod.Neoj4.port}}
        user: {{.Metadata.QueryService.Prod.Neoj4.user}}
        pwd: {{.Metadata.QueryService.Prod.Neoj4.pwd}}