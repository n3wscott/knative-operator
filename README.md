# knative-operator
WIP for Knative Operator

This is very rough. It really was trying to see how much the controller can
reconcile from semi-dynamic yaml config. 

To try it out:

- Start with a clean cluster. (for gke: `./hack/operator-cluster.sh`)
- Install the operator. (`./hack/operator-install.sh`)
- Request Istio. (`kubectl apply -f deploy/crds/example/knative_v1alpha1_knativeistio_cr.yaml`)
- Request Knative (`kubectl apply -f deploy/crds/example/knative_v1alpha1_knative_cr.yaml`)
