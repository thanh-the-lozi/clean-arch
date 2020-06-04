package entity

type Student struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}
type ListStudent []Student
