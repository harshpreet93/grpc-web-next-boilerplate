// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get information on a volume for use in other resources. This data source provides
// all of the volumes properties as configured on your DigitalOcean account. This is
// useful if the volume in question is not managed by this provider or you need to utilize
// any of the volumes data.
//
// An error is triggered if the provided volume name does not exist.
//
// ## Example Usage
//
// Get the volume:
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
// 		opt0 := "nyc3"
// 		_, err := digitalocean.LookupVolume(ctx, &digitalocean.LookupVolumeArgs{
// 			Name:   "app-data",
// 			Region: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Reuse the data about a volume to attach it to a Droplet:
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
// 		opt0 := "nyc3"
// 		exampleVolume, err := digitalocean.LookupVolume(ctx, &digitalocean.LookupVolumeArgs{
// 			Name:   "app-data",
// 			Region: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleDroplet, err := digitalocean.NewDroplet(ctx, "exampleDroplet", &digitalocean.DropletArgs{
// 			Size:   pulumi.String("s-1vcpu-1gb"),
// 			Image:  pulumi.String("ubuntu-18-04-x64"),
// 			Region: pulumi.String("nyc3"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = digitalocean.NewVolumeAttachment(ctx, "foobar", &digitalocean.VolumeAttachmentArgs{
// 			DropletId: exampleDroplet.ID(),
// 			VolumeId:  pulumi.String(exampleVolume.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("digitalocean:index/getVolume:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVolume.
type LookupVolumeArgs struct {
	// Text describing a block storage volume.
	Description *string `pulumi:"description"`
	// The name of block storage volume.
	Name string `pulumi:"name"`
	// The region the block storage volume is provisioned in.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getVolume.
type LookupVolumeResult struct {
	// Text describing a block storage volume.
	Description *string `pulumi:"description"`
	// A list of associated Droplet ids.
	DropletIds []int `pulumi:"dropletIds"`
	// Filesystem label currently in-use on the block storage volume.
	FilesystemLabel string `pulumi:"filesystemLabel"`
	// Filesystem type currently in-use on the block storage volume.
	FilesystemType string `pulumi:"filesystemType"`
	// The provider-assigned unique ID for this managed resource.
	Id     string  `pulumi:"id"`
	Name   string  `pulumi:"name"`
	Region *string `pulumi:"region"`
	// The size of the block storage volume in GiB.
	Size int `pulumi:"size"`
	// A list of the tags associated to the Volume.
	Tags []string `pulumi:"tags"`
	Urn  string   `pulumi:"urn"`
}
