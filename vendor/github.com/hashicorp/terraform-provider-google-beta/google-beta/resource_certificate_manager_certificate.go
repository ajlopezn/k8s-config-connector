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
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func certManagerDefaultScopeDiffSuppress(_, old, new string, diff *schema.ResourceData) bool {
	if old == "" && new == "DEFAULT" || old == "DEFAULT" && new == "" {
		return true
	}
	return false
}

func resourceCertificateManagerCertificate() *schema.Resource {
	return &schema.Resource{
		Create: resourceCertificateManagerCertificateCreate,
		Read:   resourceCertificateManagerCertificateRead,
		Update: resourceCertificateManagerCertificateUpdate,
		Delete: resourceCertificateManagerCertificateDelete,

		Importer: &schema.ResourceImporter{
			State: resourceCertificateManagerCertificateImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `A user-defined name of the certificate. Certificate names must be unique
The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
and all following characters must be a dash, underscore, letter or digit.`,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `A human-readable description of the resource.`,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: `Set of label tags associated with the Certificate resource.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"managed": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Configuration and state of a Managed Certificate.
Certificate Manager provisions and renews Managed Certificates
automatically, for as long as it's authorized to do so.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_authorizations": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: `Authorizations that will be used for performing domain authorization`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"domains": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Description: `The domains for which a managed SSL certificate will be generated.
Wildcard domains are only supported with DNS challenge resolution`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"authorization_attempt_info": {
							Type:     schema.TypeList,
							Computed: true,
							Description: `Detailed state of the latest authorization attempt for each domain
specified for this Managed Certificate.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"details": {
										Type:     schema.TypeString,
										Computed: true,
										Description: `Human readable explanation for reaching the state. Provided to help
address the configuration issues.
Not guaranteed to be stable. For programmatic access use 'failure_reason' field.`,
									},
									"domain": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `Domain name of the authorization attempt.`,
									},
									"failure_reason": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `Reason for failure of the authorization attempt for the domain.`,
									},
									"state": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `State of the domain for managed certificate issuance.`,
									},
								},
							},
						},
						"provisioning_issue": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `Information about issues with provisioning this Managed Certificate.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"details": {
										Type:     schema.TypeString,
										Computed: true,
										Description: `Human readable explanation about the issue. Provided to help address
the configuration issues.
Not guaranteed to be stable. For programmatic access use 'reason' field.`,
									},
									"reason": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `Reason for provisioning failures.`,
									},
								},
							},
						},
						"state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `A state of this Managed Certificate.`,
						},
					},
				},
				ExactlyOneOf: []string{"self_managed", "managed"},
			},
			"scope": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: certManagerDefaultScopeDiffSuppress,
				Description: `The scope of the certificate.

DEFAULT: Certificates with default scope are served from core Google data centers.
If unsure, choose this option.

EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates,
served from non-core Google data centers.
Currently allowed only for managed certificates.`,
				Default: "DEFAULT",
			},
			"self_managed": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `Certificate data for a SelfManaged Certificate.
SelfManaged Certificates are uploaded by the user. Updating such
certificates before they expire remains the user's responsibility.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"certificate_pem": {
							Type:       schema.TypeString,
							Optional:   true,
							Deprecated: "Deprecated in favor of `pem_certificate`",
							Description: `**Deprecated** The certificate chain in PEM-encoded form.

Leaf certificate comes first, followed by intermediate ones if any.`,
							Sensitive:    true,
							ExactlyOneOf: []string{"self_managed.0.certificate_pem", "self_managed.0.pem_certificate"},
						},
						"pem_certificate": {
							Type:     schema.TypeString,
							Optional: true,
							Description: `The certificate chain in PEM-encoded form.

Leaf certificate comes first, followed by intermediate ones if any.`,
							ExactlyOneOf: []string{"self_managed.0.certificate_pem", "self_managed.0.pem_certificate"},
						},
						"pem_private_key": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  `The private key of the leaf certificate in PEM-encoded form.`,
							ExactlyOneOf: []string{"self_managed.0.private_key_pem", "self_managed.0.pem_private_key"},
						},
						"private_key_pem": {
							Type:         schema.TypeString,
							Optional:     true,
							Deprecated:   "Deprecated in favor of `pem_private_key`",
							Description:  `**Deprecated** The private key of the leaf certificate in PEM-encoded form.`,
							Sensitive:    true,
							ExactlyOneOf: []string{"self_managed.0.private_key_pem", "self_managed.0.pem_private_key"},
						},
					},
				},
				Sensitive:    true,
				ExactlyOneOf: []string{"self_managed", "managed"},
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceCertificateManagerCertificateCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandCertificateManagerCertificateDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandCertificateManagerCertificateLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	scopeProp, err := expandCertificateManagerCertificateScope(d.Get("scope"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("scope"); !isEmptyValue(reflect.ValueOf(scopeProp)) && (ok || !reflect.DeepEqual(v, scopeProp)) {
		obj["scope"] = scopeProp
	}
	selfManagedProp, err := expandCertificateManagerCertificateSelfManaged(d.Get("self_managed"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("self_managed"); !isEmptyValue(reflect.ValueOf(selfManagedProp)) && (ok || !reflect.DeepEqual(v, selfManagedProp)) {
		obj["selfManaged"] = selfManagedProp
	}
	managedProp, err := expandCertificateManagerCertificateManaged(d.Get("managed"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("managed"); !isEmptyValue(reflect.ValueOf(managedProp)) && (ok || !reflect.DeepEqual(v, managedProp)) {
		obj["managed"] = managedProp
	}

	url, err := replaceVars(d, config, "{{CertificateManagerBasePath}}projects/{{project}}/locations/global/certificates?certificateId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Certificate: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Certificate: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Certificate: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/global/certificates/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = certificateManagerOperationWaitTime(
		config, res, project, "Creating Certificate", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Certificate: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Certificate %q: %#v", d.Id(), res)

	return resourceCertificateManagerCertificateRead(d, meta)
}

func resourceCertificateManagerCertificateRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{CertificateManagerBasePath}}projects/{{project}}/locations/global/certificates/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Certificate: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("CertificateManagerCertificate %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Certificate: %s", err)
	}

	if err := d.Set("description", flattenCertificateManagerCertificateDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Certificate: %s", err)
	}
	if err := d.Set("labels", flattenCertificateManagerCertificateLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Certificate: %s", err)
	}
	if err := d.Set("scope", flattenCertificateManagerCertificateScope(res["scope"], d, config)); err != nil {
		return fmt.Errorf("Error reading Certificate: %s", err)
	}
	if err := d.Set("managed", flattenCertificateManagerCertificateManaged(res["managed"], d, config)); err != nil {
		return fmt.Errorf("Error reading Certificate: %s", err)
	}

	return nil
}

func resourceCertificateManagerCertificateUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Certificate: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	descriptionProp, err := expandCertificateManagerCertificateDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	labelsProp, err := expandCertificateManagerCertificateLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}

	url, err := replaceVars(d, config, "{{CertificateManagerBasePath}}projects/{{project}}/locations/global/certificates/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Certificate %q: %#v", d.Id(), obj)
	updateMask := []string{}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("labels") {
		updateMask = append(updateMask, "labels")
	}
	// updateMask is a URL parameter but not present in the schema, so replaceVars
	// won't set it
	url, err = addQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating Certificate %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating Certificate %q: %#v", d.Id(), res)
	}

	err = certificateManagerOperationWaitTime(
		config, res, project, "Updating Certificate", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceCertificateManagerCertificateRead(d, meta)
}

func resourceCertificateManagerCertificateDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Certificate: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{CertificateManagerBasePath}}projects/{{project}}/locations/global/certificates/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Certificate %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Certificate")
	}

	err = certificateManagerOperationWaitTime(
		config, res, project, "Deleting Certificate", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Certificate %q: %#v", d.Id(), res)
	return nil
}

func resourceCertificateManagerCertificateImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/global/certificates/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/global/certificates/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenCertificateManagerCertificateDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateScope(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateManaged(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["domains"] =
		flattenCertificateManagerCertificateManagedDomains(original["domains"], d, config)
	transformed["dns_authorizations"] =
		flattenCertificateManagerCertificateManagedDnsAuthorizations(original["dnsAuthorizations"], d, config)
	transformed["state"] =
		flattenCertificateManagerCertificateManagedState(original["state"], d, config)
	transformed["provisioning_issue"] =
		flattenCertificateManagerCertificateManagedProvisioningIssue(original["provisioningIssue"], d, config)
	transformed["authorization_attempt_info"] =
		flattenCertificateManagerCertificateManagedAuthorizationAttemptInfo(original["authorizationAttemptInfo"], d, config)
	return []interface{}{transformed}
}
func flattenCertificateManagerCertificateManagedDomains(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateManagedDnsAuthorizations(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return d.Get("managed.0.dns_authorizations")
}

func flattenCertificateManagerCertificateManagedState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateManagedProvisioningIssue(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["reason"] =
		flattenCertificateManagerCertificateManagedProvisioningIssueReason(original["reason"], d, config)
	transformed["details"] =
		flattenCertificateManagerCertificateManagedProvisioningIssueDetails(original["details"], d, config)
	return []interface{}{transformed}
}
func flattenCertificateManagerCertificateManagedProvisioningIssueReason(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateManagedProvisioningIssueDetails(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateManagedAuthorizationAttemptInfo(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"domain":         flattenCertificateManagerCertificateManagedAuthorizationAttemptInfoDomain(original["domain"], d, config),
			"state":          flattenCertificateManagerCertificateManagedAuthorizationAttemptInfoState(original["state"], d, config),
			"failure_reason": flattenCertificateManagerCertificateManagedAuthorizationAttemptInfoFailureReason(original["failureReason"], d, config),
			"details":        flattenCertificateManagerCertificateManagedAuthorizationAttemptInfoDetails(original["details"], d, config),
		})
	}
	return transformed
}
func flattenCertificateManagerCertificateManagedAuthorizationAttemptInfoDomain(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateManagedAuthorizationAttemptInfoState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateManagedAuthorizationAttemptInfoFailureReason(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenCertificateManagerCertificateManagedAuthorizationAttemptInfoDetails(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandCertificateManagerCertificateDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandCertificateManagerCertificateScope(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateSelfManaged(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedCertificatePem, err := expandCertificateManagerCertificateSelfManagedCertificatePem(original["certificate_pem"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCertificatePem); val.IsValid() && !isEmptyValue(val) {
		transformed["certificatePem"] = transformedCertificatePem
	}

	transformedPrivateKeyPem, err := expandCertificateManagerCertificateSelfManagedPrivateKeyPem(original["private_key_pem"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPrivateKeyPem); val.IsValid() && !isEmptyValue(val) {
		transformed["privateKeyPem"] = transformedPrivateKeyPem
	}

	transformedPemCertificate, err := expandCertificateManagerCertificateSelfManagedPemCertificate(original["pem_certificate"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPemCertificate); val.IsValid() && !isEmptyValue(val) {
		transformed["pemCertificate"] = transformedPemCertificate
	}

	transformedPemPrivateKey, err := expandCertificateManagerCertificateSelfManagedPemPrivateKey(original["pem_private_key"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedPemPrivateKey); val.IsValid() && !isEmptyValue(val) {
		transformed["pemPrivateKey"] = transformedPemPrivateKey
	}

	return transformed, nil
}

func expandCertificateManagerCertificateSelfManagedCertificatePem(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateSelfManagedPrivateKeyPem(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateSelfManagedPemCertificate(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateSelfManagedPemPrivateKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManaged(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDomains, err := expandCertificateManagerCertificateManagedDomains(original["domains"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDomains); val.IsValid() && !isEmptyValue(val) {
		transformed["domains"] = transformedDomains
	}

	transformedDnsAuthorizations, err := expandCertificateManagerCertificateManagedDnsAuthorizations(original["dns_authorizations"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDnsAuthorizations); val.IsValid() && !isEmptyValue(val) {
		transformed["dnsAuthorizations"] = transformedDnsAuthorizations
	}

	transformedState, err := expandCertificateManagerCertificateManagedState(original["state"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedState); val.IsValid() && !isEmptyValue(val) {
		transformed["state"] = transformedState
	}

	transformedProvisioningIssue, err := expandCertificateManagerCertificateManagedProvisioningIssue(original["provisioning_issue"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProvisioningIssue); val.IsValid() && !isEmptyValue(val) {
		transformed["provisioningIssue"] = transformedProvisioningIssue
	}

	transformedAuthorizationAttemptInfo, err := expandCertificateManagerCertificateManagedAuthorizationAttemptInfo(original["authorization_attempt_info"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedAuthorizationAttemptInfo); val.IsValid() && !isEmptyValue(val) {
		transformed["authorizationAttemptInfo"] = transformedAuthorizationAttemptInfo
	}

	return transformed, nil
}

func expandCertificateManagerCertificateManagedDomains(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedDnsAuthorizations(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedState(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedProvisioningIssue(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedReason, err := expandCertificateManagerCertificateManagedProvisioningIssueReason(original["reason"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedReason); val.IsValid() && !isEmptyValue(val) {
		transformed["reason"] = transformedReason
	}

	transformedDetails, err := expandCertificateManagerCertificateManagedProvisioningIssueDetails(original["details"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDetails); val.IsValid() && !isEmptyValue(val) {
		transformed["details"] = transformedDetails
	}

	return transformed, nil
}

func expandCertificateManagerCertificateManagedProvisioningIssueReason(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedProvisioningIssueDetails(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedAuthorizationAttemptInfo(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDomain, err := expandCertificateManagerCertificateManagedAuthorizationAttemptInfoDomain(original["domain"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDomain); val.IsValid() && !isEmptyValue(val) {
			transformed["domain"] = transformedDomain
		}

		transformedState, err := expandCertificateManagerCertificateManagedAuthorizationAttemptInfoState(original["state"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedState); val.IsValid() && !isEmptyValue(val) {
			transformed["state"] = transformedState
		}

		transformedFailureReason, err := expandCertificateManagerCertificateManagedAuthorizationAttemptInfoFailureReason(original["failure_reason"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedFailureReason); val.IsValid() && !isEmptyValue(val) {
			transformed["failureReason"] = transformedFailureReason
		}

		transformedDetails, err := expandCertificateManagerCertificateManagedAuthorizationAttemptInfoDetails(original["details"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDetails); val.IsValid() && !isEmptyValue(val) {
			transformed["details"] = transformedDetails
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandCertificateManagerCertificateManagedAuthorizationAttemptInfoDomain(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedAuthorizationAttemptInfoState(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedAuthorizationAttemptInfoFailureReason(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandCertificateManagerCertificateManagedAuthorizationAttemptInfoDetails(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
