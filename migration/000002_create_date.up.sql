insert into muscle_groups (id, name) values
('2e6e90ec-083f-4200-9156-8e2b5c474c2f','Грудь'),
('ed365fd5-d214-4c41-a19d-1cdf545b7f08','Ноги'),
('b8c4ced4-42d3-4da7-a49c-07429e10ee1c','Бицепс'),
('b165cdd0-bb36-4a81-91ea-859552a5b03a','Трицепс'),
('1d8d4559-36d4-4986-9aea-999cf1159042','Спина'),
('aeea1c62-9bbb-49fa-9fcd-279bee475540','Плечи');

insert into exercises (id, name, muscle_group_id) values
(uuid_in(md5(random()::text || random()::text)::cstring), 'Жим лежа', '2e6e90ec-083f-4200-9156-8e2b5c474c2f'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Жим в наклоне', '2e6e90ec-083f-4200-9156-8e2b5c474c2f'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Жим гантелей в наклоне', '2e6e90ec-083f-4200-9156-8e2b5c474c2f'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Разводка гантелей', '2e6e90ec-083f-4200-9156-8e2b5c474c2f'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Бабочка', '2e6e90ec-083f-4200-9156-8e2b5c474c2f'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Кроссовер', '2e6e90ec-083f-4200-9156-8e2b5c474c2f'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Брусья', '2e6e90ec-083f-4200-9156-8e2b5c474c2f'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Пуловер', '2e6e90ec-083f-4200-9156-8e2b5c474c2f'),

(uuid_in(md5(random()::text || random()::text)::cstring), 'Подтягивания', '1d8d4559-36d4-4986-9aea-999cf1159042'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Тяга штанги в наклоне', '1d8d4559-36d4-4986-9aea-999cf1159042'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Станова тяга', '1d8d4559-36d4-4986-9aea-999cf1159042'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Станова тяга (сумо)', '1d8d4559-36d4-4986-9aea-999cf1159042'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Нижний блок', '1d8d4559-36d4-4986-9aea-999cf1159042'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Средний блок', '1d8d4559-36d4-4986-9aea-999cf1159042'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Верхний блок', '1d8d4559-36d4-4986-9aea-999cf1159042'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Тяга гантели', '1d8d4559-36d4-4986-9aea-999cf1159042'),

(uuid_in(md5(random()::text || random()::text)::cstring), 'Подъем штанги', 'b8c4ced4-42d3-4da7-a49c-07429e10ee1c'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Подъем гантелей (молотки)', 'b8c4ced4-42d3-4da7-a49c-07429e10ee1c'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Сгибания в блоке', 'b8c4ced4-42d3-4da7-a49c-07429e10ee1c'),

(uuid_in(md5(random()::text || random()::text)::cstring), 'Жим гантели сидя', 'b165cdd0-bb36-4a81-91ea-859552a5b03a'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Франзуцский жим', 'b165cdd0-bb36-4a81-91ea-859552a5b03a'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Жим штанги узким', 'b165cdd0-bb36-4a81-91ea-859552a5b03a'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Разгибания в блоке', 'b165cdd0-bb36-4a81-91ea-859552a5b03a'),

(uuid_in(md5(random()::text || random()::text)::cstring), 'Присед', 'ed365fd5-d214-4c41-a19d-1cdf545b7f08'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Жим ногами', 'ed365fd5-d214-4c41-a19d-1cdf545b7f08'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Выпады', 'ed365fd5-d214-4c41-a19d-1cdf545b7f08'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Румныская тяга', 'ed365fd5-d214-4c41-a19d-1cdf545b7f08'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Разгибания в тренажере', 'ed365fd5-d214-4c41-a19d-1cdf545b7f08'),
(uuid_in(md5(random()::text || random()::text)::cstring), 'Сгибания в тренажере', 'ed365fd5-d214-4c41-a19d-1cdf545b7f08');
