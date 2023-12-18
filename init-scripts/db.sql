CREATE TABLE forecasts
(
    id                        SERIAL PRIMARY KEY,
    forecast_made_date        DATE,
    forecast_for_date         DATE,
    area_aac                  VARCHAR(255),
    air_temperature_minimum   INTEGER,
    air_temperature_maximum   INTEGER,
    precis                    TEXT,
    precipitation_probability INTEGER,
    precipitation_range       TEXT
);

CREATE TABLE temperature_data
(
    id                   SERIAL PRIMARY KEY,
    date                 DATE REFERENCES forecasts (forecast_for_date),
    min_temperature      INTEGER,
    max_temperature      INTEGER,
    precipitation_amount INTEGER,
    precis               TEXT,
);
