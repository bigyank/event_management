package dbmodel

// ExpenseCategory represents the different categories an expense can belong to.
type ExpenseCategory string

const (
	Venue       ExpenseCategory = "VENUE"
	Catering    ExpenseCategory = "CATERING"
	Decorations ExpenseCategory = "DECORATIONS"
	Other       ExpenseCategory = "OTHER"
)

// Expense represents an expense related to an event.
type Expense struct {
	ExpenseID   string          `db:"expense_id"`
	EventID     string          `db:"event_id"`
	ItemName    string          `db:"item_name"`
	Cost        float64         `db:"cost"`
	Description string          `db:"description"`
	Category    ExpenseCategory `db:"category"`
}

// ExpenseCategoryBreakdown represents the total expenses for each category of an event.
type ExpenseCategoryBreakdown struct {
	Category  ExpenseCategory `db:"category"`
	TotalCost float64         `db:"total_cost"`
}
