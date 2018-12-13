package models

import (
	. "github.com/smartystreets/goconvey/convey"
	"kubesphere.io/alert-kubesphere-plugin/pkg/utils/idutil"
	"testing"
	"time"
)

func TestCreateMetric(t *testing.T) {
	Convey("add metrics for a kind of resource type", t, func() {
		// get product id
		var err error
		ent, err := GetEnterprise(&Enterprise{EnterpriseName: "北京优帆科技有限公司武汉分公司"})
		So(err, ShouldBeNil)

		product, err := GetProducts(&Product{EnterpriseID: ent.EnterpriseID})
		So(err, ShouldBeNil)
		So(len(*product), ShouldEqual, 1)

		productID := (*product)[0].ProductID

		resourceTypes := GetResourceTypeByProductID(productID)

		// add various metrics for a specific resource type
		for _, tp := range *resourceTypes {
			name := tp.ResourceTypeName
			id := tp.ResourceTypeID

			addMetricsForResource(name, id)
		}
	})
}

func addMetricsForResource(name, id string) {
	switch name {
	case "clusters":
		/**
		"cluster_cpu_utilisation",
		"cluster_cpu_usage",
		"cluster_cpu_total",
		"cluster_memory_utilisation",
		"cluster_memory_available",
		*/
		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "cluster_cpu_utilisation",
			Unit:           "%",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})

		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "cluster_cpu_usage",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})

		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "cluster_cpu_total",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})

		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "cluster_memory_utilisation",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})

		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "cluster_memory_available",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})
	case "nodes":
		/**
		"node_net_utilisation",
		"node_net_bytes_transmitted",
		"node_net_bytes_received",
		"node_disk_read_iops",
		"node_disk_write_iops",
		*/
		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "node_net_utilisation",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})

		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "node_net_bytes_transmitted",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})

		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "node_net_bytes_received",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})

		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "node_disk_read_iops",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})

		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "node_disk_write_iops",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})
	case "namespaces":
		/**
		"namespace_cpu_usage",
		"namespace_memory_usage",
		"namespace_memory_usage_wo_cache",
		"namespace_net_bytes_transmitted",
		*/
		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "namespace_cpu_usage",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})

		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "namespace_memory_usage",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})

		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "namespace_memory_usage_wo_cache",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})

		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "namespace_net_bytes_transmitted",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})
	case "pods":
		/**
		"pod_cpu_usage",
		"pod_memory_usage",
		"pod_memory_usage_wo_cache",
		"pod_net_bytes_transmitted",
		"pod_net_bytes_received",
		*/
		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "pod_cpu_usage",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})

		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "pod_memory_usage",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})

		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "pod_memory_usage_wo_cache",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})

		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "pod_net_bytes_transmitted",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})

		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "pod_net_bytes_received",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})
	case "containers":
		/**
		"container_cpu_usage",
		"container_memory_usage",
		"container_memory_usage_wo_cache",
		*/
		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "container_cpu_usage",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})

		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "container_memory_usage",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})

		CreateMetric(&Metric{
			MetricID:       idutil.GetUuid36("metric-"),
			ResourceTypeID: id,
			MetricName:     "container_memory_usage_wo_cache",
			Unit:           "",
			UpdatedAt:      time.Now(),
			CreatedAt:      time.Now(),
		})
	}
}
