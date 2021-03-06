apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: pubsub
  namespace: {{.K8sNamespace}}
spec:
  type: pubsub.redis
  version: v1
  metadata:
  - name: redisHost
    value: 192.168.64.4:6379
  - name: redisPassword
    value: ""
  - name: topics
    value: UpdateUserEvent
  - name: topics
    value: CreateUserEvent
