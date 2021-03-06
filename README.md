# A pure Go implementation of Lua pattern matching

## Introduction

The package implements a subset of Lua patterns in Go. It is not an idiomatic
port, but rather uses as close to a line-by-line translation of the original C
source. This means this package incurs quite a bit of overhead by implementing
a string pointer type. I have still found the speed to be rather acceptable for
most pattern matching needs.

Currently the following function equivalents are implemented:

  * string.match as Match/MatchBytes
  * string.gmatch as Gmatch/GmatchBytes
  * string.find as Find/FindBytes
  * string.gsub as Replace/ReplaceBytes

## Installing

You can install the package using [goinstall][4]

    goinstall github.com/jnwhiteh/go-luapatterns/pattern

## Updating

If you would like to update to the latest version of luapatterns, simply use
the goinstall flag to accomplish this:

    goinstall -u=true github.com/jnwhiteh/go-luapatterns/pattern

## Using

Once you have installed the package using goinstall, you can use the package in
the following manner:

        package main
    
        import "fmt"
        import pattern "github.com/jnwhiteh/go-luapatterns/pattern"
        
        func main() {
        	str := "aaaaab"
        	pat := "(.-)(b)"
        	succ, caps := pattern.Match(str, pat)
        
        	fmt.Printf("Match('%s', '%s') => %t\n", str, pat, succ)
        	for idx, capture := range caps {
        		fmt.Printf("capture[%d] = '%s'\n", idx, capture)
        	}
        }

## Reference

The documentation for Lua patterns can be found on the [Lua Reference Manual -
Section 5.4.1][3].

## Differences

  * This implementation drops the use of the 'init' argument to the find
    function, since if you would like to start finding a pattern at a point in
    the string, you can take a sub-slice in order to do this.
  * The indices returns from Find will not match the returns from the
    equivalent Lua program due to differences in array indexing (start will be
    -1) and slices. In order to get the substring of a pattern match, you can
    take `str[startIndex:endIndex]`.
  * The `Replace` function is the string-only version of gsub. There is
    currently no function/table lookup equivalent.
  * The `Gmatch` function is not a translation of the Lua version, it is a
    simple naive implementation based on index tracking and `Find`.

## Known Issues

  * Position captures are not currently implemented, as I am unsure how to
    return those values to the caller.
  * The frontier pattern %f is not working properly

## Resources
  * [Lua 5.1 Reference Manual][2]
  * [Kahlua String Library][1]

[1]: http://github.com/krka/kahlua2/blob/master/core/src/se/krka/kahlua/stdlib/StringLib.java
[2]: http://www.lua.org/manual/5.1/manual.html 
[3]: http://www.lua.org/manual/5.1/manual.html#5.4.1
[4]: http://golang.org/cmd/goinstall/
