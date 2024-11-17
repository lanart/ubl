package validator

import (
	"embed"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	xsdvalidate "github.com/terminalstatic/go-xsd-validate"
)

//go:embed "xsd"
var xsdFiles embed.FS

type Validator struct {
	xsdhandler *xsdvalidate.XsdHandler
	xsdPath    string
}

func New() (*Validator, error) {
	xsdPath, err := extractXSDs()
	if err != nil {
		return nil, fmt.Errorf("extracting XSD's: %w", err)
	}

	err = xsdvalidate.Init()
	if err != nil {
		return nil, err
	}

	fullpath := filepath.Join(xsdPath, "maindoc/UBL-Invoice-2.1.xsd")

	v := &Validator{}
	v.xsdPath = xsdPath
	v.xsdhandler, err = xsdvalidate.NewXsdHandlerUrl(fullpath, xsdvalidate.ParsErrVerbose)
	return v, err
}

func (v *Validator) Free() {
	os.RemoveAll(v.xsdPath)
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

func extractXSDs() (string, error) {
	tempDir, err := os.MkdirTemp("", "xsd")
	if err != nil {
		return "", fmt.Errorf("create temp dir for xsd: %w", err)
	}

	err = extractPath(tempDir, "maindoc")
	if err != nil {
		return "", err
	}

	err = extractPath(tempDir, "common")
	if err != nil {
		return "", err
	}

	return tempDir, nil
}

func extractPath(tempDir, subpath string) error {
	p := filepath.Join("xsd", subpath)

	err := os.Mkdir(filepath.Join(tempDir, subpath), 0755)
	if err != nil {
		return err
	}

	entries, err := xsdFiles.ReadDir(p)
	if err != nil {
		return fmt.Errorf("read dir %v: %w", p, err)
	}

	for _, entry := range entries {
		content, err := xsdFiles.ReadFile(filepath.Join(p, entry.Name()))
		if err != nil {
			return err
		}
		err = os.WriteFile(filepath.Join(tempDir, subpath, entry.Name()), content, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
