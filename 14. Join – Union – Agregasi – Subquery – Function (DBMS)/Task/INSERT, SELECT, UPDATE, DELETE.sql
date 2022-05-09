------------------------------------------ INSERT

a. insert 5 operators pada table operators

INSERT INTO operators VALUES ('1', 'Erika', NOW(), NOW());
INSERT INTO operators VALUES ('2', 'Syifa', NOW(), NOW());
INSERT INTO operators VALUES ('3', 'Jony', NOW(), NOW());
INSERT INTO operators VALUES ('4', 'Evan', NOW(), NOW());
INSERT INTO operators VALUES ('5', 'Dion', NOW(), NOW());

b. Insert 3 product type

INSERT INTO product_types VALUES ('1', 'Soap', NOW(), NOW());
INSERT INTO product_types VALUES ('2', 'Tissue', NOW(), NOW());
INSERT INTO product_types VALUES ('3', 'Tooth phaste', NOW(), NOW());

c. Insert 2 product dengan product type id = 1, dan operators id = 3

INSERT INTO products VALUES ('1', '1', '3', 'S001', 'Nuvo', '1', NOW(), NOW());
INSERT INTO products VALUES ('2', '1', '3', 'S002', 'Lux', '1', NOW(), NOW());

d. Insert 3 product dengan product type id = 2, dan operators id = 1

INSERT INTO products VALUES ('3', '2', '1', 'T001', 'Jolly', '1', NOW(), NOW());
INSERT INTO products VALUES ('4', '2', '1', 'T002', 'Green', '1', NOW(), NOW());
INSERT INTO products VALUES ('5', '2', '1', 'T003', 'Paseo', '1', NOW(), NOW());

e. Insert 3 product dengan product type id = 3, dan operators id = 4

INSERT INTO products VALUES ('6', '3', '4', 'TP001', 'Pepsodent', '1', NOW(), NOW());
INSERT INTO products VALUES ('7', '3', '4', 'TP002', 'Maxima', '1', NOW(), NOW());
INSERT INTO products VALUES ('8', '3', '4', 'TP003', 'Shafi', '1', NOW(), NOW());

f. Insert product description pada setiap product

INSERT INTO product_descriptions VALUES ('1', 'Sabun Antibakteri', NOW(), NOW());
INSERT INTO product_descriptions VALUES ('2', 'Sabun Mewah', NOW(), NOW());
INSERT INTO product_descriptions VALUES ('3', 'Tisu Bagus dan Murah', NOW(), NOW());
INSERT INTO product_descriptions VALUES ('4', 'Tisu Murah', NOW(), NOW());
INSERT INTO product_descriptions VALUES ('5', 'Tisu Terbagus', NOW(), NOW());
INSERT INTO product_descriptions VALUES ('6', 'Pasta Gigi Nomor 1', NOW(), NOW());
INSERT INTO product_descriptions VALUES ('7', 'Pasta Gigi Murah', NOW(), NOW());
INSERT INTO product_descriptions VALUES ('8', 'Pasta Gigi Mantap', NOW(), NOW());

g. Insert 3 payment methods

INSERT INTO payment_methods VALUES ('1', 'Tunai', '1', NOW(), NOW());
INSERT INTO payment_methods VALUES ('2', 'Kartu Debit', '1', NOW(), NOW());
INSERT INTO payment_methods VALUES ('3', 'Kartu Kredit', '1', NOW(), NOW());

h. insert 5 user pada tabel user

INSERT INTO users VALUES ('1', '1', '2000-06-29', 'W', NOW(), NOW());
INSERT INTO users VALUES ('2', '1', '1999-03-1', 'M', NOW(), NOW());
INSERT INTO users VALUES ('3', '1', '2000-05-14', 'W',NOW(), NOW());
INSERT INTO users VALUES ('4', '1', '1994-12-07', 'M', NOW(), NOW());
INSERT INTO users VALUES ('5', '1', '2002-06-23', 'W',NOW(), NOW());

i. Insert 3 transaction di masing-masing user

INSERT INTO transactions VALUES ('1', '1', '1', '1', '3', '24000', NOW(), NOW());
INSERT INTO transactions VALUES ('2', '1', '1', '1', '3', '15000', NOW(), NOW());
INSERT INTO transactions VALUES ('3', '1', '1', '1', '3', '12000', NOW(), NOW());

INSERT INTO transactions VALUES ('4', '2', '1', '1', '3', '19000', NOW(), NOW());
INSERT INTO transactions VALUES ('5', '2', '1', '1', '3', '20000', NOW(), NOW());
INSERT INTO transactions VALUES ('6', '2', '1', '1', '3', '17000', NOW(), NOW());

INSERT INTO transactions VALUES ('7', '3', '1', '1', '3', '15000', NOW(), NOW());
INSERT INTO transactions VALUES ('8', '3', '1', '1', '3', '20000', NOW(), NOW());
INSERT INTO transactions VALUES ('9', '3', '1', '1', '3', '25000', NOW(), NOW());

INSERT INTO transactions VALUES ('10', '4', '1', '1', '3', '12000', NOW(), NOW());
INSERT INTO transactions VALUES ('11', '4', '1', '1', '3', '21000', NOW(), NOW());
INSERT INTO transactions VALUES ('12', '4', '1', '1', '3', '16000', NOW(), NOW());

INSERT INTO transactions VALUES ('13', '5', '1', '1', '3', '14000', NOW(), NOW());
INSERT INTO transactions VALUES ('14', '5', '1', '1', '3', '20000', NOW(), NOW());
INSERT INTO transactions VALUES ('15', '5', '1', '1', '3', '21000', NOW(), NOW());

j. Insert 3 product di masing-masing transaction

INSERT INTO transaction_details VALUES ('1', '4', 'Lunas', '1', '7000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('1', '5', 'Lunas', '1', '12000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('1', '6', 'Lunas', '1', '5000', NOW(), NOW());

INSERT INTO transaction_details VALUES ('2', '2', 'Lunas', '1', '3000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('2', '4', 'Lunas', '1', '7000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('2', '6', 'Lunas', '1', '5000', NOW(), NOW());

INSERT INTO transaction_details VALUES ('3', '1', 'Lunas', '1', '4000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('3', '7', 'Lunas', '1', '4000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('3', '8', 'Lunas', '1', '4000', NOW(), NOW());

INSERT INTO transaction_details VALUES ('4', '3', 'Lunas', '1', '10000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('4', '6', 'Lunas', '1', '5000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('4', '7', 'Lunas', '1', '4000', NOW(), NOW());

INSERT INTO transaction_details VALUES ('5', '2', 'Lunas', '1', '3000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('5', '3', 'Lunas', '1', '10000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('5', '4', 'Lunas', '1', '7000', NOW(), NOW());

INSERT INTO transaction_details VALUES ('6', '1', 'Lunas', '1', '4000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('6', '2', 'Lunas', '1', '3000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('6', '3', 'Lunas', '1', '10000', NOW(), NOW());

INSERT INTO transaction_details VALUES ('7', '1', 'Lunas', '1', '4000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('7', '4', 'Lunas', '1', '7000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('7', '8', 'Lunas', '1', '4000', NOW(), NOW());

INSERT INTO transaction_details VALUES ('8', '2', 'Lunas', '1', '3000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('8', '5', 'Lunas', '1', '12000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('8', '6', 'Lunas', '1', '5000', NOW(), NOW());

INSERT INTO transaction_details VALUES ('9', '1', 'Lunas', '1', '4000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('9', '7', 'Lunas', '1', '4000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('9', '8', 'Lunas', '1', '4000', NOW(), NOW());

INSERT INTO transaction_details VALUES ('10', '1', 'Lunas', '1', '4000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('10', '3', 'Lunas', '1', '10000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('10', '4', 'Lunas', '1', '7000', NOW(), NOW());

INSERT INTO transaction_details VALUES ('11', '7', 'Lunas', '1', '4000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('11', '3', 'Lunas', '1', '10000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('11', '4', 'Lunas', '1', '7000', NOW(), NOW());

INSERT INTO transaction_details VALUES ('12', '4', 'Lunas', '1', '7000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('12', '6', 'Lunas', '1', '5000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('12', '7', 'Lunas', '1', '4000', NOW(), NOW());

INSERT INTO transaction_details VALUES ('13', '1', 'Lunas', '1', '4000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('13', '2', 'Lunas', '1', '3000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('13', '4', 'Lunas', '1', '7000', NOW(), NOW());

INSERT INTO transaction_details VALUES ('14', '5', 'Lunas', '1', '12000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('14', '7', 'Lunas', '1', '4000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('14', '8', 'Lunas', '1', '4000', NOW(), NOW());

INSERT INTO transaction_details VALUES ('15', '1', 'Lunas', '1', '4000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('15', '5', 'Lunas', '1', '12000', NOW(), NOW());
INSERT INTO transaction_details VALUES ('15', '6', 'Lunas', '1', '5000', NOW(), NOW());

------------------------------------------ SELECT

a. Tampilkan nama user/pelanggan dengan gender laki-laki/M

 SELECT name FROM users WHERE GENDER = 'M';

b. Tampilkan product dengan id = 3

SELECT code, name, status FROM products WHERE id = 3;

c. Tampilkan data pelanggan yang created_at dalam range 7 hari kebelakang dan mempunyai nama mengandung kata 'a'

SELECT name, dob, gender, YEARWEEK(created_at) FROM users WHERE name LIKE '%a';

d. Hitung jumlah user/pelanggan dengan status gender Perempuan

SELECT gender, COUNT(id) FROM users WHERE GENDER = 'W';

e. Tampilkan data pelanggan dengan urutan sesuai nama abjad

SELECT name, dob, gender FROM users ORDER BY name ASC;

f. Tampilkan 5 data pada data product

SELECT * FROM products LIMIT 5


------------------------------------------ UPDATE

a. Ubah data product id 1 dengan nama 'product dummy'

UPDATE products SET name = 'product dummy' WHERE id = 1;

b. Update qty = 3 pada transactions detail dengan product id 1

UPDATE transaction_details SET qty = '3' WHERE product_id = 1;


------------------------------------------ DELETE

a. Delete data pada tabel product dengan id 1

DELETE FROM products WHERE id = 1

b. Delete data pada tabel product dengan product type id 1

DELETE FROM products WHERE product_type_id = 1

