package validate

import (
	"fmt"
	"io/ioutil"
	"os"

	xsdvalidate "github.com/terminalstatic/go-xsd-validate"
)

var xsdhandler *xsdvalidate.XsdHandler

func Init() {
	xsdvalidate.Init()
	var err error
	xsdhandler, err = xsdvalidate.NewXsdHandlerUrl("./xsd/maindoc/UBL-Invoice-2.1.xsd", xsdvalidate.ParsErrVerbose)
	if err != nil {
		panic(err)
	}
}

func Free() {
	xsdvalidate.Cleanup()
	xsdhandler.Free()
}

func Validate(filename string) error {

	xmlFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	inXml, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		panic(err)
	}

	err = xsdhandler.ValidateMem(inXml, xsdvalidate.ValidErrDefault)
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
