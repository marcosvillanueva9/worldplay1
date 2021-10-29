package services

import (
	"database/sql"
	"fmt"
	"github.com/marcosvillanueva9/worldplay1/cmd/internal/models"
	//_ "github.com/go-sql-driver/mysql
)

var (
	DummyCarts = make(map[int]*models.Cart)
	DummyOrders = make(map[int]*models.Order)
)

func openDBConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func CreateCart(cartID int) error {
	if _, ok := DummyCarts[cartID]; ok {
		return fmt.Errorf("This cart already exists")
	}
	/* TODO add slq implementation
	db, err := openDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")
	defer insert.Close()
	*/
	DummyCarts[cartID] = &models.Cart{}
	return nil
}

func AddProductToCart(product models.Product, cartID int) error {
	cart, ok := DummyCarts[cartID]
	if !ok {
		return fmt.Errorf("Cart not founded")
	}
	/* TODO add slq implementation
	db, err := openDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")
	defer insert.Close()
	*/
	cart.Products = append(cart.Products, product)

	return nil
}

func GetCart(cartID int) (models.Cart, error) {
	/* TODO add slq implementation
	db, err := openDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")
	defer insert.Close()
	*/
	cart, ok := DummyCarts[cartID]
	if !ok {
		return models.Cart{}, fmt.Errorf("Cart not founded")
	}

	if cart == nil {
		return models.Cart{}, fmt.Errorf("Cart not founded")
	}
	return *cart, nil
}

func CreateOrder(orderID int) error {
	/* TODO add slq implementation
	db, err := openDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")
	defer insert.Close()
	*/
	if _, ok := DummyOrders[orderID]; ok {
		return fmt.Errorf("This cart already exists")
	}
	DummyOrders[orderID] = &models.Order{}
	return nil
}

func GetOrder(orderID int) (models.Order, error) {
	/* TODO add slq implementation
	db, err := openDBConnection()
	if err != nil {
		return err
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")
	defer insert.Close()
	*/
	order, ok := DummyOrders[orderID]
	if !ok {
		return models.Order{}, fmt.Errorf("Order not founded")
	}

	if order == nil {
		return models.Order{}, fmt.Errorf("Order not founded")
	}

	return *order, nil
}