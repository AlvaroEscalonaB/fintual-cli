package models

type ConfigINI struct {
	User struct {
		Email    string `ini:"email"`
		Password string `ini:"password"`
		Token    string `ini:"token"`
	} `ini:"user"`
}
