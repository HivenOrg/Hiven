-- Users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    display_img_key TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    -- enforcing uniquesness and creating indexes
    CONSTRAINT unique_users_email UNIQUE (email),
    CONSTRAINT unique_users_phone_number UNIQUE (phone_number)
);

-- Hives table
CREATE TABLE IF NOT EXISTS hives (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    display_img_key TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Members table
CREATE TABLE IF NOT EXISTS members (
    hive_id INT NOT NULL,
    user_id INT NOT NULL,
    status VARCHAR(10) NOT NULL DEFAULT 'active',
    role VARCHAR(10) NOT NULL DEFAULT 'member',
    joined_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    
    PRIMARY KEY (hive_id, user_id),
    CONSTRAINT fk_hive FOREIGN KEY (hive_id) REFERENCES hives(id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT check_member_status CHECK (status IN ('active', 'left', 'removed')),
    CONSTRAINT check_member_role CHECK (role IN ('owner', 'member'))
);

-- Chores table
CREATE TABLE IF NOT EXISTS chores (
    id SERIAL PRIMARY KEY,
    hive_id INT NOT NULL,
    creator_id INT NOT NULL,
    chore TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_chores_hive FOREIGN KEY (hive_id) REFERENCES hives(id) ON DELETE CASCADE,
    CONSTRAINT fk_chores_user FOREIGN KEY (creator_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_chores_hive_id ON chores(hive_id);

-- Shopping items table
CREATE TABLE IF NOT EXISTS shopping_items (
    id SERIAL PRIMARY KEY,
    hive_id INT NOT NULL,
    creator_id INT NOT NULL,
    item TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_shopping_items_hive FOREIGN KEY (hive_id) REFERENCES hives(id) ON DELETE CASCADE,
    CONSTRAINT fk_shopping_items_user FOREIGN KEY (creator_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_shopping_items_hive_id ON shopping_items(hive_id);
