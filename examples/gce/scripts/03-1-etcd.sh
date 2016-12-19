#!/usr/bin/env bash

dir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

source "${dir}/helpers.bash"

set -e

KUBERNETES_HOSTS=(controller0 controller1 controller2)

for host in ${KUBERNETES_HOSTS[*]}; do
  gcloud compute copy-files helpers.bash 03-2-run-inside-vms-etcd.sh ${host}:~/
done
