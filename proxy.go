package proxy

import (
	"crypto/tls"
	"errors"
	"github.com/Congenital/log/v0.2/log"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func HttpGet(Url string, param string, proxy_addr string) ([]byte, error) {
	proxy, err := url.Parse(proxy_addr)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			DisableCompression: true,
			Proxy:              http.ProxyURL(proxy),
		},
	}

	req, err := http.NewRequest("GET", Url+"?"+param, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:37.0) Gecko/20100101 Firefox/37.0")
	req.Header.Add("Accept", "application/x-www-form-urlencoded")
	req.Header.Set("Host", "www.dianhua.cn")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("Err - " + resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return body, nil
}

func HttpPost(Url string, param string, proxy_addr string) ([]byte, error) {
	proxy, err := url.Parse(proxy_addr)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:    &tls.Config{},
			DisableCompression: true,
			Proxy:              http.ProxyURL(proxy),
		},
	}

	req, err := http.NewRequest("POST", Url+"?"+param, strings.NewReader(param))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Fedora; Linux x86_64; rv:37.0) Gecko/20100101 Firefox/37.0")
	req.Header.Add("Accept", "application/x-www-form-urlencoded")
	req.Header.Set("Host", "www.dianhua.cn")
	req.Header.Set("Connection", "keep-alive")

	resp, err := client.Do(req)

	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("Err - " + resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return body, nil
}
