package services

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func setupTestDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open a stub database connection: %v", err)
	}
	return db, mock
}

// func TestGetAllCoffees(t *testing.T) {
// 	t.Parallel()

// 	t.Run("Successful Retrieval", func(t *testing.T) {
// 		// Test retrieving a list of coffee products successfully.
// 		db, mock := setupTestDB(t)
// 		defer db.Close()

// 		expectedCoffee1 := &Coffee{
// 			ID:        "1",
// 			Name:      "TestCoffee1",
// 			Image:     "image1.jpg",
// 			Roast:     "Medium",
// 			Region:    "Brazil",
// 			Price:     5.99,
// 			GrindUnit: 1,
// 		}
// 		expectedCoffee2 := &Coffee{
// 			ID:        "2",
// 			Name:      "TestCoffee2",
// 			Image:     "image2.jpg",
// 			Roast:     "Dark",
// 			Region:    "Colombia",
// 			Price:     7.99,
// 			GrindUnit: 2,
// 		}

// 		// Define mock rows with expected data.
// 		expectedRows := sqlmock.NewRows([]string{"id", "name", "image", "roast", "region", "price", "grind_unit", "created_at", "updated_at"}).
// 			AddRow(expectedCoffee1.ID, expectedCoffee1.Name, expectedCoffee1.Image, expectedCoffee1.Roast, expectedCoffee1.Region, expectedCoffee1.Price, expectedCoffee1.GrindUnit, time.Now(), time.Now()).
// 			AddRow(expectedCoffee2.ID, expectedCoffee2.Name, expectedCoffee2.Image, expectedCoffee2.Roast, expectedCoffee2.Region, expectedCoffee2.Price, expectedCoffee2.GrindUnit, time.Now(), time.Now())

// 		mock.ExpectQuery("^SELECT").WillReturnRows(expectedRows)

// 		// Call the function and check the results.
// 		coffees, err := GetAllCoffees(db, 5*time.Second)
// 		if err != nil {
// 			t.Fatalf("GetAllCoffees error: %v", err)
// 		}

// 		if len(coffees) != 2 {
// 			t.Errorf("Expected 2 coffees, got %d", len(coffees))
// 		}

// 		// Check each coffee's attributes
// 		if coffees[0].ID != expectedCoffee1.ID || coffees[0].Name != expectedCoffee1.Name {
// 			t.Errorf("Mismatch in coffee data: expected %+v, got %+v", expectedCoffee1, coffees[0])
// 		}
// 		if coffees[1].ID != expectedCoffee2.ID || coffees[1].Name != expectedCoffee2.Name {
// 			t.Errorf("Mismatch in coffee data: expected %+v, got %+v", expectedCoffee2, coffees[1])
// 		}
// 	})

// 	t.Run("Empty Database", func(t *testing.T) {
// 		// Test when the database is empty, and no coffee products are retrieved.
// 		db, mock := setupTestDB(t)
// 		defer db.Close()

// 		mock.ExpectQuery("^SELECT").WillReturnRows(sqlmock.NewRows([]string{}))

// 		coffees, err := GetAllCoffees(db, 5*time.Second)
// 		if err != nil {
// 			t.Fatalf("GetAllCoffees error: %v", err)
// 		}

// 		if len(coffees) != 0 {
// 			t.Errorf("Expected 0 coffees, got %d", len(coffees))
// 		}
// 	})

// 	t.Run("Database Error", func(t *testing.T) {
// 		// Test when an error occurs while retrieving coffee products.
// 		db, mock := setupTestDB(t)
// 		defer db.Close()

// 		mock.ExpectQuery("^SELECT").WillReturnError(sql.ErrNoRows)

// 		if _, err := GetAllCoffees(db, 5*time.Second); err == nil {
// 			t.Error("Expected an error, but got nil")
// 		}
// 	})

// 	t.Run("Timeout Error", func(t *testing.T) {
// 		// Test when a database query times out.
// 		db, mock := setupTestDB(t)
// 		defer db.Close()

// 		mock.ExpectQuery("^SELECT").WillReturnError(context.DeadlineExceeded)

// 		if _, err := GetAllCoffees(db, 1*time.Second); err == nil {
// 			t.Error("Expected a timeout error, but got nil")
// 		}
// 	})

// 	t.Run("Scan Error", func(t *testing.T) {
// 		// Test when an error occurs while scanning results.
// 		db, mock := setupTestDB(t)
// 		defer db.Close()

// 		expectedRows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "TestCoffee1")

// 		mock.ExpectQuery("^SELECT").WillReturnRows(expectedRows)

// 		if _, err := GetAllCoffees(db, 5*time.Second); err == nil {
// 			t.Error("Expected a scan error, but got nil")
// 		}
// 	})
// }

// func TestCreateCoffee(t *testing.T) {
// 	t.Parallel()

// 	t.Run("Successful Creation", func(t *testing.T) {
// 		// Test creating a new coffee product successfully.

// 		db, mock := setupTestDB(t)
// 		defer db.Close()

// 		inputCoffee := Coffee{
// 			Name:      "TestCoffee",
// 			Image:     "test.jpg",
// 			Region:    "Kenya",
// 			Roast:     "Light",
// 			Price:     9.99,
// 			GrindUnit: 1,
// 		}

// 		expectedCoffee := &Coffee{
// 			ID:        "1",
// 			Name:      inputCoffee.Name,
// 			Image:     inputCoffee.Image,
// 			Roast:     inputCoffee.Roast,
// 			Region:    inputCoffee.Region,
// 			Price:     inputCoffee.Price,
// 			GrindUnit: inputCoffee.GrindUnit,
// 		}

// 		mock.ExpectQuery("^INSERT INTO coffees").WithArgs(inputCoffee.Name, inputCoffee.Image, inputCoffee.Region, inputCoffee.Roast, inputCoffee.Price, inputCoffee.GrindUnit, sqlmock.AnyArg(), sqlmock.AnyArg()).
// 			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedCoffee.ID))

// 		createdCoffee, err := CreateCoffee(db, inputCoffee, 5*time.Second)
// 		if err != nil {
// 			t.Fatalf("CreateCoffee error: %v", err)
// 		}

// 		if createdCoffee.ID != expectedCoffee.ID || createdCoffee.Name != expectedCoffee.Name {
// 			t.Errorf("Mismatch in coffee data: expected %+v, got %+v", expectedCoffee, createdCoffee)
// 		}
// 	})

// 	t.Run("Database Error", func(t *testing.T) {
// 		// Test when an error occurs while creating a coffee product.
// 		db, mock := setupTestDB(t)
// 		defer db.Close()

// 		mock.ExpectQuery("^INSERT INTO coffees").WillReturnError(sql.ErrNoRows)

// 		if _, err := CreateCoffee(db, Coffee{}, 5*time.Second); err == nil {
// 			t.Error("Expected an error, but got nil")
// 		}

// 	})

// }

func TestGetAllCoffees(t *testing.T) {
	t.Parallel()

	t.Run("Successful Retrieval", func(t *testing.T) {
		// Test retrieving a list of coffee products successfully.
		db, mock := setupTestDB(t)
		defer db.Close()

		expectedCoffee1 := &Coffee{
			ID:        "1",
			Name:      "TestCoffee1",
			Image:     "image1.jpg",
			Roast:     "Medium",
			Region:    "Brazil",
			Price:     5.99,
			GrindUnit: 1,
		}
		expectedCoffee2 := &Coffee{
			ID:        "2",
			Name:      "TestCoffee2",
			Image:     "image2.jpg",
			Roast:     "Dark",
			Region:    "Colombia",
			Price:     7.99,
			GrindUnit: 2,
		}

		// Define mock rows with expected data.
		expectedRows := sqlmock.NewRows([]string{"id", "name", "image", "roast", "region", "price", "grind_unit", "created_at", "updated_at"}).
			AddRow(expectedCoffee1.ID, expectedCoffee1.Name, expectedCoffee1.Image, expectedCoffee1.Roast, expectedCoffee1.Region, expectedCoffee1.Price, expectedCoffee1.GrindUnit, time.Now(), time.Now()).
			AddRow(expectedCoffee2.ID, expectedCoffee2.Name, expectedCoffee2.Image, expectedCoffee2.Roast, expectedCoffee2.Region, expectedCoffee2.Price, expectedCoffee2.GrindUnit, time.Now(), time.Now())

		mock.ExpectQuery("^SELECT").WillReturnRows(expectedRows)

		// Create a Models instance with the database connection.
		models := New(db)

		// Call the function and check the results.
		coffees, err := models.Coffee.GetAllCoffees()
		if err != nil {
			t.Fatalf("GetAllCoffees error: %v", err)
		}

		if len(coffees) != 2 {
			t.Errorf("Expected 2 coffees, got %d", len(coffees))
		}

		// Check each coffee's attributes
		if coffees[0].ID != expectedCoffee1.ID || coffees[0].Name != expectedCoffee1.Name {
			t.Errorf("Mismatch in coffee data: expected %+v, got %+v", expectedCoffee1, coffees[0])
		}
		if coffees[1].ID != expectedCoffee2.ID || coffees[1].Name != expectedCoffee2.Name {
			t.Errorf("Mismatch in coffee data: expected %+v, got %+v", expectedCoffee2, coffees[1])
		}
	})

	t.Run("Empty Database", func(t *testing.T) {
		// Test when the database is empty, and no coffee products are retrieved.
		db, mock := setupTestDB(t)
		defer db.Close()

		mock.ExpectQuery("^SELECT").WillReturnRows(sqlmock.NewRows([]string{}))

		// Create a Models instance with the database connection.
		models := New(db)

		coffees, err := models.Coffee.GetAllCoffees()
		if err != nil {
			t.Fatalf("GetAllCoffees error: %v", err)
		}

		if len(coffees) != 0 {
			t.Errorf("Expected 0 coffees, got %d", len(coffees))
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		// Test when an error occurs while retrieving coffee products.
		db, mock := setupTestDB(t)
		defer db.Close()

		mock.ExpectQuery("^SELECT").WillReturnError(sql.ErrNoRows)

		// Create a Models instance with the database connection.
		models := New(db)

		if _, err := models.Coffee.GetAllCoffees(); err == nil {
			t.Error("Expected an error, but got nil")
		}
	})

	t.Run("Timeout Error", func(t *testing.T) {
		// Test when a database query times out.
		db, mock := setupTestDB(t)
		defer db.Close()

		mock.ExpectQuery("^SELECT").WillReturnError(context.DeadlineExceeded)

		// Create a Models instance with the database connection.
		models := New(db)

		if _, err := models.Coffee.GetAllCoffees(); err == nil {
			t.Error("Expected a timeout error, but got nil")
		}
	})

	t.Run("Scan Error", func(t *testing.T) {
		// Test when an error occurs while scanning results.
		db, mock := setupTestDB(t)
		defer db.Close()

		expectedRows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "TestCoffee1")

		mock.ExpectQuery("^SELECT").WillReturnRows(expectedRows)

		// Create a Models instance with the database connection.
		models := New(db)

		if _, err := models.Coffee.GetAllCoffees(); err == nil {
			t.Error("Expected a scan error, but got nil")
		}
	})

}

func TestCreateCoffee(t *testing.T) {
	t.Parallel()
	t.Run("Successful Creation", func(t *testing.T) {
		// Test creating a new coffee product successfully.

		db, mock := setupTestDB(t)
		defer db.Close()

		inputCoffee := Coffee{
			Name:      "TestCoffee",
			Image:     "test.jpg",
			Region:    "Kenya",
			Roast:     "Light",
			Price:     9.99, // Update the expected price value to match the inputCoffee
			GrindUnit: 1,
		}

		expectedCoffee := &Coffee{
			ID:        "1",
			Name:      inputCoffee.Name,
			Image:     inputCoffee.Image,
			Roast:     inputCoffee.Roast,
			Region:    inputCoffee.Region,
			Price:     inputCoffee.Price, // Ensure that the expected price matches inputCoffee
			GrindUnit: inputCoffee.GrindUnit,
		}

		// Update the expected query to use the correct Price value.
		mock.ExpectQuery("^INSERT INTO coffees").WithArgs(inputCoffee.Name, inputCoffee.Image, inputCoffee.Region, inputCoffee.Roast, inputCoffee.Price, inputCoffee.GrindUnit, sqlmock.AnyArg(), sqlmock.AnyArg()).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedCoffee.ID))

		// Create a Models instance with the database connection.
		models := New(db)

		createdCoffee, err := models.Coffee.CreateCoffee(inputCoffee)
		if err != nil {
			t.Fatalf("CreateCoffee error: %v", err)
		}

		if createdCoffee.ID != expectedCoffee.ID || createdCoffee.Name != expectedCoffee.Name {
			t.Errorf("Mismatch in coffee data: expected %+v, got %+v", expectedCoffee, createdCoffee)
		}
	})

	t.Run("Database Error", func(t *testing.T) {
		// Test when an error occurs while creating a coffee product.
		db, mock := setupTestDB(t)
		defer db.Close()

		mock.ExpectQuery("^INSERT INTO coffees").WillReturnError(sql.ErrNoRows)

		// Create a Models instance with the database connection.
		models := New(db)

		if _, err := models.Coffee.CreateCoffee(Coffee{}); err == nil {
			t.Error("Expected an error, but got nil")
		}
	})

	t.Run("Timeout Error", func(t *testing.T) {
		// Test when a database query times out.
		db, mock := setupTestDB(t)
		defer db.Close()

		mock.ExpectQuery("^INSERT INTO coffees").WillReturnError(context.DeadlineExceeded)

		// Create a Models instance with the database connection.
		models := New(db)

		if _, err := models.Coffee.CreateCoffee(Coffee{}); err == nil {
			t.Error("Expected a timeout error, but got nil")
		}
	})

	t.Run("Insert Error", func(t *testing.T) {
		// Test when an error occurs while inserting a coffee product.
		db, mock := setupTestDB(t)
		defer db.Close()

		mock.ExpectQuery("^INSERT INTO coffees").WillReturnError(errors.New("insert error"))

		// Create a Models instance with the database connection.
		models := New(db)

		if _, err := models.Coffee.CreateCoffee(Coffee{}); err == nil {
			t.Error("Expected an error, but got nil")
		}
	})

}
