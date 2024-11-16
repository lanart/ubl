package validator

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	xsdvalidate "github.com/terminalstatic/go-xsd-validate"
)

type Validator struct {
	xsdhandler *xsdvalidate.XsdHandler
}

func New(xsdPath string) (*Validator, error) {
	err := xsdvalidate.Init()
	if err != nil {
		return nil, err
	}

	fullpath := filepath.Join(xsdPath, "maindoc/UBL-Invoice-2.1.xsd")

	v := &Validator{}
	v.xsdhandler, err = xsdvalidate.NewXsdHandlerUrl(fullpath, xsdvalidate.ParsErrVerbose)
	return v, err
}

func (v *Validator) Free() {
	v.xsdhandler.Free()
	xsdvalidate.Cleanup()
}

func (v *Validator) Validate(filename string) error {

	xmlFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer xmlFile.Close()
	inXml, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return err
	}

	return v.ValidateBytes(inXml)
}

func (v *Validator) ValidateBytes(xml []byte) error {
	err := v.xsdhandler.ValidateMem(xml, xsdvalidate.ValidErrDefault)
	if err != nil {
		switch err.(type) {
		case xsdvalidate.ValidationError:
			fmt.Println(err)
			fmt.Printf("Error in line: %d\n", err.(xsdvalidate.ValidationError).Errors[0].Line)
			fmt.Println(err.(xsdvalidate.ValidationError).Errors[0].Message)
		default:
			fmt.Println(err)
		}
	}

	return err
}
