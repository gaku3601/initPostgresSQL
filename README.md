# 🎉initPostgreSQL🎉
## これは何？
PostgresSQLのデータを全て消して、指定フォルダのSQLスクリプトを全てぶち込む  
ただそれだけのcommand lineツール  

## どうやって使うの

以下のような感じで使う
```
initDB -s /Users/gaku/src/github.com/gaku3601/initPostgresSQL/testSql/ -d database -h localhost -p 5432 -pw password -u user
```

オプション内容

```
initDB -h
flag needs an argument: -h
Usage of /var/folders/q4/37490pt55dqcmksm16ycywxr0000gn/T/go-build081428788/b001/exe/main:
  -d string
        database名
  -h string
        host名
  -p int
        port番号
  -pw string
        password
  -s string
        SqlFolder Path
  -u string
        user name
exit status 2
```

## 実行ファイル
ここにある  
https://github.com/gaku3601/initPostgresSQL/releases/tag/0.0.1
