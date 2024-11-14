package models

type ReportResponse struct {
	Period           string
	TotalSpent       float32
	TransactionCount int
	Categories       []CategoryReport
}

type CategoryReport struct {
	Name  string
	Total float32
	Count int
}
type RequiredCategoryReport struct {
	Name  string
	Total float32
    Limit float32
	Count int
}
type BudgetReport struct {
	BudgetName string
	Period     string
	LeftDays   float32
	Limit      float32
	TotalSpent float32
	TransactionCount int
    RequiredCategories []RequiredCategoryReport
	Categories       []CategoryReport
}

type GetPeriodSummary struct{
	UserID string
	Period *string
	Start *string
	End *string
}