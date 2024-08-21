package adapter

import (
	"encoding/json"

	db "github.com/darwishdev/devkit-api-base/common/db/gen"
	abcv1 "github.com/darwishdev/devkit-api-base/common/pb/abc/v1"

	"golang.org/x/crypto/bcrypt"
)

// func (a *AccountsAdapter) OnGoingReservationFromReservation(resp *abcv1.CustomerReservation) *abcv1.CustomerOngoingReservation {
// 	units := make([]*abcv1.CustomerReservationUnit, 0)
// 	for _, record := range resp.Details {
// 		units = append(units, &abcv1.CustomerReservationUnit{
// 			Quantity: record.Quantity,
// 			Unit:     record,
// 		})
// 	}

// 	return &abcv1.CustomerOngoingReservation{
// 		Adults:                   resp.Adults,
// 		DateTo:                   resp.DateTo,
// 		Children:                 resp.Children,
// 		OwnerId:                  resp.OwnerId,
// 		Subtotal:                 resp.Subtotal,
// 		DateFrom:                 resp.DateFrom,
// 		TotalTax:                 resp.TotalTax,
// 		CreatedAt:                resp.CreatedAt,
// 		DateToId:                 resp.DateToId,
// 		DeletedAt:                resp.DeletedAt,
// 		OwnerName:                resp.OwnerName,
// 		UpdatedAt:                resp.UpdatedAt,
// 		CustomerId:               resp.CustomerId,
// 		PaidAmount:               resp.PaidAmount,
// 		PropertyId:               resp.PropertyId,
// 		DateFromId:               resp.DateFromId,
// 		CustomerName:             resp.CustomerName,
// 		RepresentativeOwnerName:  resp.RepresentativeOwnerName,
// 		PropertyName:             resp.PropertyName,
// 		ReservationId:            resp.ReservationId,
// 		CancelationFee:           resp.CancelationFee,
// 		ReservationCode:          resp.ReservationCode,
// 		ReservationStatus:        resp.ReservationStatus,
// 		CustomerNationalId:       resp.CustomerNationalId,
// 		ReservationStatusId:      resp.ReservationStatusId,
// 		RepresentativeOwnerId:    resp.RepresentativeOwnerId,
// 		RepresentativeOwnerEmail: resp.RepresentativeOwnerEmail,
// 		Units:                    units,
// 	}
// }
// func (a *AccountsAdapter) CustomerReservationsPrepare(resp []*abcv1.ReservationsViewEntity) []*abcv1.ReservationsViewEntity {
// 	for _, record := range resp {
// 		actions := []*abcv1.ReservationAction{}
// 		if record.ReservationStatusId <= 3 {
// 			if record.PaidAmount == 0 {
// 				actions = append(actions, &abcv1.ReservationAction{
// 					ActionName:          "cancel",
// 					ReservationStatusId: 4,
// 				})
// 			} else {
// 				actions = append(actions, &abcv1.ReservationAction{
// 					ActionName:          "refund",
// 					ReservationStatusId: 5,
// 				})
// 			}

// 		}

//			record.Actions = actions
//		}
//		return resp
//	}
func (a *AccountsAdapter) CustomerLoginGrpcFromSql(resp *db.CustomerFindRow) (*abcv1.CustomerLoginResponse, error) {
	customer := &abcv1.CustomerEntity{
		CustomerId:    resp.CustomerID,
		CustomerCode:  resp.CustomerCode,
		CustomerName:  resp.CustomerName,
		CustomerImage: resp.CustomerImage.String,
		CustomerEmail: resp.CustomerEmail,
		CustomerPhone: resp.CustomerPhone.String,
		Birthdate:     resp.Birthdate.Time.Format("2006-01-02"),
		CreatedAt:     resp.CreatedAt.Time.Format(a.dateFormat),
		UpdatedAt:     resp.UpdatedAt.Time.Format(a.dateFormat),
		DeletedAt:     resp.DeletedAt.Time.Format(a.dateFormat),
	}

	onGoingReservation := &abcv1.ReservationsViewEntity{}

	if len(resp.OngoingReservation) > 0 {
		if err := json.Unmarshal(resp.OngoingReservation, &onGoingReservation); err != nil {
			return nil, err
		}
	}
	onGoingPayment := &abcv1.PaymentEntity{}

	if len(resp.OngoingPayment) > 0 {
		if err := json.Unmarshal(resp.OngoingPayment, &onGoingPayment); err != nil {
			return nil, err
		}
	}
	response := &abcv1.CustomerLoginResponse{
		Customer: customer,
	}

	reservations := make([]*abcv1.ReservationsViewEntity, 0)
	if len(resp.Reservations) > 0 {
		if err := json.Unmarshal(resp.Reservations, &reservations); err != nil {
			return nil, err
		}
	}

	if onGoingPayment.PaymentId > 0 {
		response.OngoingPayment = onGoingPayment
	}
	if onGoingReservation.ReservationId > 0 {
		response.OngoingReservation = onGoingReservation
	}
	response.Reservations = reservations
	return response, nil
}

func (a *AccountsAdapter) CustomersListGrpcFromSql(resp *[]db.AccountsSchemaCustomer) (*abcv1.CustomersListResponse, error) {
	records := make([]*abcv1.CustomerEntity, 0)
	deletedRecords := make([]*abcv1.CustomerEntity, 0)
	for _, v := range *resp {
		record := a.CustomerEntityGrpcFromSql(&v)
		if v.DeletedAt.Valid {
			deletedRecords = append(deletedRecords, record)
		} else {
			records = append(records, record)
		}
	}
	return &abcv1.CustomersListResponse{
		Records:        records,
		DeletedRecords: deletedRecords,
	}, nil
}

func (a *AccountsAdapter) CustomerResetPasswordSqlFromGrpc(req *abcv1.CustomerResetPasswordRequest) *db.CustomerResetPasswordParams {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.CustomerPassword), bcrypt.DefaultCost)

	return &db.CustomerResetPasswordParams{
		CustomerEmail:    req.CustomerEmail,
		CustomerPassword: string(hashedPassword),
	}

}

func (a *AccountsAdapter) CustomerCreateUpdateSqlFromGrpc(req *abcv1.CustomerCreateUpdateRequest) *db.CustomerCreateUpdateParams {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.CustomerPassword), bcrypt.DefaultCost)
	return &db.CustomerCreateUpdateParams{
		CustomerID:         req.CustomerId,
		CustomerName:       req.CustomerName,
		Birthdate:          req.Birthdate,
		CustomerImage:      req.CustomerImage,
		CustomerEmail:      req.CustomerEmail,
		CustomerPhone:      req.CustomerPhone,
		CustomerPassword:   string(hashedPassword),
		CustomerNationalID: req.CustomerNationalId,
	}
}
