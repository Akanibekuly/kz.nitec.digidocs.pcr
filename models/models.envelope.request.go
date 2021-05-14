package models

import "encoding/xml"

type Data struct {
	Text     string `xml:",chardata"`
	Ns6      string `xml:"xmlns:ns6,attr"`
	Xsi      string `xml:"xmlns:xsi,attr"`
	Type     string `xml:"xsi:type,attr"`
	Iin      string `xml:"iin"`
	Login    string `xml:"login"`
	Password string `xml:"password"`
}

type RequestData struct {
	Text string `xml:",chardata"`
	Data *Data  `xml:"data"`
}

type SenderCred struct {
	Text     string `xml:",chardata"`
	SenderId string `xml:"senderId"`
	Password string `xml:"password"`
}

type RequestInfo struct {
	Text        string      `xml:",chardata"`
	MessageId   string      `xml:"messageId"`
	ServiceId   string      `xml:"serviceId"`
	MessageDate string      `xml:"messageDate"`
	Sender      *SenderCred `xml:"sender"`
}

type Request struct {
	Text    string       `xml:",chardata"`
	ReqInfo *RequestInfo `xml:"requestInfo"`
	ReqData *RequestData `xml:"requestData"`
}

type SendMessageRequest struct {
	Text string   `xml:",chardata"`
	Ns2  string   `xml:"xmlns:ns2,attr"`
	Ns3  string   `xml:"xmlns:ns3,attr"`
	Ns4  string   `xml:"xmlns:ns4,attr"`
	Req  *Request `xml:"request"`
}

type BodyRequest struct {
	Text        string              `xml:",chardata"`
	SendMessage *SendMessageRequest `xml:"ns2:SendMessage"`
}

type EnvelopeRequest struct {
	XMLName xml.Name     `xml:"soap:Envelope"`
	Text    string       `xml:",chardata"`
	Xmlns   string       `xml:"xmlns:soap,attr"`
	Body    *BodyRequest `xml:"soap:Body"`
}
