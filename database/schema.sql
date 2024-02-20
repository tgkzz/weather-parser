CREATE TABLE IF NOT EXISTS weather(
    id SERIAL PRIMARY KEY,
    city VARCHAR(255) NOT NULL,
    temp NUMERIC NOT NULL,
    tempFahrenheit NUMERIC,
    tempKelvin NUMERIC,
    main VARCHAR(255),
    description TEXT,
    last_updated TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS server(
    lastUpdated TIMESTAMP NOT NULL
);
