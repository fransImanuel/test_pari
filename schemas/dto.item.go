package schemas

type CreateItemRequest struct {
	Name        string `json:"name"  binding:"required" `
	Description string `json:"description"  binding:"required" `
	Price       int64  `json:"price"  binding:"required" `
	Qty         int64  `json:"qty"  binding:"required" `
}

type UpdateItemRequest struct {
	Name        string `json:"name"  binding:"required" `
	Description string `json:"description"  binding:"required" `
	Price       int64  `json:"price"  binding:"required" `
	Qty         int64  `json:"qty"  binding:"required" `
	IsActive    bool   `json:"is_active" `
}
