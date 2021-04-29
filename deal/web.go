package deal

import (
	"GiveMeMail/src/global"
	"net/http"
	"net/url"
	"time"
)

//Used to set up proxy
//params: ProxyUrl:the url of Proxy
//return Value: http.Client Object
func ProxyClient() http.Client{
	client:=http.Client{Timeout:30*time.Second}
	if global.ProxyUrl!="none" {
		Proxy:= func(_ *http.Request) (*url.URL, error) {
			return url.Parse(global.ProxyUrl)
		}
		transport := &http.Transport{Proxy: Proxy}
		client=http.Client{Transport: transport,Timeout:30*time.Second}
	}
	return client
}

//request repeatedly
//params: req:just a *http.Request  |client: the Client Object to make request|num retry frequency
//return value: the http.Response Object or panic
func RequestMore(req *http.Request,client http.Client,num int) *http.Response{
	var resp *http.Response
	var err error

	for i:=0;i<num;i++{
		if i==4{
			resp,err = client.Do(req)
			if err!=nil{
				return nil
			}
		}else{
			resp, err = client.Do(req)
			if err == nil {
				break
			}
		}
	}

	return resp
}