-- name: SettingsUpdate :one
SELECT
FROM
    settings_bulk_create(sqlc.arg('keys')::text[], sqlc.arg('values')::text[]);

-- name: SettingsFindForUpdate :many
SELECT
    setting_value,
    setting_key,
    setting_type
FROM
    settings s
    JOIN setting_types t ON t.setting_type_id = s.setting_type_id;

 


-- name: IconsInputList :many
select 
    icon_id ,
    icon_name ,
    icon_content
   FROM 
 icons  ;



-- name: TaxTemplatesList :many
select 
tt.tax_template_id,
tt.tax_type_id,
tt.rate,
ta.code,
ta.desc_en,
ta.desc_ar
 from tax_template_tax_types tt join tax_types ta on tt.tax_type_id = ta.tax_type_id;