// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// > **NOTE on `maxKeys`:** Retrieving very large numbers of keys can adversely affect this provider's performance.
//
// The bucket-objects data source returns keys (i.e., file names) and other metadata about objects in a Spaces bucket.
func GetSpacesBucketObjects(ctx *pulumi.Context, args *GetSpacesBucketObjectsArgs, opts ...pulumi.InvokeOption) (*GetSpacesBucketObjectsResult, error) {
	var rv GetSpacesBucketObjectsResult
	err := ctx.Invoke("digitalocean:index/getSpacesBucketObjects:getSpacesBucketObjects", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSpacesBucketObjects.
type GetSpacesBucketObjectsArgs struct {
	// Lists object keys in this Spaces bucket
	Bucket string `pulumi:"bucket"`
	// A character used to group keys (Default: none)
	Delimiter *string `pulumi:"delimiter"`
	// Encodes keys using this method (Default: none; besides none, only "url" can be used)
	EncodingType *string `pulumi:"encodingType"`
	// Maximum object keys to return (Default: 1000)
	MaxKeys *int `pulumi:"maxKeys"`
	// Limits results to object keys with this prefix (Default: none)
	Prefix *string `pulumi:"prefix"`
	// The slug of the region where the bucket is stored.
	Region string `pulumi:"region"`
}

// A collection of values returned by getSpacesBucketObjects.
type GetSpacesBucketObjectsResult struct {
	Bucket string `pulumi:"bucket"`
	// List of any keys between `prefix` and the next occurrence of `delimiter` (i.e., similar to subdirectories of the `prefix` "directory"); the list is only returned when you specify `delimiter`
	CommonPrefixes []string `pulumi:"commonPrefixes"`
	Delimiter      *string  `pulumi:"delimiter"`
	EncodingType   *string  `pulumi:"encodingType"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of strings representing object keys
	Keys    []string `pulumi:"keys"`
	MaxKeys *int     `pulumi:"maxKeys"`
	// List of strings representing object owner IDs
	Owners []string `pulumi:"owners"`
	Prefix *string  `pulumi:"prefix"`
	Region string   `pulumi:"region"`
}
