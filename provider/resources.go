package stripe

import (
	"fmt"
	"path/filepath"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	pf "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tfbridgetokens "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
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
	}

	prov.MustComputeTokens(tfbridgetokens.SingleModule("stripe_", mainMod,
		tfbridgetokens.MakeStandard(mainPkg)))
	prov.MustApplyAutoAliases()
	prov.SetAutonaming(255, "-")

	return prov
}
