package handler

import (
	"context"
	"kraneapi/graph/model"
	service "kraneapi/pkg/api/services"
)

// CreateExpense creates a new expense for an event.
func CreateExpense(ctx context.Context, input model.CreateExpenseInput) (*model.Expense, error) {
	id, err := service.CreateExpense(ctx, input)
	if err != nil {
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
	id, err := service.UpdateExpense(ctx, input)
	if err != nil {
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

	eventExpenses, err := service.GetAllEventExpenses(ctx, eventID)
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

	expenseBreakdowns, err := service.GetEventExpensesBreakdown(ctx, eventID)
	if err != nil {
		return nil, err
	}

	var breakdowns []*model.ExpenseCategoryBreakdown
	for _, result := range expenseBreakdowns {
		breakdown := model.ExpenseCategoryBreakdown{
			Category:  model.ExpenseCategory(result.Category),
			TotalCost: result.TotalCost,
		}
		breakdowns = append(breakdowns, &breakdown)
	}

	return breakdowns, nil
}
