package types

type UserGet struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

type UserGetUsername struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Mobile   int    `json:"mobile"`
	UserType int    `json:"user_type"`
	Password string `json:"password"`
}

type Restuarant struct {
	Name     string `json:"name"`
	Location string `json:"location"`
	IsVeg    bool   `json:"is_veg"`
}

type RestaurantGet struct {
	ID int `json:"id"`
}
type RestaurantGetName struct {
	Name string `json:"name"`
}

type FoodItem struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	IsVeg       bool   `json:"is_veg"`
	RestID      int    `json:"rest_id"`
	Ingredients string `json:"ingredients"`
	IsRegular   bool   `json:"is_regular"`
}

type CostBody struct {
	OrderID int `json:"order_id"`
}

type FoodItemGetResp struct {
	ID          int
	Name        string
	Price       int
	IsVeg       bool
	RestID      int
	Ingredients string
	IsRegular   bool
}

type FoodItemGet struct {
	ID int `json:"id"`
}

type FoodRating struct {
	FoodID   int    `json:"food_id"`
	UserID   int    `json:"user_id"`
	Rating   int    `json:"rating"`
	Comments string `json:"comments"`
}

type RestuarantRating struct {
	ID       int    `json:"id"`
	RestID   int    `json:"rest_id"`
	UserID   int    `json:"user_id"`
	Rating   int    `json:"rating"`
	Comments string `json:"comments"`
}

type Order struct {
	UserID      int          `json:"user_id"`
	RestID      int          `json:"rest_id"`
	IsPaid      bool         `json:"is_paid"`
	IsCash      bool         `json:"is_cash"`
	OrderStatus int          `json:"order_status"`
	OrderItems  []OrderItems `json:"order_items"`
}

type OrderResp struct {
	ID          int
	UserID      int
	RestID      int
	IsPaid      bool
	IsCash      bool
	OrderStatus int
	OrderItems  []OrderItems
}

type OrderItems struct {
	ID       int `json:"id"`
	OrderID  int `json:"order_id"`
	FoodID   int `json:"food_id"`
	Quantity int `json:"quantity"`
}

type GetOrder struct {
	ID int `json:"id"`
}

type OrderGet struct {
	ID          int
	UserID      int
	RestID      int
	IsPaid      bool
	IsCash      bool
	Time        string
	OrderStatus int
	OrderItems  []OrderItems
}

type OrderGetUser struct {
	UserID int `json:"user_id"`
}

type OrderGetRestaurant struct {
	RestID int `json:"rest_id"`
}

type OrderGetStatus struct {
	OrderStatus int `json:"order_status"`
}

type OrderStatus struct {
	OrderID int `json:"order_id"`
	Status  int `json:"status"`
}
