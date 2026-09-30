module github.com/stripe/terraform-provider-stripe/shim

go 1.27.1

require (
	github.com/hashicorp/terraform-plugin-framework v1.18.0
	github.com/stripe/terraform-provider-stripe v0.3.0
)

require (
	github.com/fatih/color v1.18.0 // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/hashicorp/terraform-plugin-framework-validators v0.19.0 // indirect
	github.com/hashicorp/terraform-plugin-go v0.30.0 // indirect
	github.com/hashicorp/terraform-plugin-log v0.10.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/stripe/stripe-go/v86 v86.0.0-local // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
)

replace (
	github.com/stripe/stripe-go/v86 => ../../upstream/stripe-go
	github.com/stripe/terraform-provider-stripe => ../../upstream
)
