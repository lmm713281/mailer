package mailer

import (
	"io"
	"net/mail"
)

// Email email struct
type Email struct {
	TO, CC, BCC   []mail.Address
	From, ReplyTo *mail.Address
	Subject       string
	Headers       mail.Header
	Attachments   []Attachment
	Template      string // template name
	Layout        string // application
}

// Attachment attachment struct
type Attachment struct {
	FileName string
	Inline   bool
	MimeType string
	Content  io.Reader
}

// Merge merge email struct and create a new one
func (email Email) Merge(e Email) Email {
	if len(e.TO) > 0 {
		email.TO = e.TO
	}

	if len(e.CC) > 0 {
		email.CC = e.CC
	}

	if len(e.BCC) > 0 {
		email.BCC = e.BCC
	}

	if e.From == nil {
		email.From = e.From
	}

	if e.ReplyTo == nil {
		email.ReplyTo = e.ReplyTo
	}

	if e.Subject != "" {
		email.Subject = e.Subject
	}
	if e.Headers != nil {
		email.Headers = e.Headers
	}

	if len(e.Attachments) > 0 {
		email.Attachments = e.Attachments
	}

	if e.Template != "" {
		email.Template = e.Template
	}

	if e.Layout != "" {
		email.Layout = e.Layout
	}

	return email
}
