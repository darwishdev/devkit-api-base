BEGIN;


SELECT plan(1);
 
-- Test 1: Insert new property with all required fields
    -- select function_returns('properties_schema', 'property_create_update', Array['Test Property', '123 Main St', 'https://example.com/property', 1 , true], 'uuid');
 select ok((select * from properties_schema.property_create_update(
	 property_name := 'Test Property',
  address_line := '123 Main St',
  location_url := 'https://example.com/property',
  property_type_id := 1,
  is_breakfast_available := TRUE
 )) is not null);
-- Finish tests and clean up (if needed)
SELECT * FROM finish();
ROLLBACK;