apiVersion: dapr.io/v1alpha1
kind: Component
metadata:
  name: applogger
  namespace: {{.K8sNamespace}}
spec:
  type: applogger.mongodb
  version: v1
  metadata:
  - name: host
    value: 192.168.64.8:27018,192.168.64.8:27019,192.168.64.8:27020
  - name: replica-set
    value: mongors
  - name: databaseName
    value: "dapr_applog"
  - name: username
    value: "dapr"
  - name: password
    value: "123456"
