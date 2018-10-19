/*
Copyright 2018 The KubeSphere Authors.
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

package prometheus

import (
	"strings"

	"alert-kubesphere-plugin/pkg/models"
	"github.com/emicklei/go-restful"
)

func MakeWorkloadRule(metricsName string, resourceName models.ResourceName) string {
	// kube_pod_info{created_by_kind="DaemonSet",created_by_name="fluent-bit",endpoint="https-main",
	// host_ip="192.168.0.14",instance="10.244.114.187:8443",job="kube-state-metrics",
	// namespace="kube-system",node="i-k89a62il",pod="fluent-bit-l5vxr",
	// pod_ip="10.244.114.175",service="kube-state-metrics"}
	rule := `kube_pod_info{created_by_kind="$1",created_by_name=$2,namespace="$3"}`
	kind := resourceName.Workload.Kind
	name := resourceName.Workload.Name
	namespace := resourceName.Namespace

	// kind alertnatives values: Deployment StatefulSet ReplicaSet DaemonSet
	kind = strings.ToLower(kind)

	switch kind {
	case "deployment":
		kind = "ReplicaSet"
		if name != "" {
			name = "~\"" + name + ".*\""
		} else {
			name = "~\".*\""
		}
		rule = strings.Replace(rule, "$1", kind, -1)
		rule = strings.Replace(rule, "$2", name, -1)
		rule = strings.Replace(rule, "$3", namespace, -1)
		return rule
	case "replicaset":
		kind = "ReplicaSet"
	case "statefulset":
		kind = "StatefulSet"
	case "daemonset":
		kind = "DaemonSet"
	}

	if name == "" {
		name = "~\".*\""
	} else {
		name = "\"" + name + "\""
	}

	rule = strings.Replace(rule, "$1", kind, -1)
	rule = strings.Replace(rule, "$2", name, -1)
	rule = strings.Replace(rule, "$3", namespace, -1)
	return rule
}

func MakeWorkspacePromQL(metricsName string, resourceName models.ResourceName) string {
	promql := RulePromQLTmplMap[metricsName]
	promql = strings.Replace(promql, "$1", resourceName.Workspace, -1)
	return promql
}

func MakeContainerPromQL(metricsName string, resourceName models.ResourceName) string {
	nsName := resourceName.Namespace
	poName := resourceName.Pod
	containerName := resourceName.Container
	// metricsName container_cpu_utilisation  container_memory_utilisation container_memory_utilisation_wo_cache
	var promql = ""
	if containerName == "" {
		// all containers maybe use filter
		metricsName += "_all"
		promql = RulePromQLTmplMap[metricsName]
		promql = strings.Replace(promql, "$1", nsName, -1)
		promql = strings.Replace(promql, "$2", poName, -1)
		//containerFilter := strings.Trim(request.QueryParameter("containers_filter"), " ")
		containerFilter := ""
		if containerFilter == "" {
			containerFilter = ".*"
		}
		promql = strings.Replace(promql, "$3", containerFilter, -1)
		return promql
	}
	promql = RulePromQLTmplMap[metricsName]

	promql = strings.Replace(promql, "$1", nsName, -1)
	promql = strings.Replace(promql, "$2", poName, -1)
	promql = strings.Replace(promql, "$3", containerName, -1)
	return promql
}

func MakePodPromQL(metricsName string, resourceName models.ResourceName) string {
	nsName := resourceName.Namespace
	nodeID := ""
	podName := resourceName.Pod
	podFilter := ""
	var promql = ""
	if nsName != "" {
		// get pod metrics by namespace
		if podName != "" {
			// specific pod
			promql = RulePromQLTmplMap[metricsName]
			promql = strings.Replace(promql, "$1", nsName, -1)
			promql = strings.Replace(promql, "$2", podName, -1)

		} else {
			// all pods
			metricsName += "_all"
			promql = RulePromQLTmplMap[metricsName]
			if podFilter == "" {
				podFilter = ".*"
			}
			promql = strings.Replace(promql, "$1", nsName, -1)
			promql = strings.Replace(promql, "$2", podFilter, -1)
		}
	} else if nodeID != "" {
		// get pod metrics by nodeid
		metricsName += "_node"
		promql = RulePromQLTmplMap[metricsName]
		promql = strings.Replace(promql, "$3", nodeID, -1)
		if podName != "" {
			// specific pod
			promql = strings.Replace(promql, "$2", podName, -1)
		} else {
			promql = strings.Replace(promql, "$2", podFilter, -1)
		}
	}
	return promql
}

func MakeNamespacePromQL(metricsName string, resourceName models.ResourceName) string {
	var recordingRule = RulePromQLTmplMap[metricsName]
	recordingRule = strings.Replace(recordingRule, "$1", resourceName.Namespace, -1)
	return recordingRule
}

func MakeNodeorClusterRule(request *restful.Request, metricsName string) string {
	nodeID := request.PathParameter("node_id")
	var rule = RulePromQLTmplMap[metricsName]

	if strings.Contains(request.SelectedRoutePath(), "monitoring/cluster") {
		// cluster
		return rule
	} else {
		// node
		nodesFilter := strings.Trim(request.QueryParameter("nodes_filter"), " ")
		if nodesFilter == "" {
			nodesFilter = ".*"
		}
		if strings.Contains(metricsName, "disk_size") || strings.Contains(metricsName, "pod") || strings.Contains(metricsName, "usage") {
			// disk size promql
			if nodeID != "" {
				nodesFilter = "{" + "node" + "=" + "\"" + nodeID + "\"" + "}"
			} else {
				nodesFilter = "{" + "node" + "=~" + "\"" + nodesFilter + "\"" + "}"
			}
			rule = strings.Replace(rule, "$1", nodesFilter, -1)
		} else {
			// cpu, memory, network, disk_iops rules
			if nodeID != "" {
				// specific node
				rule = rule + "{" + "node" + "=" + "\"" + nodeID + "\"" + "}"
			} else {
				// all nodes or specific nodes filted with re2 syntax
				rule = rule + "{" + "node" + "=~" + "\"" + nodesFilter + "\"" + "}"
			}
		}
	}
	return rule
}
