package Controller

import (
	"errors"
	"main/intializers"
	"main/models"
	"net/http"
	"strconv"

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

type TransferRequest struct {
	FromID int     `json:"from_id" binding:"required"`
	ToID   int     `json:"to_id" binding:"required"`
	Amount float64 `json:"amount" binding:"required,gt=0"`
}

func TransferHandler(c *gin.Context) {
	var req TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := transfer(req.FromID, req.ToID, req.Amount); err != nil {
		// You can check for specific errors if you want to return different status codes
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
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
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var acc models.Account
	if err := intializers.DB.First(&acc, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "account not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"account": acc})
}
