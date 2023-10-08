BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS "Settings" (
	"id"	INTEGER,
	"listenip"	TEXT NOT NULL,
	"port"	INTEGER NOT NULL,
	"allowip"	TEXT,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "Proxys" (
	"proxyid"	INTEGER,
	"proxyaddr"	TEXT NOT NULL,
	"proxyauth"	TEXT,
	"lastusage"	DATETIME,
	"lastcheck"	DATETIME,
	"status"	TEXT,
	PRIMARY KEY("proxyid")
);
CREATE TABLE IF NOT EXISTS "Result" (
	"resultid"	INTEGER,
	"maintargetid"	INTEGER,
	"secondtargetid"	INTEGER,
	"foundserviceid"	INTEGER,
	"resultdata"	TEXT,
	FOREIGN KEY("foundserviceid") REFERENCES "Foundservice"("foundserviceid"),
	FOREIGN KEY("maintargetid") REFERENCES "Main_source"("maintargetid"),
	FOREIGN KEY("secondtargetid") REFERENCES "Second_source"("id"),
	PRIMARY KEY("resultid")
);
CREATE TABLE IF NOT EXISTS "Masscan" (
	"id"	INTEGER,
	"profile"	TEXT NOT NULL,
	"string"	TEXT NOT NULL,
	"help"	TEXT,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "Ngrok" (
	"id"	INTEGER,
	"profile"	TEXT NOT NULL,
	"string"	TEXT NOT NULL,
	"help"	INTEGER,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "Hydra" (
	"id"	INTEGER,
	"profile"	TEXT NOT NULL,
	"string"	TEXT NOT NULL,
	"help"	TEXT,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "Passlist" (
	"id"	INTEGER,
	"name"	TEXT NOT NULL,
	"path"	TEXT NOT NULL,
	"help"	TEXT,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "Nmap" (
	"id"	INTEGER,
	"profile"	TEXT NOT NULL,
	"string"	TEXT NOT NULL,
	"help"	TEXT,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "Tools_list" (
	"id"	INTEGER NOT NULL,
	"name"	TEXT NOT NULL,
	"toolsid"	INTEGER NOT NULL,
	"toolscfg"	TEXT NOT NULL,
	"info"	TEXT,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "Logs" (
	"id"	INTEGER,
	"date"	DATETIME DEFAULT CURRENT_TIMESTAMP,
	"info"	TEXT,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "Userlist" (
	"id"	INTEGER NOT NULL,
	"name"	TEXT NOT NULL,
	"path"	TEXT NOT NULL,
	"help"	TEXT,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "Targets" (
	"id"	INTEGER NOT NULL,
	"parentid"	INTEGER,
	"name"	TEXT,
	"ip"	INTEGER,
	"domain"	TEXT,
	"country"	TEXT,
	"lastscan"	TEXT
);
CREATE TABLE IF NOT EXISTS "Find_services" (
	"id"	INTEGER NOT NULL,
	"targetsid"	INTEGER NOT NULL,
	"ports"	TEXT,
	"name"	TEXT,
	"details"	TEXT,
	"date"	INTEGER
);
CREATE TABLE IF NOT EXISTS "Vpn" (
	"id"	INTEGER NOT NULL,
	"name"	TEXT NOT NULL,
	"config"	TEXT NOT NULL,
	"help"	TEXT,
	PRIMARY KEY("id")
);
CREATE TABLE IF NOT EXISTS "Users" (
	"id"	INTEGER,
	"username"	TEXT NOT NULL UNIQUE,
	"password"	TEXT NOT NULL,
	"role"	TEXT NOT NULL,
	"lastaccess"	DATETIME,
	"create"	DATETIME DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY("id")
);
INSERT INTO "Nmap" VALUES (1,'Windows','nmap -T4 -F','Быстрое сканирование основных портов 1.');
INSERT INTO "Nmap" VALUES (2,'Ub-Quik1','nmap -v -A','Бла бла бла Linux нахуй Help');
INSERT INTO "Tools_list" VALUES (1,'Nmap',1,'100','IP and Port Scann tools');
INSERT INTO "Tools_list" VALUES (2,'Masscan',1,'101','IP and Port Scann tools');
INSERT INTO "Tools_list" VALUES (3,'Hydra',2,'200','Bruteforce and password attac tools');
INSERT INTO "Users" VALUES (1,'zip','Siwrol609','admin',NULL,'2023-10-06 04:18:11');
COMMIT;
