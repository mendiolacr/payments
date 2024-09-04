CREATE DATABASE payment_platform;

USE payment_platform;

-- Table for Customers
CREATE TABLE Customers (
    customer_id INT IDENTITY(1,1) PRIMARY KEY, -- Unique identifier for each customer (auto-incremented)
    name VARCHAR(255) NOT NULL, -- Name of the customer
);

-- Table for Merchants
CREATE TABLE Merchants (
    merchant_id INT IDENTITY(1,1) PRIMARY KEY, -- Unique identifier for each merchant (auto-incremented)
    name VARCHAR(255) NOT NULL, -- Name of the merchant
);

-- Table for Payments
CREATE TABLE Payments (
    payment_id INT IDENTITY(1,1) PRIMARY KEY, -- Unique identifier for each payment (auto-incremented)
    customer_id INT, -- Reference to the customer making the payment (foreign key)
    merchant_id INT, -- Reference to the merchant receiving the payment (foreign key)
    amount DECIMAL(10,2) NOT NULL, -- Amount of the payment
    currency VARCHAR(10) NOT NULL, -- Currency of the payment
    status VARCHAR(20) NOT NULL, -- Status of the payment
    created_at DATETIME DEFAULT GETDATE(), -- Timestamp of the payment
    FOREIGN KEY (customer_id) REFERENCES Customers(customer_id), -- Foreign key constraint to Customers
    FOREIGN KEY (merchant_id) REFERENCES Merchants(merchant_id) -- Foreign key constraint to Merchants
);

-- Table for Refunds
CREATE TABLE Refunds (
    refund_id INT IDENTITY(1,1) PRIMARY KEY, -- Unique identifier for each refund (auto-incremented)
    payment_id INT, -- Reference to the payment being refunded (foreign key)
    amount DECIMAL(10,2) NOT NULL, -- Amount of the refund
    status VARCHAR(20) NOT NULL, -- Status of the refund
    created_at DATETIME DEFAULT GETDATE(), -- Timestamp of the refund
    FOREIGN KEY (payment_id) REFERENCES Payments(payment_id) -- Foreign key constraint to Payments
);

-- Table for Transactions
CREATE TABLE Transactions (
    transaction_id INT IDENTITY(1,1) PRIMARY KEY, -- Unique identifier for each transaction (auto-incremented)
    payment_id INT, -- Reference to the payment associated with the transaction (foreign key)
    transaction_type VARCHAR(20) NOT NULL, -- Type of transaction 
    amount DECIMAL(10,2) NOT NULL, -- Amount of the transaction
    status VARCHAR(20) NOT NULL, -- Status of the transaction 
    created_at DATETIME DEFAULT GETDATE(), -- Timestamp of the transaction
    FOREIGN KEY (payment_id) REFERENCES Payments(payment_id) -- Foreign key constraint to Payments
);

-- Insert test costumers
INSERT INTO Customers (name)
VALUES
    ( 'Customer One'),
    ( 'Customer Two'),
    ( 'Customer Three');

-- Insert test Merchants
INSERT INTO Merchants (name)
VALUES
    ( 'Merchant One'),
    ( 'Merchant Two'),
    ( 'Merchant Three');


	select * from Transactions	
	select * from Payments