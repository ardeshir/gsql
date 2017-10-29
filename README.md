# gsql

```javascript

Using the MySQL Go Driver, and testing GC!

We're using gc on a amd64 machine
with Go version go1.8.3
Number of Goroutines:  2

mem.Alloc:  73016
mem.TotalAlloc:  73016
mem.HeapAlloc:  73016
mem.NumGC:  0
------
mem.Alloc:  100074208
mem.TotalAlloc:  1000153432
mem.HeapAlloc:  100074208
mem.NumGC:  10
------
mem.Alloc:  66760
mem.TotalAlloc:  2000234568
mem.HeapAlloc:  66760
mem.NumGC:  20
------
```

```javascript
[root@ardeshir]:~/gosrc/github.com/ardeshir/gsql
0> GODEBUG=gctrace=1 go run main.go                                                                                                                                                      
gc 1 @0.039s 0%: 0.024+0.72+0.042 ms clock, 0.049+0.23/0.090/0.61+0.084 ms cpu, 4->4->0 MB, 5 MB goal, 2 P
gc 2 @0.052s 0%: 0.017+0.97+0.031 ms clock, 0.034+0.21/0.097/0.85+0.063 ms cpu, 4->4->0 MB, 5 MB goal, 2 P
gc 3 @0.066s 1%: 0.013+1.5+0.034 ms clock, 0.027+0.19/0.62/0.71+0.068 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
gc 4 @0.079s 1%: 0.010+1.4+0.026 ms clock, 0.020+0.17/0.76/0.67+0.053 ms cpu, 4->4->1 MB, 5 MB goal, 2 P
# command-line-arguments
gc 1 @0.004s 4%: 0.018+3.9+0.026 ms clock, 0.036+0.25/0.44/3.4+0.052 ms cpu, 4->4->3 MB, 5 MB goal, 2 P
# command-line-arguments
gc 1 @0.001s 27%: 0.16+5.8+0.029 ms clock, 0.33+0.68/2.8/0.094+0.059 ms cpu, 4->5->4 MB, 5 MB goal, 2 P
gc 2 @0.013s 27%: 0.010+3.7+0.17 ms clock, 0.020+2.0/3.2/0.42+0.34 ms cpu, 7->10->9 MB, 9 MB goal, 2 P
gc 3 @0.028s 19%: 0.013+9.9+0.068 ms clock, 0.027+0.28/4.8/5.0+0.13 ms cpu, 12->13->12 MB, 19 MB goal, 2 P
gc 4 @0.050s 17%: 0.011+18+0.024 ms clock, 0.022+0.11/8.5/10+0.049 ms cpu, 19->20->19 MB, 25 MB goal, 2 P
gc 5 @0.100s 16%: 0.013+34+0.040 ms clock, 0.027+0.048/20/14+0.080 ms cpu, 34->38->36 MB, 39 MB goal, 2 P
Using the MySQL Go Driver, and testing GC!
mem.Alloc:  73576
mem.TotalAlloc:  73576
mem.HeapAlloc:  73576
mem.NumGC:  0
------
gc 1 @0.002s 4%: 0.018+0.10+0.030 ms clock, 0.037+0.063/0.035/0.086+0.061 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 2 @0.071s 0%: 0.016+0.092+0.026 ms clock, 0.032+0.052/0.032/0.074+0.052 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 3 @0.078s 0%: 0.009+0.088+0.54 ms clock, 0.018+0.041/0/0.073+1.0 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 4 @0.085s 0%: 0.009+0.085+0.028 ms clock, 0.019+0.041/0/0.069+0.056 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 5 @0.091s 0%: 0.010+0.075+0.026 ms clock, 0.020+0.044/0.024/0.060+0.052 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 6 @0.200s 0%: 0.013+0.088+0.023 ms clock, 0.026+0.078/0/0.071+0.047 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 7 @0.207s 0%: 0.010+0.084+0.027 ms clock, 0.020+0.040/0/0.072+0.055 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 8 @0.214s 0%: 0.013+0.082+0.038 ms clock, 0.026+0.077/0/0.042+0.076 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 9 @0.221s 0%: 0.009+0.084+0.028 ms clock, 0.019+0.039/0/0.069+0.056 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 10 @0.227s 0%: 0.009+0.085+0.028 ms clock, 0.018+0/0.061/0.062+0.057 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
mem.Alloc:  66544
mem.TotalAlloc:  1000153912
mem.HeapAlloc:  66544
mem.NumGC:  10
------
gc 11 @0.234s 0%: 0.008+0.079+0.22 ms clock, 0.016+0.042/0.021/0.062+0.44 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 12 @5.241s 0%: 0.009+0.17+0.023 ms clock, 0.019+0.051/0.028/0.077+0.046 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 13 @10.247s 0%: 0.010+0.19+0.016 ms clock, 0.020+0.054/0.032/0.062+0.033 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 14 @15.254s 0%: 0.012+0.15+0.016 ms clock, 0.025+0.051/0.034/0.061+0.033 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 15 @20.261s 0%: 0.011+0.10+0.025 ms clock, 0.022+0.055/0.025/0.071+0.051 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 16 @25.267s 0%: 0.014+0.088+0.051 ms clock, 0.028+0.079/0/0.073+0.10 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 17 @30.274s 0%: 0.012+0.098+0.17 ms clock, 0.024+0.058/0.027/0.085+0.34 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 18 @35.280s 0%: 0.039+0.19+0.023 ms clock, 0.079+0.053/0.027/0.073+0.047 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 19 @40.287s 0%: 0.010+0.099+0.15 ms clock, 0.021+0.063/0.022/0.084+0.30 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
gc 20 @45.293s 0%: 0.032+0.11+0.022 ms clock, 0.064+0.088/0/0.082+0.044 ms cpu, 95->95->0 MB, 96 MB goal, 2 P
mem.Alloc:  67096
mem.TotalAlloc:  2000235112
mem.HeapAlloc:  67096
mem.NumGC:  20
------
```
