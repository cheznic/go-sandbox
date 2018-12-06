// Package docs is an example of how to easily produce documentation for go artifacts.  To view this from the command line simply use the following command.
//    godoc github.com/[github id]/[path to package folder]/[package folder name]
//    godoc github.com/cheznic/go-programming/labs/documentation
// Another way to view is to use the inbuilt http server:
//    godoc -http=:8080
// Select the [packages] button on the landing page and navigate to your package.
package docs 

// One is a silly little function with no parameter or return value
func One() {
}

// Two is another silly function with one parameter and one return value
func Two(s string) string {
	return ""
}
