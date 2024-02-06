package main

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		configProvider := NewJSONConfigDataProvider("config.json")

		// Get the configuration data
		configData, err := configProvider.GetConfigData()
		if err != nil {
			return err
		}

		// Convert configData to pulumi.StringMap
		pulumiConfigData := pulumi.StringMap{}
		for k, v := range configData {
			pulumiConfigData[k] = pulumi.String(v)
		}

		// Create a ConfigMap with the above configuration data
		cfmap, err := v1.NewConfigMap(ctx, "generated-configmap", &v1.ConfigMapArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.String("generated-configmap"),
				Namespace: pulumi.String("default"),
			},
			Data: pulumiConfigData,
		})
		if err != nil {
			return err
		}

		ctx.Export("new configmap created", cfmap.Metadata.Name())
		// Return without error
		return nil
	})
}
