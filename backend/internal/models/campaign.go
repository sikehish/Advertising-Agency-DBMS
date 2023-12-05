package models

import (
	"time"
)

type Campaign struct {
	CampaignID       int        `gorm:"primaryKey" json:"campaignID"`
    ClientID         int        `json:"clientID"`
    Name             string     `json:"name"`
    Budget           float64    `json:"budget"`
    StartDate        time.Time  `gorm:"type:date" json:"startDate"`
    EndDate          time.Time  `gorm:"type:date" json:"endDate"`
	CreativeDirector string         `json:"creativeDirector"`
	// CreatedAt        time.Time      `json:"createdAt"`
	// UpdatedAt        time.Time      `json:"updatedAt"`
	// DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
