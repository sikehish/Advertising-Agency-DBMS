package models

type Advertisement struct {
	AdvertisementID int    `gorm:"primaryKey" json:"advertisementId"`
	CampaignID      int    `json:"campaignId"`
	Type            string `json:"type"`
	Content         string `json:"content"`
	CreativeTeam    string `json:"creativeTeam"`
	// CreatedAt        time.Time       `json:"createdAt"`
	// UpdatedAt        time.Time       `json:"updatedAt"`
	// DeletedAt        gorm.DeletedAt  `gorm:"index" json:"-"`
}
