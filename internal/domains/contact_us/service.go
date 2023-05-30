package contact_us

// ContactUsService is the type which implements ContactUsServiceInterface
type ContactUsService struct {
	repo ContactUsRepositoryInterface
}

// NewContactUsService and return it
func NewContactUsService(repo ContactUsRepositoryInterface) ContactUsServiceInterface {
	return &ContactUsService{
		repo: repo,
	}
}

// CreateContactUs and store it in db
func (c *ContactUsService) CreateContactUs(title, email, description string) (*ContactUs, error) {
	contact := &ContactUs{
		Email:       email,
		Title:       title,
		Description: description,
	}

	err := c.repo.CreateContactUs(contact)
	return contact, err
}
