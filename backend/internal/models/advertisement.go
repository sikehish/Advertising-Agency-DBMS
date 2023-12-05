package models

type Advertisement struct {
	AdvertisementID int    `gorm:"primaryKey" json:"advertisementID"`
	CampaignID      int    `json:"campaignID"`
	Type            string `json:"type"`
	Content         string `json:"content"`
	CreativeTeam    string `json:"creativeTeam"`
	// CreatedAt        time.Time       `json:"createdAt"`
	// UpdatedAt        time.Time       `json:"updatedAt"`
	// DeletedAt        gorm.DeletedAt  `gorm:"index" json:"-"`
}

type AdvertisementPlacement struct {
	PlacementID      int     `gorm:"primaryKey" json:"placementID"`
	AdvertisementID  int     `json:"advertisementID"`
	PlacementDetails string  `json:"placementDetails"`
	Cost             float64 `json:"cost"`
	Duration         int     `json:"duration"`
	// CreatedAt         time.Time  `json:"createdAt"`
	// UpdatedAt         time.Time  `json:"updatedAt"`
	// DeletedAt         gorm.DeletedAt  `gorm:"index" json:"-"`
}