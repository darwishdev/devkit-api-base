
CREATE TABLE dim_date
(
  date_id              SERIAl primary key,
  next_day_id int,
  FOREIGN KEY (next_day_id) REFERENCES dim_date(date_id),
  date_actual              DATE NOT NULL,
  next_day_actual              DATE ,
  epoch                    BIGINT NOT NULL,
  day_suffix               VARCHAR(4) NOT NULL,
  day_name                 VARCHAR(9) NOT NULL,
  day_of_week              INT NOT NULL,
  day_of_month             INT NOT NULL,
  day_of_quarter           INT NOT NULL,
  day_of_year              INT NOT NULL,
  week_of_month            INT NOT NULL,
  week_of_year             INT NOT NULL,
  week_of_year_iso         CHAR(10) NOT NULL,
  month_actual             INT NOT NULL,
  month_name               VARCHAR(9) NOT NULL,
  month_name_abbreviated   CHAR(3) NOT NULL,
  quarter_actual           INT NOT NULL,
  quarter_name             VARCHAR(9) NOT NULL,
  year_actual              INT NOT NULL,
  first_day_of_week        DATE NOT NULL,
  last_day_of_week         DATE NOT NULL,
  first_day_of_month       DATE NOT NULL,
  last_day_of_month        DATE NOT NULL,
  first_day_of_quarter     DATE NOT NULL,
  last_day_of_quarter      DATE NOT NULL,
  first_day_of_year        DATE NOT NULL,
  last_day_of_year         DATE NOT NULL,
  mmyyyy                   CHAR(6) NOT NULL,
  mmddyyyy                 CHAR(10) NOT NULL,
  weekend_indr             BOOLEAN NOT NULL
);


CREATE TABLE tags(
    tag varchar NOT NULL UNIQUE
);
 
CREATE TABLE setting_types(
    setting_type_id serial PRIMARY KEY,
    setting_type varchar(20) NOT NULL UNIQUE
);

CREATE TABLE settings(
    setting_id serial PRIMARY KEY,
    setting_type_id int NOT NULL,
    setting_key varchar(100) NOT NULL UNIQUE,
    setting_value text NOT NULL,
    updated_at timestamp

); 
CREATE TABLE tax_types(
    tax_type_id serial PRIMARY KEY,
    code varchar NOT NULL unique,
    desc_en varchar NOT NULL unique,
    desc_ar varchar NOT NULL unique
);  

CREATE TABLE tax_templates(
    tax_template_id serial PRIMARY KEY,
    tax_template varchar NOT NULL unique ,
    is_applied boolean not null DEFAULT false,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp
);  

CREATE TABLE tax_template_tax_types(
    tax_template_id int not null,
    FOREIGN KEY (tax_template_id) REFERENCES tax_templates(tax_template_id),
    tax_type_id int not null,
    FOREIGN KEY (tax_type_id) REFERENCES tax_types(tax_type_id),
    rate real NOT NULL 
);  
CREATE TABLE icons(
    icon_id serial PRIMARY KEY,
    icon_name varchar(200) NOT NULL UNIQUE,
    icon_content text  NOT NULL
); 
ALTER TABLE settings
    ADD FOREIGN KEY (setting_type_id) REFERENCES setting_types(setting_type_id);
 

CREATE INDEX d_date_date_actual_idx
  ON dim_date(date_actual);