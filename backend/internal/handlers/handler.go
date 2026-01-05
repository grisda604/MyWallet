package handlers

import (
	"github.com/Mywallet/internal/database"
	"github.com/Mywallet/internal/models"
	"github.com/Mywallet/internal/services"
	"github.com/gofiber/fiber/v2"
)

func GetTransactions(c *fiber.Ctx) error {
	var transactions []models.Transaction
	database.DB.Preload("Wallet").Preload("Category").Preload("TargetWallet").Order("date desc").Find(&transactions)
	return c.JSON(fiber.Map{"transactions": transactions})
}

func CreateTransaction(c *fiber.Ctx) error {
	transaction := new(models.Transaction)
	if err := c.BodyParser(transaction); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}

	err := services.CreateTransaction(transaction)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(201).JSON(fiber.Map{"message": "Transaction created successfully"})

}

func UpdateTransaction(c *fiber.Ctx) error {
	id := c.Params("id")

	oldTx := new(models.Transaction)
	if err := database.DB.First(&oldTx, id).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}
	transaction := new(models.Transaction)
	if err := c.BodyParser(transaction); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}
	err := services.UpdateTransaction(oldTx, transaction)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(transaction)
}

func DeleteTransaction(c *fiber.Ctx) error {
	id := c.Params("id")
	transaction := new(models.Transaction)
	if err := database.DB.First(&transaction, id).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}
	err := services.DeleteTransaction(transaction)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Transaction deleted successfully"})
}

func GetWallets(c *fiber.Ctx) error {
	var wallets []models.Wallet
	database.DB.Find(&wallets)
	return c.JSON(wallets)
}

func CreateWallet(c *fiber.Ctx) error {
	wallet := new(models.Wallet)
	if err := c.BodyParser(wallet); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}
	database.DB.Create(&wallet)
	return c.JSON(wallet)
}

func UpdateWallet(c *fiber.Ctx) error {
	id := c.Params("id")
	wallet := new(models.Wallet)
	if err := database.DB.First(&wallet, id).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}
	if err := c.BodyParser(wallet); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}
	database.DB.Save(&wallet)
	return c.JSON(wallet)
}

func DeleteWallet(c *fiber.Ctx) error {
	id := c.Params("id")
	wallet := new(models.Wallet)
	if err := database.DB.First(&wallet, id).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}
	err := services.DeleteWallet(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Wallet deleted successfully"})
}

func GetCategories(c *fiber.Ctx) error {
	var categories []models.Category
	database.DB.Find(&categories)
	return c.JSON(categories)
}

func CreateCategory(c *fiber.Ctx) error {
	Category := new(models.Category)
	if err := c.BodyParser(Category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}
	err := services.CreateCategory(Category)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(Category)
}

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	Category := new(models.Category)
	if err := c.BodyParser(Category); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}
	updateData := new(models.Category)
	if err := database.DB.First(&updateData, id).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}
	database.DB.Save(&updateData)
	return c.JSON(updateData)
}

func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	Category := new(models.Category)
	if err := database.DB.First(&Category, id).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid data"})
	}

	err := services.DeleteCategory(id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Category deleted successfully"})
}
