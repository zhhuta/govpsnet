//connection.go
package govpsnet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Params map[string]interface{}

type Client struct {
	Email   string
	Api_Key string
	URL     string
}

func NewConnection(email, api_key string) (*Client, error) {
	cl := &Client{
		Email:   email,
		Api_Key: api_key,
		URL:     "https://api.vps.net",
	}
	return cl, nil
}

func (c *Client) get(route string, v interface{}) error {
	r_route := fmt.Sprintf("%s/%s", c.URL, route)
	//fmt.Println(r_route)
	request, err := http.NewRequest("GET", r_route, nil)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	err = c.PerformReqest(request, v)
	if err != nil {
		return err
	}
	return nil
}
func (c *Client) post(route string, op interface{}, v interface{}) error {
	r_route := fmt.Sprintf("%s/%s", c.URL, route)
	fmt.Sprintf("%#v,%#v", op, v)
	fmt.Println(r_route)

	if op != nil {

		m_json, err := json.Marshal(op)
		if err != nil {
			return err
		}
		request, err := http.NewRequest("POST", r_route, bytes.NewReader(m_json))
		if err != nil {
			return err
		}
		request.Header.Set("Content-Type", "application/json")
		err = c.PerformReqest(request, v)
		if err != nil {
			return err
		}
	} else { //don't sending data over POST method

		request, err := http.NewRequest("POST", r_route, nil)
		if err != nil {
			return err
		}
		request.Header.Set("Content-Type", "application/json")
		err = c.PerformReqest(request, v)
		if err != nil {
			return err
		}
	}
	return nil
}
func (c *Client) put(route string, op, v interface{}) error {
	r_route := fmt.Sprintf("%s/%s", c.URL, route)
	m_json, err := json.Marshal(op)
	if err != nil {
		return err
	}
	request, err := http.NewRequest("PUT", r_route, bytes.NewReader(m_json))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	err = c.PerformReqest(request, v)
	if err != nil {
		return err
	}
	return nil
}
func (c *Client) delete(route string) error {
	r_route := fmt.Sprintf("%s/%s", c.URL, route)

	request, err := http.NewRequest("DELETE", r_route, nil)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	err = c.PerformReqest(request, nil)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) PerformReqest(req *http.Request, v interface{}) error {
	conn := &http.Client{}
	req.SetBasicAuth(c.Email, c.Api_Key)
	resp, err := conn.Do(req)
	if err != nil {
		return fmt.Errorf("govpsnet: Error during performing request: %s", err)
	}
	err = decode_resons(resp, v)
	if err != nil {
		return err
	}
	return nil
}

func decode_resons(resp *http.Response, v interface{}) error {
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return fmt.Errorf("Something Went Wrong: %#s:%#s", resp.StatusCode, resp.Body)
	}
	if v != nil {
		//json2unmashal, _ := ioutil.ReadAll(resp.Body)
		//resp.Body.Close()
		//err := json.Unmarshal(json2unmashal, &v)
		err := json.NewDecoder(resp.Body).Decode(&v)
		if err != nil {
			return fmt.Errorf("govpsnet: Error during decoding json into struct: %s", err)
		}
	}
	return nil
}
