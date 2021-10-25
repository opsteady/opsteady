1. kubectl apply -f test.yaml
2. Check the efs controller/node pods for successful creation of volume
3. kubectl get pvc
4. kubectl exec efs-reader -- cat /data/out | head -n 2
5. kubectl exec efs-writer -- cat /data/out | head -n 2
6. kubectl delete -f test.yaml
