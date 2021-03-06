apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: statestore
  namespace: {{.K8sNamespace}}
spec:
  type: state.redis
  version: v1
  metadata:
  - name: redisHost
    value: 192.168.64.4:6379
  - name: redisPassword
    value: ""
  - name: actorStateStore
    value: "true"
