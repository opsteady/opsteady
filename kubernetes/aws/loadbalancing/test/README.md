1. kubectl apply -f echoserver-deployment.yaml && kubectl apply -f echoserver-service.yaml && kubectl apply -f echoserver-ingress.yaml
2. Check the logs of the AWS load balancer to see that it created the load balancer without errors.
3. Check in the AWS management console what the DNS A record is of the load balancer. Resolve it to an IP address with dig.
4. Add a local host entry for http://echoserver.example.com/ that points to the load balancer DNS A record IP address.
5. Hit the above url and see the output.
6. kubectl delete -f echoserver-deployment.yaml && kubectl delete -f echoserver-service.yaml && kubectl delete -f echoserver-ingress.yaml
7. Check that the loadbalancer is cleaned up.
