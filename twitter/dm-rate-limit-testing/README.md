Twitter DM送信APIのレートリミットに関する実験
----

Twitter DM送信 APIのレートリミットについて、細かい検証。

をしようと思ったが、いざ実際に動かして連続でDMを送り合うと`error 403 {"errors":[{"code":226,"message":"This request looks like it might be automated. To protect our users from spam and other malicious activity, we can't complete this action right now. Please try again later."}]}` で死んでまともに実験出来なくて死んだ。
あんまりやるとDeveloper AccountがBANされちゃいそうなので実験中止。
