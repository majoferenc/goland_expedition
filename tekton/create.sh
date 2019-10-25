kubectl apply -f tekton/pipeline-account.yaml
kubectl apply -f tekton/pipeline/build-and-deploy-pipeline.yaml
kubectl apply -f tekton/resources/goland-expedition-git.yaml
kubectl apply -f tekton/tasks/deploy-using-helm.yaml
kubectl apply -f tekton/tasks/source-to-image.yaml
kubectl create -f tekton/run/goland-expedition-pipeline-run.yaml