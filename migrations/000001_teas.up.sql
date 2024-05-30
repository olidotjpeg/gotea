CREATE TABLE teas
(
    id BIGSERIAL PRIMARY KEY,
    teaName VARCHAR(255) NOT NULL,
    teaType VARCHAR(255) NOT NULL,
    shopName VARCHAR(255),
    shopLocation VARCHAR(255),
    temperature INTEGER,
    portionWeight INTEGER,
    containerWeight INTEGER,
    initialWeight INTEGER,
    brewingDuration INTEGER,
    color VARCHAR(255),
    size VARCHAR(255),
    inUse BOOLEAN,
    blendDescription VARCHAR(255)
)