# sqd [![Build Status](https://travis-ci.org/Komei22/sqd.svg?branch=master)](https://travis-ci.org/Komei22/sqd)

sqd is suspicious query detection

## How to detection
sqd detect suspicious query using list file, whitelist or blacklist.
List file is defined the structure of query.
The structure of query is the string replaced all literal with place holder token in the query.
For example, `SELECT * FROM articles WHERE id = 1` convert to `SELECT * FROM articles WHERE id = ?` and define list file.
sqd check the structure of query and decide suspicious query.
sqd support MySQL and PostgreSQL query.

## Installation

```
go get -u github.com/Komei22/sqd
```

## Usage
### Suspicious query detection
`sqd` detect suspicious query using list file, blacklist or whitelist.

```
$ sqd -q "DROP TABLE articles" -B blacklist -d mysql
DROP TABLE articles

$ sqd -q "DROP TABLE articles" -W whitelist -d mysql
DROP TABLE articles
```

### Whitelisting
`sqd create` convert the query with the structure of query and output it.

```
$ cat query.log
INSERT INTO articles (title, content, created_at, updated_at) VALUES ('hoge', 'fuga', '2018-11-01 09:53:37', '2018-11-01 09:53:37')
SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT 30
DELETE FROM articles WHERE articles.id = 1

$ cat query.log | sqd create -d mysql
DELETE FROM articles WHERE articles.id = ?
INSERT INTO articles (title, content, created_at, updated_at) VALUES (?, ?, ?, ?)
SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
```

## File format
Query log are written one query in one line such as the following example.

Example of query log
```
SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT 1
DELETE FROM\narticles\nWHERE\narticles.id = 1
SELECT * FROM articles
DROP TABLE articles
```

List file are written one structure of query in one line such as the following example.

Example of list file
```
SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
DELETE FROM articles WHERE articles.id = ?
SELECT * FROM articles
DROP TABLE articles
```
