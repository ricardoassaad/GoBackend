package UsersController

import (
	User "github.com/ricardoassaad/GoBackend/models/User"
)

func Index() string {
	return User.All()
}
