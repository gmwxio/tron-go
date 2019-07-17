module github.com/wxio/tron-go

go 1.12

replace golang.org/x/tools => github.com/wxio/tools v0.1.0

//replace golang.org/x/tools => /home/garym/devel/wxio/tools
//replace github.com/wxio/goantlr => /home/garym/devel/wxio/goantlr

require (
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/golang/protobuf v1.3.1
	github.com/golangq/q v1.0.7
	github.com/jpillora/opts v1.0.5
	github.com/wxio/goantlr v1.0.3
	golang.org/x/tools v0.0.0-00010101000000-000000000000
)
