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
