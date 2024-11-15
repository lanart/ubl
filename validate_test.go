package ubl_test

import (
	"testing"

	"github.com/lanart/ubl"
)

func TestValidate(t *testing.T) {
	ubl.Validate("testdata/invoice.xml")
}
