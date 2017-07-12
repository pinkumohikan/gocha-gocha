limonade ベンチマーク
====

https://limonade-php.github.io/

ベンチマーク結果
----

条件

```
$ php71 -v
PHP 7.1.7 (cli) (built: Jul  6 2017 12:18:04) ( NTS )
Copyright (c) 1997-2017 The PHP Group
Zend Engine v3.1.0, Copyright (c) 1998-2017 Zend Technologies
    with Zend OPcache v7.1.7, Copyright (c) 1999-2017, by Zend Technologies
```

```
$ siege -c 300 -t 10s http://133.18.25.181
** SIEGE 4.0.2
** Preparing 300 concurrent users for battle.
The server is now under siege...
Lifting the server siege...
Transactions:		       21531 hits
Availability:		      100.00 %
Elapsed time:		        9.56 secs
Data transferred:	        0.60 MB
Response time:		        0.13 secs
Transaction rate:	     2252.20 trans/sec
Throughput:		        0.06 MB/sec
Concurrency:		      290.06
Successful transactions:       21531
Failed transactions:	           0
Longest transaction:	        7.08
Shortest transaction:	        0.02
```
