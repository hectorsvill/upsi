# upsi - Multi URL status checker 
```shell
➜  upsi git:(main) ✗ ./upsi
Usage:
./upsi -s "list of urls separated by commas"
./upsi -f "file location:list of urls separated by new line"
➜  upsi git:(main) ✗ ./upsi -s "http://www.facebook.com,http://www.youtube.com,http://www.yahoo.com,http://www.amazon.com"
URL: http://www.yahoo.com                       | Status: 429 Too Many Requests
URL: http://www.amazon.com                      | Status: 200 OK
URL: http://www.youtube.com                     | Status: 200 OK
URL: http://www.facebook.com                    | Status: 200 OK
➜  upsi git:(main) ✗ ./upsi -f test.txt
URL: http://www.googleadservices.com            | Status: 404 Not Found
URL: http://www.googleweblight.com              | Status: 404 Not Found
URL: http://www.onet.pl                         | Status: 200 OK
URL: http://www.answers.yahoo.com               | Status: 429 Too Many Requests
➜  upsi git:(main) ✗
```
