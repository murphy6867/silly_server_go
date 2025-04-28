package user

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

// func (svc *UserService) SubscriptionService(ctx context.Context, body *UpdateSubscriptionDTO) error {
// 	if event := body.Event; event == "" {
// 		return utils.NewDomainError(204, "event is required")
// 	}

// 	if userID := body.Data.UserID; userID == "" {
// 		return utils.NewDomainError(204, "user id is required")
// 	}

// 	updateStatus, err := MappingUpdateSubscription(body)
// 	if err != nil {
// 		return err
// 	}

// 	return svc.repo.UpdateSubscriptionChirpyRed(ctx, updateStatus)
// }
