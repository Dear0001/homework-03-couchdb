package models

type Attachment struct {
	ContentType string `json:"content_type,omitempty"`
	RevPos      int    `json:"revpos,omitempty"`
	Digest      string `json:"digest,omitempty"`
	Length      int64  `json:"length,omitempty"`
	Stub        bool   `json:"stub,omitempty"`
}

type Document struct {
	ID          string                `json:"_id,omitempty"`
	Rev         string                `json:"_rev,omitempty"`
	Name        string                `json:"name,omitempty"`
	Gender      string                `json:"gender,omitempty"`
	Age         int                   `json:"age,omitempty"`
	Class       string                `json:"class,omitempty"`
	Majors      string                `json:"majors,omitempty"`
	Attachments map[string]Attachment `json:"_attachments,omitempty"`
}

type Documents struct {
	ID          string                `json:"_id,omitempty"`
	Rev         string                `json:"_rev,omitempty"`
	Name        string                `json:"name,omitempty"`
	Gender      string                `json:"gender,omitempty"`
	Age         int                   `json:"age,omitempty"`
	Class       string                `json:"class,omitempty"`
	Majors      string                `json:"majors,omitempty"`
}

type RequestDoc struct {
	ID     string `json:"_id,omitempty"`
	Name   string `json:"name,omitempty"`
	Gender string `json:"gender,omitempty"`
	Age    int    `json:"age,omitempty"`
	Class  string `json:"class,omitempty"`
	Majors string `json:"majors,omitempty"`
}

type RequestUpdateDoc struct {
	Name   string `json:"name,omitempty"`
	Gender string `json:"gender,omitempty"`
	Age    int    `json:"age,omitempty"`
	Class  string `json:"class,omitempty"`
	Majors string `json:"majors,omitempty"`
}
