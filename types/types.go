package types

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Mobile   int64  `json:"mobile"`
	UserType int    `json:"user_type"`
	Password string `json:"password"`
}

type UserDetails struct {
	UID      int64  `bun:"uid,pk,notnull"`
	Username string `bun:"username,notnull"`
	Email    string `bun:"email,notnull"`
	Mobile   int64  `bun:"mobile,notnull"`
	UserType int    `bun:"user_type,notnull"`
}

type UserUpdateReq struct {
	OriginalUsername string `json:"originalName"`
	Username         string `json:"name"`
	Email            string `json:"email"`
	Mobile           int64  `json:"mobile"`
}

type AuthDetails struct {
	UserID   int64  `bun:"user_id,pk,notnull"`
	Password string `bun:"password,notnull"`
}

type UserIdReq struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
}

type UserUsernameReq struct {
	Username string `json:"name"`
	Password string `json:"password"`
}

type UserIdReq2 struct {
	Username string `json:"name"`
}

type RestuarantReq struct {
	Name       string `json:"name"`
	Location   string `json:"location"`
	IsVeg      bool   `json:"is_veg"`
	VendorName string `json:"vendor_name"`
	Password   string `json:"password"`
}

type RestuarantIdReq struct {
	ID int `json:"id"`
}

type RestuarantNameReq struct {
	Name string `json:"name"`
}

type RestaurantData struct {
	UID        int64  `bun:"uid,pk,notnull"`
	Name       string `bun:"name,notnull"`
	Location   string `bun:"location,notnull"`
	IsVeg      bool   `bun:"is_veg,notnull"`
	VendorName string `bun:"vendor_name,notnull"`
	StatusId   int    `bun:"status_id, notnull"`
}

type Food struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	IsVeg       bool   `json:"is_veg"`
	RestID      int64  `json:"rest_id"`
	Ingredients string `json:"ingredients"`
	IsRegular   bool   `json:"is_regular"`
}

type FoodItems struct {
	UID           int64  `bun:"uid,pk,notnull"`
	RestuarantId  int64  `bun:"restuarant_id,notnull"`
	Item          string `bun:"item,notnull"`
	Ingredients   string `bun:"ingredients,notnull"`
	IsVeg         bool   `bun:"is_veg,notnull"`
	IsRegular     bool   `bun:"is_regular,notnull"`
	Price         int    `bun:"price,notnull"`
	IsAvailable   bool   `bun:"available,notnull"`
	AvailableTime int    `bun:"available_time,notnull"`
}

type FoodItemsRestaurantReq struct {
	RestID int `json:"rest_id"`
}

type FoodItemsRestaurantNameReq struct {
	Name int `json:"name"`
}

type OrderList struct {
	UID           int64  `bun:"uid,pk,notnull"`
	UserId        int64  `bun:"user_id,notnull"`
	RestaurantID  int64  `bun:"rest_id,notnull"`
	IsPaid        bool   `bun:"is_paid,notnull"`
	IsCash        bool   `bun:"is_cash,notnull"`
	TimeCreated   string `bun:"timestamp,notnull"`
	OrderStatusId int    `bun:"order_status,notnull"`
}

func (OrderList) TableName() string {
	return "order_list"
}

type OrderDetails struct {
	UID      int64 `bun:"uid,pk,notnull"`
	OrderId  int64 `bun:"order_id,notnull"`
	FoodId   int64 `bun:"food_id,notnull"`
	Quantity int   `bun:"quantity,notnull"`
}

type OrderItemReq struct {
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
}

type OrderItemReq2 struct {
	Item     string `json:"item"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

type OrderIDReq struct {
	ID int `json:"id"`
}

type OrderRestIDReq struct {
	RestID int `json:"id"`
}

type OrderUserIDReq struct {
	UserID int `json:"id"`
}

type Order struct {
	UserId        int64          `json:"user_id"`
	RestuarantID  int64          `json:"rest_id"`
	IsPaid        bool           `json:"is_paid"`
	IsCash        bool           `json:"is_cash"`
	TimeCreated   string         `json:"timestamp"`
	OrderStatusId int            `json:"order_status"`
	OrderItems    []OrderItemReq `json:"order_items"`
}

type OrderVendor struct {
	UID           int64           `josn:"uid"`
	UserName      string          `json:"username"`
	RestuarantID  int64           `json:"rest_id"`
	IsPaid        bool            `json:"is_paid"`
	IsCash        bool            `json:"is_cash"`
	TimeCreated   string          `json:"timestamp"`
	OrderStatusId int             `json:"order_status"`
	OrderItems    []OrderItemReq2 `json:"order_items"`
}

type CostBody struct {
	OrderID int `json:"order_id"`
}

type OrderStatusReq struct {
	OrderID int `json:"order_id"`
	Status  int `json:"status"`
}
