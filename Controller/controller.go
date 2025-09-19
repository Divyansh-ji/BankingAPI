package Controller

import (
	"errors"
	"main/intializers"
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Postingfrom(c *gin.Context) {

	var fromaccc models.Account

	err := c.ShouldBindJSON(&fromaccc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	postone := models.Account{ID: fromaccc.ID, Owner: fromaccc.Owner, Balance: fromaccc.Balance}
	intializers.DB.Create(&postone)
	c.JSON(http.StatusCreated, gin.H{"postone": postone})

}

func Postingto(c *gin.Context) {

	var toacc models.Account

	err := c.ShouldBindJSON(&toacc)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	postone := models.Account{ID: toacc.ID, Owner: toacc.Owner, Balance: toacc.Balance}
	intializers.DB.Create(&postone)
	c.JSON(http.StatusCreated, gin.H{"postone": postone})

}

func transfer(fromID int, toID int, amount float64) error {

	return intializers.DB.Transaction(func(tx *gorm.DB) error {
		var fromacc models.Account
		var toacc models.Account
		err := tx.First(&fromacc, fromID).Error
		if err != nil {
			return err
		}

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&fromacc, fromID).Error; err != nil {
			return err
		}
		if fromacc.Balance < amount {
			return errors.New("insufficient funds")

		}
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			First(&toacc, fromID).Error; err != nil {
			return err
		}
		fromacc.Balance = fromacc.Balance - amount
		toacc.Balance = toacc.Balance + amount

		if err := tx.Save(&toacc).Error; err != nil {
			return err

		}
		if err := tx.Save(&fromacc).Error; err != nil {
			return err

		}
		return nil

	})
}
func Get(c *gin.Context) {
	id := c.Param("id")

	var fromacc []models.Account

	if err := intializers.DB.Find(&fromacc, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	c.JSON(http.StatusOK, gin.H{"accounts": fromacc})

}
