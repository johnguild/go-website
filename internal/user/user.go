package user

import "errors"

// type User struct {
// 	Fullname string
// }

type Credentials struct {
	Email    string
	Password string
	Fullname string
}

func FindMatch(email string, password string) (*Credentials, error) {
	var credentials = []Credentials{
		{Email: "myemail@gmail.com", Password: "r@nd0m1234", Fullname: "John Doe"},
		{Email: "hot@gmail.com", Password: "woasdf123", Fullname: "Jane Doe"},
		{Email: "dune1234@gmail.com", Password: "asdfqwed2", Fullname: "Mark Doe"},
	}

	for _, c := range credentials {
		if c.Email == email && c.Password == password {
			// return the proper user data here
			return &c, nil
		}
	}

	return nil, errors.New("User Not Found")
}
