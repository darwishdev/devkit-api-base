-- name: CustomerFind :one
with customer_data as (
  SELECT
    c.customer_id,
    c.customer_code,
    c.customer_name,
    c.customer_image,
    c.customer_email,
    c.customer_phone,
    c.customer_password,
    c.customer_national_id,
    c.birthdate,
    c.created_at,
    c.updated_at,
    c.deleted_at
FROM
    accounts_schema.customers c
WHERE
    lower(c.customer_email) = lower(sqlc.arg('search_key'))
    OR lower(c.customer_phone) = lower(sqlc.arg('search_key'))
    OR lower(c.customer_code) = lower(sqlc.arg('search_key'))
    OR c.customer_id = sqlc.arg('customer_id')
) , ongoing_reservation as (
  select r.* from reservations_schema.reservations_view r join customer_data c on r.customer_id = c.customer_id and r.reservation_status_id <= 3
)  , ongoing_payment as (
  select p.* from reservations_schema.payments p join customer_data c on p.customer_id = c.customer_id and p.payment_status_id = 1
)  , 
reservations as (
  select r.* from reservations_schema.reservations_view r join customer_data c on r.customer_id = c.customer_id where r.tense = 'future'
)  
SELECT
    c.customer_id,
    c.customer_code,
    c.customer_name,
    c.customer_image,
    c.customer_email,
    c.customer_phone,
    c.birthdate,

    c.customer_password,
    c.customer_national_id,
    c.created_at,
    c.updated_at,
    c.deleted_at,
   (select to_jsonb(main) from ongoing_reservation main ) ongoing_reservation,
   (select to_jsonb(main) from ongoing_payment main ) ongoing_payment,
   ( SELECT Jsonb_agg(nested_reservations) FROM (select * from reservations) nested_reservations) reservations
FROM
    customer_data c;




   
 
-- name: CustomerResetPassword :exec
UPDATE
    accounts_schema.customers
SET
    customer_password =sqlc.arg('customer_password')
WHERE
    customer_email = sqlc.arg('customer_email');

-- name: CustomersList :many
SELECT
    customer_id,
    customer_code,
    customer_name,
    customer_image,
    customer_email,
    customer_phone,
    customer_password,
    birthdate,
    customer_national_id,
    created_at,
    updated_at,
    deleted_at
FROM
    accounts_schema.customers;
-- name: CustomerDeleteRestore :exec
select u.* 
from accounts_schema.customers u 
join toggle_delete_restore('accounts_schema.customers', 'customer_id', sqlc.arg('records')::int[]) d 
on d.record_id = u.customer_id;
-- name: CustomerCreateUpdate :one
select * from accounts_schema.customer_create_update(
    in_customer_id => sqlc.arg('customer_id')::int,
    in_customer_name => sqlc.arg('customer_name')::varchar(200),
    in_customer_image => sqlc.arg('customer_image')::varchar(200),
    in_customer_email => sqlc.arg('customer_email')::varchar(200),
    in_birthdate => sqlc.arg('birthdate')::varchar(200),
    in_customer_phone => sqlc.arg('customer_phone')::varchar(200),
    in_customer_password => sqlc.arg('customer_password')::varchar(200),
    in_customer_national_id => sqlc.arg('customer_national_id')::varchar(30)
); 





