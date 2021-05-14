package models

import "encoding/xml"

type EnvelopeResponse struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Xsd     string   `xml:"xsd,attr"`
	Xsi     string   `xml:"xsi,attr"`
	S       string   `xml:"s,attr"`
	Body    struct {
		Text                string `xml:",chardata"`
		SendMessageResponse struct {
			Text     string `xml:",chardata"`
			Xsi      string `xml:"xsi,attr"`
			Xsd      string `xml:"xsd,attr"`
			Xmlns    string `xml:"xmlns,attr"`
			Response struct {
				Text         string `xml:",chardata"`
				Xmlns        string `xml:"xmlns,attr"`
				ResponseInfo struct {
					Text         string `xml:",chardata"`
					ResponseDate string `xml:"responseDate"`
					Status       struct {
						Text    string `xml:",chardata"`
						Code    string `xml:"code"`
						Message string `xml:"message"`
					} `xml:"status"`
				} `xml:"responseInfo"`
				ResponseData struct {
					Text string `xml:",chardata"`
					Data struct {
						Text   string `xml:",chardata"`
						Q1     string `xml:"q1,attr"`
						Type   string `xml:"type,attr"`
						Result struct {
							Text  string `xml:",chardata"`
							Covid struct {
								Text    string `xml:",chardata"`
								Key     string `xml:"Key"`
								Patient struct {
									Text                     string `xml:",chardata"`
									IIN                      string `xml:"IIN"`
									IsResident               string `xml:"IsResident"`
									Birthday                 string `xml:"Birthday"`
									Gender                   string `xml:"Gender"`
									FirstName                string `xml:"FirstName"`
									LastName                 string `xml:"LastName"`
									MiddleName               string `xml:"MiddleName"`
									AddressOfActualResidence string `xml:"AddressOfActualResidence"`
									PlaceOfStudyOrWork       string `xml:"PlaceOfStudyOrWork"`
									Phone                    string `xml:"Phone"`
								} `xml:"Patient"`
								HasSymptomsCOVID string `xml:"HasSymptomsCOVID"`
								ProbeStatus      string `xml:"ProbeStatus"`
								CollectedTime    string `xml:"CollectedTime"`
								ProtocolDate     string `xml:"ProtocolDate"`
								ResearchResults  string `xml:"ResearchResults"`
								CreatedAt        string `xml:"CreatedAt"`
							} `xml:"covid"`
						} `xml:"result"`
					} `xml:"data"`
				} `xml:"responseData"`
			} `xml:"response"`
		} `xml:"SendMessageResponse"`
	} `xml:"Body"`
}
