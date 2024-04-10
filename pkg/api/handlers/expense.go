package handler

import (
	"context"
	"kraneapi/graph/model"
	"kraneapi/pkg/api/dbmodel"
	"kraneapi/pkg/db"

	"github.com/doug-martin/goqu/v9"
)

// CreateExpense creates a new expense for an event.
func CreateExpense(ctx context.Context, input model.CreateExpenseInput) (*model.Expense, error) {
	insert := db.GetDB().Insert("expenses").
		Cols("event_id", "item_name", "cost", "description", "category").
		Vals(goqu.Vals{input.EventID, input.ItemName, input.Cost, input.Description, input.Category}).
		Returning(goqu.C("expense_id")).
		Executor()

	var id string
	if _, err := insert.ScanVal(&id); err != nil {
		return nil, err
	}

	return &model.Expense{
		ExpenseID:   id,
		EventID:     input.EventID,
		ItemName:    input.ItemName,
		Cost:        input.Cost,
		Description: input.Description,
		Category:    input.Category,
	}, nil
}

// UpdateExpense updates an existing expense.
func UpdateExpense(ctx context.Context, input model.UpdateExpenseInput) (*model.Expense, error) {
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
		return nil, err
	}

	return &model.Expense{
		ExpenseID:   id,
		EventID:     input.EventID,
		ItemName:    input.ItemName,
		Cost:        input.Cost,
		Description: input.Description,
		Category:    input.Category,
	}, nil
}

func GetAllEventExpenses(ctx context.Context, eventID string) ([]*model.Expense, error) {

	var eventExpenses []*dbmodel.Expense

	err := db.GetDB().From("expenses").
		Where(goqu.C("event_id").Eq(eventID)).
		Select("*").ScanStructs(&eventExpenses)

	if err != nil {
		return nil, err
	}

	var expenses []*model.Expense
	for _, expenseModel := range eventExpenses {
		expense := model.Expense{
			ExpenseID:   expenseModel.ExpenseID,
			EventID:     expenseModel.EventID,
			ItemName:    expenseModel.ItemName,
			Cost:        expenseModel.Cost,
			Description: expenseModel.Description,
			Category:    model.ExpenseCategory(expenseModel.Category),
		}
		expenses = append(expenses, &expense)
	}

	return expenses, nil
}

func GetEventExpensesBreakdown(ctx context.Context, eventID string) ([]*model.ExpenseCategoryBreakdown, error) {

	query := db.GetDB().From("expenses").
		Where(goqu.C("event_id").Eq(eventID)).
		Select(goqu.C("category"), goqu.SUM("cost").As("total_cost")).
		GroupBy(goqu.C("category")).
		Executor()

	// Execute the query and scan the results into a slice of maps
	var expense []*dbmodel.ExpenseCategoryBreakdown
	if err := query.ScanStructs(&expense); err != nil {
		return nil, err
	}

	// Convert the results to model.ExpenseCategoryBreakdown objects
	var breakdowns []*model.ExpenseCategoryBreakdown
	for _, result := range expense {
		breakdown := model.ExpenseCategoryBreakdown{
			Category:  model.ExpenseCategory(result.Category),
			TotalCost: result.TotalCost,
		}
		breakdowns = append(breakdowns, &breakdown)
	}

	return breakdowns, nil
}
