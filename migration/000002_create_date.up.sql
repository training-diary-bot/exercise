insert into muscle_groups (id, name) values
('2e6e90ec-083f-4200-9156-8e2b5c474c2f','Грудь'),
('ed365fd5-d214-4c41-a19d-1cdf545b7f08','Ноги'),
('b8c4ced4-42d3-4da7-a49c-07429e10ee1c','Бицепс'),
('b165cdd0-bb36-4a81-91ea-859552a5b03a','Трицепс'),
('1d8d4559-36d4-4986-9aea-999cf1159042','Спина'),
('aeea1c62-9bbb-49fa-9fcd-279bee475540','Плечи');

insert into exercise (id, name, muscle_group_id) values
('8dcbd8a6-c1da-450e-a595-00a7c727aea8', 'Жим лежа', '2e6e90ec-083f-4200-9156-8e2b5c474c2f');