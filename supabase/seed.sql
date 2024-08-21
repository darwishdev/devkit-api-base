

alter database postgres
set timezone to 'Africa/Cairo';
CREATE POLICY "allow all 1ffg0oo_0" ON storage.objects FOR
SELECT
    TO public USING (bucket_id = 'images');

CREATE POLICY "allow all 1ffg0oo_1" ON storage.objects FOR
INSERT
    TO public WITH CHECK (bucket_id = 'images');



INSERT INTO tax_types (code, desc_en, desc_ar ) VALUES 
('T1', 'Value added tax', 'ضريبه القيمه المضافه');

INSERT INTO tax_templates (tax_template, is_applied ) VALUES 
('default', true);

INSERT INTO tax_template_tax_types (
    tax_template_id,
    tax_type_id,
    rate
 ) VALUES 
(1,1 ,15);

INSERT INTO
    properties_schema.cities(city_name, city_image)
values
    ('North Coast', 'images/cities_northcoast.webp'),
    ('Hurghada', 'images/cities_hurghada.webp'),
    (
        'El Ain El Sokhna',
        'images/cities_sharmalsheikh.webp'
    );

INSERT INTO
    properties_schema.locations(location_name, location_image, city_id)
VALUES
    (
        'El Alamein',
        'images/locations_elalamein.webp',
        1
    ),
    (
        'New Alamein',
        'images/locations_newelalamein.webp',
        1
    ),
    (
        'Vilages Road',
        'images/locations_villagesroad.webp',
        2
    ),
    (
        'Al Mamsha El Seyahi',
        'images/locations_mamsha.webp',
        2
    ),
    ('Porto Sokhna', 'images/locations_portoelsokhna.webp', 3),
    ('Telal', 'images/locations_telalelsokhna.webp', 3);

insert into
    properties_schema.property_categories (property_category_name)
values
    ('hotels'),
    ('others');

insert into
    properties_schema.compounds (compound_name)
values
    ('marasi'),
    ('Marina'),
    ('NAC Lagoons' );

insert into
    properties_schema.reservable_unit_types(
        reservable_unit_type_name,
        property_category_id,
        rooms_count
    )
values
    ('Deluxe Family', 1, 1),
    ('Deluxe King', 1, 1),
    ('Premium King', 1, 1),
    ('Deluxe Triple', 1, 1),
    ('Deluxe twin', 1, 1),
    ('Royal Suite', 1, 4),
    ('one bedroom apartment', 2, 1),
    ('two-bedroom apartment', 2, 2),
    ('three-bedroom apartment', 2, 3),
    ('Villa', 2, 5);
INSERT INTO properties_schema.amenity_groups 
( amenity_group_name, property_category_id)
values
( 'General Facilities', 1), 
( 'Room Amenities', 1),
( 'Business Facilities', 1),
( 'General Amenities', 1),
( 'Health and Wellness', 1),
( 'Activities', 1),
( 'Safety and Security', 2),
( 'Cleaning Services', 2);

INSERT INTO
    properties_schema.amenity_value_types (amenity_value_type)
VALUES
    ('unknown'),
    ('toggle'),
    ('number'),
    ('text');

INSERT INTO properties_schema.amenities ( amenity_group_id, amenity_name, amenity_icon, amenity_input_label, amenity_value_type_id) VALUES
( 1, 'Restaurant', 'Restaurant', 'Has Restaurant', 2),
( 1, 'Bar', 'Bar', 'Has Bar', 2),
( 1, 'Breakfast in the Room', 'Breakfast', 'Offers Breakfast in Room', 2),
( 1, 'Room Service', 'Room_Service', 'Has Room Service', 2),

( 1, 'Airport Shuttle', 'Airport', 'Provides Airport Shuttle', 2),
( 1, 'Bicycle Rental', 'Bicycle', 'Provides Bicycle Rental', 2),
( 1, 'Car Hire', 'Car', 'Provides Car Hire', 2),
( 1, '24-hour Front Desk', 'Front_office', 'Has 24-hour Front Desk', 2),
( 1, 'Concierge Service', 'Concierge', 'Has Concierge Service', 2),
( 1, 'Tour Desk', 'Tour', 'Has Tour Desk', 2),
( 1, 'Garden', 'Garden', 'Has Garden', 2),
( 1,'Terrace', 'Terrace', 'Has Terrace', 2),
( 1,'Library', 'Library', 'Has Library', 2),
( 1, 'Babysitting/Child Services', 'Babysitting', 'Provides Babysitting/Child Services', 2),
( 1, 'Childrens Playground', 'Playground', 'Has Childrens Playground', 2),
( 1, 'Kids Club', 'Kids_club', 'Has Kids Club', 2),
( 1, 'Shops (on site)', 'Shops', 'Has Shops On Site', 2),
( 1, 'Gift Shop', 'Gift', 'Has Gift Shop', 2),

( 2, 'Air Conditioning', 'ac', 'Has Air Conditioning', 2),
( 2, 'Heating', 'Heating', 'Has Heating', 2),
( 2, 'Designated Smoking Area', 'Smoking', 'Has Designated Smoking Area', 2),
( 2, 'Flat-screen TV', 'TV', 'Has Flat-screen TV', 2),
( 2, 'Cable Channels', 'Cable', 'Has Cable Channels', 2),
( 2, 'Satellite Channels', 'Satellite', 'Has Satellite Channels', 2),
( 2, 'Minibar', 'Minibar', 'Has Minibar', 2),
( 2, 'Hairdryer', 'Hairdryer', 'Has Hairdryer', 2),
( 2, 'Iron', 'Iron', 'Has Iron', 2),
( 2, 'Balcony', 'Balcony', 'Has Balcony', 2),
( 2, 'Patio', 'Patio', 'Has Patio', 2),
( 2, 'Sea View', 'Sea', 'Has Sea View', 2),
( 3, 'Meeting/Banquet Facilities', 'Meeting', 'Has Meeting/Banquet Facilities', 2),
( 3, 'Business Centre', 'Business', 'Has Business Centre', 2),
( 3, 'Fax/Photocopying', 'camera', 'Has Fax/Photocopying Service', 2),
( 4, 'Non-smoking Rooms', 'Non-smoking', 'Has Non-smoking Rooms', 2),
( 4, 'Facilities for Disabled Guests', 'Disabled', 'Has Facilities for Disabled Guests', 2),
( 4, 'Family Rooms', 'Family', 'Has Family Rooms', 2),
( 4, 'Lift', 'Lift', 'Has Lift', 2),
( 5, 'Sauna', 'Sauna', 'Has Sauna', 2),
( 5, 'Spa and Wellness Centre', 'Spa', 'Has Spa and Wellness Centre', 2),
( 5, 'Massage', 'Massage', 'Has Massage Service', 2),
( 5, 'Hot Tub/Jacuzzi', 'Jacuzzi', 'Has Hot Tub/Jacuzzi', 2),
( 6, 'Fitness Centre', 'Fitness', 'Has Fitness Centre', 2),
( 6, 'Golf Course (within 3 km)', 'Golf', 'Has Nearby Golf Course', 2),
( 6, 'Fishing', 'Fishing', 'Has Fishing Facility', 2),
( 6, 'Skiing', 'Skiing', 'Has Skiing Facility', 2),
( 7, 'Safety Deposit Box', 'Safety', 'Has Safety Deposit Box', 2),
( 7, '24-hour Security', 'Security', 'Has 24-hour Security', 2),
( 7, 'CCTV in Common Areas', 'CCTV', 'Has CCTV in Common Areas', 2),
( 7, 'Fire Extinguishers', 'Fire_ext', 'Has Fire Extinguishers', 2),
( 8, 'Laundry', 'Laundry', 'Has Laundry Service', 2),
( 8, 'Dry Cleaning', 'Dry_clean', 'Has Dry Cleaning Service', 2),
( 8, 'Ironing Service', 'Ironing', 'Has Ironing Service', 2);


INSERT INTO 
properties_schema.unit_amenities 
(unit_id, unit_type, amenity_id, amenity_value)
VALUES
  -- Unit 1 (includes duplicates and non-boolean value)
  (1, 'property', 23, TRUE),
  (1, 'property', 16, TRUE),
  (1, 'property', 18, TRUE),
  (1, 'property', 32, TRUE),
  (1, 'property',  2, TRUE),
  (2, 'property', 10, TRUE),
  (2, 'property',  7, TRUE),
  (2, 'property', 21, TRUE),
  (2, 'property', 35, TRUE),
  (2, 'property', 12, TRUE),
  (3, 'property', 29, TRUE),
  (3, 'property',  5, TRUE),
  (3, 'property', 19, TRUE),
  (3, 'property', 37, TRUE),
  (3, 'property', 30, TRUE),

  -- Unit 7 (includes duplicates and non-boolean value)
  (4, 'property',  1, TRUE),
  (4, 'property', 11, TRUE),
  (4, 'property', 22, TRUE),
  (4, 'property',  8, TRUE),
  (4, 'property', 20, TRUE),
  (5, 'property', 27, TRUE),
  (5, 'property',  4, TRUE),
  (5, 'property', 14, TRUE),
  (5, 'property', 24, TRUE),
  (5, 'property', 31, TRUE),
  (5, 'property',  3, TRUE),
  (6, 'property', 17, TRUE),
  (6, 'property', 26, TRUE),
  (6, 'property',  9, TRUE),
  (6, 'property',  6, TRUE),
  (6, 'property', 13, TRUE),
  (6, 'property', 25, TRUE),
  (6, 'property', 15, TRUE),
  (7, 'property', 34, TRUE),
  (7, 'property', 38, TRUE),

  -- Additional data for unit 1 and 7 (non-boolean)
  (7, 'property', 3, FALSE),
  (7, 'property', 28, FALSE),

  -- Additional units with random amenities (adjust count as needed)
  (7, 'property', 17, TRUE),
  (7, 'property', 25, TRUE),
  (7, 'property', 9, TRUE),
  (7, 'property', 33, TRUE),
  (1, 'reservable_unit', 45, TRUE),
  (1, 'reservable_unit', 48, TRUE),
  (1, 'reservable_unit', 43, TRUE),
  (1, 'reservable_unit', 40, TRUE),
  (1, 'reservable_unit', 42, TRUE),
  (1, 'reservable_unit', 47, TRUE),
  (2, 'reservable_unit', 41, TRUE),
  (2, 'reservable_unit', 39, TRUE),
  (2, 'reservable_unit', 44, TRUE),
  (2, 'reservable_unit', 49, TRUE),
  (2, 'reservable_unit', 46, TRUE),
  (3, 'reservable_unit', 42, TRUE),
  (3, 'reservable_unit', 40, TRUE),
  (3, 'reservable_unit', 47, TRUE),
  (3, 'reservable_unit', 45, TRUE),
  (4, 'reservable_unit', 49, TRUE),
  (4, 'reservable_unit', 43, TRUE),
  (4, 'reservable_unit', 41, TRUE),
  (4, 'reservable_unit', 39, TRUE),
  (4, 'reservable_unit', 40, TRUE),
  (5, 'reservable_unit', 48, TRUE),
  (5, 'reservable_unit', 44, TRUE),
  (5, 'reservable_unit', 42, TRUE),
  (5, 'reservable_unit', 46, TRUE),
  (5, 'reservable_unit', 41, TRUE),
  (6, 'reservable_unit', 47, TRUE),
  (6, 'reservable_unit', 45, TRUE),
  (6, 'reservable_unit', 40, TRUE),
  (6, 'reservable_unit', 39, TRUE),
  (6, 'reservable_unit', 43, TRUE),
  (7, 'reservable_unit', 49, TRUE),
  (7, 'reservable_unit', 41, TRUE),
  (7, 'reservable_unit', 44, TRUE),
  (7, 'reservable_unit', 42, TRUE),
  (7, 'reservable_unit', 48, TRUE),
  (8, 'reservable_unit', 46, TRUE),
  (8, 'reservable_unit', 40, TRUE),
  (8, 'reservable_unit', 45, TRUE),
  (8, 'reservable_unit', 47, TRUE),
  (8, 'reservable_unit', 39, TRUE),
  (9, 'reservable_unit', 44, TRUE),
  (9, 'reservable_unit', 42, TRUE),
  (9, 'reservable_unit', 41, TRUE),
  (9, 'reservable_unit', 49, TRUE),
  (9, 'reservable_unit', 48, TRUE),
  (10, 'reservable_unit', 40, TRUE),
  (10, 'reservable_unit', 43, TRUE),
  (10, 'reservable_unit', 47, TRUE),
  (10, 'reservable_unit', 45, TRUE),
  (10, 'reservable_unit', 46, TRUE),
  (11, 'reservable_unit', 41, TRUE),
  (11, 'reservable_unit', 49, TRUE),
  (11, 'reservable_unit', 42, TRUE),
  (12, 'reservable_unit', 46, TRUE),
  (12, 'reservable_unit', 40, TRUE),
  (12, 'reservable_unit', 45, TRUE),
  (12, 'reservable_unit', 47, TRUE),
  (12, 'reservable_unit', 39, TRUE),
  (13, 'reservable_unit', 49, TRUE),
  (13, 'reservable_unit', 41, TRUE),
  (13, 'reservable_unit', 44, TRUE),
  (13, 'reservable_unit', 42, TRUE),
  (14, 'reservable_unit', 49, TRUE),
  (14, 'reservable_unit', 43, TRUE),
  (14, 'reservable_unit', 41, TRUE),
  (14, 'reservable_unit', 39, TRUE),
   (15, 'reservable_unit', 48, TRUE),
  (15, 'reservable_unit', 44, TRUE),
  (15, 'reservable_unit', 42, TRUE),
  (15, 'reservable_unit', 46, TRUE),
  (15, 'reservable_unit', 41, TRUE);

INSERT INTO
    properties_schema.bed_types (bed_type, bed_length, bed_width)
VALUES
    ('extra-large double bed', 80, 76),
    ('single bed', 75, 30);

INSERT INTO
    properties_schema.property_types(property_type_name, property_category_id)
VALUES
    ('hotel', 1),
    ('motel', 1),
    ('inn', 1),
    ('chalet', 2),
    ('apartment', 2),
    ('tent', 2),
    ('villa', 2);

INSERT INTO
    reservations_schema.reservation_statuses(reservation_status, reservation_status_color)
VALUES
    ('pending', 'grey'),
    ('awaiting payment', 'orange'),
    ('pending payment', 'purple'),
    ('guaranteed', 'green'),
    ('canceled', 'red'),
    ('refunded', 'pink'),
    ('declined', 'yellow');




INSERT INTO
    reservations_schema.payment_statuses(payment_status, payment_status_color)
VALUES
    ('pending', 'grey'),
    ('confirmed', 'green'),
    ('canceled', 'red'),
    ('refunded', 'pink'),
    ('declined', 'yellow');




INSERT INTO
    reservations_schema.payment_methods(payment_method , payment_method_template , payment_method_text)
VALUES
    ('vodafone cash' , '<strong>01022052546</strong>' , '01022052546' ),
    ('instapay' , 'use this link  <a href="https://ipn.eg/S/darwishdev/instapay/4hekPv"> Link </a> to create instapay transaction' , 'https://ipn.eg/S/darwishdev/instapay/4hekPv' ),
    ('pay on arrival' , '' , '');

select
    dates_seed('2024-01-01');

INSERT INTO
    accounts_schema.owners(
        owner_name,
        owner_image,
        owner_email,
        owner_password,
        owner_phone,
        owner_national_id,
        representative_owner_id
    )
VALUES
    ('Mr mohamed ahmed','/images/avatar1.webp','ahmed.ashraf.devv@gmail.com','123456','012312312322','29501023201952', null),
    ('Mr Waleed Mohamed','/images/avatar2.webp','ahmad.ashraf.dev@gmail.com','123456','012312312321','29501023201951' ,null),
    ('Mr Kareem Ayman','/images/avatar3.webp','workwithmelon@gmail.com','123456','012312312323','29501023201953',  1),
    ('Mohamed El Gendy','/images/avatar3.webp','m.rashad.n@gmail.com','123456','01064453990','28212251701795',  null);
-- Insert multiple cancelation Policies in one statement
INSERT INTO
    rates_schema.cancelation_policies (
        cancelation_policy_name,
        days_before_cancel,
        is_cancelation_fee_percent,
        cancelation_fee
    )
VALUES
    ('Flexible cancelation Policy', 1, false, 0.0),
    ('Non-Refundable Policy', 0, true, 100.0),
    ('7-Day cancelation Policy', 7, true, 50.0),
    ('14-Day cancelation Policy', 14, true, 25.0),
    ('30-Day cancelation Policy', 30, true, 10.0);

-- Insert multiple Rate Plan Types in one statement
INSERT INTO
    rates_schema.rate_plan_types (
        rate_plan_type_name,
        pay_in_advance,
        is_increasing_price,
        flexible_to_ammend,
        days_before_checkin,
        cancelation_policy_id,
        minimum_stay,
        maximum_stay,
        is_breakfast_included,
        is_dinner_included,
        is_lunch_included
    )
VALUES
    ('Standard Rate', false, false, true, 1, 1, 1, null, false, false, false),
    ('Non-Refundable Rate', true, false, false, 0, 2, 1, 7, false, false, false),
    ('Bed and Breakfast', false, false, true, 1, 1, 1, null, true, false, false),
    ('Half Board', false, false, true, 1, 1, 1, null, true, true, false),
    ('Full Board', false, false, true, 1, 1, 1, null, true, true, true),
    ('Early Bird', true, true, false, 30, 2, 2, null, false, false, false),
    ('Last Minute', false, true, false, 0, 2, 1, null, false, false, false),
    ('Weekly', false, false, true, 1, 1, 7, 29, false, false, false),
    ('Monthly', false, false, true, 1, 1, 30, null, false, false, false);
 INSERT INTO
    rates_schema.rate_plans (
        rate_plan_type_id,
        rate_plan_name,
        pay_in_advance,
        flexible_to_ammend,
        days_before_checkin,
        cancelation_policy_id,
        minimum_stay,
        maximum_stay,
        parent_rate_plan_id,
        is_related_price_percent,
        related_price_value,
        color,
        is_breakfast_included,
        is_dinner_included,
        is_lunch_included
    )
VALUES 
    (1, 'Standard Rate', false, true, 1, 1, 1, null, null, false, null, 'rgba(100, 149, 237, .7)' ,false, false, false),
    (2, 'Non-Refundable Rate', true, false, 0, 2, 1, 30, 1, true, 80.0, 'rgba(255, 165, 0, .7)' ,false, false, false),
    (3, 'Bed and Breakfast', false, true, 1, 1, 1, null, null, false, null, 'rgba(252, 186, 3, .7)	' ,true, false, false),
    (4, 'Half Board', false, true, 1, 1, 1, null, null, false, null, 'rgba(143, 188, 143, .7)', true, true, false),
    (5, 'Full Board', false, true, 1, 1, 1, null, null, false, null, 'rgba(230, 126, 34, .7)' ,true, true, true),
    (6, 'Early Bird', true, false, 30, 2, 2, null, 1, true, 90.0, 'rgba(34, 180, 169, .7)', false, false, false),
    (7, 'Last Minute', false, false, 0, 2, 1, null, 1, true, 85.0,  'rgba(231, 76, 60, .7)', false, false, false),
   (8, 'Weekly', false, true, 1, 1, 7, 29, 1, true, 30, 'rgba(102, 179, 222, .7)', false, false, false),
    (9, 'Monthly', false, true, 1, 1, 30,null , 1, true, 40,  'rgba(255, 193, 7, .7)', false, false, false);

-- -- Insert multiple properties in one statement
-- Insert multiple properties in one statement
INSERT INTO
    properties_schema.properties (
        property_name,
        address_line,
        star_rating,
        location_url,
        property_type_id,
        location_id,
        compound_id,
        owner_id,
        property_image,
        property_images,
        property_description,
        checkin_time_from,
        checkin_time_to,
        checkout_time_from,
        checkout_time_to
    )
VALUES
    (
        'Telal sokhna Villa',
        'The nearest airport is Cairo International Airport, 161 km from Telal sokhna villa.',
        0,
        'https://maps.app.goo.gl/snRFij93pVGYpc3L8',
        7,
        5,
        NULL,
        3,
        'properties/510290051.jpg',
        'properties/510290051.jpg,properties/510290082.jpg,properties/510290087.jpg,properties/510290090.jpg,properties/510290092.jpg,properties/510290094.jpg,properties/510290097.jpg,properties/510290099.jpg,properties/510290100.jpg,properties/510290102.jpg,properties/510290111.jpg,properties/510290122.jpg,properties/510290132.jpg,properties/510290144.jpg,properties/510290148.jpg',
        'Featuring air-conditioned accommodation with a private pool, garden view and a balcony, Telal sokhna villa is located in Ain Sokhna. The spacious holiday home features a terrace, 3 bedrooms, a living room and a well-equipped kitchen. Free private parking is available at the holiday home. The nearest airport is Cairo International Airport, 161 km from Telal sokhna villa.',
        '12:00',
        '14:00',
        '12:00',
        '14:00' 
    ),
    (
        'Vida Marina Resort Marassi',
        'Sedy Abdelrahman',
        4,
        'https://maps.app.goo.gl/zfwKztgaNneLHnp27',
        1,
        2,
        NULL,
        3,
        'properties/558805311.jpg',
        'properties/558805311.jpg,properties/559324005.jpg,properties/559324013.jpg,properties/559324013.jpg,properties/559324022.jpg,properties/563378774.jpg,properties/563378910.jpg,properties/563378964.jpg,properties/563379114.jpg,properties/563379274.jpg,properties/563379374.jpg',
        'Located in El Alamein, 34 km from Porto Marina, Vida Marina Resort Marassi provides air-conditioned rooms and a bar. Among the facilities of this property are a restaurant, room service and a 24-hour front desk, along with free WiFi throughout the property. Rooms are fitted with a balcony. At the hotel, each room includes a terrace. A buffet breakfast is available daily at Vida Marina Resort Marassi. The nearest airport is Borg el Arab International Airport, 112 km from the accommodation.',
        '12:00',
        '14:00',
        '12:00',
        '14:00' 
    ),
    (
        'Marassi Boutique Hotel',
        '400 metres from Safi Beach',
        3,
        'https://maps.app.goo.gl/JWT9vJztQ7aujJLu9',
        2,
        3,
        1,
        3,
        'properties/548116096.jpg',
        'properties/548116096.jpg,properties/548116215.jpg,properties/548116244.jpg,properties/548116252.jpg,properties/548128851.jpg,properties/548128959.jpg,properties/559341500.jpg,properties/559341503.jpg,properties/559341505.jpg,properties/559341508.jpg,properties/559341512.jpg',
        'Located 400 metres from Safi Beach, Marassi Boutique Hotel-Marina2 offers 3-star accommodation in El Alamein and features a restaurant. The property is non-smoking and is situated 35 km from Porto Marina. The nearest airport is Borg el Arab International Airport, 112 km from the hotel.',
       '12:00',
        '14:00',
        '12:00',
        '14:00' 
    ),
    (
        'Tawila Island Resort',
        '4KM from Hurghada',
        5,
        'https://maps.app.goo.gl/Zpv9E46iNfpbvS6p7',
        1,
        4,
        NULL,
        3,
        'properties/548116096.jpg',
        'properties/501590320.jpg,properties/501590349.jpg,properties/501590382.jpg,properties/501590386.jpg,properties/503898209.jpg,properties/503898398.jpg,properties/503898563.jpg,properties/503898582.jpg,properties/503899122.jpg,properties/503899869.jpg,properties/503899984.jpg,properties/503899986.jpg,properties/503899996.jpg,properties/503900041.jpg,properties/503900594.jpg',
        'Tawila Island Resort features an outdoor swimming pool, fitness centre, a garden and terrace in Hurghada. Among the various facilities are a bar and a private beach area. The accommodation offers a 24-hour front desk, airport transfers, a kids'' club and free WiFi throughout the property. At the resort all rooms come with air conditioning, a seating area, a flat-screen TV with satellite channels, a safety deposit box and a private bathroom with a bidet, free toiletries and a hairdryer. At Tawila Island Resort each room comes with bed linen and towels. The daily breakfast offers à la carte, continental or American options. At the accommodation you will find a restaurant serving international cuisine. Vegetarian, dairy-free and halal options can also be requested.',
        '12:00',
        '14:00',
        '12:00',
        '14:00' 
    ),
    (
        'Marina Hills',
        'Marina hills block 75',
        5,
        'http://localhost:5173/properties',
        4,
        1,
        2,
        2,
        'properties/469602831.jpg',
        'properties/469602831.jpg,properties/470127527.jpg,properties/470127571.jpg,properties/470127762.jpg,properties/470127786.jpg,properties/470127832.jpg,properties/470127872.jpg,properties/470127958.jpg,properties/470127966.jpg,properties/470127972.jpg,properties/470127990.jpg,properties/470128143.jpg,properties/470128175.jpg,properties/487927693.jpg,properties/487927744.jpg',
        'A very clean family friendly well organized chalet',
        '05:00:00',
        '10:00:00',
        '06:11:00',
        '10:00:00' 
    ),
    (
        'Rhactus Hotel',
        'New Alamein',
        5,
        'https://www.google.com/maps/dir/30.0450331,31.388456/marina+north+coast+location/@30.3890578,27.5267996,7z/data=!3m1!4b1!4m9!4m8!1m1!4e1!1m5!1m1!1s0x145fedf0f6120207:0x390b26abcc8c518f!2m2!1d28.966375!2d30.851088?entry=ttu',
        1,
        2,
        NULL,
        1,
        'properties/307187926.jpg',
        'properties/307187926.jpg,properties/307228874.jpg,properties/307228896.jpg,properties/375109488.jpg,properties/375113537.jpg,properties/375132129.jpg,properties/375341520.jpg,properties/375349121.jpg,properties/375349277.jpg,properties/397150902.jpg',
        '5 stars hotel in new alamein with a water view',
        '05:00:00',
        '11:00:00',
        '04:00:00',
        '08:00:00' 
    ),
    (
        'Apartment In Porto El Sokhna',
        'Ein El Sokhna, Atakka',
        0,
        'https://maps.app.goo.gl/xgVxpJeYm7ypMYDk8',
        5,
        5,
        NULL,
        3,
        'properties/355307852.jpg',
        'properties/355307852.jpg,properties/355307853.jpg,properties/355307856.jpg,properties/355307867.jpg,properties/355307871.jpg,properties/355307873.jpg,properties/355307879.jpg,properties/355307884.jpg,properties/355307886.jpg,properties/355307889.jpg,properties/449011381.jpg,properties/449011392.jpg,properties/449011394.jpg',
        'Porto Sokhna is a 4-star hotel offering spacious, self-catering and serviced accommodations, personalized service, and resort-style facilities in Ain Sokhna. It is just steps away from beaches along the Red Sea. Apartments, up to a floor space of 500 m², are fully furnished and include private balconies, separate seating areas, LCD TVs, and luxury bathrooms. Porto Sokhna Beach Resort, on-site restaurant offers a rich breakfast buffet and traditional regional meals are prepared for dinner. Free parking is offered to all guests.',
        '12:00',
        '14:00',
        '12:00',
        '14:00' 
    ),
    (
        'Lagoon View @ NAC',
        'NAC',
        4,
        'https://maps.app.goo.gl/14qz2NbhbhPGpcjp9',
        7,
        2,
        3,
        4,
        'properties/lagoon-main.webp',
        'properties/lagoon-01.webp,properties/lagoon-02.webp,properties/lagoon-03.webp,properties/lagoon-04.webp,properties/lagoon-05.webp',
        '3 bedroom villa.',
        '14:00',
        '14:00',
        '12:00',
        '12:00' 
    );


INSERT INTO
    properties_schema.reservable_units (
        reservable_unit_name,
        reservable_unit_description,
        minimum_guests_number,
        maximum_guests_number,
        unit_area,
        reservable_unit_type_id,
        property_id,
        bathrooms_count,
        is_closed,
        base_price,
        reservable_unit_image,
        reservable_unit_images
    )
VALUES
    (
        'Rhactus Deluxe Family',
        'Spacious family room with modern amenities',
        1,
        4,
        100,
        1,
        6,
        2,
        false,
        7649.00,
        'rooms/307187993.jpg',
        'rooms/307187993.jpg,rooms/307228874.jpg,rooms/397149691.jpg'
    ),
    (
        'Rhactus Deluxe King',
        'Luxurious king-size bed with a stunning view',
        1,
        2,
        80,
        2,
        6,
        1,
        false,
        10560.00,
        'rooms/307187791.jpg',
        'rooms/307187791.jpg,rooms/307187967.jpg,rooms/307188019.jpg,rooms/307228804.jpg,rooms/307228874.jpg,rooms/307228896.jpg'
    ),
    (
        'Rhactus Premium King',
        'Elegant room featuring a premium king-size bed',
        1,
        2,
        75,
        3,
        6,
        1,
        false,
        11320.00,
        'rooms/307187791.jpg',
        'rooms/307187791.jpg,rooms/307187967.jpg,rooms/307188019.jpg,rooms/307228804.jpg,rooms/307228874.jpg'
    ),
    (
        'Rhactus Deluxe Triple',
        'Comfortable room ideal for groups or families',
        1,
        3,
        90,
        4,
        6,
        2,
        false,
        10500.00,
        'rooms/307228874.jpg',
        'rooms/307228874.jpg,rooms/397149144.jpg'
    ),
    (
        'Rhactus Deluxe twin',
        'Cozy room with twin beds perfect for friends',
        1,
        2,
        70,
        5,
        6,
        1,
        false,
        9200.00,
        'rooms/307187791.jpg',
        'rooms/307187791.jpg,rooms/307187967.jpg,rooms/307188019.jpg,rooms/307228829.jpg,rooms/307228874.jpg'
    ),
    (
        'Marina Hills one bedroom apartment',
        'Modern one-bedroom apartment with stunning views of Marina Hills. This apartment features a spacious living area, fully equipped kitchen, and a private balcony.',
        1,
        2,
        60,
        7,
        5,
        1,
        false,
        9200.00,
        'rooms/470127571.jpg',
        'rooms/470127571.jpg,rooms/470127762.jpg,rooms/470127786.jpg,rooms/470127832.jpg,rooms/470127864.jpg,rooms/470127872.jpg,rooms/470127958.jpg,rooms/470127966.jpg,rooms/470127972.jpg,rooms/470127990.jpg,rooms/470128018.jpg,rooms/470128143.jpg,rooms/470128153.jpg,rooms/470128175.jpg,rooms/470128189.jpg,rooms/Photo_1715957595657.jpeg'
    ),
    (
        'Vida Deluxe - Garden View',
        'Spacious deluxe room with a picturesque garden view at Vida Hotel. This room features modern amenities and elegant decor.',
        1,
        2,
        60,
        1,
        2,
        1,
        false,
        6500.00,
        'rooms/558805311.jpg',
        'rooms/558805311.jpg, rooms/559324005.jpg,rooms/559324022.jpg,rooms/563378910.jpg'
    ),
    (
        'Vida Executive - Garden View',
        'Luxurious executive room with a serene garden view at Vida Hotel. This room offers spacious accommodations and premium amenities.',
        1,
        2,
        60,
        1,
        2,
        1,
        false,
        9890.00,
        'rooms/558805311.jpg',
        'rooms/558805311.jpg,rooms/559324022.jpg,rooms/563378910.jpg'
    ),
    (
        'Marassi boutique - Double Room',
        'Cozy double room with modern amenities at Marassi Boutique Hotel. Ideal for couples or solo travelers seeking comfort and convenience.',
        1,
        2,
        60,
        1,
        3,
        1,
        false,
        4760.00,
        'rooms/548066614.jpg',
        'rooms/548066614.jpg'
    ),
    (
        'Marassi boutique - Royal Suite',
        'Spacious junior suite with luxurious amenities at Marassi Boutique Hotel. Perfect for guests seeking a refined and elegant stay.',
        1,
        2,
        60,
        6,
        3,
        1,
        false,
        15430.00,
        'rooms/548066625.jpg',
        'rooms/548066625.jpg'
    ),
    (
        'Tawila - Standard Bungalow',
        'Comfortable standard bungalow with scenic views at Tawila Resort. This bungalow offers a cozy retreat amidst nature.',
        1,
        2,
        60,
        1,
        4,
        1,
        false,
        34570.00,
        'rooms/501590406.jpg',
        'rooms/501590406.jpg,rooms/501590440.jpg,rooms/503898334.jpg,rooms/503898397.jpg,rooms/503898398.jpg,rooms/503898497.jpg,rooms/503898563.jpg,rooms/503898573.jpg'
    ),
    (
        'Tawila - Deluxe Bungalow',
        'Luxurious deluxe bungalow with premium amenities at Tawila Resort. This bungalow provides a lavish and comfortable stay experience.',
        1,
        2,
        60,
        1,
        4,
        1,
        false,
        45390.00,
        'rooms/501590382.jpg',
        'rooms/501590382.jpg,rooms/501590440.jpg,rooms/503898474.jpg,rooms/503898497.jpg,rooms/503898582.jpg,rooms/503899752.jpg,rooms/503899834.jpg,rooms/503899932.jpg,rooms/503900594.jpg'
    ),
    (
        'Tawila - Superior Bungalow',
        'Exclusive superior bungalow offering unparalleled luxury at Tawila Resort. This bungalow features upscale amenities and breathtaking views.',
        1,
        2,
        60,
        1,
        4,
        2,
        false,
        75610.00,
        'rooms/501590386.jpg',
        'rooms/501590386.jpg,rooms/501590418.jpg,rooms/501590440.jpg,rooms/501590455.jpg,rooms/503898474.jpg,rooms/503899122.jpg,rooms/503899850.jpg,rooms/503899932.jpg,rooms/503900550.jpg,rooms/503900594.jpg'
    ),
    (
        'Porto El Sokhna - Two Bedroom Apartment',
        'Spacious two-bedroom apartment with stunning views at Porto El Sokhna. This apartment offers comfortable accommodations for families or groups.',
        1,
        4,
        120,
        8,
        7,
        2,
        false,
        3570.00,
        'rooms/472447068.jpg',
        'rooms/472447068.jpg'
    ),
    (
        'Telal El Sokhna Villa',
        'Spacious villa located in Telal El Sokhna, offering stunning views and luxurious amenities. This villa features multiple bedrooms, a private pool, and a spacious outdoor area.',
        8,
        8,
        200,
        10,
        1,
        4,
        false,
        8790.00,
        'rooms/510290082.jpg',
        'rooms/510290082.jpg,rooms/510290136.jpg,rooms/510290147.jpg'
    ),
    (
        'Lagoon View @ NAC',
        ' a private pool, and a spacious outdoor area.Spacious villa located in Telal El Sokhna, offering stunning views and luxurious amenities. This villa features multiple bedrooms, a private pool, and a spacious outdoor area.',
        6,
        6,
        145,
        9,
        8,
        4,
        false,
        2000.00,
        'rooms/lagoon-01.webp',
        'rooms/lagoon-02.webp,rooms/lagoon-03.webp'
    );


insert into properties_schema.reservable_unit_rooms(
    reservable_unit_id,
    reservable_unit_room_name
) values (
    1,
    'RH-DEL-F-1'
) ,
(
    2,
    'RH-DEL-K-1'
) ,
(
    3,
    'RH-PER-K-1'
) ,
(
    4,
    'RH-DEL-TR-1'
) ,
(
    5,
    'RH-DEL-TW-1'
) ,
(
    6,
    'MAR-OBA-1'
) ,
(
    7,
    'VID-DEK-GV-1'
) ,
(
    8,
    'VID-EXEC-GV-1'
) ,
(
    9,
    'MAR-DBL-1'
),
(
    10,
    'MAR-RS-1'
),
(
    11,
    'TA-ST-1'
),
(
    12,
    'TA-DEL-1'
),
(
    13,
    'TA-SP-1'
),
(
    14,
    'POR-TBA-1'
),
(
    14,
    'POR-TBA-2'
),
(
    15,
    'TEL-VI-1'
),
(
    15,
    'TEL-VI-2'
),
(
    15,
    'TEL-VI-3'
),
(
    15,
    'TEL-VI-4'
),
(
    15,
    'TEL-VI-5'
),
(
    16,
    'Lagoon View @ NAC ROOM NO: 1'
),
(
    16,
    'Lagoon View @ NAC ROOM NO: 2'
),
(
    16,
    'Lagoon View @ NAC ROOM NO: 3'
);

INSERT INTO properties_schema.reservable_unit_room_beds (
  reservable_unit_room_id, bed_type_id, bed_count
) VALUES
    -- Rhactus Deluxe Family
    (1, 1, 1),  -- 1 extra-large double bed
    (1, 2, 2),  -- 2 single beds

    -- Rhactus Deluxe King
    (2, 1, 1),  -- 1 extra-large double bed

    -- Rhactus Premium King
    (3, 1, 1),  -- 1 extra-large double bed

    -- Rhactus Deluxe Triple
    (4, 1, 1),  -- 1 extra-large double bed
    (4, 2, 1),  -- 1 single bed

    -- Rhactus Deluxe Twin
    (5, 2, 2),  -- 2 single beds

    -- Marina Hills one bedroom apartment
    (6, 1, 1),  -- 1 extra-large double bed

    -- Vida Deluxe - Garden View
    (7, 1, 1),  -- 1 extra-large double bed

    -- Vida Executive - Garden View
    (8, 1, 1),  -- 1 extra-large double bed

    -- Marassi boutique - Double Room
    (9, 1, 1),  -- 1 extra-large double bed

    -- Marassi boutique - Royal Suite
    (10, 1, 1), -- 1 extra-large double bed

    -- Tawila - Standard Bungalow
    (11, 1, 1), -- 1 extra-large double bed

    -- Tawila - Deluxe Bungalow
    (12, 1, 1), -- 1 extra-large double bed

    -- Tawila - Superior Bungalow
    (13, 1, 1), -- 1 extra-large double bed

    -- Porto El Sokhna - Two Bedroom Apartment
    (14, 1, 1), -- 1 extra-large double bed
    (15, 1, 1), -- 1 extra-large double bed

    -- Telal El Sokhna Villa
    (16, 1, 1), -- 1 extra-large double bed
     -- Telal El Sokhna Villa
    (17, 2, 2), -- 2 single beds
     -- Telal El Sokhna Villa
    (18, 2, 2), -- 2 single beds
     -- Telal El Sokhna Villa
    (19, 2, 2), -- 2 single beds
     -- Telal El Sokhna Villa
    (20, 2, 2), -- 2 single beds
    (21, 1, 1), -- 1 extra-large double bed
    (22, 2, 2), -- 2 single beds
    (23, 2, 2); -- 2 single beds
insert into rates_schema.reservable_unit_rate_plans(reservable_unit_id , rate_plan_id)
select ru.reservable_unit_id , rp.rate_plan_id  
from properties_schema.reservable_units ru
join rates_schema.rate_plans rp 
on rp.rate_plan_id < 3 or rp.rate_plan_id > 6
order by ru.reservable_unit_id ;

 

 
with reservable_unit_rate_plans as (
  select  
    rurp.rate_plan_id, 
    ru.reservable_unit_id, 
    ru.maximum_guests_number, 
    ru.base_price, 
    random_between(1, 30) available_count,
    (ru.maximum_guests_number > 1 and rut.property_category_id = 1) is_hotel,
    -- rurp.base_price
    rut.property_category_id, 
    rut.reservable_unit_type_id 
  from rates_schema.reservable_unit_rate_plans rurp 
  join properties_schema.reservable_units ru on rurp.reservable_unit_id = ru.reservable_unit_id
  join properties_schema.reservable_unit_types rut on ru.reservable_unit_type_id = rut.reservable_unit_type_id
), date_groups as (
  select 
    date_id, 
    (floor((row_number() over(order by date_id) - 1 ) / 4)) + 1 as group_id
  from dim_date 
  where date_actual >=  current_date
) ,inventory as (
  select 
    rurp.reservable_unit_id, 
    rurp.rate_plan_id,
    group_id,
    IIF(rurp.is_hotel , rurp.available_count , 1 ) available_count,
    rurp.base_price + (random_between((-200)::int, (500)::int)::real) price_1,
    IIF(rurp.is_hotel ,  rurp.base_price + (random_between((500)::int, (1000)::int)::real)  , null ) price_2,
    IIF(rurp.is_hotel and rurp.maximum_guests_number > 2  ,  rurp.base_price + (random_between((1000)::int, (1500)::int)::real)  , null ) price_3,
    IIF(rurp.is_hotel and rurp.maximum_guests_number > 3  ,  rurp.base_price + (random_between((1500)::int, (2000)::int)::real)  , null ) price_4,
    IIF(rurp.is_hotel and rurp.maximum_guests_number > 4  ,  rurp.base_price + (random_between((2000)::int, (2500)::int)::real)  , null ) price_5,
    IIF(rurp.is_hotel and rurp.maximum_guests_number > 5  ,  rurp.base_price + (random_between((3000)::int, (3500)::int)::real)  , null ) price_6,
    IIF(rurp.is_hotel and rurp.maximum_guests_number > 6  ,  rurp.base_price + (random_between((4000)::int, (4500)::int)::real)  , null ) price_7,
    IIF(rurp.is_hotel and rurp.maximum_guests_number > 7  ,  rurp.base_price + (random_between((5000)::int, (5500)::int)::real)  , null ) price_8,
  rurp.is_hotel
  from reservable_unit_rate_plans rurp
  cross join date_groups dg 
  group by rurp.reservable_unit_id, rurp.rate_plan_id , rurp.available_count , group_id ,    rurp.available_count,
    rurp.base_price,
    rurp.available_count,
    rurp.is_hotel , rurp.maximum_guests_number
  order by  reservable_unit_id , rate_plan_id , group_id
 
) , inventory_prepared as (
  select * from inventory i join date_groups dg on i.group_id = dg.group_id
)
insert into rates_schema.inventory (
  reservable_unit_id,
  rate_plan_id,
  available_count,
  price_1,
  price_2,
  price_3,
  price_4,
  price_5,
  price_6,
  price_7,
  price_8,
  date_id
)  select 
reservable_unit_id,
rate_plan_id,
available_count,
price_1,
price_2,
price_3,
price_4,
price_5,
price_6,
price_7,
price_8,
date_id  from inventory_prepared  order by reservable_unit_id , rate_plan_id , date_id;