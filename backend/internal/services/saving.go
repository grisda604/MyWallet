package services

import (
	"errors"

	"github.com/Mywallet/internal/database"
	"github.com/Mywallet/internal/models"
)

func GetSavingsGoals() ([]models.SavingsGoal, error) {
	var savings []models.SavingsGoal
	if err := database.DB.Find(&savings).Error; err != nil {
		return nil, err
	}
	return savings, nil
}

func GetSavingsGoalByID(id string) (models.SavingsGoal, error) {
	var savings models.SavingsGoal
	if err := database.DB.First(&savings, id).Error; err != nil {
		return models.SavingsGoal{}, err
	}
	return savings, nil
}

func CreateSavingsGoal(savings *models.SavingsGoal) error {
	if savings.Name == "" {
		return errors.New("name is required")
	}
	if savings.TargetAmount < 0 {
		return errors.New("target amount must be greater than 0")
	}
	if savings.CurrentAmount < 0 {
		return errors.New("current amount must be greater than 0")
	}
	if savings.Deadline == nil {
		return errors.New("deadline is required")
	}
	if savings.Note == "" {
		return errors.New("note is required")
	}
	return database.DB.Create(&savings).Error
}

func UpdateSavingsGoal(id string, updatedate models.SavingsGoal) error {
	var savings models.SavingsGoal
	if err := database.DB.First(&savings, id).Error; err != nil {
		return err
	}
	if updatedate.Name != "" {
		savings.Name = updatedate.Name
	}
	if updatedate.TargetAmount != 0 {
		savings.TargetAmount = updatedate.TargetAmount
	}
	if updatedate.CurrentAmount != 0 {
		savings.CurrentAmount = updatedate.CurrentAmount
	}
	if updatedate.Deadline != nil {
		savings.Deadline = updatedate.Deadline
	}
	if updatedate.Note != "" {
		savings.Note = updatedate.Note
	}
	return database.DB.Save(&savings).Error
}

func DeleteSavingGoal(id string) error {
	var savings models.SavingsGoal
	if err := database.DB.First(&savings, id).Error; err != nil {
		return err
	}
	return database.DB.Delete(&savings).Error
}
