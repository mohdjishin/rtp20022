// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 with prefix 're'
package admn_006_001_01

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD Elements

type Document struct {
	XMLName      xml.Name
	AdmnEchoResp EchoResponse `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 AdmnEchoResp"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Document) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.AdmnEchoResp, xml.StartElement{Name: xml.Name{Local: "re:AdmnEchoResp"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// XSD ComplexType declarations

type BranchAndFinancialInstitutionIdentification4ADMN struct {
	FinInstnId FinancialInstitutionIdentification7ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 FinInstnId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification4ADMN) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FinInstnId, xml.StartElement{Name: xml.Name{Local: "re:FinInstnId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type ClearingSystemMemberIdentification2ADMN struct {
	MmbId Min11Max11Text `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 MmbId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification2ADMN) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MmbId, xml.StartElement{Name: xml.Name{Local: "re:MmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type DocumentTCH struct {
	AdmnEchoResp EchoResponseTCH `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 AdmnEchoResp"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v DocumentTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.AdmnEchoResp, xml.StartElement{Name: xml.Name{Local: "re:AdmnEchoResp"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type EchoResp struct {
	InstgAgt     BranchAndFinancialInstitutionIdentification4ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 InstgAgt"`
	InstdAgt     BranchAndFinancialInstitutionIdentification4ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 InstdAgt"`
	OrgnlInstrId Max35Text                                        `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 OrgnlInstrId"`
	FnctnCd      EchoCode                                         `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 FnctnCd"`
	TxSts        TransactionIndividualStatus3CodeEcho             `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 TxSts"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v EchoResp) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.InstgAgt, xml.StartElement{Name: xml.Name{Local: "re:InstgAgt"}})
	e.EncodeElement(v.InstdAgt, xml.StartElement{Name: xml.Name{Local: "re:InstdAgt"}})
	e.EncodeElement(v.OrgnlInstrId, xml.StartElement{Name: xml.Name{Local: "re:OrgnlInstrId"}})
	e.EncodeElement(v.FnctnCd, xml.StartElement{Name: xml.Name{Local: "re:FnctnCd"}})
	e.EncodeElement(v.TxSts, xml.StartElement{Name: xml.Name{Local: "re:TxSts"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type EchoResponse struct {
	GrpHdr       GrpHdr   `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 GrpHdr"`
	EchoResponse EchoResp `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 EchoResponse"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v EchoResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.GrpHdr, xml.StartElement{Name: xml.Name{Local: "re:GrpHdr"}})
	e.EncodeElement(v.EchoResponse, xml.StartElement{Name: xml.Name{Local: "re:EchoResponse"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type EchoResponseTCH struct {
	GrpHdr       GrpHdrTCH `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 GrpHdr"`
	EchoResponse EchoResp  `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 EchoResponse"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v EchoResponseTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.GrpHdr, xml.StartElement{Name: xml.Name{Local: "re:GrpHdr"}})
	e.EncodeElement(v.EchoResponse, xml.StartElement{Name: xml.Name{Local: "re:EchoResponse"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type FinancialInstitutionIdentification7ADMN struct {
	ClrSysMmbId ClearingSystemMemberIdentification2ADMN `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 ClrSysMmbId"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionIdentification7ADMN) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.ClrSysMmbId, xml.StartElement{Name: xml.Name{Local: "re:ClrSysMmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type GrpHdr struct {
	MsgId   Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 MsgId"`
	CreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 CreDtTm"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GrpHdr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MsgId, xml.StartElement{Name: xml.Name{Local: "re:MsgId"}})
	e.EncodeElement(v.CreDtTm, xml.StartElement{Name: xml.Name{Local: "re:CreDtTm"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type GrpHdrTCH struct {
	MsgId   Max35TextTCH    `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 MsgId"`
	CreDtTm rtp.ISODateTime `xml:"urn:iso:std:iso:20022:tech:xsd:admn.006.001.01 CreDtTm"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GrpHdrTCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MsgId, xml.StartElement{Name: xml.Name{Local: "re:MsgId"}})
	e.EncodeElement(v.CreDtTm, xml.StartElement{Name: xml.Name{Local: "re:CreDtTm"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// XSD SimpleType declarations

type EchoCode string

const EchoCode731 EchoCode = "731"

type Max35Text string

type Max35TextTCH string

type Min11Max11Text string

type TransactionIndividualStatus3CodeEcho string

const TransactionIndividualStatus3CodeEchoActc TransactionIndividualStatus3CodeEcho = "ACTC"
