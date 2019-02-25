package gotwilio

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type AccountRecord struct {
	Status          string `json:"status"`
	DateUpdated     string `json:"date_updated"`
	AuthToken       string `json:"auth_token"`
	FriendlyName    string `json:"friendly_name"`
	OwnerAccountSid string `json:"owner_account_sid"`
	URI             string `json:"uri"`
	Sid             string `json:"sid"`
	DateCreated     string `json:"date_created"`
	Type            string `json:"type"`
	SubresourceUris struct {
		Addresses             string `json:"addresses"`
		Conferences           string `json:"conferences"`
		SigningKeys           string `json:"signing_keys"`
		Transcriptions        string `json:"transcriptions"`
		ConnectApps           string `json:"connect_apps"`
		Sip                   string `json:"sip"`
		AuthorizedConnectApps string `json:"authorized_connect_apps"`
		Usage                 string `json:"usage"`
		Keys                  string `json:"keys"`
		Applications          string `json:"applications"`
		Recordings            string `json:"recordings"`
		ShortCodes            string `json:"short_codes"`
		Calls                 string `json:"calls"`
		Notifications         string `json:"notifications"`
		IncomingPhoneNumbers  string `json:"incoming_phone_numbers"`
		Queues                string `json:"queues"`
		Messages              string `json:"messages"`
		OutgoingCallerIds     string `json:"outgoing_caller_ids"`
		AvailablePhoneNumbers string `json:"available_phone_numbers"`
		Balance               string `json:"balance"`
	} `json:"subresource_uris"`
}

// UsageResponse contains information about account usage.
type AccountsResponse struct {
	PageSize        int             `json:"page_size"`
	Page            int             `json:"page"`
	Accounts        []AccountRecord `json:"accounts"`
	Start           int             `json:"start"`
	End             int             `json:"end"`
	NextPageURI     string          `json:"next_page_uri"`
	PreviousPageURI string          `json:"previous_page_uri"`
	FirstPageURI    string          `json:"first_page_uri"`
	URI             string          `json:"uri"`
}

func (twilio *Twilio) GetAccounts(pagesize, page int) (*AccountsResponse,
	*Exception, error) {
	formValues := url.Values{}

	formValues.Set("PageSize", strconv.Itoa(pagesize))
	formValues.Set("Page", strconv.Itoa(page))

	var accountsResponse *AccountsResponse
	var exception *Exception
	twilioUrl := twilio.BaseUrl + "/Accounts.json"

	res, err := twilio.get(twilioUrl + "?" + formValues.Encode())
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	if res.StatusCode != http.StatusOK {
		exception = new(Exception)
		err = json.Unmarshal(responseBody, exception)
		return nil, exception, err
	}

	accountsResponse = new(AccountsResponse)
	err = json.Unmarshal(responseBody, accountsResponse)
	return accountsResponse, nil, err
}

type CallLogsResponse struct {
	PageSize        int       `json:"page_size"`
	Page            int       `json:"page"`
	Start           int       `json:"start"`
	Calls           []CallLog `json:"calls"`
	End             int       `json:"end"`
	PreviousPageURI string    `json:"previous_page_uri"`
	NextPageURI     string    `json:"next_page_uri"`
	FirstPageURI    string    `json:"first_page_uri"`
	URI             string    `json:"uri"`
}
type CallLog struct {
	Sid             string      `json:"sid"`
	DateCreated     string      `json:"date_created"`
	DateUpdated     string      `json:"date_updated"`
	ParentCallSid   interface{} `json:"parent_call_sid"`
	AccountSid      string      `json:"account_sid"`
	To              string      `json:"to"`
	ToFormatted     string      `json:"to_formatted"`
	From            string      `json:"from"`
	FromFormatted   string      `json:"from_formatted"`
	PhoneNumberSid  string      `json:"phone_number_sid"`
	Status          string      `json:"status"`
	StartTime       string      `json:"start_time"`
	EndTime         string      `json:"end_time"`
	Duration        string      `json:"duration"`
	Price           string      `json:"price"`
	PriceUnit       string      `json:"price_unit"`
	Direction       string      `json:"direction"`
	AnsweredBy      interface{} `json:"answered_by"`
	Annotation      interface{} `json:"annotation"`
	APIVersion      string      `json:"api_version"`
	ForwardedFrom   string      `json:"forwarded_from"`
	GroupSid        interface{} `json:"group_sid"`
	CallerName      interface{} `json:"caller_name"`
	URI             string      `json:"uri"`
	SubresourceUris struct {
		Notifications string `json:"notifications"`
		Recordings    string `json:"recordings"`
	} `json:"subresource_uris"`
}

func (twilio *Twilio) GetCallUsage(pagesize, page int, uri string) (*CallLogsResponse,
	*Exception, error) {
	formValues := url.Values{}

	formValues.Set("PageSize", strconv.Itoa(pagesize))
	formValues.Set("Page", strconv.Itoa(page))

	var calllogsResponse *CallLogsResponse
	var exception *Exception
	twilioUrl := twilio.BaseUrl + uri

	res, err := twilio.get(twilioUrl)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	if res.StatusCode != http.StatusOK {
		exception = new(Exception)
		err = json.Unmarshal(responseBody, exception)
		return nil, exception, err
	}

	calllogsResponse = new(CallLogsResponse)
	err = json.Unmarshal(responseBody, calllogsResponse)
	return calllogsResponse, nil, err
}
