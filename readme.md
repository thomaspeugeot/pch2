PCH2: Parallel convex hull (in 2 dimensions)
============================================

Hardware/Software
-----------------
* CPU: Intel Xeon E5-1620 (3.6GHz, 4 cores, hyper-threading disabled)
* Memory: 16GB
* OS: Windows 7.1 64-bit
* Go: 1.4.2 windows/amd64

Results
-------
I executed the Go program 7 times for each value of ncpu={1,2,3,4}. The median execution times and the speedup factors are summarized below.
(n=<number of points>,l=<lower hull size>,u=<upper hull size>).

10,000 random points on the unit square (n=10000,l=13,u=11):

	ncpu=1		1278ms		x1
	ncpu=2		 660ms		x1.94
	ncpu=3		 458ms		x2.79
	ncpu=4		 372ms		x3.44

10,000 random points on the unit disc(n=10000,l=33,u=35):

	ncpu=1		1309ms		x1
	ncpu=2		 722ms		x1.81
	ncpu=3		 523ms		x2.50
	ncpu=4		 449ms		x2.92
