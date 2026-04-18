INSERT INTO users (name, nick, email, password)
values ('User 1', 'user_1', 'user.1@teste.com', '$2a$10$rKhC02cSRUJIv8zKOb.2ieUBRAlRDqPFPtduDpcUjPb8XXsYll3PO'),
       ('User 2', 'user_2', 'user.2@teste.com', '$2a$10$rKhC02cSRUJIv8zKOb.2ieUBRAlRDqPFPtduDpcUjPb8XXsYll3PO'),
       ('User 3', 'user_3', 'user.3@teste.com', '$2a$10$rKhC02cSRUJIv8zKOb.2ieUBRAlRDqPFPtduDpcUjPb8XXsYll3PO'),
       ('User 4', 'user_4', 'user.4@teste.com', '$2a$10$rKhC02cSRUJIv8zKOb.2ieUBRAlRDqPFPtduDpcUjPb8XXsYll3PO'),
       ('User 5', 'user_5', 'user.5@teste.com', '$2a$10$rKhC02cSRUJIv8zKOb.2ieUBRAlRDqPFPtduDpcUjPb8XXsYll3PO');

INSERT INTO followers (user_id, follower_id)
values (2, 1),
       (3, 1),
       (1, 2),
       (3, 2),
       (1, 3),
       (2, 3);