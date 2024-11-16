package validate

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	xsdvalidate "github.com/terminalstatic/go-xsd-validate"
)

var xsdhandler *xsdvalidate.XsdHandler

func Init(xsdPath string) error {
	err := xsdvalidate.Init()
	if err != nil {
		return err
	}

	fullpath := filepath.Join(xsdPath, "maindoc/UBL-Invoice-2.1.xsd")

	xsdhandler, err = xsdvalidate.NewXsdHandlerUrl(fullpath, xsdvalidate.ParsErrVerbose)
	return err
}

func Free() {
	xsdvalidate.Cleanup()
	xsdhandler.Free()
}

func File(filename string) error {

	xmlFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer xmlFile.Close()
	inXml, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return err
	}

	return Bytes(inXml)
}

func Bytes(xml []byte) error {
	err := xsdhandler.ValidateMem(xml, xsdvalidate.ValidErrDefault)
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
