# simple-httpapp
Simple HTTP Application that supports URIs`GET /status` and `GET /healthz`.

# Pre-Requisites
- Docker
- [kind-cluster](https://github.com/abhide/kind-clusters)

# Deployment
```bash
➜  simple-httpapp git:(master) ✗ make all 
go fmt ./
docker build -t simple-httpapp:latest ./
Sending build context to Docker daemon  70.66kB
Step 1/8 : FROM golang:alpine3.12
 ---> 8d4cbc6fcb0f
Step 2/8 : WORKDIR /go/src/github.com/abhide/simple-httpapp/
 ---> Using cache
 ---> e1833883d05a
Step 3/8 : COPY main.go .
 ---> Using cache
 ---> d3bc92526e45
Step 4/8 : RUN go build -o simple-httpapp ./main.go
 ---> Using cache
 ---> 575fbf5b356c
Step 5/8 : FROM alpine:3.12
 ---> 88dd2752d2ea
Step 6/8 : WORKDIR /root/
 ---> Using cache
 ---> 53679f3ecd74
Step 7/8 : COPY --from=0 /go/src/github.com/abhide/simple-httpapp/simple-httpapp .
 ---> Using cache
 ---> dc0cdee8d2c2
Step 8/8 : CMD ["./simple-httpapp"]
 ---> Using cache
 ---> d38d68577505
Successfully built d38d68577505
Successfully tagged simple-httpapp:latest
kind load docker-image simple-httpapp:latest --name=cluster01
kubectl create namespace v1 || true
namespace/v1 created
kubectl apply -f k8s/simpleapp-v1.yaml -n v1
configmap/simple-httpapp-v1-config created
deployment.apps/simple-httpapp-v1 created
service/simple-httpapp-v1-svc created
kubectl create namespace v2 || true
namespace/v2 created
kubectl apply -f k8s/simpleapp-v2.yaml -n v2
configmap/simple-httpapp-v2-config created
deployment.apps/simple-httpapp-v2 created
service/simple-httpapp-v2-svc created

➜  simple-httpapp git:(master) ✗ kubectl get pods -n v1
NAME                                 READY   STATUS    RESTARTS   AGE
simple-httpapp-v1-67c9574f65-bkn8k   1/1     Running   0          29s
simple-httpapp-v1-67c9574f65-k8c2z   1/1     Running   0          29s

➜  simple-httpapp git:(master) ✗ kubectl get pods -n v2
NAME                                 READY   STATUS    RESTARTS   AGE
simple-httpapp-v2-54f58fffc9-s99bj   1/1     Running   0          31s
simple-httpapp-v2-54f58fffc9-vgn6t   1/1     Running   0          31s

➜  simple-httpapp git:(master) ✗ kubectl get svc -n v2 
NAME                    TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
simple-httpapp-v2-svc   ClusterIP   10.96.146.89   <none>        8080/TCP   35s

➜  simple-httpapp git:(master) ✗ kubectl get svc -n v1
NAME                    TYPE        CLUSTER-IP    EXTERNAL-IP   PORT(S)    AGE
simple-httpapp-v1-svc   ClusterIP   10.96.59.91   <none>        8080/TCP   37s
```

# Delete deployment
```bash
➜  simple-httpapp git:(master) ✗ make clean-ns
kubectl delete namespace v1
namespace "v1" deleted
kubectl delete namespace v2
namespace "v2" deleted
```