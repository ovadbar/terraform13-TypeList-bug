package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceDummy() *schema.Resource {
	return &schema.Resource{
		Create: resourceDummyCreate,
		Read:   resourceDummyRead,
		Update: resourceDummyUpdate,
		Delete: resourceDummyDelete,
		Schema: map[string]*schema.Schema{
			"test": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"testlist": {
							Type:     schema.TypeSet,
							Elem:     &schema.Schema{Type: schema.TypeString},
							Optional: true,
							Set:      schema.HashString,
						},
					},
				},
			},
			"test_list": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}

func resourceDummyCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId("test")
	return nil
}

func resourceDummyRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceDummyUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDummyDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
