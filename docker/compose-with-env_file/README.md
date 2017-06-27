docker-composeのenv_file設定に 'export ' prefixを含むenvファイルを指定した場合の挙動を調べる
====

概要
----
dotenvを読み込むライブラリに、 'export ' prefixをいい感じに無視して後続のKEY=VALUEの部分のみを環境変数に設定してくれるものがあるらしく、docker-composeの env_file でそういうファイルを指定した場合にどういう挙動をするか調べる


結果
----
'export ' prefixを含む変数は **消失する**

```
$ cat hoge.env
HOGE=HOGE
AGE=SAGE
export DON=GURI
```

```
$ docker-compose -f docker-compose.yml run env-tashikametai
/ # printenv
HOGE=HOGE
HOSTNAME=9d7932cde3eb
SHLVL=1
HOME=/root
TERM=xterm
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
PWD=/
AGE=SAGE
```

note
----
* alpineのdocker image
    * https://hub.docker.com/_/alpine/
* reference: docker-compose file-v2 env_file
    * https://docs.docker.com/compose/compose-file/compose-file-v2/#env_file
