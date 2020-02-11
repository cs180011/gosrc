package main

import (
	"flag"
	"fmt"
	"net/url"
)

// URLValue type
type URLValue struct {
	URL *url.URL
}

func (v URLValue) String() string {
	if v.URL != nil {
		return v.URL.String()
	}
	return ""
}

// Set for URLValue Needed for flag.Value interface satisfaction
func (v URLValue) Set(s string) error {
	if u, err := url.Parse(s); err != nil {
		return err
	} else {
		*v.URL = *u
	}
	return nil
}

type mystr string
type mystruct struct {
	mystr string
}

//func (v *string) Set(s string) error {
	//v = mystr(s)
//	v = s
	//v = "zslzsl"
	//return nil
//	return nil
//}

//func (v string) String() string {
	//return "Sample"
//	return v
//}

func (v *mystruct) Set(s string) error {
	//v = mystr(s)
	v.mystr = s
	//v = "zslzsl"
	//return nil
	return nil
}

func (v mystruct) String() string {
	//return "Sample"
	return v.mystr
}

func (v mystr) Set(s string) error {
	//v = mystr(s)
	//v = mystr("zslzsl")
	v = "zslzsl"
	//return nil
	return nil
}

func (v mystr) String() string {
	//return "Sample"
	return string(v)
}

//var thisstr string = "foo"
var mystrval mystr = "foo"
var mystrval2 mystruct;
//var uu = &mystrval2{}

var u = &url.URL{}
var v = "zrl"

func main() {
	fs := flag.NewFlagSet("ExampleValue", flag.ExitOnError)
	fs.Var(&URLValue{u}, "url", "URL to parse")
	//type mystr string
	//var mystrval mystr = "foo"
	//fs.Parse([]string{"url", "https://golang.org/pkg/flag/"})
	wordPtr := flag.String("url", "foo", "a string")
	zzzPtr := flag.String("zzz", "foo", "a string")
	aaaPtr := flag.String("aaa", "foo", "a string")
	//flag.Var(mystr{v}, "zrl", "a string")
	var z = mystruct{v};
	//flag.Var(&mystruct{v}, "zrl", "a string")
	//flag.Var(&thisstr, "yrl", "a string")
	flag.Var(&z, "zrl", "a string")
	//flag.Var(&mystrval, "zrl", "a string")
	flag.Parse()
	fmt.Printf("{word: %s}\n", *wordPtr)
	fmt.Printf("{zzz: %s}\n", *zzzPtr)
	fmt.Printf("{aaa: %s}\n", *aaaPtr)
	fmt.Printf("{mystrval: %s}\n", z.mystr)
	//fmt.Printf("{mystrval: %s}\n", mystrval)
	//fmt.Printf("{mystrval: %s}\n", mystruct.mystr)
	//fmt.Printf(`{scheme: %q, host: %q, path: %q}`, u.Scheme, u.Host, u.Path)
	flag.Parse()
	fmt.Printf("{word: %s}\n", *wordPtr)
	fmt.Printf("{zzz: %s}\n", *zzzPtr)
	fmt.Printf("{aaa: %s}\n", *aaaPtr)
	fmt.Printf("{mystrval: %s}\n", z.mystr)
	//fmt.Printf("{mystrval: %s}\n", mystrval)
	//fmt.Printf("{mystrval: %s}\n", mystrval)


}
