if [ -z "${AZURE_SUBSCRIPTION_ID}" ]; then
  echo "Please set environment variable AZURE_SUBSCRIPTION_ID"
  exit 0
fi

if [ -z "${AZURE_TENANT_ID}" ]; then
  echo "Please set environment variable AZURE_TENANT_ID"
  exit 0
fi


if [ -z "${AZURE_CLIENT_ID}" ]; then
  echo "Please set environment variable AZURE_CLIENT_ID"
  exit 0
fi


if [ -z "${AZURE_CLIENT_SECRET}" ]; then
  echo "Please set environment variable AZURE_CLIENT_SECRET"
  exit 0
fi

if [ -z "${CLUSTER_NAME}" ]; then
  echo "Please set environment variable CLUSTER_NAME"
  exit 0
fi

k3d cluster create mycluster

export AZURE_LOCATION="eastus"
export AZURE_ENVIRONMENT="AzurePublicCloud"
export AZURE_SUBSCRIPTION_ID_B64="$(echo -n "$AZURE_SUBSCRIPTION_ID" | base64 | tr -d '\n')"
export AZURE_TENANT_ID_B64="$(echo -n "$AZURE_TENANT_ID" | base64 | tr -d '\n')"
export AZURE_CLIENT_ID_B64="$(echo -n "$AZURE_CLIENT_ID" | base64 | tr -d '\n')"
export AZURE_CLIENT_SECRET_B64="$(echo -n "$AZURE_CLIENT_SECRET" | base64 | tr -d '\n')"

export EXP_CLUSTER_RESOURCE_SET=true


export PWD="$(pwd)"
mkdir -p ~/.cluster-api
cat samples/clusterctl.yaml | envsubst > ~/.cluster-api/clusterctl.yaml

clusterctl init --infrastructure maas --bootstrap k0s --control-plane k0s

kubectl wait --for=condition=Available --timeout=5m -n capi-system deployment/capi-controller-manager
kubectl wait --for=condition=Available --timeout=5m -n capi-k0s-control-plane-system deployment/capi-k0s-control-plane-controller-manager
kubectl wait --for=condition=Available --timeout=5m -n capz-system deployment/capz-controller-manager
kubectl wait --for=condition=Available --timeout=5m -n capi-k0s-bootstrap-system deployment/capi-k0s-bootstrap-controller-manager


cat samples/maas/k0s-template.yaml | envsubst > samples/maas/k0s-cluster.yaml
kubectl create configmap maas-ccm-addon --from-file=samples/maas/maas-ccm.yaml
kubectl create configmap maas-cn-addon --from-file=samples/maas/maas-cn.yaml
kubectl apply -f samples/maas/k0s-template.yaml
kubectl apply -f samples/maas/resource-set.yaml


echo "once the cluster is up run clusterctl get kubeconfig $CLUSTER_NAME > k0s.yaml or kubectl scale kzeroscontrolplane $CLUSTER_NAME-control-plane --replicas 3 for HA"