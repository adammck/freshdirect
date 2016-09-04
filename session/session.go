package session

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Session struct {
	ident string
}

func New(ident string) *Session {
	return &Session{
		ident: ident,
	}
}

func Login(username, password string) (*Session, error) {
	panic("not implemented: Login")
}

func (s *Session) Post(path string, data interface{}) error {

	buf, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	values := url.Values{}
	values.Set("data", string(buf))

	client := &http.Client{}
	u := fmt.Sprintf("https://www.freshdirect.com/api/%s", path)
	req, err := http.NewRequest("POST", u, strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}

	// There are a shitload of cookies used on the site, but only the session ID
	// is required. Even the user ID is optional.

	req.AddCookie(&http.Cookie{
		Name:  "JSESSIONID",
		Value: s.ident,
	})

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	//fmt.Println("req:")
	//spew.Dump(req)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	//fmt.Println("res:")
	//resbuf, _ := ioutil.ReadAll(res.Body)
	//spew.Dump(resbuf)

	if res.StatusCode != 200 {
		return fmt.Errorf("api error: %d", res.StatusCode)
	}

	return nil
}
