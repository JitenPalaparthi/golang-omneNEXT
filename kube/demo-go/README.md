 kubectl -n demo annotate deployment/demo-app-deployment "kubernetes.io/change-cause=$CAUSE" --overwrite
 194 kubectl apply -f .\deployment.yaml

 podman run -p 9000:9000 -p 9001:9001 quay.io/minio/minio server /data --console-address ":9001"