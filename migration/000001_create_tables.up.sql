CREATE TABLE exercises (
                           id UUID PRIMARY KEY,
                           name TEXT NOT NULL,
                           muscle_group_id UUID
);

CREATE TABLE muscle_groups (
                               id UUID PRIMARY KEY,
                               name TEXT NOT NULL
);

ALTER TABLE exercises
    ADD CONSTRAINT fk_exercises_muscle_group_id
        FOREIGN KEY (muscle_group_id)
            REFERENCES muscle_groups(id);