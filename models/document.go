package models

import "io"

// Attachment struct for handling file attachments
type Attachment struct {
	Filename    string    `json:"filename,omitempty"`
	ContentType string    `json:"content_type,omitempty"`
	Content     io.Reader `json:"-"`
	Size        int64     `json:"size,omitempty"`
}

type Document struct {
	ID         string     `json:"_id,omitempty"`
	Rev        string     `json:"_rev,omitempty"`
	Name       string     `json:"name,omitempty"`
	Gender     string     `json:"gender,omitempty"`
	Age        int        `json:"age,omitempty"`
	Class      string     `json:"class,omitempty"`
	Majors     string     `json:"majors,omitempty"`
	Attachment Attachment `json:"attachment,omitempty"`
}

// RequestDoc struct for representing document creation requests
type RequestDoc struct {
	ID         string     `json:"_id,omitempty"`
	Name       string     `json:"name,omitempty"`
	Gender     string     `json:"gender,omitempty"`
	Age        int        `json:"age,omitempty"`
	Class      string     `json:"class,omitempty"`
	Majors     string     `json:"majors,omitempty"`
	Attachment Attachment `json:"attachment,omitempty"`
}

type RequestUpdateDoc struct {
	Name       string     `json:"name,omitempty"`
	Gender     string     `json:"gender,omitempty"`
	Age        int        `json:"age,omitempty"`
	Class      string     `json:"class,omitempty"`
	Majors     string     `json:"majors,omitempty"`
	Attachment Attachment `json:"attachment,omitempty"`
}
