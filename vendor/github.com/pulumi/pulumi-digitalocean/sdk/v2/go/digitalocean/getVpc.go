// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Retrieve information about a VPC for use in other resources.
//
// This data source provides all of the VPC's properties as configured on your
// DigitalOcean account. This is useful if the VPC in question is not managed by
// this provider or you need to utilize any of the VPC's data.
//
// VPCs may be looked up by `id` or `name`. Specifying a `region` will
// return that that region's default VPC.
//
// ## Example Usage
// ### VPC By Name
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v2/go/digitalocean"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "example-network"
// 		_, err := digitalocean.LookupVpc(ctx, &digitalocean.LookupVpcArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Reuse the data about a VPC to assign a Droplet to it:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v2/go/digitalocean"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "example-network"
// 		exampleVpc, err := digitalocean.LookupVpc(ctx, &digitalocean.LookupVpcArgs{
// 			Name: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = digitalocean.NewDroplet(ctx, "exampleDroplet", &digitalocean.DropletArgs{
// 			Size:    pulumi.String("s-1vcpu-1gb"),
// 			Image:   pulumi.String("ubuntu-18-04-x64"),
// 			Region:  pulumi.String("nyc3"),
// 			VpcUuid: pulumi.String(exampleVpc.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupVpc(ctx *pulumi.Context, args *LookupVpcArgs, opts ...pulumi.InvokeOption) (*LookupVpcResult, error) {
	var rv LookupVpcResult
	err := ctx.Invoke("digitalocean:index/getVpc:getVpc", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpc.
type LookupVpcArgs struct {
	// The unique identifier of an existing VPC.
	Id *string `pulumi:"id"`
	// The name of an existing VPC.
	Name *string `pulumi:"name"`
	// The DigitalOcean region slug for the VPC's location.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getVpc.
type LookupVpcResult struct {
	// The date and time of when the VPC was created.
	CreatedAt string `pulumi:"createdAt"`
	// A boolean indicating whether or not the VPC is the default one for the region.
	Default bool `pulumi:"default"`
	// A free-form text field describing the VPC.
	Description string `pulumi:"description"`
	// The unique identifier for the VPC.
	Id string `pulumi:"id"`
	// The range of IP addresses for the VPC in CIDR notation.
	IpRange string `pulumi:"ipRange"`
	// The name of the VPC.
	Name string `pulumi:"name"`
	// The DigitalOcean region slug for the VPC's location.
	Region string `pulumi:"region"`
	// The uniform resource name (URN) for the VPC.
	Urn string `pulumi:"urn"`
}
