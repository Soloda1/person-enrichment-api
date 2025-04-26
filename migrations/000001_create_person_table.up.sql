CREATE TABLE IF NOT EXISTS Person (
    person_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY ,
    name TEXT NOT NULL,
    surname TEXT NOT NULL,
    patronymic TEXT,
    age INT,
    gender TEXT,
    national TEXT[],
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
)