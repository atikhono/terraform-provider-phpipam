package phpipam

import (
	"errors"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourcePHPIPAMFirstFreeSubnet() *schema.Resource {
	return &schema.Resource{
		Read: dataSourcePHPIPAMFirstFreeSubnetRead,
		Schema: map[string]*schema.Schema{
			"master_subnet_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"subnet_mask": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"subnet_address": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourcePHPIPAMFirstFreeSubnetRead(d *schema.ResourceData, meta interface{}) error {
	c := meta.(*ProviderPHPIPAMClient).subnetsController
	out, err := c.GetFirstFreeSubnet(d.Get("master_subnet_id").(int), d.Get("subnet_mask").(int))
	if err != nil {
		return err
	}
	if out == "" {
		return errors.New("Cannot allocate subnet")
	}

	d.SetId(out)
	d.Set("subnet_address", strings.Split(out, "/")[0])

	return nil
}
