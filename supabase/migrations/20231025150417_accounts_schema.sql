CREATE SCHEMA accounts_schema;
 
CREATE TABLE accounts_schema.permissions(
    permission_id serial PRIMARY KEY,
    permission_function varchar(200) NOT NULL UNIQUE,
    permission_name varchar(200) NOT NULL,
    permission_description varchar(200),
    permission_group varchar(200) NOT NULL
);

CREATE TABLE accounts_schema.roles(
    role_id serial PRIMARY KEY,
    role_name varchar(200) NOT NULL UNIQUE,
    role_description varchar(200),
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp 
);

CREATE TABLE accounts_schema.role_permissions(
    role_id int NOT NULL,
    permission_id int NOT NULL,
    PRIMARY KEY (role_id, permission_id)
);

CREATE TABLE accounts_schema.users(
    user_id serial PRIMARY KEY,
    user_name varchar(200) NOT NULL,
    user_code varchar(20) UNIQUE NOT NULL,
    user_image varchar(200),
    user_email varchar(200) UNIQUE NOT NULL,
    user_phone varchar(200) UNIQUE,
    user_password varchar(200) NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE accounts_schema.user_roles(
    user_id int NOT NULL,
    role_id int NOT NULL,
    PRIMARY KEY (user_id, role_id)
);

CREATE TABLE accounts_schema.user_permissions(
    user_permission_id serial PRIMARY KEY,
    user_id int NOT NULL,
    permission_id int NOT NULL
);
   
CREATE TABLE accounts_schema.navigation_bars(
    navigation_bar_id serial PRIMARY KEY,
    menu_key varchar(200) UNIQUE NOT NULL,
    label varchar(200) NOT NULL,
    label_ar varchar(200),
    icon_id int NOT NULL,
    "route" varchar(200) UNIQUE,
    parent_id int,
    permission_id int
);

CREATE TABLE accounts_schema.owners(
    owner_id serial PRIMARY KEY,
    owner_name varchar(200) NOT NULL,
    owner_image varchar(200),
    owner_email varchar(200) UNIQUE NOT NULL,
    owner_phone varchar(200) UNIQUE,
    owner_password varchar(200) NOT NULL,
    owner_national_id varchar(30) NOT NULL UNIQUE,
    representative_owner_id int,
    FOREIGN KEY (representative_owner_id) REFERENCES accounts_schema.owners(owner_id),
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE accounts_schema.customers(
    customer_id serial PRIMARY KEY,
    customer_code varchar(20) UNIQUE NOT NULL,
    customer_name varchar(200) NOT NULL,
    customer_image varchar(200),
    customer_email varchar(200) NOT NULL UNIQUE,
    customer_phone varchar(200) UNIQUE,
    customer_password varchar(200) NOT NULL,
    birthdate date,
    customer_national_id varchar(30)  UNIQUE,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp
);
-- Alter tables within the users schema
ALTER TABLE accounts_schema.role_permissions
    ADD FOREIGN KEY (role_id) REFERENCES accounts_schema.roles(role_id),
    ADD FOREIGN KEY (permission_id) REFERENCES accounts_schema.permissions(permission_id);

ALTER TABLE accounts_schema.user_roles
    ADD FOREIGN KEY (user_id) REFERENCES accounts_schema.users(user_id),
    ADD FOREIGN KEY (role_id) REFERENCES accounts_schema.roles(role_id);

ALTER TABLE accounts_schema.user_permissions
    ADD FOREIGN KEY (user_id) REFERENCES accounts_schema.users(user_id),
    ADD FOREIGN KEY (permission_id) REFERENCES accounts_schema.permissions(permission_id);

ALTER TABLE accounts_schema.navigation_bars
    ADD FOREIGN KEY (parent_id) REFERENCES accounts_schema.navigation_bars(navigation_bar_id),
    ADD FOREIGN KEY (icon_id) REFERENCES icons(icon_id),
    ADD FOREIGN KEY (permission_id) REFERENCES accounts_schema.permissions(permission_id);
 