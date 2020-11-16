module github.com/dhinojosac/smartbins/sbin_client

go 1.15

replace github.com/dhinojosac/smartbins/sbin_pb => ../sbin_pb

require (
	github.com/dhinojosac/smartbins/sbin_pb v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.33.2
)
