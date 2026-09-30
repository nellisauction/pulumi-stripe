package shim

import (
	tfpf "github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/stripe/terraform-provider-stripe/internal/provider"
)

func NewProvider() tfpf.Provider {
	return provider.New("dev")()
}
