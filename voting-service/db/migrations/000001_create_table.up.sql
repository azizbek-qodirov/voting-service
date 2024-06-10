CREATE TABLE IF NOT EXISTS party (
    id UUID PRIMARY KEY,
    name VARCHAR(50),
    slogan VARCHAR(50),
    open_date TIMESTAMP NOT NULL,
    description TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS public (
    id UUID PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    birthday TIMESTAMP NOT NULL,
    gender VARCHAR(1) NOT NULL,
    nation VARCHAR(30) NOT NULL,
    party_id UUID REFERENCES party(id)
);

CREATE TABLE IF NOT EXISTS election (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    date TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS candidate (
    id UUID PRIMARY KEY,
    election_id UUID  UNIQUE REFERENCES election(id),
    public_id UUID  UNIQUE,
    party_id UUID
);

CREATE TABLE IF NOT EXISTS public_vote (
    id UUID PRIMARY KEY,
    election_id UUID  UNIQUE REFERENCES election(id),
    public_id UUID UNIQUE
);

CREATE TABLE IF NOT EXISTS vote (
    id UUID PRIMARY KEY,
    election_id UUID  UNIQUE REFERENCES election(id),
    candidate_id UUID REFERENCES candidate(id)
);