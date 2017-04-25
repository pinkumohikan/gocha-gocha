install-yum-package
====

概要
----
* packageリソースを使ってパッケージをインストールしてみる
* requireを使ってリソースの依存を明示してみる


実行結果
----

* packageリソースを使ってパッケージをインストールしてみる

```
$ sudo make
puppet apply --modulepath=./modules/ --verbose manifests/site.pp
Notice: Compiled catalog for marlin.pinkumohikan.com in environment production in 0.34 seconds
Info: Applying configuration version '1493084881'
Notice: /Stage[main]/Sl::Install/Package[sl]/ensure: created
Info: Creating state file /var/lib/puppet/state/state.yaml
Notice: Finished catalog run in 5.14 seconds
```

* requireを使ってリソースの依存を明示してみる
    * emacs depends to vimの記述をした状態で、vimのリソース定義を削除して動かしてみる

```
$ sudo make
puppet apply --modulepath=./modules/ --verbose manifests/site.pp
Notice: Compiled catalog for marlin.pinkumohikan.com in environment production in 0.34 seconds
Error: Could not find dependency Package[vim] for Package[emacs] at /home/h-shinoda/gocha-gocha/puppet/install-yum-package/modules/sl/manifests/init.pp:15
make: *** [apply] エラー 1
```


メモ
----

* allow_virtual絡みでワーニングが出た
    * http://s-tajima.hateblo.jp/entry/2014/07/30/161215 を読んで、 `allow_virtual => true` にて対応

```
Warning: The package type's allow_virtual parameter will be changing its default value from false to true in a future release. If you do not want to allow virtual packages, please explicitly set allow_virtual to false.
   (at /usr/share/ruby/vendor_ruby/puppet/type.rb:816:in `set_default')
```
* デフォルト無効のyumリポジトリを使う方法
    * http://blog.hifumi.info/2014/12/29/puppet-package-enablerepo-by-install_options/ を読んで、 `install_options` にて対応


参考資料
----
* http://s-tajima.hateblo.jp/entry/2014/07/30/161215
* http://blog.hifumi.info/2014/12/29/puppet-package-enablerepo-by-install_options/
* https://docs.puppet.com/puppet/latest/types/package.html
