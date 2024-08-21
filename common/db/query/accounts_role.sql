-- name: PermissionsList :many
WITH permissions_groups AS (
    SELECT
        p.permission_group
    FROM
        accounts_schema.permissions p
    GROUP BY
        p.permission_group
    order by p.permission_group
) 
    SELECT
        g.permission_group,
(
            SELECT
                Jsonb_agg(nested_permissions)
            FROM (
                SELECT
                    np.permission_id,
                    np.permission_name,
                    np.permission_function,
                    np.permission_description
                FROM
                    accounts_schema.permissions np
                WHERE
                    np.permission_group = g.permission_group) nested_permissions) permissions
        FROM
            permissions_groups g  ;

-- name: RoleCreate :one
INSERT INTO accounts_schema.roles(role_name, role_description)
    VALUES ($1, $2)
RETURNING
    *;

-- name: RolePermissionsBulkCreate :copyfrom
INSERT INTO accounts_schema.role_permissions(role_id, permission_id)
    VALUES ($1, $2);

-- name: RolePermissionsClear :exec
DELETE FROM accounts_schema.role_permissions
WHERE role_id = $1;

-- name: RoleUpdate :one
UPDATE
    accounts_schema.roles
SET
    role_name = $1,
    role_description = $2,
    updated_at = now()
WHERE
    role_id = $3
RETURNING
    *;

-- name: RoleDeleteRestore :exec
UPDATE
    accounts_schema.roles
SET
    deleted_at = CASE WHEN deleted_at IS NULL THEN
        now()
    ELSE
        NULL
    END
WHERE
    role_id = ANY (sqlc.arg('role_ids')::int[]);

-- name: RoleFindForUpdate :one
WITH rolePermissions AS (
    SELECT
        ARRAY_AGG(permission_id) permissions
    FROM
        accounts_schema.role_permissions
    WHERE
        role_id = $1
)
SELECT
    Row_to_json(role_record)::jsonb role_row
FROM (
    SELECT
        r.role_id,
        r.role_name,
        r.role_description,
        permissions
    FROM
        accounts_schema.roles r
        JOIN rolePermissions ON TRUE
    WHERE
        r.role_id = $1) role_record;

-- -- name: RoleFind :one
-- WITH current_record AS (
--     SELECT
--         user_id,
--         user_name,
--         user_image,
--         user_email,
--         user_phone,
--         user_password,
--         created_at,
--         updated_at,
--         deleted_at,
--         role_id,
--         role_name,
--         role_description,
--         created_at role_created_at,
--         updated_at role_updated_at,
--         deleted_at role_deleted_at,
--         permission_id,
--         permission_function,
--         permission_name,
--         permission_description,
--         permission_group,
--         role_permission_id,
--         role_permission_function,
--         role_permission_name,
--         role_permission_description,
--         role_permission_group
--     FROM
--         accounts_schema.users_view
--     WHERE
--         role_id = $1
-- ),
-- current_role_permissions AS (
--     SELECT
--         rp.role_id,
-- (
--             SELECT
--                 Jsonb_agg(nested_permissions)
--             FROM (
--                 SELECT
--                     rpp.role_permission_id,
--                     rpp.role_permission_name,
--                     rpp.role_permission_function,
--                     rpp.role_permission_description,
--                     rpp.role_permission_group
--                 FROM
--                     current_record rpp
--                 WHERE
--                     rpp.role_permission_id IS NOT NULL
--                 GROUP BY
--                     rpp.role_permission_id,
--                     rpp.role_permission_name,
--                     rpp.role_permission_function,
--                     rpp.role_permission_description,
--                     rpp.role_permission_group) nested_permissions) permissions
--         FROM
--             current_record rp
--         GROUP BY
--             rp.role_id
-- ),
-- current_role_users AS (
--     SELECT
--         rp.role_id,
-- (
--             SELECT
--                 Jsonb_agg(nested_users)
--             FROM (
--                 SELECT
--                     ru.user_id,
--                     ru.user_name,
--                     ru.user_image,
--                     ru.user_email,
--                     ru.user_phone,
--                     ru.user_password,
--                     ru.created_at,
--                     ru.updated_at,
--                     ru.deleted_at
--                 FROM
--                     current_record ru
--                 WHERE
--                     ru.user_id IS NOT NULL
--                 GROUP BY
--                     ru.user_id,
--                     ru.user_name,
--                     ru.user_image,
--                     ru.user_email,
--                     ru.user_phone,
--                     ru.user_password,
--                     ru.created_at,
--                     ru.updated_at,
--                     ru.deleted_at) nested_users) users
--         FROM
--             current_record rp
--         GROUP BY
--             rp.role_id
-- )
-- SELECT
--     Row_to_json(role_record)::jsonb role_row
-- FROM (
--     SELECT
--         r.role_id,
--         r.role_name,
--         r.role_description,
--         r.role_created_at,
--         r.role_updated_at,
--         r.role_deleted_at,
--         crp.permissions,
--         cru.users
--     FROM
--         current_record r
--         JOIN current_role_permissions crp ON r.role_id = crp.role_id
--         JOIN current_role_users cru ON r.role_id = cru.role_id
--     GROUP BY
--         r.role_id,
--         r.role_name,
--         r.role_description,
--         r.role_created_at,
--         r.role_updated_at,
--         r.role_deleted_at,
--         crp.permissions,
--         cru.users) role_record;
-- name: RolesList :many
WITH role_record AS (
    SELECT
        role_id,
        role_name,
        role_description,
        created_at,
        deleted_at
    FROM
        accounts_schema.roles
),
permissions_count AS (
    SELECT
        rp.role_id,
        COUNT(rp.permission_id) permissions_count
    FROM
        accounts_schema.role_permissions rp
    GROUP BY
        rp.role_id
),
users_count AS (
    SELECT
        ur.role_id,
        COUNT(ur.user_id) users_count
    FROM
        accounts_schema.user_roles ur
    GROUP BY
        ur.role_id
)
SELECT
    r.role_id,
    r.role_name,
    r.role_description,
    COALESCE(pc.permissions_count, 0),
    COALESCE(uc.users_count, 0),
    r.created_at,
    r.deleted_at
FROM
    role_record r
    LEFT JOIN permissions_count pc ON r.role_id = pc.role_id
    LEFT JOIN users_count uc ON r.role_id = uc.role_id;

-- name: RolesInputList :many
SELECT
    role_id,
    role_name
FROM
    accounts_schema.roles
WHERE
    deleted_at IS NULL;

