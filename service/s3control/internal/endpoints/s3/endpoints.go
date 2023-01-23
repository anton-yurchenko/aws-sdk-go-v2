// Code generated by smithy-go-codegen DO NOT EDIT.

package endpoints

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	endpoints "github.com/aws/aws-sdk-go-v2/internal/endpoints/v2"
	"github.com/aws/smithy-go/logging"
	"regexp"
)

// Options is the endpoint resolver configuration options
type Options struct {
	// Logger is a logging implementation that log events should be sent to.
	Logger logging.Logger

	// LogDeprecated indicates that deprecated endpoints should be logged to the
	// provided logger.
	LogDeprecated bool

	// ResolvedRegion is used to override the region to be resolved, rather then the
	// using the value passed to the ResolveEndpoint method. This value is used by the
	// SDK to translate regions like fips-us-east-1 or us-east-1-fips to an alternative
	// name. You must not set this value directly in your application.
	ResolvedRegion string

	// DisableHTTPS informs the resolver to return an endpoint that does not use the
	// HTTPS scheme.
	DisableHTTPS bool

	// UseDualStackEndpoint specifies the resolver must resolve a dual-stack endpoint.
	UseDualStackEndpoint aws.DualStackEndpointState

	// UseFIPSEndpoint specifies the resolver must resolve a FIPS endpoint.
	UseFIPSEndpoint aws.FIPSEndpointState
}

func (o Options) GetResolvedRegion() string {
	return o.ResolvedRegion
}

func (o Options) GetDisableHTTPS() bool {
	return o.DisableHTTPS
}

func (o Options) GetUseDualStackEndpoint() aws.DualStackEndpointState {
	return o.UseDualStackEndpoint
}

func (o Options) GetUseFIPSEndpoint() aws.FIPSEndpointState {
	return o.UseFIPSEndpoint
}

func transformToSharedOptions(options Options) endpoints.Options {
	return endpoints.Options{
		Logger:               options.Logger,
		LogDeprecated:        options.LogDeprecated,
		ResolvedRegion:       options.ResolvedRegion,
		DisableHTTPS:         options.DisableHTTPS,
		UseDualStackEndpoint: options.UseDualStackEndpoint,
		UseFIPSEndpoint:      options.UseFIPSEndpoint,
	}
}

// Resolver S3 endpoint resolver
type Resolver struct {
	partitions endpoints.Partitions
}

// ResolveEndpoint resolves the service endpoint for the given region and options
func (r *Resolver) ResolveEndpoint(region string, options Options) (endpoint aws.Endpoint, err error) {
	if len(region) == 0 {
		return endpoint, &aws.MissingRegionError{}
	}

	opt := transformToSharedOptions(options)
	return r.partitions.ResolveEndpoint(region, opt)
}

// New returns a new Resolver
func New() *Resolver {
	return &Resolver{
		partitions: defaultPartitions,
	}
}

var partitionRegexp = struct {
	Aws      *regexp.Regexp
	AwsCn    *regexp.Regexp
	AwsIso   *regexp.Regexp
	AwsIsoB  *regexp.Regexp
	AwsUsGov *regexp.Regexp
}{

	Aws:      regexp.MustCompile("^(us|eu|ap|sa|ca|me|af)\\-\\w+\\-\\d+$"),
	AwsCn:    regexp.MustCompile("^cn\\-\\w+\\-\\d+$"),
	AwsIso:   regexp.MustCompile("^us\\-iso\\-\\w+\\-\\d+$"),
	AwsIsoB:  regexp.MustCompile("^us\\-isob\\-\\w+\\-\\d+$"),
	AwsUsGov: regexp.MustCompile("^us\\-gov\\-\\w+\\-\\d+$"),
}

var defaultPartitions = endpoints.Partitions{
	{
		ID: "aws",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "s3.dualstack.{region}.amazonaws.com",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"s3v4"},
			},
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "s3-fips.{region}.amazonaws.com",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"s3v4"},
			},
			{
				Variant: endpoints.FIPSVariant | endpoints.DualStackVariant,
			}: {
				Hostname:          "s3-fips.dualstack.{region}.amazonaws.com",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"s3v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "s3.{region}.amazonaws.com",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"s3v4"},
			},
		},
		RegionRegex:    partitionRegexp.Aws,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			endpoints.EndpointKey{
				Region: "af-south-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "af-south-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.af-south-1.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "ap-east-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "ap-east-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.ap-east-1.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "ap-northeast-1",
			}: endpoints.Endpoint{
				Hostname:          "s3.ap-northeast-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region:  "ap-northeast-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "s3.dualstack.ap-northeast-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region: "ap-northeast-2",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "ap-northeast-2",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.ap-northeast-2.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "ap-northeast-3",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "ap-northeast-3",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.ap-northeast-3.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "ap-south-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "ap-south-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.ap-south-1.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "ap-south-2",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "ap-south-2",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.ap-south-2.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "ap-southeast-1",
			}: endpoints.Endpoint{
				Hostname:          "s3.ap-southeast-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region:  "ap-southeast-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "s3.dualstack.ap-southeast-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region: "ap-southeast-2",
			}: endpoints.Endpoint{
				Hostname:          "s3.ap-southeast-2.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region:  "ap-southeast-2",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "s3.dualstack.ap-southeast-2.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region: "ap-southeast-3",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "ap-southeast-3",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.ap-southeast-3.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "ap-southeast-4",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "ap-southeast-4",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.ap-southeast-4.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "aws-global",
			}: endpoints.Endpoint{
				Hostname:          "s3.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-1",
				},
			},
			endpoints.EndpointKey{
				Region: "ca-central-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "ca-central-1",
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname: "s3-fips.ca-central-1.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region:  "ca-central-1",
				Variant: endpoints.FIPSVariant | endpoints.DualStackVariant,
			}: {
				Hostname: "s3-fips.dualstack.ca-central-1.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region:  "ca-central-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.ca-central-1.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "eu-central-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "eu-central-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.eu-central-1.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "eu-central-2",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "eu-central-2",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.eu-central-2.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "eu-north-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "eu-north-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.eu-north-1.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "eu-south-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "eu-south-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.eu-south-1.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "eu-south-2",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "eu-south-2",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.eu-south-2.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "eu-west-1",
			}: endpoints.Endpoint{
				Hostname:          "s3.eu-west-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region:  "eu-west-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "s3.dualstack.eu-west-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region: "eu-west-2",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "eu-west-2",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.eu-west-2.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "eu-west-3",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "eu-west-3",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.eu-west-3.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "fips-ca-central-1",
			}: endpoints.Endpoint{
				Hostname: "s3-fips.ca-central-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "ca-central-1",
				},
				Deprecated: aws.TrueTernary,
			},
			endpoints.EndpointKey{
				Region: "fips-us-east-1",
			}: endpoints.Endpoint{
				Hostname: "s3-fips.us-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-1",
				},
				Deprecated: aws.TrueTernary,
			},
			endpoints.EndpointKey{
				Region: "fips-us-east-2",
			}: endpoints.Endpoint{
				Hostname: "s3-fips.us-east-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-2",
				},
				Deprecated: aws.TrueTernary,
			},
			endpoints.EndpointKey{
				Region: "fips-us-west-1",
			}: endpoints.Endpoint{
				Hostname: "s3-fips.us-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-1",
				},
				Deprecated: aws.TrueTernary,
			},
			endpoints.EndpointKey{
				Region: "fips-us-west-2",
			}: endpoints.Endpoint{
				Hostname: "s3-fips.us-west-2.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-west-2",
				},
				Deprecated: aws.TrueTernary,
			},
			endpoints.EndpointKey{
				Region: "me-central-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "me-central-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.me-central-1.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "me-south-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "me-south-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.me-south-1.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "s3-external-1",
			}: endpoints.Endpoint{
				Hostname:          "s3-external-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-1",
				},
			},
			endpoints.EndpointKey{
				Region: "sa-east-1",
			}: endpoints.Endpoint{
				Hostname:          "s3.sa-east-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region:  "sa-east-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "s3.dualstack.sa-east-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region: "us-east-1",
			}: endpoints.Endpoint{
				Hostname:          "s3.us-east-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region:  "us-east-1",
				Variant: endpoints.FIPSVariant | endpoints.DualStackVariant,
			}: {
				Hostname:          "s3-fips.dualstack.us-east-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region:  "us-east-1",
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "s3-fips.us-east-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region:  "us-east-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "s3.dualstack.us-east-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region: "us-east-2",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "us-east-2",
				Variant: endpoints.FIPSVariant | endpoints.DualStackVariant,
			}: {
				Hostname: "s3-fips.dualstack.us-east-2.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region:  "us-east-2",
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname: "s3-fips.us-east-2.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region:  "us-east-2",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.us-east-2.amazonaws.com",
			},
			endpoints.EndpointKey{
				Region: "us-west-1",
			}: endpoints.Endpoint{
				Hostname:          "s3.us-west-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region:  "us-west-1",
				Variant: endpoints.FIPSVariant | endpoints.DualStackVariant,
			}: {
				Hostname:          "s3-fips.dualstack.us-west-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region:  "us-west-1",
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "s3-fips.us-west-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region:  "us-west-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "s3.dualstack.us-west-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region: "us-west-2",
			}: endpoints.Endpoint{
				Hostname:          "s3.us-west-2.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region:  "us-west-2",
				Variant: endpoints.FIPSVariant | endpoints.DualStackVariant,
			}: {
				Hostname:          "s3-fips.dualstack.us-west-2.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region:  "us-west-2",
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "s3-fips.us-west-2.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			endpoints.EndpointKey{
				Region:  "us-west-2",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "s3.dualstack.us-west-2.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
		},
	},
	{
		ID: "aws-cn",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "s3.dualstack.{region}.amazonaws.com.cn",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"s3v4"},
			},
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "s3-fips.{region}.amazonaws.com.cn",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"s3v4"},
			},
			{
				Variant: endpoints.FIPSVariant | endpoints.DualStackVariant,
			}: {
				Hostname:          "s3-fips.{region}.api.amazonwebservices.com.cn",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"s3v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "s3.{region}.amazonaws.com.cn",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"s3v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsCn,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			endpoints.EndpointKey{
				Region: "cn-north-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "cn-north-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.cn-north-1.amazonaws.com.cn",
			},
			endpoints.EndpointKey{
				Region: "cn-northwest-1",
			}: endpoints.Endpoint{},
			endpoints.EndpointKey{
				Region:  "cn-northwest-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname: "s3.dualstack.cn-northwest-1.amazonaws.com.cn",
			},
		},
	},
	{
		ID: "aws-iso",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "s3-fips.{region}.c2s.ic.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"s3v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "s3.{region}.c2s.ic.gov",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"s3v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsIso,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			endpoints.EndpointKey{
				Region: "us-iso-east-1",
			}: endpoints.Endpoint{
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"s3v4"},
			},
			endpoints.EndpointKey{
				Region: "us-iso-west-1",
			}: endpoints.Endpoint{},
		},
	},
	{
		ID: "aws-iso-b",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "s3-fips.{region}.sc2s.sgov.gov",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"s3v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "s3.{region}.sc2s.sgov.gov",
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"s3v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsIsoB,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			endpoints.EndpointKey{
				Region: "us-isob-east-1",
			}: endpoints.Endpoint{},
		},
	},
	{
		ID: "aws-us-gov",
		Defaults: map[endpoints.DefaultKey]endpoints.Endpoint{
			{
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:          "s3.dualstack.{region}.amazonaws.com",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"s3", "s3v4"},
			},
			{
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:          "s3-fips.{region}.amazonaws.com",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"s3", "s3v4"},
			},
			{
				Variant: endpoints.FIPSVariant | endpoints.DualStackVariant,
			}: {
				Hostname:          "s3-fips.dualstack.{region}.amazonaws.com",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"s3", "s3v4"},
			},
			{
				Variant: 0,
			}: {
				Hostname:          "s3.{region}.amazonaws.com",
				Protocols:         []string{"https"},
				SignatureVersions: []string{"s3", "s3v4"},
			},
		},
		RegionRegex:    partitionRegexp.AwsUsGov,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			endpoints.EndpointKey{
				Region: "fips-us-gov-east-1",
			}: endpoints.Endpoint{
				Hostname: "s3-fips.us-gov-east-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-east-1",
				},
				Deprecated: aws.TrueTernary,
			},
			endpoints.EndpointKey{
				Region: "fips-us-gov-west-1",
			}: endpoints.Endpoint{
				Hostname: "s3-fips.us-gov-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-west-1",
				},
				Deprecated: aws.TrueTernary,
			},
			endpoints.EndpointKey{
				Region: "us-gov-east-1",
			}: endpoints.Endpoint{
				Hostname:  "s3.us-gov-east-1.amazonaws.com",
				Protocols: []string{"http", "https"},
			},
			endpoints.EndpointKey{
				Region:  "us-gov-east-1",
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:  "s3-fips.us-gov-east-1.amazonaws.com",
				Protocols: []string{"http", "https"},
			},
			endpoints.EndpointKey{
				Region:  "us-gov-east-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:  "s3.dualstack.us-gov-east-1.amazonaws.com",
				Protocols: []string{"http", "https"},
			},
			endpoints.EndpointKey{
				Region: "us-gov-west-1",
			}: endpoints.Endpoint{
				Hostname:  "s3.us-gov-west-1.amazonaws.com",
				Protocols: []string{"http", "https"},
			},
			endpoints.EndpointKey{
				Region:  "us-gov-west-1",
				Variant: endpoints.FIPSVariant,
			}: {
				Hostname:  "s3-fips.us-gov-west-1.amazonaws.com",
				Protocols: []string{"http", "https"},
			},
			endpoints.EndpointKey{
				Region:  "us-gov-west-1",
				Variant: endpoints.DualStackVariant,
			}: {
				Hostname:  "s3.dualstack.us-gov-west-1.amazonaws.com",
				Protocols: []string{"http", "https"},
			},
		},
	},
}
