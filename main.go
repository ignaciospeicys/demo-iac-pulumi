package main

import (
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Configuration data to be stored in the ConfigMap.
		configData := map[string]string{
			"key1": "value1",
			"key2": "value2",
			"key3": "value3",
		}

		// Convert configData to pulumi.StringMap
		pulumiConfigData := pulumi.StringMap{}
		for k, v := range configData {
			pulumiConfigData[k] = pulumi.String(v)
		}

		// Create a ConfigMap with the above configuration data.
		_, err := v1.NewConfigMap(ctx, "my-configmap", &v1.ConfigMapArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.String("my-configmap"),
				Namespace: pulumi.String("default"),
			},
			Data: pulumiConfigData,
		})
		if err != nil {
			return err
		}

		// Return without error.
		return nil
	})
}
