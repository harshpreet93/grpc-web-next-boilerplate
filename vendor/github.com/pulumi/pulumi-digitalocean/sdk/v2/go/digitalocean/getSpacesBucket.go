// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get information on a Spaces bucket for use in other resources. This is useful if the Spaces bucket in question
// is not managed by this provider or you need to utilize any of the bucket's data.
//
// ## Example Usage
//
// Get the bucket by name:
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
// 		example, err := digitalocean.LookupSpacesBucket(ctx, &digitalocean.LookupSpacesBucketArgs{
// 			Name:   "my-spaces-bucket",
// 			Region: "nyc3",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("bucketDomainName", example.BucketDomainName)
// 		return nil
// 	})
// }
// ```
func LookupSpacesBucket(ctx *pulumi.Context, args *LookupSpacesBucketArgs, opts ...pulumi.InvokeOption) (*LookupSpacesBucketResult, error) {
	var rv LookupSpacesBucketResult
	err := ctx.Invoke("digitalocean:index/getSpacesBucket:getSpacesBucket", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSpacesBucket.
type LookupSpacesBucketArgs struct {
	// The name of the Spaces bucket.
	Name string `pulumi:"name"`
	// The slug of the region where the bucket is stored.
	Region string `pulumi:"region"`
}

// A collection of values returned by getSpacesBucket.
type LookupSpacesBucketResult struct {
	// The FQDN of the bucket (e.g. bucket-name.nyc3.digitaloceanspaces.com)
	BucketDomainName string `pulumi:"bucketDomainName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the Spaces bucket
	Name string `pulumi:"name"`
	// The slug of the region where the bucket is stored.
	Region string `pulumi:"region"`
	// The uniform resource name of the bucket
	Urn string `pulumi:"urn"`
}
