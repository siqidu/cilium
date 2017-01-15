//
// Copyright 2016 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package daemon

import (
	"fmt"

	"github.com/cilium/cilium/api/v1/models"
	"github.com/cilium/cilium/api/v1/server/restapi"

	ctx "golang.org/x/net/context"
)

func (d *Daemon) GetHealthz(params restapi.GetHealthzParams) *restapi.GetHealthzOK {
	sr := models.StatusResponse{}

	sr.Nodeaddress6 = d.conf.NodeAddress.String()

	if info, err := d.kvClient.Status(); err != nil {
		sr.Kvstore = &models.Status{State: models.StatusStateFailure, Msg: fmt.Sprintf("Err: %s - %s", err, info)}
	} else {
		sr.Kvstore = &models.Status{State: models.StatusStateOk, Msg: info}
	}

	if _, err := d.dockerClient.Info(ctx.Background()); err != nil {
		sr.Docker = &models.Status{State: models.StatusStateFailure, Msg: err.Error()}
	} else {
		sr.Docker = &models.Status{State: models.StatusStateOk, Msg: ""}
	}

	if d.conf.IsK8sEnabled() {
		if v, err := d.k8sClient.ServerVersion(); err != nil {
			sr.Kubernetes = &models.Status{State: models.StatusStateOk, Msg: err.Error()}
		} else {
			sr.Kubernetes = &models.Status{State: models.StatusStateOk, Msg: v.String()}
		}
	} else {
		sr.Kubernetes = &models.Status{State: models.StatusStateDisabled}
	}

	if sr.Kvstore.State != models.StatusStateOk {
		sr.Cilium = &models.Status{State: sr.Kvstore.State, Msg: "Kvstore service is not ready!"}
	} else if sr.Docker.State != models.StatusStateOk {
		sr.Cilium = &models.Status{State: sr.Docker.State, Msg: "Docker service is not ready!"}
	} else if d.conf.IsK8sEnabled() && sr.Kubernetes.State != models.StatusStateOk {
		sr.Cilium = &models.Status{State: sr.Kubernetes.State, Msg: "Kubernetes service is not ready!"}
	}

	// TODO Create a logstash status in its runnable function
	//Logstash   Status `json:"logstash"`

	//sr.IPAMStatus = d.DumpIPAM()

	return restapi.NewGetHealthzOK().WithPayload(&sr)
}
