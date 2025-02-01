package run

import (
	"bytes"
	"github.com/xanderazuaje/xocket/parsing"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func AddFormData(test *parsing.Test, body *bytes.Buffer) error {
	if test.Header == nil {
		test.Header = http.Header{}
	}
	switch test.Form.Type {
	case "multipart":
		return setMultipart(test, body)
	default:
		encodedForm := test.Form.Values.Encode()
		_, err := body.WriteString(encodedForm)
		if err != nil {
			return err
		}
		test.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	return nil
}

func setMultipart(test *parsing.Test, body *bytes.Buffer) error {
	w := multipart.NewWriter(body)
	for key, values := range test.Form.Values {
		for _, v := range values {
			_ = w.WriteField(key, v)
		}
	}
	for key, filePath := range test.Form.Files {
		fw, err := w.CreateFormFile(key, filepath.Base(filePath))
		if err != nil {
			return err
		}
		fileData, err := os.Open(filePath)
		if err != nil {
			return err
		}
		defer fileData.Close()
		_, err = io.Copy(fw, fileData)
		if err != nil {
			return err
		}
	}
	err := w.Close()
	if err != nil {
		return err
	}
	test.Header.Set("Content-Type", w.FormDataContentType())
	return nil
}
