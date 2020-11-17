// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DigitalOcean Cloud Firewall resource. This can be used to create,
// modify, and delete Firewalls.
type Firewall struct {
	pulumi.CustomResourceState

	// A time value given in ISO8601 combined date and time format
	// that represents when the Firewall was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The list of the IDs of the Droplets assigned
	// to the Firewall.
	DropletIds pulumi.IntArrayOutput `pulumi:"dropletIds"`
	// The inbound access rule block for the Firewall.
	// The `inboundRule` block is documented below.
	InboundRules FirewallInboundRuleArrayOutput `pulumi:"inboundRules"`
	// The Firewall name
	Name pulumi.StringOutput `pulumi:"name"`
	// The outbound access rule block for the Firewall.
	// The `outboundRule` block is documented below.
	OutboundRules FirewallOutboundRuleArrayOutput `pulumi:"outboundRules"`
	// An list of object containing the fields, "dropletId",
	// "removing", and "status".  It is provided to detail exactly which Droplets
	// are having their security policies updated.  When empty, all changes
	// have been successfully applied.
	PendingChanges FirewallPendingChangeArrayOutput `pulumi:"pendingChanges"`
	// A status string indicating the current state of the Firewall.
	// This can be "waiting", "succeeded", or "failed".
	Status pulumi.StringOutput `pulumi:"status"`
	// The names of the Tags assigned to the Firewall.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewFirewall registers a new resource with the given unique name, arguments, and options.
func NewFirewall(ctx *pulumi.Context,
	name string, args *FirewallArgs, opts ...pulumi.ResourceOption) (*Firewall, error) {
	if args == nil {
		args = &FirewallArgs{}
	}
	var resource Firewall
	err := ctx.RegisterResource("digitalocean:index/firewall:Firewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewall gets an existing Firewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallState, opts ...pulumi.ResourceOption) (*Firewall, error) {
	var resource Firewall
	err := ctx.ReadResource("digitalocean:index/firewall:Firewall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Firewall resources.
type firewallState struct {
	// A time value given in ISO8601 combined date and time format
	// that represents when the Firewall was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The list of the IDs of the Droplets assigned
	// to the Firewall.
	DropletIds []int `pulumi:"dropletIds"`
	// The inbound access rule block for the Firewall.
	// The `inboundRule` block is documented below.
	InboundRules []FirewallInboundRule `pulumi:"inboundRules"`
	// The Firewall name
	Name *string `pulumi:"name"`
	// The outbound access rule block for the Firewall.
	// The `outboundRule` block is documented below.
	OutboundRules []FirewallOutboundRule `pulumi:"outboundRules"`
	// An list of object containing the fields, "dropletId",
	// "removing", and "status".  It is provided to detail exactly which Droplets
	// are having their security policies updated.  When empty, all changes
	// have been successfully applied.
	PendingChanges []FirewallPendingChange `pulumi:"pendingChanges"`
	// A status string indicating the current state of the Firewall.
	// This can be "waiting", "succeeded", or "failed".
	Status *string `pulumi:"status"`
	// The names of the Tags assigned to the Firewall.
	Tags []string `pulumi:"tags"`
}

type FirewallState struct {
	// A time value given in ISO8601 combined date and time format
	// that represents when the Firewall was created.
	CreatedAt pulumi.StringPtrInput
	// The list of the IDs of the Droplets assigned
	// to the Firewall.
	DropletIds pulumi.IntArrayInput
	// The inbound access rule block for the Firewall.
	// The `inboundRule` block is documented below.
	InboundRules FirewallInboundRuleArrayInput
	// The Firewall name
	Name pulumi.StringPtrInput
	// The outbound access rule block for the Firewall.
	// The `outboundRule` block is documented below.
	OutboundRules FirewallOutboundRuleArrayInput
	// An list of object containing the fields, "dropletId",
	// "removing", and "status".  It is provided to detail exactly which Droplets
	// are having their security policies updated.  When empty, all changes
	// have been successfully applied.
	PendingChanges FirewallPendingChangeArrayInput
	// A status string indicating the current state of the Firewall.
	// This can be "waiting", "succeeded", or "failed".
	Status pulumi.StringPtrInput
	// The names of the Tags assigned to the Firewall.
	Tags pulumi.StringArrayInput
}

func (FirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallState)(nil)).Elem()
}

type firewallArgs struct {
	// The list of the IDs of the Droplets assigned
	// to the Firewall.
	DropletIds []int `pulumi:"dropletIds"`
	// The inbound access rule block for the Firewall.
	// The `inboundRule` block is documented below.
	InboundRules []FirewallInboundRule `pulumi:"inboundRules"`
	// The Firewall name
	Name *string `pulumi:"name"`
	// The outbound access rule block for the Firewall.
	// The `outboundRule` block is documented below.
	OutboundRules []FirewallOutboundRule `pulumi:"outboundRules"`
	// The names of the Tags assigned to the Firewall.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a Firewall resource.
type FirewallArgs struct {
	// The list of the IDs of the Droplets assigned
	// to the Firewall.
	DropletIds pulumi.IntArrayInput
	// The inbound access rule block for the Firewall.
	// The `inboundRule` block is documented below.
	InboundRules FirewallInboundRuleArrayInput
	// The Firewall name
	Name pulumi.StringPtrInput
	// The outbound access rule block for the Firewall.
	// The `outboundRule` block is documented below.
	OutboundRules FirewallOutboundRuleArrayInput
	// The names of the Tags assigned to the Firewall.
	Tags pulumi.StringArrayInput
}

func (FirewallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallArgs)(nil)).Elem()
}