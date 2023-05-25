package huaweiapi

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type HuaweiAPI struct {
	host       string
	token      string
	session_id string
}

type SignalInfo struct {
	XMLName     xml.Name `xml:"response"`
	Band        string   `xml:"band"`
	UlBandwidth string   `xml:"ulbandwidth"`
	DlBandwidth string   `xml:"dlbandwidth"`
}

type DeviceInformation struct {
	XMLName         xml.Name `xml:"response"`
	DeviceName      string   `xml:"DeviceName"`
	SerialNumber    string   `xml:"SerialNumber"`
	Imei            string   `xml:"Imei"`
	Imsi            string   `xml:"Imsi"`
	HardwareVersion string   `xml:"HardwareVersion"`
	SoftwareVersion string   `xml:"SoftwareVersion"`
	WebUIVersion    string   `xml:"WebUIVersion"`
	SupportMode     string   `xml:"supportmode"`
	WorkMode        string   `xml:"workmode"`
}

type NetNetMode struct {
	XMLName     xml.Name `xml:"response"`
	NetworkMode string   `xml:"NetworkMode"`
	NetworkBand string   `xml:"NetworkBand"`
	LTEBand     string   `xml:"LTEBand"`
}

type HuaweiAPIIface interface {
	httpGetXml(url string, responseXml interface{}) error
	Connect(host string) error
	DeviceInformation() (*DeviceInformation, error)
	DeviceSignal() (*SignalInfo, error)
	NetNetMode() (*NetNetMode, error)
}

func New() *HuaweiAPI {
	var h HuaweiAPI
	return &h
}

func (h *HuaweiAPI) httpGetXml(url string, responseXml interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	if h.token != "" {
		req.Header.Set("__RequestVerificationToken", h.token)
	}
	if h.session_id != "" {
		req.AddCookie(&http.Cookie{Name: "SessionID", Value: h.session_id})
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("Bad http code: %d", resp.StatusCode))
	}

	answer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	//fmt.Println(url)
	//fmt.Println(string(answer))

	err = xml.Unmarshal(answer, responseXml)
	if err != nil {
		return err
	}

	return nil
}

func (h *HuaweiAPI) Connect(host string) error {
	h.host = host
	h.session_id = ""
	h.token = ""
	var xmlResp struct {
		XMLName xml.Name `xml:"response"`
		SesInfo string   `xml:"SesInfo"`
		TokInfo string   `xml:"TokInfo"`
	}
	err := h.httpGetXml(fmt.Sprintf("http://%s/api/webserver/SesTokInfo", h.host), &xmlResp)

	if err != nil {
		return err
	}

	h.session_id = xmlResp.SesInfo
	h.token = xmlResp.TokInfo

	return nil
}

func (h *HuaweiAPI) DeviceInformation() (*DeviceInformation, error) {
	var rv DeviceInformation
	err := h.httpGetXml(fmt.Sprintf("http://%s/api/device/information", h.host), &rv)

	if err != nil {
		return nil, err
	}

	return &rv, nil
}

func (h *HuaweiAPI) DeviceSignal() (*SignalInfo, error) {
	var rv SignalInfo
	err := h.httpGetXml(fmt.Sprintf("http://%s/api/device/signal", h.host), &rv)

	if err != nil {
		return nil, err
	}

	return &rv, nil
}

func (h *HuaweiAPI) NetNetMode() (*NetNetMode, error) {
	var rv NetNetMode
	err := h.httpGetXml(fmt.Sprintf("http://%s/api/net/net-mode", h.host), &rv)

	if err != nil {
		return nil, err
	}

	return &rv, nil

}
