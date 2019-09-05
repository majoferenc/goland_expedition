# Goland Expedition
Experimental microservice with Go. API fetching Elastichsearch index for data.

### Features
 - GraphQL API with GraphiQL playground
 - Skaffold CI/CD
 - Draft CI/CD
 - Docker image
 - Docker-compose ready
 - K8s Helm chart

### Deployment

You can run Goland API automatic Docker image build and K8s Helm chart push after every source change with:

```sh
skaffold dev
```

If you want also publish Docker image to your container registry you can use:

```sh
draft up
```
