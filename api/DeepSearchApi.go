package api

import (
	"GiveMeMail/src/deal"
	"GiveMeMail/src/global"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

//DeepSearchApi Follow up the url returned by google and search for the mailbox in-depth on the page
//跟踪google返回的网址，并在页面上深入搜索邮箱
func DeepSearchApi(parsemap map[string]string) map[string]string{
	ProxyUrl:=global.ProxyUrl
	deepParseMap:=make(chan map[string]string,len(parsemap))
	parseurl:= make(chan string,len(parsemap))

	//协程计数器
	//Goroutine counter
	wg := sync.WaitGroup{}
	for one,_:=range parsemap{
		//将url需要访问的url放入通道中
		//Put the url to be visited into the channel
		parseurl <- one
	}

	wg.Add(len(parsemap)/3)
	for i := 0; i < len(parsemap)/3; i++ {
		go worker(
			parseurl,
			ProxyUrl,
			deepParseMap,
			&wg,
		)
	}

	wg.Wait()
	res  := make(map[string]string)
	//把通道中的数据转存到要返回的map中
	//Dump the data in the channel to the map to be returned
	for ;;{
		select {
		case one :=<-deepParseMap:
			for oneUrl,content := range one {
				res[oneUrl]=content
			}
		default:
			return res
		}
	}

}

// RequsetAndParse Send request and return html
func RequsetAndParse(req *http.Request,deepMap chan map[string]string,client http.Client,url string){
	resp,err:=client.Do(req)
	//Put the returned body of each URL visit into the channel
	//把每一个url访问的返回体放入通道中
	if err==nil{
		body,_:=ioutil.ReadAll(resp.Body)
		m := map[string]string{url:string(body)}
		deepMap <- m
	}else{
		m := map[string]string{url:"NONE"}
		deepMap <- m
	}
}
//worker Represents a coroutine, constantly getting data from the channel to execute the RequsetAndParse function
func worker(parseurl chan string,ProxyUrl string,deepParseMap chan map[string]string,wg *sync.WaitGroup){
	for i := 1; i >0 ; i++ {
		select {
		//从通道中读取一个url用于访问
		//Read a url from the channel for Request
		case OneUrl :=<-parseurl:
			client:=deal.ProxyClient()
			urlRequest,_:=url.Parse(OneUrl)
			req,err:=http.NewRequest("GET",urlRequest.String(),nil)
			if err!=nil{panic("error")}
			req.Header.Add("User-agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36")
			RequsetAndParse(req,deepParseMap,client,OneUrl)
		default:
			wg.Done()
			return
		}
	}
}