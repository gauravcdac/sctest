CREATE DATABASE test;

\c test;

CREATE TABLE ransomware_data (
    id SERIAL PRIMARY KEY,
    names VARCHAR[],
    extensions VARCHAR(1000),
    extensionPattern VARCHAR(1000),
    ransomNoteFilenames VARCHAR(1000),
    comment VARCHAR(1000),
    encryptionAlgorithm VARCHAR(1000),
    decryptor VARCHAR(1000),
    resources VARCHAR[],
    screenshots VARCHAR(1000),
    microsoftDetectionName VARCHAR(1000),
    microsoftInfo VARCHAR(1000),
    sandbox VARCHAR(1000),
    iocs VARCHAR(1000),
    snort VARCHAR(1000)
);
