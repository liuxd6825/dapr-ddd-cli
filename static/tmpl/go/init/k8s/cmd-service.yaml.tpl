kind: Service
apiVersion: v1
metadata:
  name: {{.CommandServiceName}}
  namespace: {{.K8sNamespace}}
  labels:
    app: {{.CommandServiceName}}
spec:
  type: NodePort
  selector:
    app: {{.CommandServiceName}}
  ports:
    - port: 8080
      targetPort: 8080
      nodePort: 30001

---
kind: Deployment
apiVersion: apps/v1
metadata:
  name: {{.CommandServiceName}}
  namespace: {{.K8sNamespace}}
  labels:
    app: {{.CommandServiceName}}
spec:
  replicas: 1
  selector:
    matchLabels:
      app: {{.CommandServiceName}}
  template:
    metadata:
      labels:
        app: {{.CommandServiceName}}
      annotations:
        dapr.io/enabled: "true"
        dapr.io/app-id: "{{.CommandServiceName}}"
        dapr.io/app-port: "8080"
    spec:
      # hostNetwork: true
      containers:
        - name: {{.CommandServiceName}}
          image: {{.CommandImage}}
          ports:
          - containerPort: 8080

---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: {{.CommandServiceName}}-ingress
  namespace: {{.K8sNamespace}}
spec:
  rules:
  - http:
      paths:
      - path: /darp-cmd
        pathType: Prefix
        backend:
          service:
            name: {{.CommandServiceName}}
            port:
              number: 8080