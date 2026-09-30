package stripe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	pf "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tfbridgetokens "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	pschema "github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
	stripeshim "github.com/stripe/terraform-provider-stripe/shim"

	"github.com/nellisauction/pulumi-stripe/provider/pkg/version"
)

const (
	mainPkg = "stripe"
	mainMod = "index"
)

//go:embed cmd/pulumi-resource-stripe/bridge-metadata.json
var metadata []byte

func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:                 pf.ShimProvider(stripeshim.NewProvider()),
		Name:              "stripe",
		Version:           version.Version,
		DisplayName:       "Stripe",
		Publisher:         "NellisAuction",
		PluginDownloadURL: "github://api.github.com/nellisauction/pulumi-stripe",
		Description:       "A Pulumi package for managing Stripe resources.",
		Keywords:          []string{"pulumi", "stripe", "category/infrastructure"},
		License:           "Apache-2.0",
		Homepage:          "https://github.com/nellisauction/pulumi-stripe",
		Repository:        "https://github.com/nellisauction/pulumi-stripe",
		GitHubOrg:         "stripe",
		UpstreamRepoPath:  "./upstream",
		Config: map[string]*tfbridge.SchemaInfo{
			"api_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"STRIPE_API_KEY"},
				},
			},
			"stripe_account": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"STRIPE_ACCOUNT"},
				},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName:          "@nellisauction/pulumi-stripe",
			RespectSchemaVersion: true,
		},
		Python: (func() *tfbridge.PythonInfo {
			i := &tfbridge.PythonInfo{RespectSchemaVersion: true}
			i.PyProject.Enabled = true
			return i
		})(),
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/nellisauction/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
			RespectSchemaVersion:           true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences:    map[string]string{"Pulumi": "3.*"},
			Namespaces:           map[string]string{mainPkg: "Stripe"},
		},
		MetadataInfo:                   tfbridge.NewProviderMetadata(metadata),
		EnableZeroDefaultSchemaVersion: true,
		EnableAccurateBridgePreview:    true,
		SchemaPostProcessor: func(spec *pschema.PackageSpec) {
			rewriteDocLinks(spec)
			unmarkScalarSecrets(spec)
		},
	}

	prov.MustComputeTokens(tfbridgetokens.SingleModule("stripe_", mainMod,
		tfbridgetokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}

// Upstream descriptions link to Stripe docs with root-relative paths, which tfgen resolves against terraform.io.
func rewriteDocLinks(spec *pschema.PackageSpec) {
	b, err := json.Marshal(spec)
	contract.AssertNoErrorf(err, "marshal package schema")
	b = bytes.ReplaceAll(b, []byte("https://www.terraform.io/"), []byte("https://docs.stripe.com/"))
	var out pschema.PackageSpec
	contract.AssertNoErrorf(json.Unmarshal(b, &out), "unmarshal package schema")
	*spec = out
}

// tfgen marks every write-only attribute secret, and the nodejs codegen drops false and 0 for secret inputs.
func unmarkScalarSecrets(spec *pschema.PackageSpec) {
	for _, r := range spec.Resources {
		props := []map[string]pschema.PropertySpec{r.InputProperties, r.Properties}
		if r.StateInputs != nil {
			props = append(props, r.StateInputs.Properties)
		}
		for _, m := range props {
			for k, p := range m {
				switch p.Type {
				case "boolean", "integer", "number":
					if p.Secret {
						p.Secret = false
						m[k] = p
					}
				}
			}
		}
	}
}
