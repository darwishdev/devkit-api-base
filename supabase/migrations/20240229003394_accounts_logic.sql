CREATE OR REPLACE FUNCTION accounts_schema.user_find_permissions(in_user_id int)
    RETURNS table(
        permission_id int,
        permission_function varchar(200),
        permission_group varchar(200),
        permission_name varchar(200)
    )
    LANGUAGE plpgsql
    AS $$
BEGIN
   return query (
        SELECT 
            rp.permission_id ,
            p.permission_function,
            p.permission_group,
            p.permission_name
        FROM accounts_schema.user_roles ur 
            JOIN accounts_schema.role_permissions rp ON rp.role_id = ur.role_id 
            JOIN accounts_schema.permissions p on p.permission_id = rp.permission_id
            and ur.user_id = in_user_id
        UNION distinct
        SELECT 
            up.permission_id ,
            p.permission_function,
            p.permission_group,
            p.permission_name
        FROM 
            accounts_schema.user_permissions up
            JOIN accounts_schema.permissions p on p.permission_id = up.permission_id
        WHERE
            user_id = in_user_id
        
        order by permission_id
   );
END
$$; 



CREATE OR REPLACE FUNCTION accounts_schema.user_create_update(
    in_user_id int,
    in_user_name varchar(200),
    in_user_image varchar(200),
    in_user_email varchar(200),
    in_user_phone varchar(200),
    in_user_password varchar(200),
    in_permissions int[],
    in_roles int[]
)
    RETURNS setof accounts_schema.users
    LANGUAGE plpgsql
    AS $$
    declare v_user_id int;
BEGIN
    IF IsNull(in_user_id) THEN
        INSERT INTO accounts_schema.users (
            user_name,
            user_code,
            user_image,
            user_email,
            user_phone,
            user_password
        ) VALUES (
            in_user_name,
            get_table_code('users'),
            in_user_image,
            in_user_email,
            in_user_phone,
            in_user_password
        ) RETURNING user_id INTO v_user_id;


        INSERT INTO accounts_schema.user_roles (
            user_id,
            role_id
        ) select v_user_id , unnest(in_roles);
        INSERT INTO accounts_schema.user_permissions (
            user_id,
            permission_id
        ) select v_user_id , unnest(in_permissions);
    ELSE
        v_user_id := in_user_id;
        UPDATE
            accounts_schema.users
        SET
            user_name = IsNullReplace(in_user_name , user_name),
            user_code = IsNullReplace(get_table_code('users') , user_code),
            user_image = IsNullReplace(in_user_image , user_image),
            user_email = IsNullReplace(in_user_email , user_email),
            user_phone = IsNullReplace(in_user_phone , user_phone),
            user_password = IsNullReplace(in_user_password ,user_password )
        WHERE
            user_id = v_user_id;


        DELETE FROM accounts_schema.user_roles where user_id = v_user_id;
        DELETE FROM accounts_schema.user_permissions where user_id = v_user_id;
        INSERT INTO accounts_schema.user_roles (
            user_id,
            role_id
        ) select v_user_id , unnest(in_roles);
        INSERT INTO accounts_schema.user_permissions (
            user_id,
            permission_id
        ) select v_user_id , unnest(in_permissions);
    END IF;



    return query 
        select * from accounts_schema.users where user_id = v_user_id;
     
END
$$; 








CREATE OR REPLACE FUNCTION accounts_schema.role_create_update(
    in_role_id int,
    in_role_name varchar(200),
    in_role_description varchar(200),
    in_permissions int[]
)
    RETURNS setof accounts_schema.roles
    LANGUAGE plpgsql
    AS $$
    declare v_role_id int;
BEGIN
    IF IsNull(in_role_id) THEN
        INSERT INTO accounts_schema.roles (
            role_name,
            role_description 
        ) VALUES (
            in_role_name,
            in_role_description
        ) RETURNING role_id INTO v_role_id;
    ELSE
        UPDATE
            accounts_schema.roles
        SET
            role_name = in_role_name,
            role_description = in_role_description
        WHERE
            role_id = v_role_id;


        DELETE FROM accounts_schema.role_permissions where role_id = v_role_id;
        
        INSERT INTO accounts_schema.role_permissions (
            role_id,
            permission_id
        ) select v_role_id , unnest(in_permissions);
    END IF;


    return query  
        select * from accounts_schema.users where user_id = v_user_id;
      
END
$$; 

CREATE OR REPLACE FUNCTION accounts_schema.customer_create_update(
    in_customer_id int,
    in_customer_name varchar(200),
    in_customer_image varchar(200),
    in_customer_email varchar(200),
    in_birthdate varchar(200),
    in_customer_phone varchar(200),
    in_customer_password varchar(200),
    in_customer_national_id varchar(30) = null
)
    RETURNS setof accounts_schema.customers
    LANGUAGE plpgsql
    AS $$
    declare v_customer_id int;
BEGIN
    IF IsNull(in_customer_id) THEN
        INSERT INTO accounts_schema.customers (
            customer_name,
            customer_code,
            customer_image,
            birthdate,
            customer_email,
            customer_phone,
            customer_password,
            customer_national_id
        ) VALUES (
            in_customer_name,
            get_table_code('customers'),
            in_customer_image,
            TO_DATE(in_birthdate::text , 'YYYYMMDD'),
            in_customer_email,
            in_customer_phone,
            in_customer_password,
            IsNullReplace(in_customer_national_id , null)
        ) RETURNING customer_id INTO v_customer_id;
    ELSE
        v_customer_id := in_customer_id;
        UPDATE
            accounts_schema.customers
        SET
            customer_name = IsNullReplace(in_customer_name , customer_name),
            customer_code = IsNullReplace(get_table_code('customers') , customer_code),
            customer_image = IsNullReplace(in_customer_image , customer_image),
            birthdate =TO_DATE(in_birthdate::text , 'YYYYMMDD'),
            customer_email = IsNullReplace(in_customer_email , customer_email),
            customer_phone = IsNullReplace(in_customer_phone , customer_phone),
            customer_password = IsNullReplace(in_customer_password ,customer_password ),
            customer_national_id = IsNullReplace(in_customer_national_id ,customer_national_id )
        WHERE
            customer_id = v_customer_id;
    END IF;
    return query 
        select * from accounts_schema.customers where customer_id = v_customer_id;
END
$$; 












CREATE OR REPLACE FUNCTION accounts_schema.users_delete_restore(
    in_user_ids int[] 
)
    RETURNS setof accounts_schema.users
    LANGUAGE plpgsql
    AS $$
    declare v_role_id int;
BEGIN
drop table if exists temp_roles;

select role_id into v_role_id from accounts_schema.user_roles where user_id = any(in_user_ids) and role_id = 1;


if v_role_id is not null THEN
 RAISE EXCEPTION 'admin users can not be deleted'; 
end if;
return query
select u.* 
from accounts_schema.users u 
join toggle_delete_restore('accounts_schema.users', 'user_id', in_user_ids) d 
on d.record_id = u.user_id;
     
END
$$; 



