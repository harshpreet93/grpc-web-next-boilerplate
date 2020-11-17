// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get information on a domain. This data source provides the name, TTL, and zone
// file as configured on your DigitalOcean account. This is useful if the domain
// name in question is not managed by this provider or you need to utilize TTL or zone
// file data.
//
// An error is triggered if the provided domain name is not managed with your
// DigitalOcean account.
//
// ## Example Usage
//
// Get the zone file for a domain:
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
// 		example, err := digitalocean.LookupDomain(ctx, &digitalocean.LookupDomainArgs{
// 			Name: "example.com",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("domainOutput", example.ZoneFile)
// 		return nil
// 	})
// }
// ```
func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	var rv LookupDomainResult
	err := ctx.Invoke("digitalocean:index/getDomain:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDomain.
type LookupDomainArgs struct {
	// The name of the domain.
	Name string `pulumi:"name"`
}

// A collection of values returned by getDomain.
type LookupDomainResult struct {
	// The uniform resource name of the domain
	// * `zoneFile`: The zone file of the domain.
	DomainUrn string `pulumi:"domainUrn"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	Name     string `pulumi:"name"`
	Ttl      int    `pulumi:"ttl"`
	ZoneFile string `pulumi:"zoneFile"`
}