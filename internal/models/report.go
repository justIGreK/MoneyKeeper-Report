package models 

type ReportResponse struct {
    UserID          string      `json:"user_id"`
    Period          string      `json:"period"`
    TotalSpent      float32     `json:"total_spent"`
    TransactionCount int        `json:"transaction_count"`
    Categories      []CategoryReport `json:"categories"`
}

type CategoryReport struct {
    Name  string  `json:"name"`
    Total float32 `json:"total"`
	Count int    `json:"transaction_count"`
}
