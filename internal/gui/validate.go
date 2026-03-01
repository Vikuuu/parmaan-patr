package gui

import "github.com/Vikuuu/invoice_generator/internal/validation"

func validateCompanyDetail(name, gst, addr string) error {
	var err error
	if err = validation.ValidateString(name); err != nil {
		return err
	}
	if err = validation.ValidateGST(gst); err != nil {
		return err
	}
	if err = validation.ValidateString(gst); err != nil {
		return err
	}

	return nil
}
