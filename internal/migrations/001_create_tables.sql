-- +goose Up
PRAGMA foreign_keys = ON;

CREATE TABLE IF NOT EXISTS company (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    gst TEXT NOT NULL,
    address TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS payment_detail (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    acc_holder TEXT NOT NULL,
    acc_number INTEGER NOT NULL,
    ifsc TEXT NOT NULL,
    branch TEXT NOT NULL,
    bank_name TEXT NOT NULL,
    virtual_payment_addr TEXT NOT NULL,
    fk_company_id INTEGER NOT NULL,
    FOREIGN KEY (fk_company_id)
        REFERENCES company (id)
            ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS shipping_address (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    address TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS item (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    hsn INTEGER NOT NULL,
    price INTEGER NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS item;
DROP TABLE IF EXISTS shipping_address;
DROP TABLE IF EXISTS payment_detail;
DROP TABLE IF EXISTS company;
PARGMA foreign_keys = OFF;
