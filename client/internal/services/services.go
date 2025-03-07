package services

type Services struct {
	UserService UserServiceClient
}

func NewServices(userService UserServiceClient) *Services {
	return &Services{
		UserService: userService,
	}
}
