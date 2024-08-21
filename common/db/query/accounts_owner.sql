
-- name: OwnerFind :one
SELECT
    owner_id,
    owner_name,
    owner_image,
    owner_email,
    owner_phone,
    owner_password,
    owner_national_id,
    created_at,
    updated_at,
    deleted_at,
    properties
FROM
    accounts_schema.owners_view
WHERE
    owner_id = $1;
-- name: OwnerResetPassword :exec
UPDATE
    accounts_schema.owners
SET
    owner_password = $2
WHERE
    owner_email = $1;

-- name: OwnersList :many
SELECT
    owner_id,
    owner_name,
    owner_image,
    owner_email,
    owner_phone,
    owner_password,
    owner_national_id,
    representative_owner_id,
    created_at,
    updated_at,
    deleted_at
FROM
    accounts_schema.owners ;


-- name: OwnersInputList :many
SELECT
    owner_id record_value,
    owner_name record_label,
    concat( owner_email ,'(' , owner_national_id , ')')::varchar  note
FROM
    accounts_schema.owners where deleted_at is null;
-- name: OwnerFindByEmailOrCode :one
SELECT
    owner_id,
    owner_name,
    owner_image,
    owner_email,
    owner_phone,
    owner_password,
    owner_national_id,
    created_at,
    updated_at,
    deleted_at,
    properties
FROM
    accounts_schema.owners_view
WHERE
    owner_email = $1
    OR owner_phone = $1;
 
-- name: OwnerDeleteRestore :exec
UPDATE
    accounts_schema.owners
SET
    deleted_at = IIF(deleted_at IS NULL, now(), NULL)
WHERE
    owner_id = ANY (sqlc.arg('records')::int[]);

-- name: OwnerCreate :one
INSERT INTO accounts_schema.owners(
    owner_name,
    owner_image,
    owner_email,
    owner_password,
    owner_phone,
    owner_national_id,
    representative_owner_id)
    VALUES ($1, $2, $3, $4, $5 , $6 , $7)
RETURNING
    *;

-- name: OwnerUpdate :one
UPDATE
    accounts_schema.owners
SET
    owner_name = $2,
    owner_image = $3,
    owner_email = $4,
    owner_phone = $5,
    owner_national_id = $6,
    representative_owner_id = $7,
    owner_password = sqlc.arg('owner_password')
WHERE
    owner_id = $1
RETURNING
    *;
  
-- name: OwnerFindForUpdate :one
SELECT
    owner_id,
    owner_name,
    owner_image,
    owner_email,
    owner_phone,
    owner_national_id,
    representative_owner_id
FROM
    accounts_schema.owners
WHERE
    owner_id = $1;