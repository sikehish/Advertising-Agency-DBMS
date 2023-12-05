package models

import (
	"time"
)

type PerformanceMetric struct {
    MetricID          int        `gorm:"primaryKey" json:"metricID"`
    AdvertisementID   int        `json:"advertisementID"`
    Impressions       int        `json:"impressions"`
    Clicks            int        `json:"clicks"`
    Conversions       int        `json:"conversions"`
    Date              time.Time  `gorm:"type:date" json:"date"`
    // CreatedAt         time.Time  `json:"createdAt"`
    // UpdatedAt         time.Time  `json:"updatedAt"`
    // DeletedAt         gorm.DeletedAt  `gorm:"index" json:"-"`
}