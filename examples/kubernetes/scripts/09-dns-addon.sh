#!/usr/bin/env bash

dir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

source "${dir}/helpers.bash"

set -e

kubectl create -f https://raw.githubusercontent.com/cilium/cilium/changing-k8s-installation/examples/kubernetes/deployments/kubedns-svc.yaml

kubectl create -f https://raw.githubusercontent.com/cilium/cilium/changing-k8s-installation/examples/kubernetes/deployments/kubedns-rc.yaml

kubectl create -f https://raw.githubusercontent.com/cilium/cilium/changing-k8s-installation/examples/kubernetes/deployments/kube-dashboard.yaml

kubectl --namespace=kube-system get svc
kubectl --namespace=kube-system get pods
