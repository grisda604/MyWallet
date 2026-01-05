package services

import (
	"errors"

	"github.com/Mywallet/internal/database"
	"github.com/Mywallet/internal/models"
)

func GetCategories() ([]models.Category, error) {
	var categories []models.Category
	database.DB.Find(&categories)
	return categories, nil
}

func GetCategoryById(id string) (models.Category, error) {
	var category models.Category
	database.DB.First(&category, id)
	return category, nil
}

func CreateCategory(category *models.Category) error {
	if category.Name == "" {
		return errors.New("name is required")
	}
	if category.Type == "" {
		return errors.New("type is required")
	}
	if category.Type != "income" && category.Type != "expense" {
		return errors.New("type must be income or expense")
	}
	database.DB.Create(&category)
	return nil
}

func UpdateCategory(id string, updateData *models.Category) error {
	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		return errors.New("category not found")
	}
	// validate
	if updateData.Name != "" {
		category.Name = updateData.Name
	}
	if updateData.Type != "" {
		if updateData.Type != "income" && updateData.Type != "expense" {
			return errors.New("type must be income or expense")
		}
		category.Type = updateData.Type
	}
	if updateData.Color != "" {
		category.Color = updateData.Color
	}
	if updateData.Icon != "" {
		category.Icon = updateData.Icon
	}
	return database.DB.Save(&category).Error
}

func DeleteCategory(id string) error {
	var category models.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		return errors.New("category not found")
	}
	return database.DB.Delete(&category).Error
}
