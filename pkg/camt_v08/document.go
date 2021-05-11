// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v08

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt00400108 struct {
	XMLName xml.Name
	Attrs   []utils.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
	RtrAcct ReturnAccountV08 `xml:"RtrAcct"`
}

func (doc DocumentCamt00400108) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt00400108) NameSpace() string {
	return utils.DocumentCamt00400108NameSpace
}

func (doc DocumentCamt00400108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
		RtrAcct ReturnAccountV08 `xml:"RtrAcct"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt00500108 struct {
	XMLName xml.Name
	Attrs   []utils.Attr      `xml:",any,attr,omitempty" json:",omitempty"`
	GetTx   GetTransactionV08 `xml:"GetTx"`
}

func (doc DocumentCamt00500108) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt00500108) NameSpace() string {
	return utils.DocumentCamt00500108NameSpace
}

func (doc DocumentCamt00500108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr      `xml:",any,attr,omitempty" json:",omitempty"`
		GetTx   GetTransactionV08 `xml:"GetTx"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt00600108 struct {
	XMLName xml.Name
	Attrs   []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
	RtrTx   ReturnTransactionV08 `xml:"RtrTx"`
}

func (doc DocumentCamt00600108) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt00600108) NameSpace() string {
	return utils.DocumentCamt00600108NameSpace
}

func (doc DocumentCamt00600108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
		RtrTx   ReturnTransactionV08 `xml:"RtrTx"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt00700108 struct {
	XMLName xml.Name
	Attrs   []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
	ModfyTx ModifyTransactionV08 `xml:"ModfyTx"`
}

func (doc DocumentCamt00700108) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt00700108) NameSpace() string {
	return utils.DocumentCamt00700108NameSpace
}

func (doc DocumentCamt00700108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
		ModfyTx ModifyTransactionV08 `xml:"ModfyTx"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt00800108 struct {
	XMLName xml.Name
	Attrs   []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
	CclTx   CancelTransactionV08 `xml:"CclTx"`
}

func (doc DocumentCamt00800108) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt00800108) NameSpace() string {
	return utils.DocumentCamt00800108NameSpace
}

func (doc DocumentCamt00800108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr         `xml:",any,attr,omitempty" json:",omitempty"`
		CclTx   CancelTransactionV08 `xml:"CclTx"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt01000108 struct {
	XMLName xml.Name
	Attrs   []utils.Attr   `xml:",any,attr,omitempty" json:",omitempty"`
	RtrLmt  ReturnLimitV08 `xml:"RtrLmt"`
}

func (doc DocumentCamt01000108) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt01000108) NameSpace() string {
	return utils.DocumentCamt01000108NameSpace
}

func (doc DocumentCamt01000108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName xml.Name
		Attrs   []utils.Attr   `xml:",any,attr,omitempty" json:",omitempty"`
		RtrLmt  ReturnLimitV08 `xml:"RtrLmt"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt02600108 struct {
	XMLName    xml.Name
	Attrs      []utils.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
	UblToApply UnableToApplyV08 `xml:"UblToApply"`
}

func (doc DocumentCamt02600108) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02600108) NameSpace() string {
	return utils.DocumentCamt02600108NameSpace
}

func (doc DocumentCamt02600108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName    xml.Name
		Attrs      []utils.Attr     `xml:",any,attr,omitempty" json:",omitempty"`
		UblToApply UnableToApplyV08 `xml:"UblToApply"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt02700108 struct {
	XMLName   xml.Name
	Attrs     []utils.Attr       `xml:",any,attr,omitempty" json:",omitempty"`
	ClmNonRct ClaimNonReceiptV08 `xml:"ClmNonRct"`
}

func (doc DocumentCamt02700108) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02700108) NameSpace() string {
	return utils.DocumentCamt02700108NameSpace
}

func (doc DocumentCamt02700108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName   xml.Name
		Attrs     []utils.Attr       `xml:",any,attr,omitempty" json:",omitempty"`
		ClmNonRct ClaimNonReceiptV08 `xml:"ClmNonRct"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt03700108 struct {
	XMLName       xml.Name
	Attrs         []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
	DbtAuthstnReq DebitAuthorisationRequestV08 `xml:"DbtAuthstnReq"`
}

func (doc DocumentCamt03700108) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03700108) NameSpace() string {
	return utils.DocumentCamt03700108NameSpace
}

func (doc DocumentCamt03700108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName       xml.Name
		Attrs         []utils.Attr                 `xml:",any,attr,omitempty" json:",omitempty"`
		DbtAuthstnReq DebitAuthorisationRequestV08 `xml:"DbtAuthstnReq"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt05200108 struct {
	XMLName          xml.Name
	Attrs            []utils.Attr                   `xml:",any,attr,omitempty" json:",omitempty"`
	BkToCstmrAcctRpt BankToCustomerAccountReportV08 `xml:"BkToCstmrAcctRpt"`
}

func (doc DocumentCamt05200108) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05200108) NameSpace() string {
	return utils.DocumentCamt05200108NameSpace
}

func (doc DocumentCamt05200108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName          xml.Name
		Attrs            []utils.Attr                   `xml:",any,attr,omitempty" json:",omitempty"`
		BkToCstmrAcctRpt BankToCustomerAccountReportV08 `xml:"BkToCstmrAcctRpt"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt05300108 struct {
	XMLName       xml.Name
	Attrs         []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
	BkToCstmrStmt BankToCustomerStatementV08 `xml:"BkToCstmrStmt"`
}

func (doc DocumentCamt05300108) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05300108) NameSpace() string {
	return utils.DocumentCamt05300108NameSpace
}

func (doc DocumentCamt05300108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName       xml.Name
		Attrs         []utils.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
		BkToCstmrStmt BankToCustomerStatementV08 `xml:"BkToCstmrStmt"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}

type DocumentCamt05400108 struct {
	XMLName               xml.Name
	Attrs                 []utils.Attr                             `xml:",any,attr,omitempty" json:",omitempty"`
	BkToCstmrDbtCdtNtfctn BankToCustomerDebitCreditNotificationV08 `xml:"BkToCstmrDbtCdtNtfctn"`
}

func (doc DocumentCamt05400108) Validate() error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() != attr.Value {
			return utils.NewErrInvalidNameSpace()
		}
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05400108) NameSpace() string {
	return utils.DocumentCamt05400108NameSpace
}

func (doc DocumentCamt05400108) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			doc.XMLName.Space = ""
		}
	}
	α := struct {
		XMLName               xml.Name
		Attrs                 []utils.Attr                             `xml:",any,attr,omitempty" json:",omitempty"`
		BkToCstmrDbtCdtNtfctn BankToCustomerDebitCreditNotificationV08 `xml:"BkToCstmrDbtCdtNtfctn"`
	}(doc)
	if len(doc.XMLName.Local) > 0 {
		start.Name.Local = doc.XMLName.Local
	}
	return e.EncodeElement(&α, start)
}
