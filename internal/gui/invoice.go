package gui

import (
	"log/slog"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"github.com/Vikuuu/invoice_generator/internal/database"
	"github.com/Vikuuu/invoice_generator/internal/generator"
)

func (c *Config) generateInvoicePage(a fyne.App, w fyne.Window) *widget.Form {
	slog.Info("GUI:", "msg", "Generate invoice page")
	companies, err := c.dbListCompany()
	if err != nil {
		slog.Error("DB:", "msg", err)
	}
	companiesMap := map[string]database.Company{}
	companySelectEntryData := make([]string, 0, len(companies))
	for _, company := range companies {
		companySelectEntryData = append(companySelectEntryData, company.Name)
		companiesMap[company.Name] = company
	}

	fromCompany := widget.NewSelectEntry(companySelectEntryData)
	toCompany := widget.NewSelectEntry(companySelectEntryData)

	dateEntry := widget.NewDateEntry()
	currentDate := time.Now()
	dateEntry.SetDate(&currentDate)

	items, err := c.dbListItem()
	if err != nil {
		slog.Error("DB:", "msg", err)
	}
	itemsMap := map[string]database.Item{}
	itemsSelectEntryData := make([]string, 0, len(items))
	for _, item := range items {
		itemsSelectEntryData = append(itemsSelectEntryData, item.Name)
		itemsMap[item.Name] = item
	}

	itemEntry := widget.NewSelectEntry(itemsSelectEntryData)
	qtyEntry := widget.NewEntry()

	shipAddrs, err := c.dbListShippingAddress()
	if err != nil {
		slog.Error("DB:", "msg", err)
	}
	shipAddrsMap := map[string]database.ShippingAddress{}
	shipAddrSelectEntryData := make([]string, 0, len(shipAddrs))
	for _, shipAddr := range shipAddrs {
		shipAddrSelectEntryData = append(shipAddrSelectEntryData, shipAddr.Name)
		shipAddrsMap[shipAddr.Name] = shipAddr
	}
	shipTo := widget.NewSelectEntry(shipAddrSelectEntryData)

	paymentDetails, err := c.dbListPaymentDetail()
	if err != nil {
		slog.Error("DB:", "msg", err)
	}
	paymentDetailsMap := map[string]database.PaymentDetail{}
	paymentDetailSelectEntryData := make([]string, 0, len(paymentDetails))
	for _, paymentDetail := range paymentDetails {
		paymentDetailSelectEntryData = append(
			paymentDetailSelectEntryData,
			paymentDetail.AccHolder,
		)
		paymentDetailsMap[paymentDetail.AccHolder] = paymentDetail
	}
	paymentTo := widget.NewSelectEntry(paymentDetailSelectEntryData)

	form := &widget.Form{}
	form.Append("Invoice From", fromCompany)
	form.Append("Invoice To", toCompany)
	form.Append("Date", dateEntry)
	form.Append("Item", itemEntry)
	form.Append("Qty", qtyEntry)
	form.Append("Ship to", shipTo)
	form.Append("Payment to", paymentTo)

	form.OnSubmit = func() {
		slog.Info("User Input:", "invoiceFrom", fromCompany.Text)

		fromCompany := companiesMap[fromCompany.Text]
		toCompany := companiesMap[toCompany.Text]
		date := dateEntry.Date
		item := itemsMap[itemEntry.Text]
		qty, _ := strconv.Atoi(qtyEntry.Text)
		shipAddr := shipAddrsMap[shipTo.Text]
		pay := paymentDetailsMap[paymentTo.Text]

		if err := generator.ValidateInvoiceData(fromCompany, toCompany, date, item, qty, shipAddr, pay); err != nil {
			slog.Error("Validation:", "msg", err)
		}

		slog.Info(
			"Invoice data from invoice.go",
			"from company",
			fromCompany,
			"to company",
			toCompany,
			"date",
			date,
			"item",
			item,
			"qty",
			qty,
			"shipAddr",
			shipAddr,
			"pay",
			pay,
		)

		invoiceData := generator.InvoiceDataMap(
			fromCompany,
			toCompany,
			date,
			item,
			qty,
			shipAddr,
			pay,
		)

		// TO-DO: Save the invoice data into database
		// Generate the invoice PDF
		generator.GenerateInvoice(invoiceData, c.TypstBinPath)
	}

	return form
}
