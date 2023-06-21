// Code generated by GoComply XSD2Go for Moov; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:admi.004.001.02
package admi_004_001_02

import (
	"encoding/xml"

	"github.com/moov-io/rtp20022/pkg/rtp"
)

// XSD Elements

type Document struct {
	XMLName      xml.Name
	SysEvtNtfctn SystemEventNotificationV02 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 SysEvtNtfctn"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Document) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.SysEvtNtfctn, xml.StartElement{Name: xml.Name{Local: "ne:SysEvtNtfctn"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// XSD ComplexType declarations

type Event2 struct {
	EvtCd    Max4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtCd"`
	EvtParam []*Max35Text         `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtParam,omitempty"`
	EvtDesc  *Max1000Text         `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtDesc,omitempty"`
	EvtTm    *rtp.ISODateTime     `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtTm,omitempty"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Event2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.EvtCd, xml.StartElement{Name: xml.Name{Local: "ne:EvtCd"}})
	e.EncodeElement(v.EvtParam, xml.StartElement{Name: xml.Name{Local: "ne:EvtParam"}})
	e.EncodeElement(v.EvtDesc, xml.StartElement{Name: xml.Name{Local: "ne:EvtDesc"}})
	e.EncodeElement(v.EvtTm, xml.StartElement{Name: xml.Name{Local: "ne:EvtTm"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

type SystemEventNotificationV02 struct {
	EvtInf Event2 `xml:"urn:iso:std:iso:20022:tech:xsd:admi.004.001.02 EvtInf"`
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v SystemEventNotificationV02) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.EvtInf, xml.StartElement{Name: xml.Name{Local: "ne:EvtInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// XSD SimpleType declarations

type Max1000Text string

type Max35Text string

type Max4AlphaNumericText string
