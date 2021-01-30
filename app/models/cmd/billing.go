package cmd

import "github.com/Windscribe/fider/app/models/dto"

type CancelBillingSubscription struct {
}

type CreateBillingSubscription struct {
	PlanID string
}

type CreateBillingCustomer struct {
}

type DeleteBillingCustomer struct {
}

type ClearPaymentInfo struct {
}

type UpdatePaymentInfo struct {
	Input *dto.CreateEditBillingPaymentInfo
}
