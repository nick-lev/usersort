// test tool for sort user struct by selected fields
package main

import (
	"fmt"
	"sort"
)

type User struct {
	Fname string
	Lname string
	Age   int
}

type ByFname []User
type ByLname []User
type ByAge []User

func (s ByFname) Len() int {
	return len(s)
}
func (s ByLname) Len() int {
	return len(s)
}
func (s ByAge) Len() int {
	return len(s)
}

func (s ByFname) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLname) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByFname) Less(i, j int) bool {
	return (s[i].Fname < s[j].Fname)
}
func (s ByLname) Less(i, j int) bool {
	return (s[i].Lname < s[j].Lname)
}
func (s ByAge) Less(i, j int) bool {
	return (s[i].Age < s[j].Age)
}

func (u User) String() string {
	return fmt.Sprintf("%s %s %d", u.Fname, u.Lname, u.Age)
}

func main() {
	users := []User{
		{"John", "Dou", 88},
		{"Jobody", "Dou", 91},
		{"Jobody", "Know", 100},
		{"Abraham", "Link", 18},
		{"Bocka", "Owner", 22},
	}
	fmt.Println(users)
	sort.Sort(ByFname(users))
	fmt.Println(users)
	sort.Sort(ByLname(users))
	fmt.Println(users)
	sort.Sort(ByAge(users))
	fmt.Println(users)
}
