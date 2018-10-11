# sqd
suspicious query detection

## Usage
`sqd check` detect illegal query in query log file using list file(blacklist or whitelist).

```
$ sqd check -m whitelist -q query.log -l whitelist
Suspicious querys
select * from articles
drop database articles
```

## File format
Query log and list file are written one query in one line such as the following example.

``` query.log
SELECT articles.* FROM articles ORDER BY articles.id DESC LIMIT 10
DELETE FROM articles WHERE articles.id = 1
SELECT * FROM articles
DROP TABLE articles
```
