MIME-Version: 1.0
Content-Type: multipart/mixed; boundary="==MYBOUNDARY=="

--==MYBOUNDARY==
Content-Type: text/x-shellscript; charset="us-ascii"

#!/bin/bash

/etc/eks/bootstrap.sh ${CLUSTER_NAME} \
--kubelet-extra-args '--max-pods=40' \
--b64-cluster-ca  ${B64_CLUSTER_CA} \
--apiserver-endpoint ${API_SERVER_URL}
--container-runtime containerd

--==MYBOUNDARY==--
