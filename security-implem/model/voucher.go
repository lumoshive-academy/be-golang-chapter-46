package model

type Voucher struct {
	Base
	VoucherName     string   `json:"voucher_name" gorm:"type:varchar(100);not null" binding:"required" example:"Summer Sale"`                                  // Name of the voucher
	VoucherCode     string   `json:"voucher_code" gorm:"type:varchar(50);unique;not null" binding:"required" example:"SUMMER2024"`                             // Unique code for the voucher
	VoucherType     string   `json:"voucher_type" gorm:"type:varchar(50);not null" binding:"required,oneof='e-commerce' 'redeem points'" example:"e-commerce"` // Type of voucher
	Description     string   `json:"description" gorm:"type:text;not null" binding:"required" example:"Get 20% off on all items"`                              // Terms and conditions
	VoucherCategory string   `json:"voucher_category" gorm:"type:varchar(50);not null" binding:"required,oneof='Free Shipping' 'Discount'" example:"Discount"` // Category of voucher
	MinimumPurchase float64  `json:"minimum_purchase,omitempty" gorm:"type:decimal(10,2)" example:"100000"`                                                    // Minimum purchase amount
	PaymentMethods  []string `json:"payment_methods,omitempty" gorm:"-" binding:"dive,required" example:"['COD', 'Debit Card']"`                               // Payment methods
	DiscountPercent float64  `json:"discount_percent,omitempty" gorm:"type:decimal(5,2)" example:"10.00"`                                                      // Discount percentage
	DiscountAmount  float64  `json:"discount_amount,omitempty" gorm:"type:decimal(10,2)" example:"50000"`                                                      // Discount fixed amount
	StartDate       string   `json:"start_date" gorm:"type:date;not null" binding:"required,datetime=2006-01-02" example:"2024-01-01"`                         // Start date of voucher validity
	ExpiryDate      string   `json:"expiry_date" gorm:"type:date;not null" binding:"required,datetime=2006-01-02" example:"2024-12-31"`                        // Expiry date of voucher validity
	ApplicableAreas []string `json:"applicable_areas" gorm:"-" binding:"dive,required" example:"['Jakarta', 'Bali']"`                                          // Areas where the voucher applies
	IsActive        bool     `json:"is_active" gorm:"type:boolean;default:true" example:"true"`                                                                // Status of the voucher (active/inactive)
}
