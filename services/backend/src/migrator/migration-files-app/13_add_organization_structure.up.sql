CREATE TABLE responsible_region (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  identifier VARCHAR UNIQUE,
  name VARCHAR
);

CREATE TABLE responsible_behoerde (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR,
  region_id UUID REFERENCES responsible_region(id)
);

CREATE TABLE responsible_behoerde_abteilung (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR,
  behoerde_id UUID REFERENCES responsible_behoerde(id)
);

CREATE TABLE application_responsible_organization (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  application_id UUID UNIQUE REFERENCES applications(id) ON DELETE CASCADE,
  abteilung_id UUID REFERENCES responsible_behoerde_abteilung(id) DEFAULT NULL,
  behoerde_id UUID REFERENCES responsible_behoerde(id) DEFAULT NULL,
  region_id UUID REFERENCES responsible_region(id) DEFAULT NULL,
  CHECK (
    ( abteilung_id IS NOT NULL AND behoerde_id IS NULL AND region_id IS NULL ) OR
    ( abteilung_id IS NULL AND behoerde_id IS NOT NULL AND region_id IS NULL ) OR
    ( abteilung_id IS NULL AND behoerde_id IS NULL AND region_id IS NOT NULL )
  ),
  UNIQUE (application_id, abteilung_id, behoerde_id, region_id)
);

CREATE OR REPLACE VIEW application_responsible_organization_full AS
SELECT
    aro.id               AS aoro_id,          
    aro.application_id   AS application_id,
    aro.abteilung_id    AS abteilung_id,
    COALESCE(
        aro.behoerde_id,                            -- direct reference
        (SELECT ba.behoerde_id
           FROM responsible_behoerde_abteilung ba
          WHERE ba.id = aro.abteilung_id)           -- derived from abteilung
    ) AS behoerde_id,

    COALESCE(
        aro.region_id,                              -- direct
        (SELECT b.region_id
           FROM responsible_behoerde b
          WHERE b.id = aro.behoerde_id),            -- from behoerde
        (SELECT b.region_id
           FROM responsible_behoerde b
          JOIN responsible_behoerde_abteilung ba
                ON ba.behoerde_id = b.id
         WHERE ba.id = aro.abteilung_id)            -- from abteilung
    ) AS region_id

FROM application_responsible_organization aro;