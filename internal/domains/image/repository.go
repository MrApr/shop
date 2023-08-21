package image

import "gorm.io/gorm"

// MysqlImageRepo struct defines struct that handles ImageRepository's responsibilities
type MysqlImageRepo struct {
	dbConn *gorm.DB
}

// NewImageRepo creates and returns new mysql image repository
func NewImageRepo(conn *gorm.DB) ImageRepoContract {
	return &MysqlImageRepo{
		dbConn: conn,
	}
}

// GetImage fetched image and returns it
func (m *MysqlImageRepo) GetImage(id int) (*Image, error) {
	img := new(Image)
	result := m.dbConn.Where("id = ?", id).First(img)
	return img, result.Error
}

// GetImageByImgables and return it
func (m *MysqlImageRepo) GetImageByImgables(imageable ImageableContract) (*Image, error) {
	var img Image
	result := m.dbConn.Where("imageable_id = ? AND imageable_type = ?", imageable.ImageableId(), imageable.ImageableType()).First(&img)
	return &img, result.Error
}
