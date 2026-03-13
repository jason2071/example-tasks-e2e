CREATE TABLE
    IF NOT EXISTS example.tasks (
        "id" BIGSERIAL PRIMARY KEY,
        "title" text NOT NULL,
        "status" varchar(20) DEFAULT 'pending',
        "created_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
        "updated_at" timestamptz DEFAULT CURRENT_TIMESTAMP,
        CONSTRAINT "check_campaign_status" CHECK (status IN ('pending', 'doing', 'done'))
    );

-- Drop existing unique key constraint on title if exists
ALTER TABLE example.tasks
DROP CONSTRAINT IF EXISTS uq_example_tasks_ref;

-- Add unique constraint on title
ALTER TABLE example.tasks ADD CONSTRAINT uq_example_tasks_ref UNIQUE (title);

-- Add priority column
ALTER TABLE example.tasks
ADD COLUMN IF NOT EXISTS "priority" int NULL;

-- Create index on priority column
CREATE INDEX IF NOT EXISTS idx_example_tasks_priority ON example.tasks (priority);

-- Add check constraint for priority range (1 to 5)
ALTER TABLE example.tasks ADD CONSTRAINT "check_priority_range" CHECK (
    priority IS NULL
    OR (
        priority >= 1
        AND priority <= 5
    )
);