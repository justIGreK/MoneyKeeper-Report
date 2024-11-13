package models


type Budget struct {
	ID        string   
	UserID    string     
	Name      string     
	Limit     float64    
	StartDate string  
	EndDate   string  
	Category  []Category 
}

type Category struct {
	ID    string  
	Name  string  
	Limit float64 
}

