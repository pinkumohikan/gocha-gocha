fortee上で応募されているプロポーザルの枠ごとの件数見つける
----

せっかくプロポーザル応募するなら競争率が低く、登壇出来る可能性がある枠を知りたい。
そのためにはまず、どの枠にどのぐらい応募がされているのかを知る必要がある。


### 使い方

```
$ make
fetching... page=1
fetching... page=2
fetching... page=3
fetching... page=4
fetching... page=5
fetching... page=6
fetching... page=7
find ./proposal-pages -type f | xargs -n 1 ./parse.sh 2>/dev/null | sort | uniq -c
  20  LT（5分）
  37 レギュラートーク(15分)
  45 レギュラートーク(30分)
  18 レギュラートーク(45分)
```
