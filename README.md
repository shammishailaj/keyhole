# Keyhole

Peek into `mongod` for

- Write Throughputs Test
- Load test
- Monitoring
- Cluster Info
- Seed data

## Usage
```
$ build/keyhole-linux-x64 -h
  -conn int
    	nuumber of connections (default 20)
  -duration int
    	load test duration in minutes (default 6)
  -info
    	get cluster info
  -peek
    	only collect data
  -seed
    	seed a database for demo
  -ssl
    	use TLS/SSL
  -sslCAFile string
    	CA file
  -tps int
    	number of trasaction per second per connection (default 600)
  -uri string
    	Mongo
    	DB URI
  -v	verbose
  -view string
    	server status file
```

## Download
### MacOS
```
curl -L https://github.com/simagix/keyhole/blob/master/build/keyhole-osx-x64?raw=true > keyhole ; chmod +x keyhole
```
### Linux
```
curl -L https://github.com/simagix/keyhole/blob/master/build/keyhole-linux-x64?raw=true > keyhole ; chmod +x keyhole
```
### Windows
The download link is as below.

```
https://github.com/simagix/keyhole/blob/master/build/keyhole-win-x64.exe?raw=true
```

## Use Cases
### Write Throughputs Test
Test MongoDB write throughput.

```
build/keyhole-linux-x64 -uri=mongodb://localhost/?replicaSet=replset -duration=1
```

### Load Test
Load test a cluster/replica.  A default cycle last six minutes.

- Populate data in first minute
- Perform CRUD operations during the second and third minutes
- Burst test during the fourth and fifth minutes
- Perform CRUD ops in the last minute

```
build/keyhole-linux-x64 -uri=mongodb://localhost/?replicaSet=replset
```

### Monitoring
Only collects data from `db.serverStatus()`

```
build/keyhole-linux-x64 -uri=mongodb://localhost/?replicaSet=replset -peek
```

### Cluster Info
Collect cluster information:

- Sharded cluster
- Replica set
- Standalone

```
build/keyhole-linux-x64 -uri=mongodb://localhost/?replicaSet=replset -info
```

### Seed Data
Populate a small amount of data for demo.

```
build/keyhole-linux-x64 -uri=mongodb://localhost/?replicaSet=replset -seed
```

## Load Test Example
```
build/keyhole-linux-x64 -uri=mongodb://localhost/?replicaSet=replset

MongoDB URI: mongodb://localhost
Total TPS: 600 (tps) * 20 (conns) = 12000, duration = 6 (mins)
cleanup mongodb://localhost
dropping database _KEYHOLE_88800
2018-04-19T08:41:33-04:00 res:     519, virt:    5436, faults:  1476
2018-04-19T08:41:42-04:00 data:    0.0 ->  426.7, rate   42.7 MB/sec
2018-04-19T08:41:52-04:00 data:  426.7 ->  900.8, rate   47.4 MB/sec
2018-04-19T08:42:02-04:00 data:  900.8 -> 1375.0, rate   47.4 MB/sec
2018-04-19T08:42:12-04:00 data: 1375.0 -> 1849.1, rate   47.4 MB/sec
2018-04-19T08:42:22-04:00 data: 1849.1 -> 2323.2, rate   47.4 MB/sec
2018-04-19T08:42:32-04:00 data: 2323.2 -> 2797.4, rate   47.4 MB/sec
2018-04-19T08:42:33-04:00 res:    1040, virt:    5980, faults:  1476, i:  720000, q:       0, u:       0, d:       0, iops:   12000
2018-04-19T08:42:42-04:00 data: 2797.4 -> 2844.8, rate    4.7 MB/sec
2018-04-19T08:43:33-04:00 res:    1040, virt:    5980, faults:  1476, i:  105983, q: 1165667, u:  105976, d:  105973, iops:   24726
2018-04-19T08:44:33-04:00 res:    1040, virt:    5980, faults:  1476, i:  105878, q: 1164640, u:  105876, d:  105873, iops:   24704
2018-04-19T08:45:33-04:00 res:    1040, virt:    5981, faults:  1476, i:  253501, q: 2788513, u:  253504, d:  253504, iops:   59150
2018-04-19T08:46:33-04:00 res:    1040, virt:    5981, faults:  1476, i:  242957, q: 2672503, u:  242954, d:  242954, iops:   56689

Server status written to /var/folders/mv/q3097r9j5kxb59sg1btgf2s80000gp/T/keyhole_stats.19048-04-19T08-47-32

--- Analytic Summary ---
+-------------------------+-------+-------+------+--------+--------+--------+--------+--------+--------+--------+
| Date/Time               | res   | virt  | fault| Command| Delete | Getmore| Insert | Query  | Update | iops   |
|-------------------------|-------+-------|------|--------|--------|--------|--------|--------|--------|--------|
|2018-04-19T12:42:33Z     |   1040|   5980|  1476|     150|       0|       0|  720000|       0|       0|   12002|
|2018-04-19T12:43:33Z     |   1040|   5980|  1476|     198|  105972|       0|  105982|  211946|  105976|    8834|
|2018-04-19T12:44:33Z     |   1040|   5980|  1476|     139|  105874|       0|  105879|  211755|  105876|    8825|
|2018-04-19T12:45:33Z     |   1040|   5981|  1476|     141|  253502|       0|  253500|  507003|  253500|   21127|
|2018-04-19T12:46:33Z     |   1040|   5981|  1476|     142|  242956|       0|  242956|  485907|  242954|   20248|
|2018-04-19T12:47:32Z     |   1040|   5986|  1476|     146|  103077|       0|  103077|  206157|  103081|    8888|
+-------------------------+-------+-------+------+--------+--------+--------+--------+--------+--------+--------+

--- Latencies Summary ---
+-------------------------+----------+----------+----------+
| Date/Time               | reads    | writes   | commands |
|-------------------------|----------|----------|----------|
|2018-04-19T12:42:33Z     |         0|        96|       620|
|2018-04-19T12:43:33Z     |        48|        66|        40|
|2018-04-19T12:44:33Z     |        48|        66|        19|
|2018-04-19T12:45:33Z     |       101|       145|        17|
|2018-04-19T12:46:33Z     |       100|       142|        16|
|2018-04-19T12:47:32Z     |        53|        71|        16|
+-------------------------+----------+----------+----------+

--- Metrics ---
+-------------------------+----------+------------+------------+--------------+----------+----------+----------+----------+
| Date/Time               | Scanned  | ScannedObj |ScanAndOrder|WriteConflicts| Deleted  | Inserted | Returned | Updated  |
|-------------------------|----------|------------|------------|--------------|----------|----------|----------|----------|
|2018-04-19T12:42:33Z     |         0|           0|           0|             0|         0|    720000|         0|         0|
|2018-04-19T12:43:33Z     |    317926|     1377616|           0|             0|    105973|    105983|   1165667|    105976|
|2018-04-19T12:44:33Z     |    317629|     1376389|           0|             0|    105873|    105878|   1164640|    105876|
|2018-04-19T12:45:33Z     |    760511|     3295521|           0|             0|    253504|    253501|   2788513|    253504|
|2018-04-19T12:46:33Z     |    728861|     3158411|           0|             0|    242954|    242957|   2672503|    242954|
|2018-04-19T12:47:32Z     |    309234|     1339994|           0|             0|    103077|    103075|   1133840|    103077|
+-------------------------+----------+------------+------------+--------------+----------+----------+----------+----------+

--- WiredTiger Summary ---
+-------------------------+--------------------+------------------+------------------------+-------------------+--------------------+-----------------------+
| Date/Time               | MaxBytesConfigured | CurrentlyInCache | UnmodifiedPagesEvicted | TrackedDirtyBytes | PagesReadIntoCache | PagesWrittenFromCache |
|-------------------------|--------------------|------------------|------------------------|-------------------|--------------------|-----------------------|
|2018-04-19T12:42:33Z     |          1073741824|         910344103|                  296423|          105923087|                   5|                 117398|
|2018-04-19T12:43:33Z     |          1073741824|         796112450|                  301305|           20500567|                   6|                   4221|
|2018-04-19T12:44:33Z     |          1073741824|         817724406|                  301798|           55422665|                   0|                     45|
|2018-04-19T12:45:33Z     |          1073741824|         771915373|                  301798|            9613667|                   0|                     69|
|2018-04-19T12:46:33Z     |          1073741824|         794083508|                  301798|           31781837|                   0|                     70|
|2018-04-19T12:47:32Z     |          1073741824|         809118456|                  301798|           46816750|                   0|                     44|
+-------------------------+--------------------+------------------+------------------------+-------------------+--------------------+-----------------------+
cleanup mongodb://localhost
dropping database _KEYHOLE_88800
Kens-MBP:keyhole kenchen$ go run keyhole.go -view=/var/folders/mv/q3097r9j5kxb59sg1btgf2s80000gp/T/keyhole_stats.19048-04-19T08-47-32
MongoDB URI: mongodb://localhost

--- Analytic Summary ---
+-------------------------+-------+-------+------+--------+--------+--------+--------+--------+--------+--------+
| Date/Time               | res   | virt  | fault| Command| Delete | Getmore| Insert | Query  | Update | iops   |
|-------------------------|-------+-------|------|--------|--------|--------|--------|--------|--------|--------|
|2018-04-19T12:42:33Z     |   1040|   5980|  1476|     150|       0|       0|  720000|       0|       0|   12002|
|2018-04-19T12:43:33Z     |   1040|   5980|  1476|     198|  105972|       0|  105982|  211946|  105976|    8834|
|2018-04-19T12:44:33Z     |   1040|   5980|  1476|     139|  105874|       0|  105879|  211755|  105876|    8825|
|2018-04-19T12:45:33Z     |   1040|   5981|  1476|     141|  253502|       0|  253500|  507003|  253500|   21127|
|2018-04-19T12:46:33Z     |   1040|   5981|  1476|     142|  242956|       0|  242956|  485907|  242954|   20248|
|2018-04-19T12:47:32Z     |   1040|   5986|  1476|     146|  103077|       0|  103077|  206157|  103081|    8888|
+-------------------------+-------+-------+------+--------+--------+--------+--------+--------+--------+--------+

--- Latencies Summary ---
+-------------------------+----------+----------+----------+
| Date/Time               | reads    | writes   | commands |
|-------------------------|----------|----------|----------|
|2018-04-19T12:42:33Z     |         0|        96|       620|
|2018-04-19T12:43:33Z     |        48|        66|        40|
|2018-04-19T12:44:33Z     |        48|        66|        19|
|2018-04-19T12:45:33Z     |       101|       145|        17|
|2018-04-19T12:46:33Z     |       100|       142|        16|
|2018-04-19T12:47:32Z     |        53|        71|        16|
+-------------------------+----------+----------+----------+

--- Metrics ---
+-------------------------+----------+------------+------------+--------------+----------+----------+----------+----------+
| Date/Time               | Scanned  | ScannedObj |ScanAndOrder|WriteConflicts| Deleted  | Inserted | Returned | Updated  |
|-------------------------|----------|------------|------------|--------------|----------|----------|----------|----------|
|2018-04-19T12:42:33Z     |         0|           0|           0|             0|         0|    720000|         0|         0|
|2018-04-19T12:43:33Z     |    317926|     1377616|           0|             0|    105973|    105983|   1165667|    105976|
|2018-04-19T12:44:33Z     |    317629|     1376389|           0|             0|    105873|    105878|   1164640|    105876|
|2018-04-19T12:45:33Z     |    760511|     3295521|           0|             0|    253504|    253501|   2788513|    253504|
|2018-04-19T12:46:33Z     |    728861|     3158411|           0|             0|    242954|    242957|   2672503|    242954|
|2018-04-19T12:47:32Z     |    309234|     1339994|           0|             0|    103077|    103075|   1133840|    103077|
+-------------------------+----------+------------+------------+--------------+----------+----------+----------+----------+

--- WiredTiger Summary ---
+-------------------------+--------------------+------------------+------------------------+-------------------+--------------------+-----------------------+
| Date/Time               | MaxBytesConfigured | CurrentlyInCache | UnmodifiedPagesEvicted | TrackedDirtyBytes | PagesReadIntoCache | PagesWrittenFromCache |
|-------------------------|--------------------|------------------|------------------------|-------------------|--------------------|-----------------------|
|2018-04-19T12:42:33Z     |          1073741824|         910344103|                  296423|          105923087|                   5|                 117398|
|2018-04-19T12:43:33Z     |          1073741824|         796112450|                  301305|           20500567|                   6|                   4221|
|2018-04-19T12:44:33Z     |          1073741824|         817724406|                  301798|           55422665|                   0|                     45|
|2018-04-19T12:45:33Z     |          1073741824|         771915373|                  301798|            9613667|                   0|                     69|
|2018-04-19T12:46:33Z     |          1073741824|         794083508|                  301798|           31781837|                   0|                     70|
|2018-04-19T12:47:32Z     |          1073741824|         809118456|                  301798|           46816750|                   0|                     44|
+-------------------------+--------------------+------------------+------------------------+-------------------+--------------------+-----------------------+
```

## Atlas TLS/SSL Mode
An example connecting to Atlas

```
build/keyhole-osx-x64 -uri=mongodb://user:secret@cluster0-shard-00-01-nhftn.mongodb.net.:27017,cluster0-shard-00-02-nhftn.mongodb.net.:27017,cluster0-shard-00-00-nhftn.mongodb.net.:27017/test?replicaSet=Cluster0-shard-0\&authSource=admin -ssl -sslCAFile=ssl/ca.crt -info
```
