# GiveMeMail
```
  ▄████   ██▓  ██▒    ██▓ ▓██████    ███▄ ▄███▓▓██████   ███▄ ▄███▓    ██     ██▓ ██▓    
 ██▒ ▀█▒ ▓██▒ ▓██░    █▒  ▓█   ▀    ▓██▒▀█▀ ██▒▓█   ▀   ▓██▒▀█▀ ██▒▒ ██  █▄  ▓██▒▓██▒    
▒██░▄▄▄░ ▒██▒  ▓██   █▒░  ▒██████   ▓██    ▓██░▒██████  ▓██    ▓██░▒██▄▄▄▄█▄ ▒██▒▒██░    
░▓█  ██▓ ░██░   ▒██ ██░░ ▒▓█  ▄     ▒██    ▒██ ▓█  ▄    ▒██    ▒██ ░██    ██ ░██░▒██░    
░▒▓███▀▒ ░██░    ▒▀█     ░▒████▒█   ▒██▒   ░██▒▒████▒█  ▒██▒   ░██▒ ▓█   ▓██▒░██░░██████▒
 ░▒   ▒  ░▓      ░ ▐░   ░░ ▒░ ░ ░ ▒░  ░░░ ▒░ ░  ░ ▒░   ░  ░ ▒▒   ▓▒█░░▓  ░ ▒░▓  ░
  ░   ░   ▒ ░    ░ ░░    ░ ░  ░ ░     ░ ░ ░  ░  ░  ░      ░  ▒   ▒▒ ░ ▒ ░░ ░ ▒  ░
░ ░   ░   ▒ ░      ░░      ░    ░  ░      ░     ░      ░     ░   ▒    ▒ ░  ░ ░   
      ░   ░         ░      ░  ░    ░      ░  ░         ░         ░  ░ ░      ░  ░
                    ░                                                                
                    
 ```    
                    
![](https://const27blog.oss-cn-beijing.aliyuncs.com/img/QQ图片20210429140606.png)
## 简介 description
 这个工具用于收集邮箱，目前提供了四个api：Google，Skymem，Email-format 和深度搜索DeepSearchApi

深度搜索引擎会跟进每一个Google返回的url，并在其页面中查找匹配邮箱,有协程，所以非常快


English:
This tool is used to collect mailboxes and currently provides four apis: Google, Skymem, Email-format and DeepSearchApi

The deep search engine will follow up every url returned by Google and find a matching email address on its page(very fast)
 ## 参数 params
 ```
 -d Email Domain
-r MatchRule<all(default),strict>  all:Match all mailboxes，strict:Only match the mailbox specified by the domain
-p ProxyUrl
-method the method of finding emails<default,deep> deep:will use DeepSerchApi
 ```
 ## 快速上手  Get started quickly
 ```
 GiveMeMail.exe -d xxx.com
 GiveMeMail.exe -d xxx.com -p http://127.0.0.1:8080
 GiveMeMail.exe -d xxx.com -r strict -method deep
 ```
 贡献者:	Rookie-is
