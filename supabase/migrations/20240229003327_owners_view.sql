

CREATE OR REPLACE VIEW accounts_schema.owners_view AS
with properties as (
  select 
    property_id,
    property_name,
    address_line,
    star_rating,
    location_id,
    location_name,
    location_image,
    location_url,
    property_type_id,
    property_type_name,
    property_type_icon,
    compound_id,
    owner_id,
    owner_name,
    compound_name,
    property_images,
    property_description, 
    checkin_time_from,
    checkin_time_to,
    checkout_time_from,
    checkout_time_to,
    created_at,
    updated_at,
    deleted_at
  from properties_schema.properties_view
  where deleted_at is null
)  
select 
o.owner_id,
o.owner_name,
o.owner_image,
o.owner_email,
o.owner_phone,
o.owner_password,
o.owner_national_id,
o.created_at,
o.updated_at,
o.deleted_at,
    (
      SELECT 
        Jsonb_agg(nested_properties) 
      FROM 
      (
        SELECT 
          p.*
        FROM 
          properties p
        WHERE 
          o.owner_id = p.owner_id
      ) nested_properties
           
  ) AS  properties
   FROM  accounts_schema.owners o


