package types

import (
	"io"
	"net/textproto"
)

type (
	Email struct {
		Subject     string
		From        string
		To          []string
		Parts       []EmailPart
		Attachments []EmailFile
		Inlines     []EmailFile
		Tags        []string
		Headers     textproto.MIMEHeader
	}
	EmailPart struct {
		ContentType string
		Body        string
	}
	EmailFile struct {
		Name     string
		MimeType string
		Data     io.ReadCloser
	}
)

func NewEmail() *Email {
	return new(Email)
}

func (e *Email) formatContact(name, address string) string {
	return name + " <" + address + ">"
}

func (e *Email) SetSubject(subject string) *Email {
	e.Subject = subject
	return e
}

func (e *Email) SetFromAddress(address string) *Email {
	e.From = address
	return e
}

func (e *Email) SetFromContact(name, address string) *Email {
	e.From = e.formatContact(name, address)
	return e
}

func (e *Email) AddToAddress(addresses ...string) *Email {
	e.To = append(e.To, addresses...)
	return e
}

func (e *Email) AddToContact(name, address string) *Email {
	e.AddToAddress(e.formatContact(name, address))
	return e
}

func (e *Email) SetBody(contentType string, body string) *Email {
	e.Parts = []EmailPart{
		{
			ContentType: contentType,
			Body:        body,
		},
	}
	return e
}

func (e *Email) SetBodyHTML(body string) *Email {
	return e.SetBody("text/html", body)
}

func (e *Email) SetBodyPlain(body string) *Email {
	return e.SetBody("text/plain", body)
}

func (e *Email) AddTag(tags ...string) *Email {
	e.Tags = append(e.Tags, tags...)
	return e
}
