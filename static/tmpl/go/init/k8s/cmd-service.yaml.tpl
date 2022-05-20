kind: Service
apiVersion: v1
metadata:
  name: cmd-service
  namespace: dapr-system
  labels:
    app: {{.Namespace}}-cmd-service
spec:
  type: NodePort
  selector:
    app: {{.Namespace}}-cmd-service
  ports:
    - port: 8080
      targetPort: 8080
      nodePort: 30001

---
kind: Deployment
apiVersion: apps/v1
metadata:
  name: {{.Namespace}}-cmd-service
  namespace: dapr-system
  labels:
    app: cmd-service
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{.Namespace}}-cmd-service
  template:
    metadata:
      labels:
        app: {{.Namespace}}-cmd-service
      annotations:
        dapr.io/enabled: "true"
        dapr.io/app-id: "{{.Namespace}}-cmd-service"
        dapr.io/app-port: "8080"
    spec:
      # hostNetwork: true
      containers:
        - name: cmd-service
          image: 192.168.64.12/{{.Namespace}}-cmd-service:dapr-linux-arm64
          ports:
          - containerPort: 8080

---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: cmd-service-ingress
  namespace: dapr-system
spec:
  rules:
  - http:
      paths:
      - path: /{{.Namespace}}/cmd/
        pathType: Prefix
        backend:
          service:
            name: {{.Namespace}}-cmd-service
            port:
              number: 8080