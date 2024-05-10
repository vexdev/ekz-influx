package ekz

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type EkzReader struct {
	client *http.Client
}

func NewEkzReader() (*EkzReader, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Jar: jar,
	}
	return &EkzReader{
		client: client,
	}, nil
}

func (e *EkzReader) Authenticate(username string, password string) error {
	response, err := e.client.Get("https://my.ekz.ch/oauth2/authorization/cos-portal-web-app-client")
	if err != nil {
		return err
	}

	// Read the response body
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("unexpected status code for authorization: %d", response.StatusCode)
	}

	// Search in the html the regex
	regex := `https://login.ekz.ch/auth/realms/myEKZ/login-actions/authenticate\?session_code=(?P<session_code>.+)&amp;execution=(?P<execution>.+)&amp;client_id=(?P<client_id>.+)&amp;tab_id=(?P<tab_id>[^"]+)`
	re := regexp.MustCompile(regex)
	match := re.FindStringSubmatch(string(body))

	// Make a post request to authenticate
	var authUrl = fmt.Sprintf("https://login.ekz.ch/auth/realms/myEKZ/login-actions/authenticate?session_code=%s&execution=%s&client_id=%s&tab_id=%s", match[1], match[2], match[3], match[4])

	data := url.Values{}
	data.Set("username", username)
	data.Set("password", password)
	data.Set("credentialId", "")
	encodedData := data.Encode()
	response, err = e.client.Post(authUrl, "application/x-www-form-urlencoded", strings.NewReader(encodedData))
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("unexpected status code for authentication: %d", response.StatusCode)
	}

	return nil
}

func (e *EkzReader) GetConsumptionData(installationId string, from time.Time, to time.Time) (*EkzData, error) {
	consumptionUrl := consumptionDataUrl(installationId, from, to)
	req, err := e.client.Get(consumptionUrl)
	if err != nil {
		return nil, err
	}
	if req.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code for consumption data: %d", req.StatusCode)
	}

	defer req.Body.Close()
	bodyString, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	ekzData, err := EkzDataFromJson(bodyString)
	return &ekzData, err
}

func consumptionDataUrl(installationId string, from time.Time, to time.Time) string {
	return fmt.Sprintf("https://my.ekz.ch/api/portal-services/consumption-view/v1/consumption-data?installationId=%s&from=%s&to=%s&type=PK_VERB_15MIN", installationId, from.Format("2006-01-02"), to.Format("2006-01-02"))
}
