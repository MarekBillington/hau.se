
-- Seed country data
INSERT INTO "countries" (active, created_at, updated_at, name, has_states)
VALUES (true, NOW(), NOW(), 'United Kingdom', false),
(true, NOW(), NOW(), 'Germany', true),
(true, NOW(), NOW(), 'United States', true),
(true, NOW(), NOW(), 'New Zealand', false);

-- Seed user role data

INSERT INTO "user_roles" (active, created_at, updated_at, name, permissions)
VALUES (true, NOW(), NOW(), 'Admin', '{}'),
(true, NOW(), NOW(), 'Maintainer', '{}'),
(true, NOW(), NOW(), 'Viewer', '{}');
