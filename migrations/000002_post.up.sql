CREATE TABLE posts (
    id UUID NOT NULL,
    owner_id UUID NOT NULL,
    content TEXT NOT NULL,
    title TEXT NOT NULL,
    likes INT,
    dislikes INT,
    views INT, 
    category VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);

INSERT INTO posts (id, owner_id, content, title, likes, dislikes, views, category)
VALUES 
  ('d3f1d36a-58b3-4e5d-8b40-5e4ee537ee25', 'ef5071ae-6cde-4217-a4fa-80eaea17fb6f', 'Lorem ipsum dolor sit amet', 'First Post', 10, 2, 50, 'General'),
  ('5073cadb-e0bf-4b05-8bf2-528af176aafb', 'ef5071ae-6cde-4217-a4fa-80eaea17fb6f', 'Sed do labore et dolore magna aliqua.', 'Second Post', 15, 5, 70, 'Technology'),
  ('15604eef-5546-4da9-8bd5-396b7fa467ba', '2f1031d6-1412-4bf0-95fb-eabaf878aef1', 'Ut enim aliquip ex ea commodo consequat.', 'Third Post', 20, 3, 60, 'Sports'),
  ('ade86809-1728-4204-8ea2-49ee0e50bbed', '09303f4d-45e9-45ed-8bae-361e483db770', 'Duis aute  in voluptate fugiat nulla pariatur.', 'Fourth Post', 12, 1, 45, 'Travel'),
  ('c5662378-a0b6-4e82-b4c4-20104c401331', 'bb9e0115-67f1-48d6-aef0-18725be78c14', 'Excepteur  sunt in culpa anim id est laborum.', 'Fifth Post', 8, 4, 55, 'Food'),
  ('452b66f8-5cc6-413e-b6d8-76521efcf863', '9414b8fe-f86b-4d91-9d2b-4be299a6f723', 'Lorem ipsum dolor sit amet, consectetur adipiscing elit.', 'Sixth Post', 17, 3, 65, 'General'),
  ('59b054e2-d888-4a4b-8186-7686277bbc87', 'ecd68fe8-3bea-4487-8e1e-faf968e5e9a0', 'Sed do ut labore et dolore magna aliqua.', 'Seventh Post', 22, 6, 75, 'Technology'),
  ('53c033d5-8ede-4e3b-b270-dd2de5360d0a', '94e55a1a-17b9-46ef-9ef4-68b9a4b28980', 'Ut enim ad minim ex ea commodo consequat.', 'Eighth Post', 13, 2, 55, 'Sports'),
  ('2c1f96e0-4c3d-4169-9f6d-050e5cf45dd0', '94e55a1a-17b9-46ef-9ef4-68b9a4b28980', 'Duis reprehenderit fugiat nulla pariatur.', 'Ninth Post', 9, 3, 40, 'Travel'),
  ('5119ced3-3862-49e2-8400-f1a4001e51b5', 'a2bc1301-27bc-4d7d-9b11-580f47292fb0', 'Excepteur  mollit anim id est laborum.', 'Tenth Post', 16, 5, 60, 'Food'),
  ('5388968d-e45f-44f7-9acd-590ccda53691', 'f28d07fd-7e58-4984-af5f-925397c9e2c7', 'Lorem ipsum adipiscing elit.', 'Eleventh Post', 11, 1, 47, 'General'),
  ('32096fac-a0c9-4e4b-8d9f-cd872cdf5d8b', 'bb9e0115-67f1-48d6-aef0-18725be78c14', 'Sed do eiusmod  et dolore magna aliqua.', 'Twelfth Post', 18, 4, 68, 'Technology'),
  ('330f5d39-2b8b-4657-98ad-3950d2064e71', '2f1031d6-1412-4bf0-95fb-eabaf878aef1', 'Ut enim ad ullamcoex ea commodo consequat.', 'Thirteenth Post', 14, 2, 52, 'Sports'),
  ('7c8f39cc-f2b9-4cd3-be3e-3038fddf4d14', 'a2bc1301-27bc-4d7d-9b11-580f47292fb0', 'Duis velit esse cillum dolore eu fugiat nulla pariatur.', 'Fourteenth Post', 20, 3, 57, 'Travel'),
  ('ade86809-1728-4204-8ea2-49ee0e50bbed', 'f28d07fd-7e58-4984-af5f-925397c9e2c7', 'Excepteur sint occaecat cupidatat non proident', 'Fifteenth Post', 25, 6, 72, 'Food');

