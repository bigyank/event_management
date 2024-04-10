package service

import (
	"context"
	"kraneapi/graph/model"
	"kraneapi/pkg/api/dbmodel"
	"kraneapi/pkg/db"

	"github.com/doug-martin/goqu/v9"
)

// CreateExpense creates a new expense for an event in the database.
func CreateExpense(ctx context.Context, input model.CreateExpenseInput) (string, error) {
	insert := db.GetDB().Insert("expenses").
		Cols("event_id", "item_name", "cost", "description", "category").
		Vals(goqu.Vals{input.EventID, input.ItemName, input.Cost, input.Description, input.Category}).
		Returning(goqu.C("expense_id")).
		Executor()

	var id string
	if _, err := insert.ScanVal(&id); err != nil {
		return "", err
	}

	return id, nil
}

// UpdateExpense updates an existing expense in the database.
func UpdateExpense(ctx context.Context, input model.UpdateExpenseInput) (string, error) {
	update := db.GetDB().Update("expenses").
		Set(map[string]interface{}{
			"item_name":   input.ItemName,
			"cost":        input.Cost,
			"description": input.Description,
			"category":    input.Category,
		}).
		Where(goqu.C("expense_id").Eq(input.ExpenseID), goqu.C("event_id").Eq(input.EventID)).
		Returning(goqu.C("expense_id")).
		Executor()

	var id string
	if _, err := update.ScanVal(&id); err != nil {
		return "", err
	}

	return id, nil
}

// GetAllEventExpenses retrieves all expenses for a specific event from the database.
func GetAllEventExpenses(ctx context.Context, eventID string) ([]*dbmodel.Expense, error) {
	var eventExpenses []*dbmodel.Expense
	err := db.GetDB().From("expenses").
		Where(goqu.C("event_id").Eq(eventID)).
		Select("*").ScanStructs(&eventExpenses)
	if err != nil {
		return nil, err
	}
	return eventExpenses, nil
}

// GetEventExpensesBreakdown retrieves a breakdown of expenses by category for a specific event.
func GetEventExpensesBreakdown(ctx context.Context, eventID string) ([]*dbmodel.ExpenseCategoryBreakdown, error) {
	query := db.GetDB().From("expenses").
		Where(goqu.C("event_id").Eq(eventID)).
		Select(goqu.C("category"), goqu.SUM("cost").As("total_cost")).
		GroupBy(goqu.C("category")).
		Executor()

	var expense []*dbmodel.ExpenseCategoryBreakdown
	if err := query.ScanStructs(&expense); err != nil {
		return nil, err
	}

	return expense, nil
}
