package model

// Users - users struct
type Users struct {
	ID       int
	Name     string
	Password string
}

// ProductType - product type struct
type ProductType struct {
	ProductTypeID int    `gorm:"column:id"`
	Name          string `gorm:"type:varchar(255);unique"`
	Description   string `gorm:"type:varchar(255)"`
}

// Products - products struct
type Products struct {
	ID            int    `gorm:"column:id"`
	ProductTypeID int    `gorm:"association_foreignkey:ProductTypeID"`
	Name          string `gorm:"type:varchar(255);unique"`
	Price         int    `gorm:"type:int"`
	Description   string `gorm:"type:varchar(255)"`
}

// Text - chatfuel text struct
type Text struct {
	Text string `json:"text"`
}

// Message - chatfuel text struct
type Message struct {
	Message []Text `json:"messages"`
}

// FbUsers - fb users struct
type FbUsers struct {
	ID              int    `gorm:"column:id"`
	MessengerUserID string `gorm:"varchar(100);unique"`
	FirstName       string `gorm:"varchar(50)"`
	LastName        string `gorm:"varchar(50)"`
	Gender          string `gorm:"varchar(10)"`
	ProfilePicURL   string `gorm:"text"`
	Timezone        int    `gorm:"int(11)"`
	Locale          string `gorm:"varchar(10)"`
}

// ChatfuelFbUser - chatfuel fb user
type ChatfuelFbUser struct {
	MessengerUserID string `json:"messenger user id"`
	FirstName       string `json:"first name"`
	LastName        string `json:"last name"`
	Gender          string `json:"gender"`
	ProfilePicURL   string `json:"profile pic url"`
	Timezone        int    `json:"timezone"`
	Locale          string `json:"locale"`
}

// Carts - carts
type Carts struct {
	ID              int    `gorm:"column:id"`
	MessengerUserID string `gorm:"varchar(100);unique"`
	FirstName       string `gorm:"varchar(50);column:first_name"`
	ProductID       int    `gorm:"int(11);column:product_id"`
	ProductName     string `gorm:"varchar(50);column:product_name"`
	Qty             int    `gorm:"int(11);column:qty"`
	Price           int    `gorm:"int(11);column:price"`
}

// ChatfuelCarts - chatfuel cars
type ChatfuelCarts struct {
	MessengerUserID string `json:"messenger user id"`
	FirstName       string `json:"first name"`
	ProductID       int    `json:"product_id"`
	ProductName     string `json:"product_name"`
	Qty             int    `json:"qty"`
	Price           int    `json:"price"`
}

// Orders - orders
type Orders struct {
	ID              int    `gorm:"column:id"`
	MessengerUserID string `gorm:"varchar(100);unique"`
	FirstName       string `gorm:"varchar(50);column:first_name"`
	Detail          string `gorm:"text;column:detail"`
	TotalPrice      int    `gorm:"int(11);column:total_price"`
}
