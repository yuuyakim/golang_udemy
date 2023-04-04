package main

import "fmt"

type T struct {
	// 構造体の埋め込みの場合は型の省略できる
	// User User
	User
}

type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

type User struct {
	Name string
	Age  int
}

// structのスライス
type Users []*User

// 擬似的にコンストラクタを実装できる
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

// struct func
func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) SetName2(name string) {
	u.Name = name
}

func UpdateUser(user User) {
	user.Name = "A"
	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	// var user1 User
	// user1.Name = "user1"
	// user1.Age = 10
	// fmt.Println(user1)

	// user7 := new(User)
	// fmt.Println(user7)

	// user8 := &User{}
	// fmt.Println(user8)

	// UpdateUser(user1)
	// UpdateUser2(user8)
	// fmt.Println(user1)
	// fmt.Println(user8)

	// user1.SayName()
	// user1.SetName("BBBB")
	// user1.SayName()
	// user1.SetName2("BBBB")
	// user1.SayName()

	t := T{User: User{Name: "user1", Age: 10}}

	fmt.Println(t)
	fmt.Println(t.User)
	fmt.Println(t.User.Name)
	// 埋め込んで型を省略した場合のみt.User.Name のUserを省略できる
	fmt.Println(t.Name)

	t.SetName2("hogehoge")
	fmt.Println(t.Name)


	user_hoge := NewUser("hoge", 2)
	fmt.Println(*user_hoge)

	user10 := User{Name: "user10", Age: 10}
	user20 := User{Name: "user20", Age: 20}
	user30 := User{Name: "user30", Age: 30}
	user40 := User{Name: "user40", Age: 40}

	users := Users{}

	users = append(users, &user10)
	users = append(users, &user20)
	users = append(users, &user30, &user40)
	fmt.Println(users)


	for _, u := range users {
		fmt.Println(*u)
	}

	users2 := make([]*User, 0)
	users2 = append(users2, &user10, &user20)

	for _, u := range users2 {
		fmt.Println(*u)
	}

	m := map[int]User{
		1: {Name: "user1", Age: 10},
		2: {Name: "user2", Age: 20},
	}

	fmt.Println(m)

	m2 := map[User]string{
		{Name: "hoge", Age: 20}: "tokyo",
		{Name: "hoge", Age: 10}: "LA",
	}

	fmt.Println(m2)

	m3 := make(map[int]User)
	fmt.Println(m3)
	m3[1] = User{Name: "user3333"}
	fmt.Println(m3)

	for _, v := range m {
		fmt.Println(v)
	}

	var mi MyInt
	fmt.Printf("%T\n", mi)

	// i := 100
	mi.Print()
}
