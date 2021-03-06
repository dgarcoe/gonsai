
CREATE TABLE IF NOT EXISTS "bonsais"(ID INTEGER PRIMARY KEY AUTOINCREMENT, NAME TEXT NOT NULL, AGE INT NOT NULL, SPECIES TEXT NOT NULL, STYLE TEXT NOT NULL, ACQUIRED TEXT NOT NULL, PRICE REAL, IMGPATH TEXT, BTYPE TEXT);

CREATE TABLE IF NOT EXISTS "pots"(ID INTEGER PRIMARY KEY AUTOINCREMENT, NAME TEXT NOT NULL, TYPE TEXT NOT NULL, ACQUIRED TEXT NOT NULL, PRICE REAL, IMGPATH TEXT);

CREATE TABLE IF NOT EXISTS "events"(ID INTEGER PRIMARY KEY AUTOINCREMENT, BONSAIID INTEGER NOT NULL, TYPE TEXT NOT NULL, DATE TEXT, COMMENT TEXT, FOREIGN KEY(BONSAIID) REFERENCES bonsais(ID) ON DELETE CASCADE ON UPDATE NO ACTION);
