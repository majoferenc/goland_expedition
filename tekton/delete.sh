kubectl delete -f tekton/pipeline-account.yaml
kubectl delete -f tekton/pipeline/build-and-deploy-pipeline.yaml
kubectl delete -f tekton/resources/goland-expedition-git.yaml
kubectl delete -f tekton/tasks/deploy-using-helm.yaml
kubectl delete -f tekton/tasks/source-to-image.yaml
kubectl delete -f tekton/run/goland-expedition-pipeline-run.yaml