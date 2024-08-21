
CREATE OR REPLACE FUNCTION IIF(condition boolean, true_result ANYELEMENT, false_result ANYELEMENT)
    RETURNS ANYELEMENT
    LANGUAGE plpgsql
    AS $$
BEGIN
    IF condition THEN
        RETURN true_result;
    ELSE
        RETURN false_result;
    END IF;
END
$$; 
 CREATE OR REPLACE FUNCTION IsNull(in_value ANYELEMENT)
    RETURNS boolean
    LANGUAGE plpgsql
AS $$
    declare value_type varchar(30);
BEGIN

    IF in_value IS NULL THEN 
        RETURN TRUE;
    END IF;
    select pg_typeof(in_value) into value_type;
    IF value_type = 'character varying' OR value_type = 'text'  THEN
        IF in_value = '' OR in_value IS NULL THEN
            RETURN true;
        ELSE
            RETURN false;
        END IF;
    ELSIF value_type = 'integer' OR  value_type = 'real'  THEN
        IF in_value = 0 OR in_value IS NULL THEN
            RETURN true;
        ELSE
            RETURN false;
        END IF;

    ELSEIF value_type LIKE '%[]' THEN
        IF array_length(in_value , 1) IS NULL THEN
            RETURN true;
        ELSE
            RETURN false;
        END IF;
    ELSIF value_type = 'time with time zone' 
    OR  value_type = 'time without time zone'  
    OR  value_type = 'timestamp with time zone'  
    OR  value_type = 'timestamp without time zone'  
    OR  value_type = 'boolean'  
    OR  value_type = 'boolean'  
    THEN
        IF in_value IS NULL THEN
            RETURN true;
        ELSE
            RETURN false;
        END IF;
    ELSE
        -- Handle other data types if needed
        RAISE EXCEPTION 'Unsupported data type: %', pg_typeof(in_value);
    END IF;
END
$$;

CREATE OR REPLACE FUNCTION IsNullReplace(in_value ANYELEMENT , in_target_value ANYELEMENT)
    RETURNS ANYELEMENT
    LANGUAGE plpgsql
AS $$
    declare value_type varchar(30);
BEGIN
    IF IsNull(in_value) then 
      return in_target_value;
    ELSE
        RETURN in_value;
    END IF;
END
$$;

 
 
CREATE OR REPLACE FUNCTION settings_bulk_create(keys text[], vals text[])
    RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
    -- Create a temporary table to hold the new values
    CREATE TEMP TABLE temp_settings AS
    SELECT
        unnest($1) AS setting_key,
        unnest($2) AS setting_value;
    -- Update the main table based on the temporary table
    UPDATE
        settings AS s
    SET
        setting_value = t.setting_value
    FROM
        temp_settings AS t
    WHERE
        s.setting_key = t.setting_key;
    -- Drop the temporary table
    DROP TABLE temp_settings;
END
$$; 
CREATE OR REPLACE FUNCTION random_between(min_value int, max_value int)
    RETURNS int
    AS $$
DECLARE
    v_random_value int;
BEGIN
    SELECT
        floor(random() *(max_value - min_value + 1)) + min_value INTO v_random_value;
    RETURN v_random_value;
END;
$$
LANGUAGE plpgsql;
 
CREATE OR REPLACE FUNCTION random_between(min_value int, max_value int)
    RETURNS int
    AS $$
DECLARE
    v_random_value int;
BEGIN
    SELECT
        floor(random() *(max_value - min_value + 1)) + min_value INTO v_random_value;
    RETURN v_random_value;
END;
$$
LANGUAGE plpgsql;



CREATE OR REPLACE FUNCTION dates_seed(in_start_date varchar)
    RETURNS void
    LANGUAGE plpgsql
    AS $$
BEGIN
  

INSERT INTO dim_date
SELECT TO_CHAR(datum, 'yyyymmdd')::INT AS date_dim_id,
        null,
       datum AS date_actual,
       null,
       EXTRACT(EPOCH FROM datum) AS epoch,
       TO_CHAR(datum, 'fmDDth') AS day_suffix,
       TO_CHAR(datum, 'TMDay') AS day_name,
       EXTRACT(ISODOW FROM datum) AS day_of_week,
       EXTRACT(DAY FROM datum) AS day_of_month,
       datum - DATE_TRUNC('quarter', datum)::DATE + 1 AS day_of_quarter,
       EXTRACT(DOY FROM datum) AS day_of_year,
       TO_CHAR(datum, 'W')::INT AS week_of_month,
       EXTRACT(WEEK FROM datum) AS week_of_year,
       EXTRACT(ISOYEAR FROM datum) || TO_CHAR(datum, '"-W"IW-') || EXTRACT(ISODOW FROM datum) AS week_of_year_iso,
       EXTRACT(MONTH FROM datum) AS month_actual,
       TO_CHAR(datum, 'TMMonth') AS month_name,
       TO_CHAR(datum, 'Mon') AS month_name_abbreviated,
       EXTRACT(QUARTER FROM datum) AS quarter_actual,
       CASE
           WHEN EXTRACT(QUARTER FROM datum) = 1 THEN 'First'
           WHEN EXTRACT(QUARTER FROM datum) = 2 THEN 'Second'
           WHEN EXTRACT(QUARTER FROM datum) = 3 THEN 'Third'
           WHEN EXTRACT(QUARTER FROM datum) = 4 THEN 'Fourth'
           END AS quarter_name,
       EXTRACT(YEAR FROM datum) AS year_actual,
       datum + (1 - EXTRACT(ISODOW FROM datum))::INT AS first_day_of_week,
       datum + (7 - EXTRACT(ISODOW FROM datum))::INT AS last_day_of_week,
       datum + (1 - EXTRACT(DAY FROM datum))::INT AS first_day_of_month,
       (DATE_TRUNC('MONTH', datum) + INTERVAL '1 MONTH - 1 day')::DATE AS last_day_of_month,
       DATE_TRUNC('quarter', datum)::DATE AS first_day_of_quarter,
       (DATE_TRUNC('quarter', datum) + INTERVAL '3 MONTH - 1 day')::DATE AS last_day_of_quarter,
       TO_DATE(EXTRACT(YEAR FROM datum) || '-01-01', 'YYYY-MM-DD') AS first_day_of_year,
       TO_DATE(EXTRACT(YEAR FROM datum) || '-12-31', 'YYYY-MM-DD') AS last_day_of_year,
       TO_CHAR(datum, 'mmyyyy') AS mmyyyy,
       TO_CHAR(datum, 'mmddyyyy') AS mmddyyyy,
       CASE
           WHEN EXTRACT(ISODOW FROM datum) IN (6, 7) THEN TRUE
           ELSE FALSE
           END AS weekend_indr
      FROM (SELECT in_start_date::DATE + SEQUENCE.DAY AS datum
      FROM GENERATE_SERIES(0, 365) AS SEQUENCE (DAY)
      GROUP BY SEQUENCE.DAY) DQ
ORDER BY 1;



 UPDATE dim_date a
    SET next_day_id = b.date_id,
     next_day_actual = b.date_actual
    FROM dim_date b
    WHERE a.date_actual = b.date_actual - INTERVAL '1 day';

END
$$;








CREATE OR REPLACE FUNCTION toggle_delete_restore(
    p_table_name VARCHAR,
    p_id_column_name VARCHAR,
    p_records INT[]
)

    RETURNS table (
      record_id int
    )
    AS $$

    declare v_query varchar;
BEGIN
    -- Construct the dynamic query to update the deleted_at field
    v_query := concat('update ' , p_table_name , ' set deleted_at = IIF(deleted_at IS NULL, now(), NULL) where ' , p_id_column_name , ' = any($1)');

    -- Execute the dynamic query and return the affected records
    EXECUTE v_query using p_records;
    return query select  unnest(p_records) record_id;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION generate_code_series(
    in_table_name VARCHAR,
    in_column_name VARCHAR,
    in_start_series INT
)

    RETURNS varchar
    AS $$

    declare 
    v_query varchar;
    v_result_code varchar;
BEGIN
    -- Construct the dynamic query to update the deleted_at field
    v_query := concat('select (IsNullReplace(max(' , in_column_name , ')::varchar,' , in_start_series , '::varchar)::int + 1)::varchar code  from ' , in_table_name ); 
    EXECUTE v_query into v_result_code;
    return  v_result_code;
END;
$$ LANGUAGE plpgsql;



CREATE OR REPLACE FUNCTION get_table_code(
    in_table_name VARCHAR
)
    RETURNS varchar
    AS $$
    declare 
    v_table_name varchar;
    v_column_name varchar;
    v_start_series int;
BEGIN
drop table if exists code_supported_tables;  
create temp table  code_supported_tables (
    table_name varchar,
    actual_table_name varchar,
    column_name varchar,
  start_series int
);

insert into code_supported_tables (
  table_name,
  actual_table_name,
  column_name,
  start_series
) values (
  'users',
  'accounts_schema.users',
  'user_code',
  1000
), (
  'customers',
  'accounts_schema.customers',
  'customer_code',
  2000
),  (
  'payments',
  'reservations_schema.payments',
  'payment_code',
  1000
),
(
  'reservations',
  'reservations_schema.reservations',
  'reservation_code',
  1000
);


  
   select t.actual_table_name , t.column_name, t.start_series
   into v_table_name , v_column_name , v_start_series
    from code_supported_tables t where t.table_name = in_table_name;
 
  IF v_table_name IS NULL then 
    RAISE EXCEPTION 'unsupported_table';
  END IF;
  return generate_code_series(v_table_name , v_column_name , v_start_series);
 
END;
$$ LANGUAGE plpgsql;





CREATE OR REPLACE FUNCTION get_date_tense(
    in_date_from int,
    in_date_to int
)
    RETURNS varchar
    AS $$
    declare 
    v_tense varchar;
    v_current_date int;
BEGIN 

-- define the variables
v_current_date := TO_NUMBER(TO_CHAR(CURRENT_DATE, 'YYYYMMDD'), '99999999')::int ;
IF  v_current_date between in_date_from and in_date_to then
  RETURN 'present';
ELSEIF v_current_date < in_date_from then
    RETURN 'future';
ELSE 
    RETURN 'past';
end if;
END;
$$ LANGUAGE plpgsql;



CREATE OR REPLACE FUNCTION reservations_schema.get_reservation_key(
    in_invetory_ids int[],
    in_persons int 
)
    RETURNS varchar
    AS $$
BEGIN
    return  concat(array_to_string(in_invetory_ids, ',') ,'-' ,in_persons::varchar) reservation_key;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION date_diffrence(
    in_date_from_id int,
    in_date_to_id int
)
    RETURNS INT
    AS $$
    declare v_passed_days_count int;
BEGIN
    SELECT  d1.date_actual - d2.date_actual   into v_passed_days_count from dim_date d1 join dim_date d2 on d2.date_id = in_date_from_id where d1.date_id = in_date_to_id ;
    return v_passed_days_count;
END;
$$ LANGUAGE plpgsql;
