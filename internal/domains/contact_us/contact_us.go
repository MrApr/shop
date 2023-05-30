package contact_us

import "context"

// ContactUsRepositoryInterface defines set of method for every type who is contact us repository
type ContactUsRepositoryInterface interface {
	CreateContactUs(contactUs *ContactUs) error
}

// ContactUsServiceInterface defines set of method for every type who is contact us service
type ContactUsServiceInterface interface {
	CreateContactUs(title, email, description string) (*ContactUs, error)
}

// ContactUsUseCaseRepository defines set of method for every type who is contact us use-case
type ContactUsUseCaseRepository interface {
	CreateContactUs(ctx context.Context, request *CreateContactUsRequest) (*ContactUs, error)
}
