package contact_us

import "gorm.io/gorm"

// ContactUsRepository implements ContactUsRepositoryInterface
type ContactUsRepository struct {
	db *gorm.DB
}

// NewContactUsRepository and return it
func NewContactUsRepository(db *gorm.DB) ContactUsRepositoryInterface {
	return &ContactUsRepository{
		db: db,
	}
}

// CreateContactUs and store it in db
func (c *ContactUsRepository) CreateContactUs(contactUs *ContactUs) error {
	result := c.db.Create(contactUs)
	return result.Error
}
