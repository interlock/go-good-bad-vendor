Purpose
-------

This attempts to show how using the new support for `vendor` in go can lead to issues with multiple copies of the same
library being referenced.


Install
-------

`go get github.com/interlock/go-good-bad-vendor`

The Setup
---------

You will need go1.5 and need to have `GO15VENDOREXPERIMENT="1"` set in your environment.

In this contrived example, our `good` and `bad` apps can be compiled with:

`go build good_example.go` and `go build bad_example.go`

They both include `some_lib` which represents a library you may include in a real project. This library happens to include another library you also use. This is where things get tricky, I think.

Compile and run both examples from a clean checkout. They both work, yay!

Now run `./sync.sh`, which creates a vendor folder in some_lib with the good and bad lib installed. This simulates using a dependency manager like `glide`.

Now run the good and bad again. Bad will fail to compile with a message similar to this:

`./bad_example.go:18: cannot use og (type *"github.com/interlock/good-bad-vendor/vendor/github.com/interlock/some_lib/vendor/github.com/interlock/bad_lib".BadLib) as type *"github.com/interlock/good-bad-vendor/vendor/github.com/interlock/bad_lib".BadLib in argument to doBad`

Notice that the full vendor path had become part of the type for the struct, and strictly speaking, that is not the same type were expecting.

The good exampel works because interfaces, while part of a package space, are just contracts of functions to be implemented. If an underlying struct implements the interface then go is happy to pass it along. 

You can reset the example with `./reset.sh`, which removes the vendor folder from `some_lib`

What did we learn?
------------------

Some library may never be used in a way that would span between to other libaries, but if there is a chance they will it would be worth
the time to replace any publicly exposed structs with interfaces and creator functions which use an internal private struct for a default implemenation.

You get an added bonus when you do this, which is your library is likely much more mockable for testing within other projects.