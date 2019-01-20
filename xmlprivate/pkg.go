// Package privxml exists so that structs can be publicly exported for Go's build-in XML parser. We do not want these to
// be considered part of the xlsx package's public interface, so we move them into their own package where they are less
// likely to be accessed by a client programmer.
package xmlprivate
