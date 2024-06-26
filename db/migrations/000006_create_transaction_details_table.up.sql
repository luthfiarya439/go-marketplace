CREATE TABLE transaction_details (
    id INT AUTO_INCREMENT PRIMARY KEY,
    transaction_id INT,
    product_id INT,
    product_name VARCHAR(255) NOT NULL,
    product_price VARCHAR(255) NOT NULL,
    product_quantity VARCHAR(255) NOT NULL,
    FOREIGN KEY(transaction_id) REFERENCES transactions(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
    FOREIGN KEY(product_id) REFERENCES products(id)
    ON DELETE SET NULL
    ON UPDATE CASCADE
);