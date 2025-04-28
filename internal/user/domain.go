package user

// func MappingUpdateSubscription(body *UpdateSubscriptionDTO) (*User, error) {
// 	if event := body.Event; event != "user.upgraded" {
// 		return nil, utils.NewDomainError(204, "invalid event")
// 	}

// 	userID, err := uuid.Parse(body.Data.UserID)
// 	if err != nil {
// 		return nil, utils.NewDomainError(204, "user id is required")
// 	}

// 	return &User{
// 		ID:          userID,
// 		IsChirpyRed: body.Event == "user.upgraded",
// 	}, nil
// }
