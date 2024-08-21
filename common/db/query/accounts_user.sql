-- name: UserFind :one
SELECT
    user_id,
    user_name,
    user_code,
    user_image,
    user_email,
    user_phone,
    user_password,
    created_at,
    updated_at,
    deleted_at
FROM
    accounts_schema.users
WHERE deleted_at is null and (
    user_email = sqlc.arg('search_key')
    OR user_phone = sqlc.arg('search_key')
    OR user_code = sqlc.arg('search_key')
    OR user_id = sqlc.arg('user_id'));


-- name: UserFindNavigationBars :many
WITH userpermissions AS (
  SELECT 
    permission_id::int
  FROM 
    accounts_schema.user_find_permissions($1) 
)
, allowed_navigations as (
    SELECT
        navigation_bar_id,
        menu_key "key",
        label,
        label_ar,
        icon_id,
        "route",
        menu_key,
        parent_id
        from accounts_schema.navigation_bars n
        JOIN userpermissions p on n.permission_id = p.permission_id 
    union 
    SELECT
        navigation_bar_id,
        menu_key "key",
        label,
        label_ar,
        icon_id,
        "route",
        menu_key,
        parent_id
        from accounts_schema.navigation_bars n 
        where n.permission_id is null
    ORDER BY
        menu_key
) , children_permissions as (
    select * from allowed_navigations where parent_id is not null
) select 
p.navigation_bar_id,
p.menu_key "key",
p.label,
p.label_ar,
p.icon_id,
p."route",
(
    select Jsonb_Agg(nested_items) from (
        select c.* from children_permissions c where 
        c.parent_id = p.navigation_bar_id
    ) nested_items
) items
from allowed_navigations p where route is null or parent_id is null order by p.navigation_bar_id;


-- name: UserPermissionsList :many
SELECT 
    permission_id::int,
    permission_function::varchar(200),
    permission_group::varchar(200),
    permission_name::varchar(200)
FROM 
    accounts_schema.user_find_permissions($1) ;



-- name: UserRolesList :many
SELECT 
   r.role_id,
   r.role_name
FROM 
    accounts_schema.user_roles ur 
join 
    accounts_schema.roles r on r.role_id = ur.role_id 
where
    ur.user_id = $1;
 
-- name: UserResetPassword :exec
UPDATE
    accounts_schema.users
SET
    user_password = $2
WHERE
    user_email = $1;

-- name: UsersList :many
with user_roles as (
    select role_id , user_id from accounts_schema.user_roles
) SELECT
    u.user_id,
    u.user_name,
    u.user_code,
    u.user_image,
    u.user_email,
    u.user_phone,
    u.user_password,
    (select array_agg(role_id)::int[] from user_roles ur where ur.user_id = u.user_id) role_ids,
    u.created_at,
    u.updated_at,
    u.deleted_at
FROM
    accounts_schema.users u;
-- name: UserDeleteRestore :exec
select * from accounts_schema.users_delete_restore(sqlc.arg('records')::int[]);

-- name: UserCreateUpdate :one
select * from accounts_schema.user_create_update(
    sqlc.arg('user_id')::int,
    sqlc.arg('user_name')::varchar(200),
    sqlc.arg('user_image')::varchar(200),
    sqlc.arg('user_email')::varchar(200),
    sqlc.arg('user_phone')::varchar(200),
    sqlc.arg('user_password')::varchar(200),
    sqlc.arg('permissions')::int[],
    sqlc.arg('roles')::int[]
);
-- name: UserFindForUpdate :one
with user_roles as (
    select array_agg(role_id) roles from  accounts_schema.user_roles where user_id = $1 
), user_permissions as (
    select array_agg(permission_id) permissions  from  accounts_schema.user_permissions where user_id = $1 
)
select 
u.user_id , 
u.user_name , 
u.user_code , 
u.user_image , 
u.user_email , 
u.user_phone , 
'' user_password , 
ur.roles::int[] roles,
up.permissions::int[] permissions
from accounts_schema.users u
join user_roles ur on true
join user_permissions up on true
where u.user_id = $1;