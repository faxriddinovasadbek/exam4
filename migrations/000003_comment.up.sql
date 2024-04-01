CREATE TABLE comments (
    id uuid DEFAULT gen_random_uuid(),
    content TEXT NOT NULL,
    post_id UUID NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL
);


INSERT INTO comments (id, content, post_id, user_id)
VALUES 
  ('47194831-ff63-4243-b1b4-a3de532a95e9', 'Lorem ipsum dolor', 'd3f1d36a-58b3-4e5d-8b40-5e4ee537ee25', 'ef5071ae-6cde-4217-a4fa-80eaea17fb6f'),
  ('c0773676-1173-4ebe-ae7c-a321216e295c', 'Sed do eiusmod', '5073cadb-e0bf-4b05-8bf2-528af176aafb', '2f1031d6-1412-4bf0-95fb-eabaf878aef1'),
  ('f77850af-e65e-4cf1-a975-111e4776e122', 'Ut enim ad', '15604eef-5546-4da9-8bd5-396b7fa467ba', '09303f4d-45e9-45ed-8bae-361e483db770'),
  ('b2e56d52-da17-4d25-b710-c7c3bec7235d', 'Duis aute irure', '8d059a55-06b7-4dbe-a04b-de05cbfe62a3', '9414b8fe-f86b-4d91-9d2b-4be299a6f723'),
  ('c3a17e10-901d-4d3f-bd1b-67bcfe0b26c2', 'Excepteur sint occaecat', '59b054e2-d888-4a4b-8186-7686277bbc87', 'b899d2c8-f456-4c45-8bc4-00f3e47740dc'),
  ('8dac87e3-906f-4e73-8aab-55d318b17ffc', 'Lorem ipsum dolor', '59b054e2-d888-4a4b-8186-7686277bbc87', 'ecd68fe8-3bea-4487-8e1e-faf968e5e9a0'),
  ('4745d5e2-d34f-4cf4-8d02-253e7b9b5488', 'Sed do eiusmod', '7c8f39cc-f2b9-4cd3-be3e-3038fddf4d14', 'a2bc1301-27bc-4d7d-9b11-580f47292fb0'),
  ('66ad22fd-661d-4d24-b1fa-9427e852f620', 'Ut enim ad', 'd2c9c149-8b8d-428b-8910-a98bb1b029e0', 'f28d07fd-7e58-4984-af5f-925397c9e2c7'),
  ('6fd58c64-064e-4a59-ba61-79bd9fcbdd78', 'Duis aute irure', '53c033d5-8ede-4e3b-b270-dd2de5360d0a', '2f1031d6-1412-4bf0-95fb-eabaf878aef1'),
  ('6022e5d2-d93a-44b8-a8a8-21f1cbb71d52', 'Excepteur sint occaecat', 'ade86809-1728-4204-8ea2-49ee0e50bbed', 'a2bc1301-27bc-4d7d-9b11-580f47292fb0'),
  ('e69fb9a4-81d5-48aa-b50c-6c93e8ee459b', 'Lorem ipsum dolor', '330f5d39-2b8b-4657-98ad-3950d2064e71', '09303f4d-45e9-45ed-8bae-361e483db770'),
  ('e475bbca-9b48-46f4-a5dd-92719a1d27cb', 'Sed do eiusmod', '32096fac-a0c9-4e4b-8d9f-cd872cdf5d8b', '2f1031d6-1412-4bf0-95fb-eabaf878aef1'),
  ('87ec42d3-94d9-4779-9e4b-0da1d220e108', 'Ut enim ad', '5388968d-e45f-44f7-9acd-590ccda53691', '94e55a1a-17b9-46ef-9ef4-68b9a4b28980 '),
  ('e89cda39-b55b-4e63-868e-76b97c331da7', 'Duis aute irure', '2c1f96e0-4c3d-4169-9f6d-050e5cf45dd0', 'bb9e0115-67f1-48d6-aef0-18725be78c14'),
  ('1314c774-6884-47e3-b2a1-8a06b875da2d', 'Excepteur sint occaecat', '5119ced3-3862-49e2-8400-f1a4001e51b5', '9414b8fe-f86b-4d91-9d2b-4be299a6f723');





