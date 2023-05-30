package contact_us

import "context"

// ContactUsUseCase implements ContactUsUseCaseInterface
type ContactUsUseCase struct {
	sv ContactUsServiceInterface
}

// NewUseCase and return it as a pointer
func NewUseCase(sv ContactUsServiceInterface) ContactUsUseCaseInterface {
	return &ContactUsUseCase{
		sv: sv,
	}
}

// CreateContactUs and store it in db
func (c *ContactUsUseCase) CreateContactUs(ctx context.Context, request *CreateContactUsRequest) (*ContactUs, error) {
	return c.sv.CreateContactUs(request.Title, request.Email, request.Description)
}
