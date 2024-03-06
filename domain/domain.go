package domain

type UserStates struct {
	Id      string
	Message string
	Mode    string `default:""`
	Hour    int
	Minute  int
}

// type UserTime struct {
// 	Id     string

// }

// type UserMsg struct {
// 	Id      string

// }
