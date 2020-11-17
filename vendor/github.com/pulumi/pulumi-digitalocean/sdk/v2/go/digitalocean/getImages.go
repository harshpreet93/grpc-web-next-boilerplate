// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get information on images for use in other resources (e.g. creating a Droplet
// based on a snapshot), with the ability to filter and sort the results. If no filters are specified,
// all images will be returned.
//
// This data source is useful if the image in question is not managed by this provider or you need to utilize any
// of the image's data.
//
// Note: You can use the `getImage` data source to obtain metadata
// about a single image if you already know the `slug`, unique `name`, or `id` to retrieve.
//
// ## Example Usage
//
// Use the `filter` block with a `key` string and `values` list to filter images.
//
// For example to find all Ubuntu images:
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
// 		_, err := digitalocean.GetImages(ctx, &digitalocean.GetImagesArgs{
// 			Filters: []digitalocean.GetImagesFilter{
// 				digitalocean.GetImagesFilter{
// 					Key: "distribution",
// 					Values: []string{
// 						"Ubuntu",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// You can filter on multiple fields and sort the results as well:
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
// 		_, err := digitalocean.GetImages(ctx, &digitalocean.GetImagesArgs{
// 			Filters: []digitalocean.GetImagesFilter{
// 				digitalocean.GetImagesFilter{
// 					Key: "distribution",
// 					Values: []string{
// 						"Ubuntu",
// 					},
// 				},
// 				digitalocean.GetImagesFilter{
// 					Key: "regions",
// 					Values: []string{
// 						"nyc3",
// 					},
// 				},
// 			},
// 			Sorts: []digitalocean.GetImagesSort{
// 				digitalocean.GetImagesSort{
// 					Direction: "desc",
// 					Key:       "created",
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetImages(ctx *pulumi.Context, args *GetImagesArgs, opts ...pulumi.InvokeOption) (*GetImagesResult, error) {
	var rv GetImagesResult
	err := ctx.Invoke("digitalocean:index/getImages:getImages", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImages.
type GetImagesArgs struct {
	// Filter the results.
	// The `filter` block is documented below.
	Filters []GetImagesFilter `pulumi:"filters"`
	// Sort the results.
	// The `sort` block is documented below.
	Sorts []GetImagesSort `pulumi:"sorts"`
}

// A collection of values returned by getImages.
type GetImagesResult struct {
	Filters []GetImagesFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A set of images satisfying any `filter` and `sort` criteria. Each image has the following attributes:
	// - `slug`: Unique text identifier of the image.
	// - `id`: The ID of the image.
	// - `name`: The name of the image.
	// - `type`: Type of the image.
	Images []GetImagesImage `pulumi:"images"`
	Sorts  []GetImagesSort  `pulumi:"sorts"`
}