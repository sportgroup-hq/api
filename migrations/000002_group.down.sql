BEGIN;

ALTER TABLE users DROP COLUMN phone;
ALTER TABLE users DROP COLUMN date_of_birth;
ALTER TABLE users DROP COLUMN created_at;
ALTER TABLE users DROP COLUMN updated_at;
ALTER TABLE users RENAME COLUMN picture_url TO picture;

DROP TABLE groups;
DROP TABLE group_invites;
DROP TABLE group_members;

DROP TYPE group_member_type;

DROP FUNCTION update_updated_at();

COMMIT;
