-- Create Resources Table
CREATE TABLE IF NOT EXISTS resources (
    resource_id SERIAL PRIMARY KEY,
    resource_name VARCHAR(255) NOT NULL,
    resource_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    status VARCHAR(50) NOT NULL,
    last_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Create Configuration Table
CREATE TABLE IF NOT EXISTS configuration (
    config_id SERIAL PRIMARY KEY,
    resource_id INTEGER NOT NULL,
    config_key VARCHAR(255) NOT NULL,
    config_value TEXT NOT NULL,
    last_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (resource_id) REFERENCES resources(resource_id)
);

-- Create Stack Table
CREATE TABLE IF NOT EXISTS stack (
    stack_id SERIAL PRIMARY KEY,
    stack_name VARCHAR(255) UNIQUE NOT NULL
);

-- Create ResourceStack Link Table
    CREATE TABLE IF NOT EXISTS resource_stack (
    resource_stack_id SERIAL PRIMARY KEY,
    resource_id INTEGER NOT NULL,
    stack_id INTEGER NOT NULL,
    FOREIGN KEY (resource_id) REFERENCES resources(resource_id),
    FOREIGN KEY (stack_id) REFERENCES stack(stack_id),
    UNIQUE (resource_id, stack_id)
);
