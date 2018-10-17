# sqd [![Build Status](https://travis-ci.org/Komei22/sqd.svg?branch=master)](https://travis-ci.org/Komei22/sqd)
suspicious query detection

## Usage
`sqd` detect suspicious query using list file(blacklist or whitelist).

```
$ ./sqd -q "DROP TABLE articles" -B testdata/blacklist
Suspicious queries:
DROP TABLE articles

./sqd -q "DROP TABLE articles" -W testdata/whitelist
Suspicious queries:
DROP TABLE articles
```

## File format
Query log and list file are written one query in one line such as the following example.

example of query.log
``` query.log
SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT 10
DELETE FROM articles WHERE articles.id = 1
SELECT * FROM articles
DROP TABLE articles
```

example of list
```
SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT ?
DELETE FROM articles WHERE articles.id = ?
SELECT * FROM articles
DROP TABLE articles
```
