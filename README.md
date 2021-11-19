# Cluster API k0s from Eupraxia Labs and Mirantis

Cluster API (CAPI) Bootstrap Provider k0s (CABP0) is a component of [Cluster API](https://github.com/kubernetes-sigs/cluster-api/blob/master/README.md) that is responsible for generating a cloud-init script to turn a Machine into a Kubernetes Node; this implementation brings up a lightweight Kubernetes cluster distribution [k0s](https://k0sproject.io/) from Mirantis.

CABP0 is the bootstrap component of Cluster API for k0s and brings in the following CRDS and controllers:
- k0s bootstrap provider (KZeros, KZerosTemplate)

Cluster API ControlPlane provider k0s (CACP0) is a component of [Cluster API](https://github.com/kubernetes-sigs/cluster-api/blob/master/README.md) that is responsible for managing the lifecycle of control plane machines for k0s; this implementation brings up the lightweight [k0s](https://k0sproject.io/) clusters for the edge.

CACP0 is the controlplane component of Cluster API for k0s and brings in the following CRDS and controllers:
- k0s controlplane provider (KZerosControlPlane)

## Testing it out.

**Warning**: This project and documentation are in an early stage. There is an assumption that an user of this provider is already familiar with **Cluster API (CAPI)**.  


### Prerequisites

Check out the [ClusterAPI Quickstart](https://cluster-api.sigs.k8s.io/user/quick-start.html) page to see the prerequisites for Cluster API.

Three main pieces are 

1. Bootstrap cluster. In the `samples/azure/azure-setup.sh` script, We run [k0s in Docker ](https://github.com/k0sproject/k0s/blob/main/docs/k0s-in-docker.md), but feel free to use [kind](https://kind.sigs.k8s.io/) as well.
2. 'clusterctl'. Please check out [ClusterAPI Quickstart](https://cluster-api.sigs.k8s.io/user/quick-start.html) for instructions.
3. Azure Service Principals. For more information go to [CAPZ Getting Started](https://github.com/kubernetes-sigs/cluster-api-provider-azure/blob/master/docs/getting-started.md)

CABP0 has been tested only on with an Azure and AzureStackHCI environment. To try out the Azure flow, fork the repo and look at `samples/azure/azure-setup.sh`.

CACP0 is available! Sample now includes the K0s Control Plane Provider. If you run the sample script you will get a cluster with a control plane and two workers.

Then run the following to scale the control plane...
```sh
kubectl scale kzeroscontrolplane ${CLUSTER_NAME}-control-plane --replicas 3
```

### Known Issues

Leak kubeconfig after cluster deletion. This is because the bootstrap provider is generating the kubeconfig until we have a control plane provider. 

## Roadmap

* Support for External Databases
* Fix Token Logic
* Setup CAPA and CAPV samples
* Post an issue!

