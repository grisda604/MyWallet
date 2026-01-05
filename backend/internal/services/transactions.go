package services

import (
	"errors"

	"github.com/Mywallet/internal/database"
	"github.com/Mywallet/internal/models"
	"gorm.io/gorm"
)

func CreateTransaction(txData *models.Transaction) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		// 1. ตรวจสอบว่ามี Wallet นี้จริงไหม
		if txData.Amount <= 0 {
			return errors.New("amount must be greater than 0")
		}

		var wallet models.Wallet
		if err := tx.First(&wallet, txData.WalletID).Error; err != nil {
			return errors.New("wallet not found")
		}

		switch txData.Type {
		case "income":
			wallet.Balance += txData.Amount
		case "expense":
			if wallet.Balance < txData.Amount {
				return errors.New("insufficient balance")
			}
			wallet.Balance -= txData.Amount
		case "transfer":
			if txData.TargetWalletID == nil {
				return errors.New("target wallet is required for transfer")
			}
			wallet.Balance -= txData.Amount

			var targetWallet models.Wallet
			if err := tx.First(&targetWallet, *txData.TargetWalletID).Error; err != nil {
				return errors.New("target wallet not found")
			}
			targetWallet.Balance += txData.Amount

			if wallet.Balance < txData.Amount {
				return errors.New("insufficient balance")
			}

			if err := tx.Save(&targetWallet).Error; err != nil {
				return err
			}
		default:
			return errors.New("invalid transaction type")
		}
		// บันทึกยอดเงินใหม่ของ Wallet ต้นทาง
		if err := tx.Save(&wallet).Error; err != nil {
			return err
		}
		// บันทึก Transaction ลง Database
		if err := tx.Create(txData).Error; err != nil {
			return err
		}
		// ถ้าทุกอย่างผ่าน Return nil เพื่อ Commit Transaction
		return nil
	})
}

func UpdateTransaction(oldTx, newTx *models.Transaction) error {
	return database.DB.Transaction(func(dbTx *gorm.DB) error {

		// การเปลี่ยนแปลงทำการตรงกันข้ามกับการสร่้าง
		if err := reverseTransaction(dbTx, oldTx); err != nil {
			return err
		}
		// ทำการเปลี่ยนแปลงใหม่
		if err := ApplyTransaction(dbTx, newTx); err != nil {
			return err
		}
		return dbTx.Save(&newTx).Error
	})
}

func DeleteTransaction(tx *models.Transaction) error {
	return database.DB.Transaction(func(dbTx *gorm.DB) error {
		if err := reverseTransaction(dbTx, tx); err != nil {
			return err
		}
		return dbTx.Delete(tx).Error
	})
}

func reverseTransaction(dbTx *gorm.DB, tx *models.Transaction) error {
	switch tx.Type {
	case "income":
		return UpdateWalletBalance(dbTx, tx.WalletID, -tx.Amount)
	case "expense":
		return UpdateWalletBalance(dbTx, tx.WalletID, tx.Amount)
	case "transfer":
		if tx.TargetWalletID == nil {
			return errors.New("target wallet required for transfer")
		}
		if err := UpdateWalletBalance(dbTx, tx.WalletID, tx.Amount); err != nil {
			return err
		}
		if err := UpdateWalletBalance(dbTx, *tx.TargetWalletID, -tx.Amount); err != nil {
			return err
		}
	}
	return nil
}

func UpdateWalletBalance(dbTx *gorm.DB, walletID uint, amount float64) error {
	return dbTx.Model(&models.Wallet{}).Where("id = ?", walletID).Update("balance", gorm.Expr("balance + ?", amount)).Error
}

func ApplyTransaction(dbTx *gorm.DB, tx *models.Transaction) error {
	// transaction ต้องมี amount มากกว่า 0
	if tx.Amount <= 0 {
		return errors.New("amount must be greater than 0")
	}
	switch tx.Type {
	case "income":
		return UpdateWalletBalance(dbTx, tx.WalletID, tx.Amount)
	case "expense":
		return UpdateWalletBalance(dbTx, tx.WalletID, -tx.Amount)
	case "transfer":
		if tx.TargetWalletID == nil {
			return errors.New("target wallet required for transfer")
		}
		if err := UpdateWalletBalance(dbTx, tx.WalletID, tx.Amount); err != nil {
			return err
		}

		// ตรวจสอบยอดเงินพอหรือไม่
		if err := UpdateWalletBalance(dbTx, tx.WalletID, tx.Amount); err != nil {
			return err
		}
		return UpdateWalletBalance(dbTx, *tx.TargetWalletID, -tx.Amount)

	default:
		return errors.New("invalid transaction type")
	}

}

func UpdateWallet(dbTx *gorm.DB, WalletID uint, amount float64) error {
	return dbTx.Model(&models.Wallet{}).Where("id = ?", WalletID).Update("balance", gorm.Expr("balance + ?", amount)).Error
}
