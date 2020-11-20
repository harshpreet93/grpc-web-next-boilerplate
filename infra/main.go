package main

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/pulumi/pulumi-digitalocean/sdk/v2/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

type userDataVars struct {
	DoToken  string
	AppImage string
}

func synthesizeStack(stackName, appImage, tag, dc, sizeSlug, sshKey, userData string, desiredInstances int) func(ctx *pulumi.Context) error {
	return func(ctx *pulumi.Context) error {
		c := config.New(ctx, "")
		doToken := c.Require("do_token")
		image := fmt.Sprintf("%s:%s", appImage, tag)
		t := template.Must(template.New("user_data").Parse(userData))
		var out bytes.Buffer
		if err := t.Execute(&out, userDataVars{DoToken: doToken, AppImage: image}); err != nil {
			return err
		
		// TODO: set up LB
		
		for i := 0; i < desiredInstances; i++ {
			droplet, err := digitalocean.NewDroplet(ctx, fmt.Sprintf("%s-%d", stackName, i), &digitalocean.DropletArgs{
				Name:     pulumi.String(fmt.Sprintf("%s-%d", stackName, i)),
				Region:   pulumi.String(dc),
				Image:    pulumi.String("docker-18-04"),
				Size:     pulumi.String(sizeSlug),
				SshKeys:  pulumi.StringArray{pulumi.String(sshKey)},
				UserData: pulumi.StringPtr(out.String()),
			})
			if err != nil {
				return err
			}
			ctx.Export(fmt.Sprintf("appInstance-%s-%d", stackName, i), droplet.Ipv4Address)
			// TODO: ship necessary files to the instances
			// TODO: Reload services and restart
			// TODO: wait for app to be healthy
		}

		return nil
	}
}

func main() {
	stack := synthesizeStack("snacker", "registry.digitalocean.com/default/test", "latest",
		"sfo2", "s-1vcpu-1gb", "26813976", userData, 1)
	pulumi.Run(stack)
}
