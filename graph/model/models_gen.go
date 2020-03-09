// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Bike struct {
	ID        string  `json:"id"`
	CreatedAt int     `json:"createdAt"`
	UpdatedAt int     `json:"updatedAt"`
	DeletedAt *int    `json:"deletedAt"`
	Make      string  `json:"make"`
	Model     string  `json:"model"`
	Reg       string  `json:"reg"`
	Price     float64 `json:"price"`
	Active    bool    `json:"active"`
}

type BikeQuery struct {
	Meta *Meta   `json:"meta"`
	Data []*Bike `json:"data"`
}

type Meta struct {
	Cursor     int `json:"cursor"`
	TotalItems int `json:"totalItems"`
}

type NewBike struct {
	Make   string  `json:"make"`
	Model  string  `json:"model"`
	Reg    string  `json:"reg"`
	Price  float64 `json:"price"`
	Active bool    `json:"active"`
}

type NewPost struct {
	Content string `json:"content"`
	UserID  string `json:"userId"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Active   bool   `json:"active"`
}

type PaginatedQuery struct {
	Cursor   int `json:"cursor"`
	PageSize int `json:"pageSize"`
}

type Post struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	User    *User  `json:"user"`
}

type PostQuery struct {
	Meta *Meta   `json:"meta"`
	Data []*Post `json:"data"`
}

type User struct {
	ID        string `json:"id"`
	CreatedAt int    `json:"createdAt"`
	UpdatedAt int    `json:"updatedAt"`
	DeletedAt *int   `json:"deletedAt"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Active    bool   `json:"active"`
}

type UserQuery struct {
	Meta *Meta   `json:"meta"`
	Data []*User `json:"data"`
}
