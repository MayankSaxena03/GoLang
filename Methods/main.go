package main

import "fmt"

func main() {
	/* Difference between methods and function is that "method is a function inside class or struct" */
	mayank := User{"Mayank", 20, "mayank@gmail.com", true}
	fmt.Println("Is User Active : ", mayank.GetStatus())
	fmt.Println("User Email : ", mayank.Email)
	mayank.SetMail("saxenamayank@gmail.com") /* Won't Update Email */
	fmt.Println("User Email : ", mayank.Email)
	mayank.SetMail2("saxenamayank@gmail.com") /* Will Update Email */
	fmt.Println("User Email : ", mayank.Email)
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func (u User) GetStatus() bool {
	return u.Status
}

/* Below method won't update the original value of user.mail but will change the copy of the value */
func (u User) SetMail(mail string) {
	u.Email = mail
}

/* Passing reference of User */
func (u *User) SetMail2(mail string) {
	u.Email = mail
}
