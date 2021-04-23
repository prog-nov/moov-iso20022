// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v07

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt00300107 struct {
	XMLName                 *xml.Name     `json:",omitempty"`
	Xmlns                   string        `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool          `xml:",omitempty" json:",omitempty"`
	GetAcct                 GetAccountV07 `xml:"GetAcct"`
}

func (doc DocumentCamt00300107) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt00300107) NameSpace() string {
	return utils.DocumentCamt00300107NameSpace
}

func (doc DocumentCamt00300107) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetAcct GetAccountV07 `xml:"GetAcct"`
	}
	output.GetAcct = doc.GetAcct
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt00900107 struct {
	XMLName                 *xml.Name   `json:",omitempty"`
	Xmlns                   string      `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool        `xml:",omitempty" json:",omitempty"`
	GetLmt                  GetLimitV07 `xml:"GetLmt"`
}

func (doc DocumentCamt00900107) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt00900107) NameSpace() string {
	return utils.DocumentCamt00900107NameSpace
}

func (doc DocumentCamt00900107) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetLmt GetLimitV07 `xml:"GetLmt"`
	}
	output.GetLmt = doc.GetLmt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt01100107 struct {
	XMLName                 *xml.Name      `json:",omitempty"`
	Xmlns                   string         `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool           `xml:",omitempty" json:",omitempty"`
	ModfyLmt                ModifyLimitV07 `xml:"ModfyLmt"`
}

func (doc DocumentCamt01100107) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt01100107) NameSpace() string {
	return utils.DocumentCamt01100107NameSpace
}

func (doc DocumentCamt01100107) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ModfyLmt ModifyLimitV07 `xml:"ModfyLmt"`
	}
	output.ModfyLmt = doc.ModfyLmt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt01200107 struct {
	XMLName                 *xml.Name      `json:",omitempty"`
	Xmlns                   string         `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool           `xml:",omitempty" json:",omitempty"`
	DelLmt                  DeleteLimitV07 `xml:"DelLmt"`
}

func (doc DocumentCamt01200107) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt01200107) NameSpace() string {
	return utils.DocumentCamt01200107NameSpace
}

func (doc DocumentCamt01200107) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		DelLmt DeleteLimitV07 `xml:"DelLmt"`
	}
	output.DelLmt = doc.DelLmt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt01900107 struct {
	XMLName                 *xml.Name                       `json:",omitempty"`
	Xmlns                   string                          `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                            `xml:",omitempty" json:",omitempty"`
	RtrBizDayInf            ReturnBusinessDayInformationV07 `xml:"RtrBizDayInf"`
}

func (doc DocumentCamt01900107) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt01900107) NameSpace() string {
	return utils.DocumentCamt01900107NameSpace
}

func (doc DocumentCamt01900107) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		RtrBizDayInf ReturnBusinessDayInformationV07 `xml:"RtrBizDayInf"`
	}
	output.RtrBizDayInf = doc.RtrBizDayInf
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt02300107 struct {
	XMLName                 *xml.Name        `json:",omitempty"`
	Xmlns                   string           `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool             `xml:",omitempty" json:",omitempty"`
	BckpPmt                 BackupPaymentV07 `xml:"BckpPmt"`
}

func (doc DocumentCamt02300107) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02300107) NameSpace() string {
	return utils.DocumentCamt02300107NameSpace
}

func (doc DocumentCamt02300107) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		BckpPmt BackupPaymentV07 `xml:"BckpPmt"`
	}
	output.BckpPmt = doc.BckpPmt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt08700107 struct {
	XMLName                 *xml.Name                 `json:",omitempty"`
	Xmlns                   string                    `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                      `xml:",omitempty" json:",omitempty"`
	ReqToModfyPmt           RequestToModifyPaymentV07 `xml:"ReqToModfyPmt"`
}

func (doc DocumentCamt08700107) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt08700107) NameSpace() string {
	return utils.DocumentCamt08700107NameSpace
}

func (doc DocumentCamt08700107) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ReqToModfyPmt RequestToModifyPaymentV07 `xml:"ReqToModfyPmt"`
	}
	output.ReqToModfyPmt = doc.ReqToModfyPmt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}
