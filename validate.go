package ubl

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	xsdvalidate "github.com/terminalstatic/go-xsd-validate"
)

func Validate(filename string) {

	xsdvalidate.Init()
	defer xsdvalidate.Cleanup()
	xsdhandler, err := xsdvalidate.NewXsdHandlerUrl("./xsd/maindoc/UBL-Invoice-2.1.xsd", xsdvalidate.ParsErrVerbose)
	if err != nil {
		log.Printf("xsdhandler: %v", err)
		panic(err)
	}
	defer xsdhandler.Free()

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

}
