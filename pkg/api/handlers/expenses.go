package handler

import (
	"context"
	"errors"
	"kraneapi/graph/model"
	service "kraneapi/pkg/api/services"
	"kraneapi/utils"
)

// CreateExpense creates a new expense for an event.
func CreateExpense(ctx context.Context, input model.CreateExpenseInput) (*model.Expense, error) {

	// Check if the user has permission to add event sessions
	canAdd, err := utils.CheckPermission(ctx, input.EventID, "expenses", "add")
	if err != nil {
		return nil, err
	}
	if !canAdd {
		return nil, errors.New("permission denied")
	}

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

	canUpdate, err := utils.CheckPermission(ctx, input.EventID, "expenses", "update")
	if err != nil {
		return nil, err
	}
	if !canUpdate {
		return nil, errors.New("permission denied")
	}

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

	canView, err := utils.CheckPermission(ctx, eventID, "expenses", "view")
	if err != nil {
		return nil, err
	}
	if !canView {
		return nil, errors.New("permission denied")
	}

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

	canView, err := utils.CheckPermission(ctx, eventID, "expenses", "view")
	if err != nil {
		return nil, err
	}
	if !canView {
		return nil, errors.New("permission denied")
	}

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
