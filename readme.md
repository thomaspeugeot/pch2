PCH2: Parallel convex hull (in 2 dimensions)
============================================

Algorithm
---------
See [doc](https://godoc.org/github.com/reconditematter/pq#ParConvHull2q) for more info.

Hardware/Software
-----------------
* CPU: Intel Xeon E5-1620 (3.6GHz, 4 cores, hyper-threading disabled)
* Memory: 16GB
* OS: Windows 7.1 64-bit
* Go: 1.4.2 windows/amd64

Results
-------
I executed the Go program 7 times for each value of ncpu={1,2,3,4}. The median execution times and the speedup factors are summarized below.
(N=number_of_points,L=lower_hull_size,U=upper_hull_size).

10,000 random points on the unit square (N=10000,L=13,U=11):

	ncpu=1		1278ms		x1
	ncpu=2		 660ms		x1.94
	ncpu=3		 458ms		x2.79
	ncpu=4		 372ms		x3.44

10,000 random points on the unit disc (N=10000,L=33,U=35):

	ncpu=1		1309ms		x1
	ncpu=2		 722ms		x1.81
	ncpu=3		 523ms		x2.50
	ncpu=4		 449ms		x2.92

100,000 random points on the unit square (N=100000,L=13,U=17):

	ncpu=1		13330ms		x1
	ncpu=2		 6833ms		x1.95
	ncpu=3		 4678ms		x2.85
	ncpu=4		 3767ms		x3.54

100,000 random points on the unit disc (N=100000,L=77,U=78):

	ncpu=1		13782ms		x1
	ncpu=2		 7055ms		x1.95
	ncpu=3		 4831ms		x2.85
	ncpu=4		 3872ms		x3.56

1,000,000 random points on the unit square (N=1000000,L=16,U=22):

	ncpu=1		140220ms	x1
	ncpu=2		 72083ms	x1.95
	ncpu=3		 49555ms	x2.83
	ncpu=4		 39853ms	x3.52

1,000,000 random points on the unit disc (N=1000000,L=167,U=173):

	ncpu=1		143886ms	x1
	ncpu=2		 74243ms	x1.94
	ncpu=3		 51128ms	x2.81
	ncpu=4		 41145ms	x3.50

Notes
-----
This implementation is far from the state of the art; much faster libraries are available ([Qhull](http://www.qhull.org) and [CGAL](http://www.cgal.org)).
I am interested in relative performance of a Go program on different number of CPUs. In this case, a satisfactory speedup factor x3.56 was achieved on a 4-core CPU.

Addendum 1
----------
The Go Team released Go 1.5 on 19 August 2015. I installed the new version and repeated the benchmark. The median execution times and the speedup factors
are summarized below.

10,000 random points on the unit square (N=10000,L=13,U=11):

	ncpu=1		1116ms		x1
	ncpu=2		 672ms		x1.66
	ncpu=3		 492ms		x2.27
	ncpu=4		 465ms		x2.40

10,000 random points on the unit disc (N=10000,L=33,U=35):

	ncpu=1		1147ms		x1
	ncpu=2		 680ms		x1.69
	ncpu=3		 507ms		x2.26
	ncpu=4		 477ms		x2.40

100,000 random points on the unit square (N=100000,L=13,U=17):

	ncpu=1		11173ms		x1
	ncpu=2		 6146ms		x1.82
	ncpu=3		 4463ms		x2.50
	ncpu=4		 3786ms		x2.95

100,000 random points on the unit disc (N=100000,L=77,U=78):

	ncpu=1		11574ms		x1
	ncpu=2		 6328ms		x1.83
	ncpu=3		 4588ms		x2.52
	ncpu=4		 3911ms		x2.96

1,000,000 random points on the unit square (N=1000000,L=16,U=22):

	ncpu=1		119670ms	x1
	ncpu=2		 65186ms	x1.84
	ncpu=3		 47155ms	x2.54
	ncpu=4		 39666ms	x3.02

1,000,000 random points on the unit disc (N=1000000,L=167,U=173):

	ncpu=1		122756ms	x1
	ncpu=2		 67114ms	x1.83
	ncpu=3		 48572ms	x2.53
	ncpu=4		 40732ms	x3.01

Addendum 2
----------
The Go Team released Go 1.7 on 15 August 2016. I installed the new version and repeated the benchmark. The median execution times and the speedup factors
are summarized below.

10,000 random points on the unit square (N=10000,L=13,U=11):

	ncpu=1		 852ms		x1
	ncpu=2		 468ms		x1.82
	ncpu=3		 365ms		x2.33
	ncpu=4		 333ms		x2.56

10,000 random points on the unit disc (N=10000,L=33,U=35):

	ncpu=1		 873ms		x1
	ncpu=2		 482ms		x1.81
	ncpu=3		 373ms		x2.34
	ncpu=4		 340ms		x2.57

100,000 random points on the unit square (N=100000,L=13,U=17):

	ncpu=1		 8962ms		x1
	ncpu=2		 4769ms		x1.88
	ncpu=3		 3393ms		x2.64
	ncpu=4		 2791ms		x3.21

100,000 random points on the unit disc (N=100000,L=77,U=78):

	ncpu=1		 9279ms		x1
	ncpu=2		 4938ms		x1.83
	ncpu=3		 3494ms		x2.52
	ncpu=4		 2885ms		x2.96

1,000,000 random points on the unit square (N=1000000,L=16,U=22):

	ncpu=1		103133ms	x1
	ncpu=2		 53789ms	x1.92
	ncpu=3		 38233ms	x2.70
	ncpu=4		 31242ms	x3.30

1,000,000 random points on the unit disc (N=1000000,L=167,U=173):

	ncpu=1		105733ms	x1
	ncpu=2		 54906ms	x1.93
	ncpu=3		 38952ms	x2.71
	ncpu=4		 32024ms	x3.30
	