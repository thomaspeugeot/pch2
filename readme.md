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
	cnpu=4		 3872ms		x3.56

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
