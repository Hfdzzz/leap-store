package controllers

import (
	//"leap_store/app/controllers"
	//"leap_store/app/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	//"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupMockDB() (*gorm.DB, sqlmock.Sqlmock, error){
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	dialector := postgres.New(postgres.Config{
		Conn: sqlDB,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}

	return db, mock, nil
	
}

func CartHomeHandler_TestCreateCart(t *testing.T){
	// db, mock, err := SetupMockDB()
	// assert.NoError(t, err)
	
	// sqlDB, _ := db.DB()
	// defer sqlDB.Close()

	// mock.ExpectQuery("SELECT \\* FROM `carts` WHERE id = ?").WithArgs(1).WillReturnError(gorm.ErrRecordNotFound)

	// mock.ExpectBegin()
	// mock.ExpectExec("INSERT INTO `carts`").WithArgs(sqlmock.AnyArg(), 2, 100.0).WillReturnResult(sqlmock.NewResult(1, 1))
	// mock.ExpectCommit()

	// controller := CartController{DB: db}
	// cart := models.Cart{Amount_Product: 2, Total_Price:100.0}
	// err = controller.CartHomeHandler(cart)

	// assert.NoError(t, err)
	// assert.NoError(t, mock.ExpectationsWereMet())

}