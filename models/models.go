package models

type Brand struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

type Voucher struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	BrandID     uint   `gorm:"not null" json:"brand_id"`
	Name        string `gorm:"not null" json:"name"`
	CostInPoint int    `gorm:"not null" json:"cost_in_point"`
	Brand       Brand  `gorm:"foreignKey:BrandID;references:ID" json:"brand"`
}

type Transaction struct {
	ID           uint                `gorm:"primaryKey" json:"id"`
	CustomerName string              `gorm:"not null" json:"customer_name"`
	TotalPoints  int                 `gorm:"not null" json:"total_points"`
	Details      []TransactionDetail `gorm:"foreignKey:TransactionID" json:"details"`
}

type TransactionDetail struct {
	ID            uint    `gorm:"primaryKey" json:"id"`
	TransactionID uint    `gorm:"not null" json:"transaction_id"`
	VoucherID     uint    `gorm:"not null" json:"voucher_id"`
	Quantity      int     `gorm:"not null" json:"quantity"`
	TotalCost     int     `gorm:"not null" json:"total_cost"`
	Voucher       Voucher `gorm:"foreignKey:VoucherID" json:"voucher"`
}
