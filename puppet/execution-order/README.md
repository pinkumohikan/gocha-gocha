execution-order
====

概要
----
チェイニングアローを利用して実行順序を指定する場合、クラスのincludeの位置で実行順序が変わるか検証した


実行結果 (一部)
----

```
$ make apply
puppet apply --modulepath=./modules/ --verbose manifests/site.pp
Notice: Compiled catalog for marlin.pinkumohikan.com in environment production in 0.02 seconds
Info: Applying configuration version '1492495013'
Notice: install
Notice: /Stage[main]/Execution_order::Install/Notify[notify_install]/message: defined 'message' as 'install'
Notice: prepare
Notice: /Stage[main]/Execution_order::Prepare/Notify[notify_prepare]/message: defined 'message' as 'prepare'
Notice: hoge_one
Notice: /Stage[main]/Execution_order/Notify[notify_hoge_one]/message: defined 'message' as 'hoge_one'
Notice: hoge_two
Notice: /Stage[main]/Execution_order/Notify[notify_hoge_two]/message: defined 'message' as 'hoge_two'
Notice: Finished catalog run in 0.03 seconds
```


結論
----

* チェイニングアローで指定した順番でクラスが実行される
    * チェイニングアローの前後でincludeしても結果は変わらない
        * そのクラスをすべて解釈したあと、チェイニングアローに基づいて実行順を制御している模様


参考資料
----
* https://docs.puppet.com/puppet/4.10/lang_relationships.html
