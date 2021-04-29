package api

import (
	"GiveMeMail/src/deal"
	"GiveMeMail/src/global"
	"fmt"
	"github.com/antchfx/htmlquery"
	"net/http"
	"net/url"
	"strconv"
)


type GoogleApi struct {
	Name string
}

func NewGoogeApi() *GoogleApi{
	return &GoogleApi{"Google"}
}

//return name of this api
func (api GoogleApi) GetName() string{
	return api.Name
}

//main func of this api
func (api GoogleApi) Run() []string{
	Method:=global.Method

	ParseMap:=getParseMapGoogleApi()
	MailResult:= deal.MatchMail(ParseMap)
	if Method=="deep"{
		fmt.Printf("find [%d] urls|",len(ParseMap))

		DeepParseMap:=DeepSearchApi(ParseMap)

		MailResultTemp:=deal.MatchMail(DeepParseMap)
		for url,mail := range(MailResultTemp){
			MailResult[url]=mail
		}
		fmt.Printf("DeepSearchApi find [%d] emails|",len(MailResultTemp))
	}

	return MailResult
}

//Send a request to a page of search results and parse the url and content block, and save the two into a map object and return
//对搜索结果的某页发送请求并解析url与内容块，并将这两者存入一个map对象并返回
//params:  domain:the email domain   |    page:the page of Google Search Result
//return value:  Result map and stopflag
func reqAndParseGoogleApi(page int) (map[string]string,bool){
	Domain:=global.Domain

	//init two return value 初始化两个返回值
	ParseResult :=make(map[string]string)
	FlagStop := false

	//HTTP Request
	client:= deal.ProxyClient()

	//define the params of http request  定义传参参数
	Params:=url.Values{}
	Params.Set("q","\"邮箱\" | \"邮件\" | \"mail\" and site:"+Domain)
	Params.Set("num","100") //"i" means the Number of pages of the google request answer
	Params.Set("start",strconv.Itoa(page*100))

	//parse the url of https://www.google.com/search  解析url
	Url,err:=url.Parse("https://www.google.com/search")
	if err!=nil{
		panic("parse url error!")
	}

	//encode the params to url_encode  把参数进行url编码
	Url.RawQuery = Params.Encode()

	//make http request,and return the response  发起请求并返回结果
	req,_:=http.NewRequest("GET",Url.String(),nil)
	req.Header.Add("User-agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.190 Safari/537.36")

	//request repeatedly
	resp:=deal.RequestMore(req,client,5)
	if resp==nil{
		return ParseResult,true
	}

	//Parse Xpath
	//define two Xpath,first is the root node of a SearchResult Div,The second is to test whether google has an xpath that returns results
	//定义两个Xpath，第一个是一个搜索结果的的根节点，第二个是测试google是否具有返回结果的xpath
	XpathSample := "//div[@class='g'][%d]/div/div"
	XpathFlag:="//div[@class='tF2Cxc']"

	//parse the response
	//解析响应
	nodes, err := htmlquery.Parse(resp.Body)

	//Determine whether google returns search results 判断google是否返回搜索结果
	if htmlquery.Find(nodes,XpathFlag)==nil{
		FlagStop=true
	}else{
		FlagStop=false
	}

	if err != nil {
		panic("error in parse response(Parse)")
	}

	//Loop traversal every search answer Result
	//循环遍历每个搜索结果块
	for j := 1; j < 100; j++ {

		//Format the formatted Xpath just defined, and find the xpath in html
		//格式化刚才定义的格式化Xpath，并在html中找到该xpath
		XpathDiv := fmt.Sprintf(XpathSample, j)
		DivInfo := htmlquery.Find(nodes, XpathDiv)

		//Loop traversal the div to get url,title and content
		//循环遍历div以获取url，标题和内容
		for _,n := range DivInfo {
			//get the infomation of url block
			//获取URL块的信息
			TitleDivInfo := htmlquery.FindOne(n, "//div[@class='yuRUbf']/a")
			if TitleDivInfo == nil {
				break
			}

			TitleInfoUrl := htmlquery.SelectAttr(TitleDivInfo, "href")

			//Get information about the content block. The reason why there is if-else is because Google returns the content block of search results in two formats
			//获取内容块的信息.为什么会有一个if-else，这是因为google的内容块有两种格式
			ContentInfoArray := htmlquery.FindOne(n, "//div[@class='IsZvec']/span[@class='aCOpRe']/span[2]")
			ContentInfo := ""
			if ContentInfoArray == nil {
				ContentInfoTemp := htmlquery.Find(n, "//div[@class='IsZvec']/span[@class='aCOpRe']/span")
				ContentInfo = htmlquery.InnerText(ContentInfoTemp[0])
			} else {
				ContentInfoTemp := htmlquery.Find(n, "//div[@class='IsZvec']/span[@class='aCOpRe']/span[2]")
				ContentInfo = htmlquery.InnerText(ContentInfoTemp[0])
			}
			//Load the result into the map 将结果装载进map
			//fmt.Print(ContentInfo+"\n"+"|"+strconv.Itoa(page))
			ParseResult [TitleInfoUrl]=ContentInfo
		}
	}
	return ParseResult,FlagStop
}


//Summarize the analysis results of each page 将每一页的解析结果进行汇总
//params: domain:the email domain
//return value:  completed Result map
func getParseMapGoogleApi() map[string]string{
	ParseMap :=make(map[string]string)
	for i:=0;;i++{
		ParseMapTemp,StopFlag := reqAndParseGoogleApi(i)
		if StopFlag{
			break
		}else{
			for url,content := range ParseMapTemp{
				ParseMap[url] = content
			}
		}
	}
	return ParseMap
}