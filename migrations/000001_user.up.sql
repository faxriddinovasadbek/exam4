CREATE TABLE users (
    id UUID NOT NULL,
    name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    username VARCHAR(64) NOT NULL,
    email VARCHAR(64) NOT NULL,
    password TEXT NOT NULL,
    bio TEXT,
    website TEXT,
    refresh_token TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP 
);

INSERT INTO users (id, name, last_name, username, email, password, bio, website, refresh_token)
VALUES 
  ('ef5071ae-6cde-4217-a4fa-80eaea17fb6f', 'John', 'Doe', 'johndoe', 'john@example.com', 'password123', 'Software engineer', 'https://example.com', 'your_refresh_token_here'),
  ('2f1031d6-1412-4bf0-95fb-eabaf878aef1', 'Emma', 'Smith', 'emmasmith', 'emma@example.com', 'password456', 'Data scientist', 'https://emma.com', 'refresh_token_2'),
  ('09303f4d-45e9-45ed-8bae-361e483db770', 'Michael', 'Johnson', 'michaeljohnson', 'michael@example.com', 'password789', 'Web developer', 'https://michael.com', 'refresh_token_3'),
  ('bb9e0115-67f1-48d6-aef0-18725be78c14', 'Sophia', 'Williams', 'sophiawilliams', 'sophia@example.com', 'password987', 'UX/UI designer', 'https://sophia.com', 'refresh_token_4'),
  ('9414b8fe-f86b-4d91-9d2b-4be299a6f723', 'William', 'Jones', 'williamjones', 'william@example.com', 'password654', 'Software engineer', 'https://william.com', 'refresh_token_5'),
  ('b899d2c8-f456-4c45-8bc4-00f3e47740dc', 'Olivia', 'Brown', 'oliviabrown', 'olivia@example.com', 'password321', 'Frontend developer', 'https://olivia.com', 'refresh_token_6'),
  ('ecd68fe8-3bea-4487-8e1e-faf968e5e9a0', 'James', 'Davis', 'jamesdavis', 'james@example.com', 'password123', 'Backend developer', 'https://james.com', 'refresh_token_7'),
  ('94e55a1a-17b9-46ef-9ef4-68b9a4b28980', 'Charlotte', 'Miller', 'charlottemiller', 'charlotte@example.com', 'password456', 'Data analyst', 'https://charlotte.com', 'refresh_token_8'),
  ('a2bc1301-27bc-4d7d-9b11-580f47292fb0', 'Liam', 'Wilson', 'liamwilson', 'liam@example.com', 'password789', 'Systems administrator', 'https://liam.com', 'refresh_token_9'),
  ('f28d07fd-7e58-4984-af5f-925397c9e2c7', 'Ava', 'Taylor', 'avataylor', 'ava@example.com', 'password987', 'Network engineer', 'https://ava.com', 'refresh_token_10');
