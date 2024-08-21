-- Drop the unique constraint on "owner" and "currency" if it exists
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT  "owner_currency_key";

-- Drop the foreign key constraint on "owner" if it exists
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT  "accounts_owner_fkey";

-- Drop the "users" table if it exists
DROP TABLE IF EXISTS "users";
