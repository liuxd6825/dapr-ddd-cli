kind: Service
apiVersion: v1
metadata:
  name: {{.Namespace}}-query-service
  namespace: dapr-system
  labels:
    app: {{.Namespace}}-query-service
spec:
  type: NodePort
  selector:
    app: {{.Namespace}}-query-service
  ports:
    - port: 80
      targetPort: 8080
      nodePort: 30002
---
kind: Deployment
apiVersion: apps/v1
metadata:
  name: {{.Namespace}}-query-service
  namespace: dapr-system
  labels:
    app: {{.Namespace}}-query-service
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{.Namespace}}-query-service
  template:
    metadata:
      labels:
        app: {{.Namespace}}-query-service
      annotations:
        dapr.io/enabled: "true"
        dapr.io/app-id: "{{.Namespace}}-query-service"
        dapr.io/app-port: "8080"
    spec:
      # hostNetwork: true
      containers:
        - name: {{.Namespace}}-query-service
          image: 192.168.64.12/{{.Namespace}}-query-service:dapr-linux-arm64
          ports:
            - containerPort: 8080
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: query-service-ingress
  namespace: dapr-system
spec:
  rules:
    - http:
        paths:
          - path: /{{.Namespace}}/query
            pathType: Prefix
            backend:
              service:
                name: {{.Namespace}}-query-service
                port:
                  number: 8080