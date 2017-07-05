# Reiner-SQL-Benchmark

```
Yamide-MacBook-Air-2:Reiner-SQL-Brenchmark YamiOdymel$ go test -bench=. -benchmem
BenchmarkReinerInsert-4             3000            571298 ns/op            1719 B/op         49 allocs/op
BenchmarkSQLInsert-4                3000            429340 ns/op             901 B/op         17 allocs/op
BenchmarkDbrInsert-4                5000            413442 ns/op            2210 B/op         37 allocs/op
BenchmarkSQLxInsert-4               3000            444055 ns/op             902 B/op         17 allocs/op
BenchmarkGormInsert-4               2000            776838 ns/op            5319 B/op        101 allocs/op
BenchmarkXormInsert-4               3000            562341 ns/op            2921 B/op         64 allocs/op

BenchmarkReinerSelect1-4            5000            320895 ns/op            2325 B/op         59 allocs/op
BenchmarkReinerSelect10-4           5000            352664 ns/op            6197 B/op        162 allocs/op
BenchmarkReinerSelect100-4          2000            659189 ns/op           42907 B/op       1155 allocs/op
BenchmarkReinerSelect1000-4         1000           2236801 ns/op          388265 B/op      11058 allocs/op
BenchmarkReinerSelect10000-4         100          19271158 ns/op         5353186 B/op     110074 allocs/op

BenchmarkSQLSelect1-4              10000            173317 ns/op             896 B/op         23 allocs/op
BenchmarkSQLSelect10-4             10000            170967 ns/op            3680 B/op         90 allocs/op
BenchmarkSQLSelect100-4             5000            336121 ns/op           28864 B/op        723 allocs/op
BenchmarkSQLSelect1000-4            1000           1420339 ns/op          269058 B/op       7026 allocs/op
BenchmarkSQLSelect10000-4            100          10465566 ns/op         3753629 B/op      70035 allocs/op

BenchmarkDbrSelect1-4              10000            191162 ns/op            2553 B/op         47 allocs/op
BenchmarkDbrSelect10-4              5000            226188 ns/op           10451 B/op        195 allocs/op
BenchmarkDbrSelect100-4             3000            529430 ns/op           87496 B/op       1638 allocs/op
BenchmarkDbrSelect1000-4             500           3173413 ns/op          836139 B/op      16042 allocs/op
BenchmarkDbrSelect10000-4             50          29553882 ns/op         9833158 B/op     160056 allocs/op

BenchmarkSQLxSelect1-4             10000            159328 ns/op            1232 B/op         30 allocs/op
BenchmarkSQLxSelect10-4            10000            193809 ns/op            4304 B/op        106 allocs/op
BenchmarkSQLxSelect100-4            3000            376810 ns/op           32368 B/op        829 allocs/op
BenchmarkSQLxSelect1000-4           1000           1738220 ns/op          291252 B/op       8032 allocs/op
BenchmarkSQLxSelect10000-4          1000           1790335 ns/op          291252 B/op       8032 allocs/op

BenchmarkGormSelect1-4             10000            197715 ns/op            6250 B/op        101 allocs/op
BenchmarkGormSelect10-4             5000            254482 ns/op           24890 B/op        447 allocs/op
BenchmarkGormSelect100-4            2000            726107 ns/op          209236 B/op       3870 allocs/op
BenchmarkGormSelect1000-4            300           5175425 ns/op         2030803 B/op      38075 allocs/op
BenchmarkGormSelect10000-4            30          51186789 ns/op        21755784 B/op     380091 allocs/op

BenchmarkXormSelect1-4              5000            204462 ns/op            4210 B/op        121 allocs/op
BenchmarkXormSelect10-4             5000            277146 ns/op           13349 B/op        530 allocs/op
BenchmarkXormSelect100-4            2000            868688 ns/op          103358 B/op       4583 allocs/op
BenchmarkXormSelect1000-4            200           6306180 ns/op          991790 B/op      45088 allocs/op
BenchmarkXormSelect10000-4            20          68857571 ns/op        11029146 B/op     450100 allocs/op
PASS
ok      _/Users/YamiOdymel/Reiner-SQL-Brenchmark        61.423s
```