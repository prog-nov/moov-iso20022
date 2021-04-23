// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package camt_v05

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type DocumentCamt01800105 struct {
	XMLName                 *xml.Name                    `json:",omitempty"`
	Xmlns                   string                       `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                         `xml:",omitempty" json:",omitempty"`
	GetBizDayInf            GetBusinessDayInformationV05 `xml:"GetBizDayInf"`
}

func (doc DocumentCamt01800105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt01800105) NameSpace() string {
	return utils.DocumentCamt01800105NameSpace
}

func (doc DocumentCamt01800105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetBizDayInf GetBusinessDayInformationV05 `xml:"GetBizDayInf"`
	}
	output.GetBizDayInf = doc.GetBizDayInf
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt02500105 struct {
	XMLName                 *xml.Name  `json:",omitempty"`
	Xmlns                   string     `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool       `xml:",omitempty" json:",omitempty"`
	Rct                     ReceiptV05 `xml:"Rct"`
}

func (doc DocumentCamt02500105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02500105) NameSpace() string {
	return utils.DocumentCamt02500105NameSpace
}

func (doc DocumentCamt02500105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		Rct ReceiptV05 `xml:"Rct"`
	}
	output.Rct = doc.Rct
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt02600105 struct {
	XMLName                 *xml.Name        `json:",omitempty"`
	Xmlns                   string           `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool             `xml:",omitempty" json:",omitempty"`
	UblToApply              UnableToApplyV05 `xml:"UblToApply"`
}

func (doc DocumentCamt02600105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02600105) NameSpace() string {
	return utils.DocumentCamt02600105NameSpace
}

func (doc DocumentCamt02600105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		UblToApply UnableToApplyV05 `xml:"UblToApply"`
	}
	output.UblToApply = doc.UblToApply
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt02800105 struct {
	XMLName                 *xml.Name                       `json:",omitempty"`
	Xmlns                   string                          `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                            `xml:",omitempty" json:",omitempty"`
	AddtlPmtInf             AdditionalPaymentInformationV05 `xml:"AddtlPmtInf"`
}

func (doc DocumentCamt02800105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt02800105) NameSpace() string {
	return utils.DocumentCamt02800105NameSpace
}

func (doc DocumentCamt02800105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AddtlPmtInf AdditionalPaymentInformationV05 `xml:"AddtlPmtInf"`
	}
	output.AddtlPmtInf = doc.AddtlPmtInf
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt03000105 struct {
	XMLName                 *xml.Name                       `json:",omitempty"`
	Xmlns                   string                          `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                            `xml:",omitempty" json:",omitempty"`
	NtfctnOfCaseAssgnmt     NotificationOfCaseAssignmentV05 `xml:"NtfctnOfCaseAssgnmt"`
}

func (doc DocumentCamt03000105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03000105) NameSpace() string {
	return utils.DocumentCamt03000105NameSpace
}

func (doc DocumentCamt03000105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		NtfctnOfCaseAssgnmt NotificationOfCaseAssignmentV05 `xml:"NtfctnOfCaseAssgnmt"`
	}
	output.NtfctnOfCaseAssgnmt = doc.NtfctnOfCaseAssgnmt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt03500105 struct {
	XMLName                 *xml.Name                         `json:",omitempty"`
	Xmlns                   string                            `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                              `xml:",omitempty" json:",omitempty"`
	PrtryFrmtInvstgtn       ProprietaryFormatInvestigationV05 `xml:"PrtryFrmtInvstgtn"`
}

func (doc DocumentCamt03500105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03500105) NameSpace() string {
	return utils.DocumentCamt03500105NameSpace
}

func (doc DocumentCamt03500105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		PrtryFrmtInvstgtn ProprietaryFormatInvestigationV05 `xml:"PrtryFrmtInvstgtn"`
	}
	output.PrtryFrmtInvstgtn = doc.PrtryFrmtInvstgtn
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt03600105 struct {
	XMLName                 *xml.Name                     `json:",omitempty"`
	Xmlns                   string                        `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                          `xml:",omitempty" json:",omitempty"`
	DbtAuthstnRspn          DebitAuthorisationResponseV05 `xml:"DbtAuthstnRspn"`
}

func (doc DocumentCamt03600105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03600105) NameSpace() string {
	return utils.DocumentCamt03600105NameSpace
}

func (doc DocumentCamt03600105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		DbtAuthstnRspn DebitAuthorisationResponseV05 `xml:"DbtAuthstnRspn"`
	}
	output.DbtAuthstnRspn = doc.DbtAuthstnRspn
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt03900105 struct {
	XMLName                 *xml.Name           `json:",omitempty"`
	Xmlns                   string              `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                `xml:",omitempty" json:",omitempty"`
	CaseStsRpt              CaseStatusReportV05 `xml:"CaseStsRpt"`
}

func (doc DocumentCamt03900105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt03900105) NameSpace() string {
	return utils.DocumentCamt03900105NameSpace
}

func (doc DocumentCamt03900105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		CaseStsRpt CaseStatusReportV05 `xml:"CaseStsRpt"`
	}
	output.CaseStsRpt = doc.CaseStsRpt
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt04600105 struct {
	XMLName                 *xml.Name         `json:",omitempty"`
	Xmlns                   string            `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool              `xml:",omitempty" json:",omitempty"`
	GetRsvatn               GetReservationV05 `xml:"GetRsvatn"`
}

func (doc DocumentCamt04600105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt04600105) NameSpace() string {
	return utils.DocumentCamt04600105NameSpace
}

func (doc DocumentCamt04600105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		GetRsvatn GetReservationV05 `xml:"GetRsvatn"`
	}
	output.GetRsvatn = doc.GetRsvatn
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt04800105 struct {
	XMLName                 *xml.Name            `json:",omitempty"`
	Xmlns                   string               `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                 `xml:",omitempty" json:",omitempty"`
	ModfyRsvatn             ModifyReservationV05 `xml:"ModfyRsvatn"`
}

func (doc DocumentCamt04800105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt04800105) NameSpace() string {
	return utils.DocumentCamt04800105NameSpace
}

func (doc DocumentCamt04800105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		ModfyRsvatn ModifyReservationV05 `xml:"ModfyRsvatn"`
	}
	output.ModfyRsvatn = doc.ModfyRsvatn
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt04900105 struct {
	XMLName                 *xml.Name            `json:",omitempty"`
	Xmlns                   string               `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                 `xml:",omitempty" json:",omitempty"`
	DelRsvatn               DeleteReservationV05 `xml:"DelRsvatn"`
}

func (doc DocumentCamt04900105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt04900105) NameSpace() string {
	return utils.DocumentCamt04900105NameSpace
}

func (doc DocumentCamt04900105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		DelRsvatn DeleteReservationV05 `xml:"DelRsvatn"`
	}
	output.DelRsvatn = doc.DelRsvatn
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt05000105 struct {
	XMLName                 *xml.Name                  `json:",omitempty"`
	Xmlns                   string                     `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                       `xml:",omitempty" json:",omitempty"`
	LqdtyCdtTrf             LiquidityCreditTransferV05 `xml:"LqdtyCdtTrf"`
}

func (doc DocumentCamt05000105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05000105) NameSpace() string {
	return utils.DocumentCamt05000105NameSpace
}

func (doc DocumentCamt05000105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		LqdtyCdtTrf LiquidityCreditTransferV05 `xml:"LqdtyCdtTrf"`
	}
	output.LqdtyCdtTrf = doc.LqdtyCdtTrf
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt05100105 struct {
	XMLName                 *xml.Name                 `json:",omitempty"`
	Xmlns                   string                    `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                      `xml:",omitempty" json:",omitempty"`
	LqdtyDbtTrf             LiquidityDebitTransferV05 `xml:"LqdtyDbtTrf"`
}

func (doc DocumentCamt05100105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05100105) NameSpace() string {
	return utils.DocumentCamt05100105NameSpace
}

func (doc DocumentCamt05100105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		LqdtyDbtTrf LiquidityDebitTransferV05 `xml:"LqdtyDbtTrf"`
	}
	output.LqdtyDbtTrf = doc.LqdtyDbtTrf
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt05600105 struct {
	XMLName                 *xml.Name                           `json:",omitempty"`
	Xmlns                   string                              `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                                `xml:",omitempty" json:",omitempty"`
	FIToFIPmtCxlReq         FIToFIPaymentCancellationRequestV05 `xml:"FIToFIPmtCxlReq"`
}

func (doc DocumentCamt05600105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt05600105) NameSpace() string {
	return utils.DocumentCamt05600105NameSpace
}

func (doc DocumentCamt05600105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		FIToFIPmtCxlReq FIToFIPaymentCancellationRequestV05 `xml:"FIToFIPmtCxlReq"`
	}
	output.FIToFIPmtCxlReq = doc.FIToFIPmtCxlReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}

type DocumentCamt06000105 struct {
	XMLName                 *xml.Name                  `json:",omitempty"`
	Xmlns                   string                     `xml:"xmlns,attr,omitempty" json:",omitempty"`
	DisableDefaultNamespace bool                       `xml:",omitempty" json:",omitempty"`
	AcctRptgReq             AccountReportingRequestV05 `xml:"AcctRptgReq"`
}

func (doc DocumentCamt06000105) Validate() error {
	if doc.NameSpace() != doc.Xmlns {
		return utils.NewErrInvalidNameSpace()
	}
	return utils.Validate(&doc)
}

func (doc DocumentCamt06000105) NameSpace() string {
	return utils.DocumentCamt06000105NameSpace
}

func (doc DocumentCamt06000105) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var output struct {
		AcctRptgReq AccountReportingRequestV05 `xml:"AcctRptgReq"`
	}
	output.AcctRptgReq = doc.AcctRptgReq
	utils.BaseXmlElement(&start, doc.XMLName, doc.NameSpace(), doc.DisableDefaultNamespace)
	return e.EncodeElement(&output, start)
}
