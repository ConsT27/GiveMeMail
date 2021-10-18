package main

import (
	"GiveMeMail/src/api"
	"GiveMeMail/src/global"
	"flag"
	"fmt"
	"os"
	"strconv"
)


func main(){
	//set the params from console
	DomainTmp:=flag.String("d","","the domain")
	ProxyUrlTmp:=flag.String("p","none","the Url of Proxy")
	MethodTmp:=flag.String("method","default","the Url of Proxy")
	MatchRuleTmp:=flag.String("r","all","the match rule")
	flag.Parse()

	global.MatchRule=*MatchRuleTmp
	global.ProxyUrl=*ProxyUrlTmp
	global.Domain=*DomainTmp
	global.Method=*MethodTmp

	Domain:=global.Domain
	ProxyUrl:=global.ProxyUrl
	Method:=global.Method
	MatchRule:=global.MatchRule

	//Output the Params

	/**fmt.Print (" ██████╗ ██╗██╗   ██╗███████╗███╗   ███╗███████╗███╗   ███╗ █████╗ ██╗██╗     \n")
	fmt.Print ("██╔════╝ ██║██║   ██║██╔════╝████╗ ████║██╔════╝████╗ ████║██╔══██╗██║██║     \n")
	fmt.Print ("██║  ███╗██║██║   ██║█████╗  ██╔████╔██║█████╗  ██╔████╔██║███████║██║██║     \n")
	fmt.Print ("██║   ██║██║╚██╗ ██╔╝██╔══╝  ██║╚██╔╝██║██╔══╝  ██║╚██╔╝██║██╔══██║██║██║     \n")
	fmt.Print ("╚██████╔╝██║ ╚████╔╝ ███████╗██║ ╚═╝ ██║███████╗██║ ╚═╝ ██║██║  ██║██║███████╗\n")
	fmt.Print (" ╚═════╝ ╚═╝  ╚═══╝  ╚══════╝╚═╝     ╚═╝╚══════╝╚═╝     ╚═╝╚═╝  ╚═╝╚═╝╚══════╝\n")
	fmt.Print ("                                                                              \n")**/


	fmt.Printf("\n")
	fmt.Print("  ▄████   ██▓  ██▒    ██▓ ▓██████    ███▄ ▄███▓▓██████   ███▄ ▄███▓    ██     ██▓ ██▓    \n")
	fmt.Print(" ██▒ ▀█▒ ▓██▒ ▓██░    █▒  ▓█   ▀    ▓██▒▀█▀ ██▒▓█   ▀   ▓██▒▀█▀ ██▒▒ ██  █▄  ▓██▒▓██▒    \n")
	fmt.Print("▒██░▄▄▄░ ▒██▒  ▓██   █▒░  ▒██████   ▓██    ▓██░▒██████  ▓██    ▓██░▒██▄▄▄▄█▄ ▒██▒▒██░    \n")
	fmt.Print("░▓█  ██▓ ░██░   ▒██ ██░░ ▒▓█  ▄     ▒██    ▒██ ▓█  ▄    ▒██    ▒██ ░██    ██ ░██░▒██░    \n")
	fmt.Print("░▒▓███▀▒ ░██░    ▒▀█     ░▒████▒█   ▒██▒   ░██▒▒████▒█  ▒██▒   ░██▒ ▓█   ▓██▒░██░░██████▒\n")
	fmt.Print(" ░▒   ▒  ░▓      ░ ▐░   ░░ ▒░ ░ ░ ▒░  ░░░ ▒░ ░  ░ ▒░   ░  ░ ▒▒   ▓▒█░░▓  ░ ▒░▓  ░\n")
	fmt.Print("  ░   ░   ▒ ░    ░ ░░    ░ ░  ░ ░     ░ ░ ░  ░  ░  ░      ░  ▒   ▒▒ ░ ▒ ░░ ░ ▒  ░\n")
	fmt.Print("░ ░   ░   ▒ ░      ░░      ░    ░  ░      ░     ░      ░     ░   ▒    ▒ ░  ░ ░   \n")
	fmt.Print("      ░   ░         ░      ░  ░    ░      ░  ░         ░         ░  ░ ░      ░  ░\n")
	fmt.Print("                    ░                                                                 \n")

	fmt.Print("V1.1\n")
	if Domain==""{
		fmt.Print("Error:domain Empty!")
		os.Exit(0)
	}

	fmt.Printf("[+]Domain:%s\n",Domain)
	fmt.Printf("[+]Method:%s\n",Method)
	fmt.Printf("[+]MatchRule:%s\n",MatchRule)
	fmt.Printf("[+]Proxy:%s\n\n",ProxyUrl)

	//init the Result map
	MailResult := make([]string,0)

	Api:=api.NewApiManager(Domain,ProxyUrl,Method)
	MailResult=Api.Run()

	fmt.Print("\nResult:\n")
	for _,mail:=range MailResult{
		fmt.Print(mail+"\n")
	}

	fmt.Print("[+]Find "+strconv.Itoa(len(MailResult))+" emails")
}