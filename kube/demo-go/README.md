 kubectl -n demo annotate deployment/demo-app-deployment "kubernetes.io/change-cause=$CAUSE" --overwrite
 194 kubectl apply -f .\deployment.yaml