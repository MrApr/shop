package products

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

// TestCategoryRepository_GetAllCategories functionality
func TestCategoryRepository_GetAllCategories(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Connecting to database failed")

	repo := createCategoryRepository(db)
	createdCats := mockAndInsertCategories(db, 2)
	defer destructCreatedType(db, createdCats[0].TypeId)
	defer destructCreatedCategories(db, createdCats)

	fetchedCats := repo.GetAllCategories()
	assertCategories(t, createdCats, fetchedCats)
}

// createCategoryRepository for testing purpose
func createCategoryRepository(db *gorm.DB) CategoriesRepositoryInterface {
	return NewCategoryRepo(db)
}

// setupDbConnection to testing in memory database
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(Type{}, Category{})
	return db, err
}

// mockAndInsertCategories in database
func mockAndInsertCategories(db *gorm.DB, count int) []Category {
	customType := createTypeInDb(db)

	categories := make([]Category, 0, count)

	for i := 0; i < count; i++ {
		mockedCat := mockCategory(customType.Id)
		result := db.Create(mockedCat)
		if result.Error != nil {
			continue
		}
		categories = append(categories, *mockedCat)
	}

	return categories
}

// createTypeInDb mocks and insert a type in database
func createTypeInDb(db *gorm.DB) *Type {
	tmpType := mockType()
	db.Create(tmpType)
	return tmpType
}

// mockType struct for test
func mockType() *Type {
	return &Type{
		Title: "Test default type",
	}
}

// mockCategory for test
func mockCategory(typeId int) *Category {
	return &Category{
		TypeId:         typeId,
		ParentCatId:    nil,
		ParentCategory: nil,
		Title:          "Test category",
		Indent:         1,
		Order:          1,
	}
}

// assertCategories check whether they are equal or not
func assertCategories(t *testing.T, createdCats, fetchedCats []Category) {
	for index := range createdCats {
		assert.Equal(t, createdCats[index].Id, fetchedCats[index].Id, "Categories Repository test: Ids are not equal")
		assert.Equal(t, createdCats[index].Title, fetchedCats[index].Title, "Categories Repository test: titles are not equal")
		assert.Equal(t, createdCats[index].TypeId, fetchedCats[index].TypeId, "Categories Repository test: type ids are not equal")
		assert.Equal(t, createdCats[index].Indent, fetchedCats[index].Indent, "Categories Repository test: indents are not equal")
		assert.Equal(t, createdCats[index].Order, fetchedCats[index].Order, "Categories Repository test: orders are not equal")
		assert.NotNil(t, fetchedCats[index].Type, "Categories Repository test: Type is not eager loaded properly")
	}
}

// destructCreatedCategories and delete them from db
func destructCreatedCategories(db *gorm.DB, cats []Category) {
	for _, cat := range cats {
		db.Unscoped().Delete(cat)
	}
}

// destructCreatedType and deleted it from DB
func destructCreatedType(db *gorm.DB, typeId int) {
	db.Unscoped().Delete(Type{}, typeId)
}
