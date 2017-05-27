CREATE DATABASE gwp;

\c gwp

CREATE USER gwp WITH PASSWORD 'gwp';

GRANT ALL PRIVILEGES ON DATABASE gwp to gwp;

GRANT ALL PRIVILEGES ON TABLE gwp to gwp;

CREATE TABLE leads (
   ID       SERIAL PRIMARY KEY,
   CONTENT   TEXT,
    AUTHOR  VARCHAR(255)
);

INSERT INTO gwp (id, content, author) VALUES
('1','Lead','Devilbiss','Arie','20000.00'),
('2','Negotiation','Phillips','Lynda','45000.00'),
('3','Proposal','Phillips','Lynda','20000.00'),
('4','Lead','MSA','Deb','1500.00');
