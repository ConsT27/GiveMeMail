package api

import (
	"GiveMeMail/src/deal"
	"fmt"
	"strconv"
)

// this is a struct to call the apis
type ApiManager struct {
}

func NewApiManager(domain string,ProxyUrl string,Method string) *ApiManager{
	return &ApiManager{}
}

func (api ApiManager)Run() []string{
	//init the MailResult map
	MailResult := make([]string,0)

	//Execute the apis and get the data
	ApiList := []ApiInterface{NewGoogeApi(),NewSkymemApi(),NewEmailFormatApi()}

	for _,apis:= range ApiList{
		fmt.Print("[*]begin to start api:"+apis.GetName()+"... | ")
		MailResultTemp:=apis.Run()
		MailResult=append(MailResult,MailResultTemp...)
		ResultSize:=strconv.Itoa(len(MailResultTemp))
		if ResultSize=="0"{
			fmt.Print("Nothing found. May be you should consider about Proxy or Verification code,Please try again\n")
		}else{fmt.Print(apis.GetName()+" find "+ResultSize+" emails\n")}
	}

	//data dereplication
	//数据去重
	MailResult=deal.Dereplication(MailResult)
	return MailResult
}