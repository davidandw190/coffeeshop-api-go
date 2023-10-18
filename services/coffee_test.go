package services

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAllCoffees(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open a stub database connection: %v", err)
	}

	defer db.Close()

	t.Run("TestGetAllCoffees", func(t *testing.T) {
		t.Parallel()

		// Define test data
		expectedRows := sqlmock.NewRows([]string{"id", "name", "image", "roast", "region", "price", "grind_unit", "created_at", "updated_at"}).
			AddRow("1", "TestCoffee1", "image1.jpg", "Medium", "Brazil", 5.99, 1, time.Now(), time.Now()).
			AddRow("2", "TestCoffee2", "image2.jpg", "Dark", "Colombia", 7.99, 2, time.Now(), time.Now())

		// Set up the SQL mock expectations for the SELECT query
		mock.ExpectQuery("^SELECT").WillReturnRows(expectedRows)

		// Call the function and get the result
		coffees, err := GetAllCoffees(db, 5*time.Second)
		if err != nil {
			t.Fatalf("GetAllCoffees error: %v", err)
		}

		if len(coffees) != 2 {
			t.Errorf("Expected 2 coffees, got %d", len(coffees))
		}
	})

}
