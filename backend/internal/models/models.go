package models

import (
	"time"

	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Name     string  `json:"name" gorm:"not null"`
	Type     string  `json:"type" gorm:"not null"`
	Balance  float64 `json:"balance" gorm:"default:0"`
	Color    string  `json:"color"`
	Icon     string  `json:"icon"`
	Currency string  `json:"currency" gorm:"default:'THB'"`
}

type Category struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null"`
	Type  string `json:"type" gorm:"not null"`
	Color string `json:"color"`
	Icon  string `json:"icon"`

	Budgets []Budget `json:"budgets" gorm:"foreignKey:CategoryID"`
}

type Budget struct {
	gorm.Model
	CategoryID uint     `json:"category_id" gorm:"not null"`
	Category   Category `json:"category" gorm:"foreignKey:CategoryID"`

	Amount float64 `json:"amount" gorm:"not null"` // ยอดที่ตั้งไว้
	Month  string  `json:"month" gorm:"not null"`  // Format: "YYYY-MM" (2025-01)
}

type Transaction struct {
	gorm.Model
	Date   time.Time `json:"date"`
	Amount float64   `json:"amount" gorm:"not null"`
	Type   string    `json:"type" gorm:"not null"` // "income", "expense", "transfer"
	Note   string    `json:"note"`

	WalletID uint   `json:"wallet_id" gorm:"not null"`
	Wallet   Wallet `json:"wallet" gorm:"foreignKey:WalletID"`

	TargetWalletID *uint     `json:"target_wallet_id"`
	TargetWallet   *Wallet   `json:"target_wallet" gorm:"foreignKey:TargetWalletID"`
	CategoryID     *uint     `json:"category_id"`
	Category       *Category `json:"category" gorm:"foreignKey:CategoryID"`
}

type SavingsGoal struct {
	gorm.Model
	Name          string     `json:"name" gorm:"not null"`
	TargetAmount  float64    `json:"target_amount" gorm:"not null"`
	CurrentAmount float64    `json:"current_amount" gorm:"default:0"`
	Color         string     `json:"color"`
	Icon          string     `json:"icon"`
	Deadline      *time.Time `json:"deadline"` // วันที่ครบกำหนด (Optional)
}
