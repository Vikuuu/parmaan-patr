package gui

import (
	db "github.com/Vikuuu/invoice_generator/internal/database"
)

func (c *Config) dbAddCompany(name, gst, address string) error {
	arg := db.CreateCompanyParams{
		Name:    name,
		Gst:     gst,
		Address: address,
	}

	_, err := c.Queries.CreateCompany(c.Context, arg)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) dbAddPaymentMethod(
	holder string, number int64, ifsc, branch, name,
	virtualPaymentAddress string, companyId int64,
) error {
	arg := db.CreatePaymentDetailParams{
		AccHolder:          holder,
		AccNumber:          number,
		Ifsc:               ifsc,
		Branch:             branch,
		BankName:           name,
		VirtualPaymentAddr: virtualPaymentAddress,
		FkCompanyID:        companyId,
	}

	_, err := c.Queries.CreatePaymentDetail(c.Context, arg)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) dbAddItemMethod(name string, hsn, price int64) error {
	arg := db.CreateItemParams{
		Name:  name,
		Hsn:   hsn,
		Price: price,
	}

	_, err := c.Queries.CreateItem(c.Context, arg)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) dbAddShippingAddress(name, address string) error {
	arg := db.CreateShippingAddressParams{
		Name:    name,
		Address: address,
	}

	_, err := c.Queries.CreateShippingAddress(c.Context, arg)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) dbListCompany() ([]db.Company, error) {
	companies, err := c.Queries.ListCompany(c.Context)
	if err != nil {
		return nil, err
	}
	return companies, nil
}

func (c *Config) dbListPaymentDetail() ([]db.PaymentDetail, error) {
	paymentDetails, err := c.Queries.ListPaymentDetail(c.Context)
	if err != nil {
		return nil, err
	}
	return paymentDetails, nil
}

func (c *Config) dbListItem() ([]db.Item, error) {
	items, err := c.Queries.ListItem(c.Context)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (c *Config) dbListShippingAddress() ([]db.ShippingAddress, error) {
	address, err := c.Queries.ListShippingAddress(c.Context)
	if err != nil {
		return nil, err
	}

	return address, nil
}
