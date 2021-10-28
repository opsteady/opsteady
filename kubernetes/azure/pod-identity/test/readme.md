# How to test AAD pod identity

<https://github.com/Azure/aad-pod-identity/issues/798>

```bash
export IDENTITY_RESOURCE_GROUP=platform-nodes-plt1
export IDENTITY_NAME="demo"
az identity create -g ${IDENTITY_RESOURCE_GROUP} -n ${IDENTITY_NAME}
export IDENTITY_CLIENT_ID="$(az identity show -g ${IDENTITY_RESOURCE_GROUP} -n ${IDENTITY_NAME} --query clientId -otsv)"
export IDENTITY_RESOURCE_ID="$(az identity show -g ${IDENTITY_RESOURCE_GROUP} -n ${IDENTITY_NAME} --query id -otsv)"

cat <<EOF | kubectl apply -f -
apiVersion: "aadpodidentity.k8s.io/v1"
kind: AzureIdentity
metadata:
  name: ${IDENTITY_NAME}
  namespace: platform
spec:
  type: 0
  resourceID: ${IDENTITY_RESOURCE_ID}
  clientID: ${IDENTITY_CLIENT_ID}
EOF

cat <<EOF | kubectl apply -f -
apiVersion: "aadpodidentity.k8s.io/v1"
kind: AzureIdentityBinding
metadata:
  name: ${IDENTITY_NAME}-binding
  namespace: platform
spec:
  azureIdentity: ${IDENTITY_NAME}
  selector: ${IDENTITY_NAME}
EOF
```

Wait at least 2 minutes because it doesn't seem to find the role directly

```bash
cat << EOF | kubectl apply -f -
apiVersion: v1
kind: Pod
metadata:
  name: demo
  namespace: platform
  labels:
    aadpodidbinding: $IDENTITY_NAME
spec:
  containers:
  - name: demo
    image: mcr.microsoft.com/oss/azure/aad-pod-identity/demo:v1.8.4
    args:
      - --subscription-id=${SUBSCRIPTION_ID}
      - --resource-group=${IDENTITY_RESOURCE_GROUP}
      - --identity-client-id=${IDENTITY_CLIENT_ID}
  nodeSelector:
    kubernetes.io/os: linux
EOF
```

## Verify

```bash
kubectl -n platform logs demo
```

This can take some time you should eventually see something like:
`I0510 18:17:04.474610 1 main.go:50] Try decoding your token <JWT token> at https://jwt.io`

## Cleanup

```bash
kubectl -n platform delete pod demo
kubectl -n platform delete azureidentity ${IDENTITY_NAME}
kubectl -n platform delete azureidentitybinding ${IDENTITY_NAME}-binding
az identity delete -g ${IDENTITY_RESOURCE_GROUP} -n ${IDENTITY_NAME}
```
