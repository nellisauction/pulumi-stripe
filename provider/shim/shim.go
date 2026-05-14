package shim

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stripe/terraform-provider-stripe/internal/provider"
)

func NewProvider() *schema.Provider {
	return provider.New("dev")()
}
