module clzrt.io/autolabel/compute

go 1.22.3
require (
	clzrt.io/autolabel/struct v0.0.0
	clzrt.io/autolabel/storage v0.0.0
)

replace (
	clzrt.io/autolabel/struct => ../struct
	clzrt.io/autolabel/storage => ../storage
)