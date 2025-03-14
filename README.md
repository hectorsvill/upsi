# upsi - Multi URL status checker 
```shell
➜  upsi git:(main) ✗ go build
➜  upsi git:(main) ✗ ./upsi -s "http://www.facebook.com,http://www.youtube.com,http://www.yahoo.com,http://www.amazon.com"
http://www.yahoo.com - 429 Too Many Requests
http://www.youtube.com - 200 OK
http://www.facebook.com - 200 OK
http://www.amazon.com - 200 OK
➜  upsi git:(main) ✗
```
