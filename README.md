## domain-stress
English | [简体中文](./README-zh_CN.md)
### Overview 
domain-stress is a domain name-based stress test tool that you can use to stress test the service corresponding to a domain name. Specify the domain name and the total number of requests to use.  

You can use this tool when you're missing a load balancer for your backend IP pool or you only know the domain name.

At present, HTTP get, POST, PUT, DELETE and other requests are supported, you can use the -n parameter to specify the total number of requests, and the -c parameter to specify the number of concurrency per backend IP.

The currently completed request is output once per second during the pressure generation process, and the full result is output after the test is completed.

### Install
```bash
$ wget https://github.com/flowscolors/domain-stress/releases/tag/0.1.0
```


### Usage
Here are some simple examples of using domain-stress

GET Request
```bash
$ ./domain-stress stress -d ebay.com -c 2 -n 100
向目标域名 ebay.com 以每个后端IP 2 个协程的方式发送 100 个请求。每个请求超时时间为 5000 ms，重试次数为 3 次。

************************  执行计划 stat  ***************************
协程id 1  ，计划往目标 URL http://216.113.181.253/ ，发送 GET 请求数 8
协程id 2  ，计划往目标 URL http://216.113.181.253/ ，发送 GET 请求数 8
协程id 3  ，计划往目标 URL http://64.4.253.77/ ，发送 GET 请求数 8
协程id 4  ，计划往目标 URL http://64.4.253.77/ ，发送 GET 请求数 8
协程id 5  ，计划往目标 URL http://216.113.179.53/ ，发送 GET 请求数 8
协程id 6  ，计划往目标 URL http://216.113.179.53/ ，发送 GET 请求数 8
协程id 7  ，计划往目标 URL http://66.211.175.229/ ，发送 GET 请求数 8
协程id 8  ，计划往目标 URL http://66.211.175.229/ ，发送 GET 请求数 8
协程id 9  ，计划往目标 URL http://66.211.172.37/ ，发送 GET 请求数 8
协程id 10  ，计划往目标 URL http://66.211.172.37/ ，发送 GET 请求数 8
协程id 11  ，计划往目标 URL http://209.140.148.143/ ，发送 GET 请求数 8
协程id 12  ，计划往目标 URL http://209.140.148.143/ ，发送 GET 请求数 12
************************  执行计划 end  ***************************


─────┬───────┬───────┬───────┬────────┬────────┬────────┬────────┬────────┬────────┬────────
 耗时│ 并发数│ 成功数│ 失败数│   qps  │最长耗时│最短耗时│平均耗时│下载字节│字节每秒│ 返回码
─────┼───────┼───────┼───────┼────────┼────────┼────────┼────────┼────────┼────────┼────────
   1s│      0│      0│      0│    0.00│    0.00│    0.00│    0.00│        │        │
   2s│      8│      8│      0│    8.24│ 1612.91│ 1313.37│ 1456.07│        │        │200:8
   3s│     12│     16│      0│    7.44│ 2731.86│ 1170.14│ 1612.51│        │        │200:16
   4s│     12│     22│      0│    7.38│ 2731.86│ 1150.10│ 1624.94│        │        │200:22
   5s│     12│     27│      0│    7.07│ 2731.86│ 1150.10│ 1697.17│        │        │200:27
   6s│     12│     35│      0│    7.02│ 2747.76│ 1150.10│ 1708.65│        │        │200:35
   7s│     12│     41│      0│    6.96│ 2747.76│ 1098.07│ 1723.95│        │        │200:41
   8s│     12│     49│      0│    6.88│ 2747.76│ 1098.07│ 1743.67│        │        │200:49
   9s│     12│     55│      0│    6.98│ 2957.22│ 1098.07│ 1718.54│        │        │200:55
  10s│     12│     61│      0│    7.03│ 2957.22│ 1098.07│ 1707.13│        │        │200:61
第  1  次请求失败: Get "http://209.140.148.143/": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
  11s│     12│     69│      0│    7.00│ 2957.22│ 1098.07│ 1714.03│        │        │200:69
  12s│     12│     74│      0│    6.96│ 2978.68│ 1098.07│ 1723.08│        │        │200:74
  13s│     12│     81│      0│    6.74│ 7007.09│ 1098.07│ 1781.57│        │        │200:81
  14s│     12│     84│      0│    6.75│ 7007.09│ 1098.07│ 1778.41│        │        │200:84
  15s│     12│     87│      0│    6.67│ 7007.09│ 1098.07│ 1797.78│        │        │200:87
  16s│     12│     89│      0│    6.61│ 7007.09│ 1098.07│ 1814.75│        │        │200:89
  17s│     12│     91│      0│    6.60│ 7007.09│ 1098.07│ 1818.82│        │        │200:91
  18s│     12│     93│      0│    6.54│ 7007.09│ 1098.07│ 1834.62│        │        │200:93
  19s│     12│     94│      0│    6.54│ 7007.09│ 1098.07│ 1835.52│        │        │200:94
  20s│     12│     94│      0│    6.54│ 7007.09│ 1098.07│ 1835.52│        │        │200:94
  21s│     12│     96│      0│    6.50│ 7007.09│ 1098.07│ 1846.97│        │        │200:96
  22s│     12│     96│      0│    6.50│ 7007.09│ 1098.07│ 1846.97│        │        │200:96
  23s│     12│     96│      0│    6.50│ 7007.09│ 1098.07│ 1846.97│        │        │200:96
  24s│     12│     97│      0│    6.45│ 7007.09│ 1098.07│ 1861.15│        │        │200:97
  25s│     12│     97│      0│    6.45│ 7007.09│ 1098.07│ 1861.15│        │        │200:97
  26s│     12│     97│      0│    6.45│ 7007.09│ 1098.07│ 1861.15│        │        │200:97
  27s│     12│     97│      0│    6.45│ 7007.09│ 1098.07│ 1861.15│        │        │200:97
  28s│     12│     97│      0│    6.45│ 7007.09│ 1098.07│ 1861.15│        │        │200:97
第  1  次请求失败: Get "https://www.ebay.com/": context deadline exceeded (Client.Timeout exceeded while awaiting headers)
  29s│     12│     97│      0│    6.45│ 7007.09│ 1098.07│ 1861.15│        │        │200:97
  30s│     12│     97│      0│    6.45│ 7007.09│ 1098.07│ 1861.15│        │        │200:97
  31s│     12│     97│      0│    6.45│ 7007.09│ 1098.07│ 1861.15│        │        │200:97
  32s│     12│     97│      0│    6.45│ 7007.09│ 1098.07│ 1861.15│        │        │200:97
  33s│     12│     98│      0│    6.23│ 8241.65│ 1098.07│ 1926.26│        │        │200:98
  34s│     12│     98│      0│    6.23│ 8241.65│ 1098.07│ 1926.26│        │        │200:98
  35s│     12│     99│      0│    6.23│ 8241.65│ 1098.07│ 1927.45│        │        │200:99
  36s│     12│     99│      0│    6.23│ 8241.65│ 1098.07│ 1927.45│        │        │200:99
  37s│     12│     99│      0│    6.23│ 8241.65│ 1098.07│ 1927.45│        │        │200:99
  37s│     12│    100│      0│    6.20│ 8241.65│ 1098.07│ 1935.77│        │        │200:100


*************************  结果 stat  ****************************
处理协程数量: 12
请求总数: 100
总请求时间: 37.009 秒
成功请求数: 100 失败请求数: 0
*************************  结果 end   ****************************
```

POST Request : Specifies the Header message header and body message body
```bash
$ ./domain-stress stress -d www.riversouth.xyz -c 10 -n 200 -p :10081/api/v1/namespaces/default/services -m POST \
    -H Authorization: Beaer eyJhbGciOiJSUzI1NiIsImtpZCI6ImFlU3h1YjEwSTZ6dThScjBOX1JHX0RFVk5EWnMwVmlSanZjdWN1NTU1MXMifQ.eyJpc3MiOiJrdWJlcm5ldGVzL3NlcnZpY2VhY2NvdW50Iiwia3ViZXJuZXRlcy5pby9zZXJ2aWNlYWNjb3VudC9uYW1lc3BhY2UiOiJrdWJlLXN5c3RlbSIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VjcmV0Lm5hbWUiOiJhZG1pbi10b2tlbi12bmN3ZCIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50Lm5hbWUiOiJhZG1pbiIsImt1YmVybmV0ZXMuaW8vc2VydmljZWFjY291bnQvc2VydmljZS1hY2NvdW50LnVpZCI6IjEwYWNlODM5LTg2ZjgtNGRiOC1hMWUwLTdhZTk4MmZiZDM1NCIsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDprdWJlLXN5c3RlbTphZG1pbiJ9.qOhvRaHRiQv0ZA7axVYLWPvr7-I_hYDM5NLbl_vSimhP9GuZ1J9KoGHEvK_MG7WHfvPpJrUD_RFDTfJ2e0xi8u7eIlHP2kD1cCkn6JWgPwu1SG_NT2HcXus2mtQdlKQOXB0ZpOqmaur1DSExXiUydzNweIaitVsLYfBA8ZTb8BLweDFXgq91CusditgvgzF8105kxhN5-ILtalwxUq-gAqh3miq_O3c1M68tRUfp1uoQj-vyeqztwXTr6lrHfNCtyrvFyyMiyIiHeGZgf6gVDGsE26daq5UJTpCRy5DyiFNw70K4PD3ywMoPXNpmF3KFvcUJbfNzkxM0_zNpQ2zE6w \
    -b { "kind": "Service", "apiVersion": "v1", "metadata": { "name": "nginx", "labels": { "app": "nginx" }, "namespace": "default" }, "spec": { "selector": { "app": "nginx" }, "type": "NodePort", "ports": [ { "name": "nginx", "nodePort": 30000, "port": 80, "protocol": "TCP", "targetPort": 80 } ] } }
```

Timeout :  Specifies a timeout of 1000ms.The default is 5000ms
```bash
$ ./domain-stress stress -d ebay.com -c 3 -n 120 -t 1000
```

Retry Number : Specifies that the number of retries is 1.The default is 3
```bash
$ ./domain-stress stress -d ebay.com -c 3 -n 120 -r 1
```

Test Command  : Test delay, based on ICMP protocol.Some firewalls or hosts may ban pinging
```bash
$ ./domain-stress test --domain www.baidu.com
Your target domian is: www.baidu.com 
The ips bebind are : 
36.152.44.96
36.152.44.95
The Delay for ips by ping : 
&{3 3 0 0 36.152.44.96 36.152.44.96 [24.013634ms 33.595654ms 34.765174ms] 24.013634ms 34.765174ms 30.791487ms 4.816389ms}
&{3 3 0 0 36.152.44.95 36.152.44.95 [39.162182ms 38.494679ms 18.895278ms] 18.895278ms 39.162182ms 32.184047ms 9.400529ms}
```
### Options
Input Parameters
```bash
$ ./domain-stress stress -h
Stress all ip after domain

Usage:
  domain-stress stress [flags]

Flags:
  -b, --body string           HTTP POST方式传送数据body消息体
  -c, --concurrency uint      并发数 (default 1)
  -d, --domain string         目标域名
  -f, --form string           请求格式 (default "http")
  -H, --headers stringArray   自定义头信息传递给服务器 示例:-H 'Content-Type: application/json'
  -h, --help                  help for stress
  -m, --method string         请求方法 (default "GET")
  -M, --mode string           负载均衡模式 (default "Average")
  -p, --path string           请求路径 (default "/")
  -r, --retry uint            重试次数 (default 3)
  -t, --timeout uint          超时时间 单位ms (default 5000)
  -n, --total uint            请求数(单个并发/协程) (default 1)
```


### Content to be developed
× At present, the protocol mode only supports http, and the subsequent grpc and websocket protocols need to be supported   
× At present, the backend IP of Execution plan is an evenly distributed request, and it needs to support allocation according to the backend IP latency.

### License
This software is released under the [Apache 2.0 license](https://www.apache.org/licenses/LICENSE-2.0).