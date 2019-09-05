package hac

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

const (
	HACLogin       string = "https://homeaccess.katyisd.org/HomeAccess/Account/LogOn?ReturnUrl=%2fHomeAccess%2f"
	HACAssignments string = "https://homeaccess.katyisd.org/HomeAccess/Content/Student/Assignments.aspx"
)

type Class struct {
	ClassName string
	ClassAvg  uint8
}

type HAC struct {
	UserName string
	Password string
	Database string

	client *http.Client
}

func NewHAC(UserName, Password, Database string) (*HAC, error) {
	form := url.Values{
		"Database":              {Database},
		"LogOnDetails.UserName": {UserName},
		"LogOnDetails.Password": {Password},
	}

	cookie, _ := cookiejar.New(nil)
	client := &http.Client{Jar: cookie}
	response, err := client.PostForm(HACLogin, form)
	if err != nil || response.StatusCode != 200 {
		return nil, errors.New("Request failed")
	}
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	regex := regexp.MustCompile(`unsuccessful`)
	if len(regex.FindAllString(string(body), -1)) != 0 {
		return nil, errors.New("Unsuccessful login")
	}

	return &HAC{UserName, Password, Database, client}, nil
}

func (hac *HAC) Classes() ([]string, error) {
	response, err := hac.client.Get(HACAssignments)
	if err != nil || response.StatusCode != 200 {
		return nil, errors.New("Request failed")
	}
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	regex := regexp.MustCompile(`\d{4}[A-Z]{0,3}\s-\s[^<]+\n`)
	retStrings := regex.FindAllString(string(body), -1)

	for i, s := range retStrings {
		retStrings[i] = s[:len(s)-2]
	}
	return retStrings, nil
}

func (hac *HAC) Grades() ([]Class, error) {
	response, err := hac.client.Get(HACAssignments)
	if err != nil || response.StatusCode != 200 {
		return nil, errors.New("Request failed")
	}
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	regex := regexp.MustCompile(`Classwork Average \d+`)
	grades := regex.FindAllString(string(body), -1)

	regex = regexp.MustCompile(`\d{4}[A-Z]{0,3}\s-\s[^<]+\n`)
	classes := regex.FindAllString(string(body), -1)

	if len(grades) != len(classes) {
		return nil, errors.New("Incorrect Regex Pattern")
	}

	var ret []Class
	for i := 0; i < len(grades); i++ {
		avg, _ := strconv.Atoi(strings.Fields(grades[i])[2])
		ret = append(ret, Class{classes[i][:len(classes[i])-2], uint8(avg)})
	}
	return ret, nil
}
