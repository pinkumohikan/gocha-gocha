# mysqldefお試し
RidgepoleライクなDBマイグレーションツールのsqldefを使ってみる
https://github.com/k0kubun/sqldef


## 前提
* sqldef: v0.11.38, MySQL版, darwin amd64用ビルド
* MySQL: 8.0.28


## 気になりポイント
### view定義時、ソースクエリ内のテーブル/カラム名にDB名がついた完全修飾のかたちでないと不一致扱いとなる

#### 再現

(1) 下記のようなschema.sqlを用意する

```sql
CREATE TABLE orders (
    id         INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id    INT UNSIGNED NOT NULL,
    product_id INT UNSIGNED NOT NULL
);
CREATE VIEW order_count_per_user_view AS
SELECT user_id,
       COUNT(orders.id) AS order_count
FROM orders
GROUP BY user_id;
```

(2) (DBまっさらな状態で) applyする

```bash
$ ./bin/mysqldef --host 127.0.0.1 --user root --port 33060 exp1 < ./schema.sql
-- Apply --
CREATE TABLE orders (
    id         INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    user_id    INT UNSIGNED NOT NULL,
    product_id INT UNSIGNED NOT NULL
);
CREATE VIEW order_count_per_user_view AS
SELECT user_id,
       COUNT(orders.id) AS order_count
FROM orders

$ echo $?
0
```

(3) 再度applyするとviewだけ差分ありとみなされる

```bash
$ ./bin/mysqldef --host 127.0.0.1 --user root --port 33060 exp1 < ./schema.sql
-- Apply --
CREATE OR REPLACE VIEW `order_count_per_user_view` AS select user_id, COUNT(orders.id) as order_count from orders group by user_id;
```

#### 深堀り

(1) 現在のテーブル定義をexportすると差分なしになる

```bash
$ ./bin/mysqldef --host 127.0.0.1 --user root --port 33060 exp1 --export > schema.sql

$ ./bin/mysqldef --host 127.0.0.1 --user root --port 33060 exp1 < ./schema.sql
-- Nothing is modified --
```

(2) exportしたviewのDDLを比較

before:

```sql
CREATE VIEW order_count_per_user_view AS
SELECT user_id,
       COUNT(orders.id) AS order_count
FROM orders
GROUP BY user_id;
```

after:

```sql
CREATE VIEW order_count_per_user_view AS
SELECT `exp1`.`orders`.`user_id` AS `user_id`,
       count(`exp1`.`orders`.`id`) AS `order_count`
FROM `exp1`.`orders`
GROUP BY `exp1`.`orders`.`user_id`;
```

※ afterは整形済み (整形した結果差分がでないことは確認済み)

(3) viewのDDLからDB名を削ってみると差分ありになる

```bash
$ sed -i '' -e 's/`exp1`.//g' schema.sql

$ git diff
diff --git a/mysqldef/exp1/schema.sql b/mysqldef/exp1/schema.sql
index 62412fa..4177142 100644
--- a/mysqldef/exp1/schema.sql
+++ b/mysqldef/exp1/schema.sql
@@ -6,7 +6,7 @@ CREATE TABLE `orders` (
 ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

 CREATE VIEW order_count_per_user_view AS
-SELECT `exp1`.`orders`.`user_id` AS `user_id`,
-       count(`exp1`.`orders`.`id`) AS `order_count`
-FROM `exp1`.`orders`
-GROUP BY `exp1`.`orders`.`user_id`;
+SELECT `orders`.`user_id` AS `user_id`,
+       count(`orders`.`id`) AS `order_count`
+FROM `orders`
+GROUP BY `orders`.`user_id`;
```

```bash
$ ./bin/mysqldef --host 127.0.0.1 --user root --port 33060 exp1 < schema.sql
-- Apply --
CREATE OR REPLACE VIEW `order_count_per_user_view` AS select orders.user_id as user_id, count(orders.id) as order_count from orders group by orders.user_id;
```
