package main

import (
	"os"
	//"github.com/path/to/terratest"
)

// Test information
var TerraformReleasedPath = "base_ref_from_main"
var TerraformSrcPath = "../src"

// Terraform variables
var AwsAccountId = ""
var BucketName = "-" + os.Getenv("")
var CdnDomain = ""
var CloudfrontWhitelistIps = `[""]`
var ControlPlane = ""
var EnableCloudFront = ""
var IamRoles = `[""]`
var Region = ""
var ExpirationLifecycleDays = 20
var ExpirationLifecycleVersions = 5

var Vars = map[string]interface{}{
	"aws_account_id":           	 AwsAccountId,
	"bucket_name":              	 BucketName,
	"cdn_domain":               	 CdnDomain,
	"cloudfront_whitelist_ips": 	 CloudfrontWhitelistIps,
	"control_plane":            	 ControlPlane,
	"enable_cloudfront":        	 EnableCloudFront,
	"iam_roles":                	 IamRoles,
	"region":                   	 Region,
	"expiration_lifecycle_days": 	 ExpirationLifecycleDays,
	"expiration_lifecycle_versions": ExpirationLifecycleVersions,
	//"default_tags":             	 terratest.GetDefaultTags(make(map[string]string)),
}
