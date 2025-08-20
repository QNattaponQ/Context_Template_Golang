package repository_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"yourproject/repository"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	err = db.AutoMigrate(&repository.User{})
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func TestUserRepository_CreateAndFind(t *testing.T) {
	db := setupTestDB(t)
	repo := repository.NewUserRepository(db)

	ctx := context.Background()
	user := &repository.User{Name: "Alice", Email: "alice@example.com", Age: 30}

	err := repo.Create(ctx, user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)

	found, err := repo.FindByID(ctx, user.ID)
	assert.NoError(t, err)
	assert.Equal(t, "Alice", found.Name)
	assert.Equal(t, "alice@example.com", found.Email)
}
