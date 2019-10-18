# Goland Expedition
Experimental microservice with Go. API fetching Elastichsearch index for data.

### Features
 - GraphQL API with GraphiQL playground for public API
 - GRPC API for fast communication between microservices
 - Skaffold CI/CD
 - Draft CI/CD
 - Docker image
 - Docker-compose ready
 - K8s Helm chart
 
 
### GRPC interface development

Prerequisites:
```sh
go get -u google.golang.org/grpc 
go get -u github.com/golang/protobuf/protoc-gen-go 
```

When you change something at *.proto files, they need to be recompiled to *.pb.go with command:
```sh
protoc --go_out=plugins=grpc:. *.prot
```

### Test GRPC using GRPCUI

Prerequisites:
```sh
go get github.com/fullstorydev/grpcui 
go install github.com/fullstorydev/grpcui/cmd/grpcui
```

You can connect to your GRPC API via command:
```sh
 grpcui -plaintext localhost:4040
```

### Deployment

You can run Goland API automatic Docker image build and K8s Helm chart push after every source change with:

```sh
skaffold dev
```

If you want also publish Docker image to your container registry you can use:

```sh
draft up
```
