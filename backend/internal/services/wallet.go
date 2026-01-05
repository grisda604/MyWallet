package services

import (
	"errors"

	"github.com/Mywallet/internal/database"
	"github.com/Mywallet/internal/models"
)

func GetWallers() ([]models.Wallet, error) {
	var wallet []models.Wallet
	if err := database.DB.Find(&wallet).Error; err != nil {
		return nil, err
	}
	return wallet, nil
}

func GetWalletByID(id string) (models.Wallet, error) {
	var wallet models.Wallet
	if err := database.DB.First(&wallet, id).Error; err != nil {
		return models.Wallet{}, err
	}
	return wallet, nil
}

func CreateWallet(wallet *models.Wallet) error {
	if wallet.Name == "" {
		return errors.New("name is required")
	}
	if wallet.Balance < 0 {
		return errors.New("balance must be greater than 0")
	}
	var existingWallet models.Wallet
	if err := database.DB.Where("name = ?", wallet.Name).First(&existingWallet).Error; err == nil {
		return errors.New("wallet name already exists")
	}
	return database.DB.Create(&wallet).Error
}

func ModifyWallet(id string, updateData *models.Wallet) error {
	var wallet models.Wallet
	if err := database.DB.First(&wallet, id).Error; err != nil {
		return errors.New("wallet not found")
	}
	if updateData.Name != "" {
		var existingWallet models.Wallet
		if err := database.DB.Where("name = ?", updateData.Name).First(&existingWallet).Error; err == nil {
			return errors.New("wallet name already exists")
		}
	}
	// validate
	if updateData.Name != "" {
		wallet.Name = updateData.Name
	}
	return database.DB.Save(&wallet).Error
}

func DeleteWallet(id string) error {
	var wallet models.Wallet
	if err := database.DB.First(&wallet, id).Error; err != nil {
		return errors.New("wallet not found")
	}
	return database.DB.Delete(&wallet).Error
}
