package User

import (
	"fmt"
	database "github.com/ricardoassaad/GoBackend/database"
)

func getConnection() {
	_ = database.InitiateDatabaseConnection()
}

func init() {
	fmt.Println(getConnection())
}

type User struct {
	id   int
	name string
}

func All() string {
	return "test"
}
