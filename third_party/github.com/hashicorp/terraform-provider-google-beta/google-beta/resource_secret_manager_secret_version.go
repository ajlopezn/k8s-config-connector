// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"encoding/base64"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/googleapi"
)

func resourceSecretManagerSecretVersionUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	_, err := expandSecretManagerSecretVersionEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	}

	return resourceSecretManagerSecretVersionRead(d, meta)
}

func resourceSecretManagerSecretVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceSecretManagerSecretVersionCreate,
		Read:   resourceSecretManagerSecretVersionRead,
		Delete: resourceSecretManagerSecretVersionDelete,

		Importer: &schema.ResourceImporter{
			State: resourceSecretManagerSecretVersionImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Update: resourceSecretManagerSecretVersionUpdate,

		Schema: map[string]*schema.Schema{
			"secret_data": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The secret data. Must be no larger than 64KiB.`,
				Sensitive:   true,
			},

			"secret": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `Secret Manager secret resource`,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: `The current state of the SecretVersion.`,
				Default:     true,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time at which the Secret was created.`,
			},
			"destroy_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The time at which the Secret was destroyed. Only present if state is DESTROYED.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `The resource name of the SecretVersion. Format:
'projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}'`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceSecretManagerSecretVersionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	stateProp, err := expandSecretManagerSecretVersionEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("enabled"); !isEmptyValue(reflect.ValueOf(stateProp)) && (ok || !reflect.DeepEqual(v, stateProp)) {
		obj["state"] = stateProp
	}
	payloadProp, err := expandSecretManagerSecretVersionPayload(nil, d, config)
	if err != nil {
		return err
	} else if !isEmptyValue(reflect.ValueOf(payloadProp)) {
		obj["payload"] = payloadProp
	}

	url, err := replaceVars(d, config, "{{SecretManagerBasePath}}{{secret}}:addVersion")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new SecretVersion: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating SecretVersion: %s", err)
	}
	if err := d.Set("name", flattenSecretManagerSecretVersionName(res["name"], d, config)); err != nil {
		return fmt.Errorf(`Error setting computed identity field "name": %s`, err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// `name` is autogenerated from the api so needs to be set post-create
	name, ok := res["name"]
	if !ok {
		return fmt.Errorf("Create response didn't contain critical fields. Create may not have succeeded.")
	}
	if err := d.Set("name", name.(string)); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}
	d.SetId(name.(string))

	_, err = expandSecretManagerSecretVersionEnabled(d.Get("enabled"), d, config)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished creating SecretVersion %q: %#v", d.Id(), res)

	return resourceSecretManagerSecretVersionRead(d, meta)
}

func resourceSecretManagerSecretVersionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{SecretManagerBasePath}}{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("SecretManagerSecretVersion %q", d.Id()))
	}

	if err := d.Set("enabled", flattenSecretManagerSecretVersionEnabled(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecretVersion: %s", err)
	}
	if err := d.Set("name", flattenSecretManagerSecretVersionName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecretVersion: %s", err)
	}
	if err := d.Set("create_time", flattenSecretManagerSecretVersionCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecretVersion: %s", err)
	}
	if err := d.Set("destroy_time", flattenSecretManagerSecretVersionDestroyTime(res["destroyTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading SecretVersion: %s", err)
	}
	// Terraform must set the top level schema field, but since this object contains collapsed properties
	// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
	if flattenedProp := flattenSecretManagerSecretVersionPayload(res["payload"], d, config); flattenedProp != nil {
		if gerr, ok := flattenedProp.(*googleapi.Error); ok {
			return fmt.Errorf("Error reading SecretVersion: %s", gerr)
		}
		casted := flattenedProp.([]interface{})[0]
		if casted != nil {
			for k, v := range casted.(map[string]interface{}) {
				if err := d.Set(k, v); err != nil {
					return fmt.Errorf("Error setting %s: %s", k, err)
				}
			}
		}
	}

	return nil
}

func resourceSecretManagerSecretVersionDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := replaceVars(d, config, "{{SecretManagerBasePath}}{{name}}:destroy")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting SecretVersion %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "SecretVersion")
	}

	log.Printf("[DEBUG] Finished deleting SecretVersion %q: %#v", d.Id(), res)
	return nil
}

func resourceSecretManagerSecretVersionImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)

	// current import_formats can't import fields with forward slashes in their value
	if err := parseImportId([]string{"(?P<name>.+)"}, d, config); err != nil {
		return nil, err
	}

	name := d.Get("name").(string)
	secretRegex := regexp.MustCompile("(projects/.+/secrets/.+)/versions/.+$")

	parts := secretRegex.FindStringSubmatch(name)
	if len(parts) != 2 {
		panic(fmt.Sprintf("Version name doesn not fit the format `projects/{{project}}/secrets/{{secret}}/versions/{{version}}`"))
	}
	if err := d.Set("secret", parts[1]); err != nil {
		return nil, fmt.Errorf("Error setting secret: %s", err)
	}

	return []*schema.ResourceData{d}, nil
}

func flattenSecretManagerSecretVersionEnabled(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v.(string) == "ENABLED" {
		return true
	}

	return false
}

func flattenSecretManagerSecretVersionName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenSecretManagerSecretVersionCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenSecretManagerSecretVersionDestroyTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenSecretManagerSecretVersionPayload(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	transformed := make(map[string]interface{})

	// if this secret version is disabled, the api will return an error, as the value cannot be accessed, return what we have
	if d.Get("enabled").(bool) == false {
		transformed["secret_data"] = d.Get("secret_data")
		return []interface{}{transformed}
	}

	url, err := replaceVars(d, config, "{{SecretManagerBasePath}}{{name}}:access")
	if err != nil {
		return err
	}

	parts := strings.Split(d.Get("name").(string), "/")
	project := parts[1]

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	accessRes, err := sendRequest(config, "GET", project, url, userAgent, nil)
	if err != nil {
		return err
	}

	data, err := base64.StdEncoding.DecodeString(accessRes["payload"].(map[string]interface{})["data"].(string))
	if err != nil {
		return err
	}
	transformed["secret_data"] = string(data)
	return []interface{}{transformed}
}

func expandSecretManagerSecretVersionEnabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	name := d.Get("name").(string)
	if name == "" {
		return "", nil
	}

	url, err := replaceVars(d, config, "{{SecretManagerBasePath}}{{name}}")
	if err != nil {
		return nil, err
	}

	if v == true {
		url = fmt.Sprintf("%s:enable", url)
	} else {
		url = fmt.Sprintf("%s:disable", url)
	}

	parts := strings.Split(name, "/")
	project := parts[1]

	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return nil, err
	}

	_, err = sendRequest(config, "POST", project, url, userAgent, nil)
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func expandSecretManagerSecretVersionPayload(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	transformed := make(map[string]interface{})
	transformedSecretData, err := expandSecretManagerSecretVersionPayloadSecretData(d.Get("secret_data"), d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedSecretData); val.IsValid() && !isEmptyValue(val) {
		transformed["data"] = transformedSecretData
	}

	return transformed, nil
}

func expandSecretManagerSecretVersionPayloadSecretData(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	if v == nil {
		return nil, nil
	}

	return base64.StdEncoding.EncodeToString([]byte(v.(string))), nil
}