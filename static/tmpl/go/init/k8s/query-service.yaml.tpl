kind: Service
apiVersion: v1
metadata:
  name: {{.QueryServiceName}}
  namespace: {{.K8sNamespace}}
  labels:
    app: {{.QueryServiceName}}
spec:
  type: NodePort
  selector:
    app: {{.QueryServiceName}}
  ports:
    - port: 80
      targetPort: 8080
      nodePort: 30002
---
kind: Deployment
apiVersion: apps/v1
metadata:
  name: {{.QueryServiceName}}
  namespace: {{.K8sNamespace}}
  labels:
    app: {{.QueryServiceName}}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{.QueryServiceName}}
  template:
    metadata:
      labels:
        app: {{.QueryServiceName}}
      annotations:
        dapr.io/enabled: "true"
        dapr.io/app-id: "{{.QueryServiceName}}"
        dapr.io/app-port: "8080"
    spec:
      # hostNetwork: true
      containers:
        - name: {{.QueryServiceName}}
          image:{{.CommandImage}}
          ports:
            - containerPort: 8080
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: {{.QueryServiceName}}-ingress
  namespace: {{.K8sNamespace}}
spec:
  rules:
    - http:
        paths:
          - path: /darp-query
            pathType: Prefix
            backend:
              service:
                name: cmd-service
                port:
                  number: 8080