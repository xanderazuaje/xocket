package types

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type Form struct {
	Type   string
	Values url.Values
	Files  map[string]string `yaml:"_FILES"`
}

func (form *Form) GetBodyBuff(header *http.Header) (*bytes.Buffer, error) {
	body := &bytes.Buffer{}
	if *header == nil {
		*header = http.Header{}
	}
	switch form.Type {
	case "multipart":
		h, err := form.setMultipart(body)
		if err != nil {
			return nil, err
		}
		(*header).Set("Content-Type", h)
		return body, nil
	default:
		encodedForm := form.Values.Encode()
		_, err := body.WriteString(encodedForm)
		if err != nil {
			return nil, err
		}
		(*header).Set("Content-Type", "application/x-www-form-urlencoded")
		return body, nil
	}
}

func (form *Form) setMultipart(body *bytes.Buffer) (string, error) {
	w := multipart.NewWriter(body)
	for key, values := range form.Values {
		for _, v := range values {
			_ = w.WriteField(key, v)
		}
	}
	for key, filePath := range form.Files {
		fw, err := w.CreateFormFile(key, filepath.Base(filePath))
		if err != nil {
			return "", err
		}
		fileData, err := os.Open(filePath)
		if err != nil {
			return "", err
		}
		_, err = io.Copy(fw, fileData)
		if err != nil {
			return "", err
		}
		fileData.Close()
	}
	err := w.Close()
	if err != nil {
		return "", err
	}
	return w.FormDataContentType(), nil
}
