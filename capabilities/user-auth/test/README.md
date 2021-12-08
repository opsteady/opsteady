# User login trough pinniped

```bash
# chek that with your admin account you can
kubectl auth can-i get pods
# get pinniped config
pinniped get kubeconfig --oidc-listen-port 48095 --oidc-scopes offline_access,openid,profile --concierge-authenticator-name pinniped-cli --concierge-authenticator-type jwt > tmp.yaml
# things should work but you should get no as answer
kubectl auth can-i get pods --kubeconfig=tmp.yaml
```
