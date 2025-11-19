package version

import (
"regexp"

v "github.com/hashicorp/go-version"
)

// will be replaced with the release version when using goreleaser
var version = "development"

var (
VersionRegexp = regexp.MustCompile("^" + v.VersionRegexpRaw + "$")
SemverRegexp  = regexp.MustCompile("^" + v.SemverRegexpRaw + "$")
)

// Product information
const (
ProductName    = "Barrer Security Mesh VPN"
ProductURL     = "https://github.com/BarrerSoftware/barrer-security-mesh-vpn"
UpstreamName   = "NetBird"
UpstreamURL    = "https://github.com/netbirdio/netbird"
Attribution    = "Built from NetBird Open Source Project"
)

// NetbirdVersion returns the Netbird version
// Kept for backward compatibility with original codebase
func NetbirdVersion() string {
return version
}

// BarrerVersion returns the Barrer Security Mesh VPN version
func BarrerVersion() string {
return version
}

// GetProductInfo returns product name and attribution
func GetProductInfo() (product, upstream, attribution string) {
return ProductName, UpstreamName + " (" + UpstreamURL + ")", Attribution
}
