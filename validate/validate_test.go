package validate_test

import (
	"strings"
	"testing"

	"github.com/lanart/ubl/validate"
)

func TestValidate(t *testing.T) {

	validate.Init()
	defer validate.Free()

	err := validate.Validate("testdata/invoice_base_correct.xml")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	err = validate.Validate("testdata/invoice_syntax_error.xml")
	if err == nil {
		t.Errorf("expected an error but did not receive one")
	}
	if err.Error() != "Malformed xml document" {
		t.Errorf("expected 'Malformed xml document' but got %v", err)
	}

	err = validate.Validate("testdata/invoice_missing_element.xml")
	if err == nil {
		t.Errorf("expected an error but did not receive one")
	}
	if !strings.Contains(err.Error(), "This element is not expected") {
		t.Errorf("expected 'This element is not expected' but got %v", err)
	}
}
