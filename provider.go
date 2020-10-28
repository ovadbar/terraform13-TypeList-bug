package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"dummy_dummy": resourceDummy(),
		},
		ConfigureFunc: setMeta(),
	}
}

func setMeta() schema.ConfigureFunc {
	return schema.ConfigureFunc(func(data *schema.ResourceData) (interface{}, error) {
		return nil, nil
	})
}
