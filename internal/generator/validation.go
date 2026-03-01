package generator

import (
	"time"

	"github.com/Vikuuu/invoice_generator/internal/database"
)

func ValidateInvoiceData(
	from, to database.Company,
	date *time.Time,
	item database.Item,
	qty int,
	shipToAddr database.ShippingAddress,
	payTo database.PaymentDetail,
) error {
	return nil
}
