package models

import (
	"kubesphere.io/ks-alert/pkg/utils/dbutil"
	"kubesphere.io/ks-alert/pkg/utils/idutil"
	"time"
)

type Metric struct {
	MetricID       string    `gorm:"primary_key"`
	MetricName     string    `gorm:"type:varchar(50);not null;unique"`
	ResourceTypeID string    `gorm:"type:varchar(50);not null;"`
	Unit           string    `gorm:"type:varchar(5);not null;"`
	CreatedAt      time.Time `gorm:"not null;"`
	UpdatedAt      time.Time `gorm:"not null;"`
}

func CreateMetric(metric *Metric) error {
	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}

	metric.MetricID = idutil.GetUuid36("metric-")

	err = db.Model(&Metric{}).Create(metric).Error
	return err
}

func GetMetricsByResourceTypeID(resourceTypeID string) (*[]Metric, error) {
	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}
	var Metrics []Metric
	db.Model(&Metric{}).Where(&Metric{ResourceTypeID: resourceTypeID}).Find(&Metrics)
	return &Metrics, err
}

func GetMetrics(resourceType *ResourceType) (*[]Metric, error) {

	resourceTypeID := resourceType.ResourceTypeID

	if resourceTypeID == "" {
		return nil, nil
	}

	return GetMetricsByResourceTypeID(resourceTypeID)
}

func GetMetricByMetricID(metricID string) (*Metric, error) {
	db, err := dbutil.DBClient()
	if err != nil {
		panic(err)
	}
	var Metrics Metric
	db.Model(&Metric{}).Where(&Metric{MetricID: metricID}).First(&Metrics)
	return &Metrics, err
}
