package image

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math/rand"
	"testing"
)

// TestGetImage tests fetching ability of repository
func TestGetImage(t *testing.T) {
	db, err := setupDbConnection()
	assert.NoError(t, err, "Cannot connect to memory database")

	img := createTestImgInDb(db, t)

	imageRepo := NewImageRepo(db)
	fetchedImg, err := imageRepo.GetImage(img.Id)
	assert.NoError(t, err, "Cannot fetch image")
	checkImagesEqualities(t, img, fetchedImg)

}

// getRandomImgObj generates random object
func getRandomImgObj() Image {
	return Image{
		ImageableID:   rand.Int(),
		ImageableType: "testableImg",
		Path:          "aaaaaaa",
	}
}

// checkImagesEqualities
func checkImagesEqualities(t *testing.T, firstImg, secondImg *Image) {
	assert.Equal(t, firstImg.ImageableID, secondImg.ImageableID)
	assert.Equal(t, firstImg.ImageableType, secondImg.ImageableType)
	assert.Equal(t, firstImg.Path, secondImg.Path)
	assert.Equal(t, firstImg.Id, secondImg.Id)
}

// setupDbConnection and run migration
func setupDbConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(Image{})
	return db, err
}

// createTestImgInDb and stores it then returns it
func createTestImgInDb(db *gorm.DB, t *testing.T) *Image {
	img := getRandomImgObj()

	result := db.Create(&img)
	assert.NoError(t, result.Error, "Cannot create test image")
	return &img
}
