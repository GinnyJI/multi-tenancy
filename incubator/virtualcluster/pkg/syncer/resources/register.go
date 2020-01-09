/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resources

import (
	"k8s.io/client-go/informers"
	clientset "k8s.io/client-go/kubernetes"

	"github.com/kubernetes-sigs/multi-tenancy/incubator/virtualcluster/pkg/syncer/manager"
	"github.com/kubernetes-sigs/multi-tenancy/incubator/virtualcluster/pkg/syncer/resources/configmap"
	"github.com/kubernetes-sigs/multi-tenancy/incubator/virtualcluster/pkg/syncer/resources/endpoints"
	"github.com/kubernetes-sigs/multi-tenancy/incubator/virtualcluster/pkg/syncer/resources/event"
	"github.com/kubernetes-sigs/multi-tenancy/incubator/virtualcluster/pkg/syncer/resources/namespace"
	"github.com/kubernetes-sigs/multi-tenancy/incubator/virtualcluster/pkg/syncer/resources/node"
	"github.com/kubernetes-sigs/multi-tenancy/incubator/virtualcluster/pkg/syncer/resources/pod"
	"github.com/kubernetes-sigs/multi-tenancy/incubator/virtualcluster/pkg/syncer/resources/secret"
	"github.com/kubernetes-sigs/multi-tenancy/incubator/virtualcluster/pkg/syncer/resources/service"
	"github.com/kubernetes-sigs/multi-tenancy/incubator/virtualcluster/pkg/syncer/resources/serviceaccount"
)

func Register(client clientset.Interface, informerFactory informers.SharedInformerFactory, controllerManager *manager.ControllerManager) {
	namespace.Register(client.CoreV1(), informerFactory.Core().V1().Namespaces(), controllerManager)
	pod.Register(client.CoreV1(), informerFactory.Core().V1(), controllerManager)
	configmap.Register(client.CoreV1(), informerFactory.Core().V1().ConfigMaps(), controllerManager)
	secret.Register(client.CoreV1(), informerFactory.Core().V1().Secrets(), controllerManager)
	serviceaccount.Register(client.CoreV1(), informerFactory.Core().V1().ServiceAccounts(), controllerManager)
	node.Register(client.CoreV1(), informerFactory.Core().V1().Nodes(), controllerManager)
	service.Register(client.CoreV1(), informerFactory.Core().V1().Services(), controllerManager)
	endpoints.Register(client.CoreV1(), informerFactory.Core().V1().Endpoints(), controllerManager)
	event.Register(client.CoreV1(), informerFactory.Core().V1(), controllerManager)
}