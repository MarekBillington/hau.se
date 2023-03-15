
-- Seed country data
INSERT INTO "countries" (active, created_at, updated_at, name, currency, has_states)
VALUES (true, NOW(), NOW(), 'United Kingdom', 'GBP', false),
(true, NOW(), NOW(), 'Germany', 'EUR', true),
(true, NOW(), NOW(), 'United States', 'USD', true),
(true, NOW(), NOW(), 'New Zealand', 'NZD', false);

-- Seed user role data

INSERT INTO "user_roles" (active, created_at, updated_at, name, permissions)
VALUES (true, NOW(), NOW(), 'Admin', '{}'),
(true, NOW(), NOW(), 'Maintainer', '{}'),
(true, NOW(), NOW(), 'Viewer', '{}');
