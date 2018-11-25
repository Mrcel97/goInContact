<p align="center">
  <img src="https://www.devteam.space/wp-content/uploads/2017/03/gopher_head-min.png" alt="Golang_image" height="150px" />
</p>

# Go Contact Course

## Course aim:
This course will be a basic follow up of different Go introduction courses and official tutorials. All of them will be registered on the [Bibliography section](##Bibliography) as special thanks.

## Course sections:
This chapter will be updated every time I will start a new course/tutorial/test. Actually, the course structure is:

> GoInContact:
> > BumperCourse:  
> Basic introduction into Go language, this section will include different chapters following [Sentdex Tutorial](https://www.youtube.com/watch?v=G3PvTWRIhZA) on YT:  
> > 1. Introduction: Basic HelloWorld program in Go.
> > 2. Syntax
> > 3. Types
> > 4. Pointers
> > 5. WebApplication
> > 6. Structs
> > 7. Methods
> > 8. MethodsV2
> > 9. WebDevBasics
> > 10. Web processing
> > 11. Web parsing - [Modified](##Tutorials-Modifications) - [Install HTML parser](##Documentation:)


## Documentation:

An example of how to get documentation can be the 'println()' function inside fmt.
```
$ godoc fmt Println
```

Due to complete BumperCourse step 11 you will need Go HTML parser package, you can get it doing:
```
$ go get golang.org/x/net/html
```

## Tutorials-Modifications

Some tutorials have outdated steps which can make the tutorial impossible to test or learn. In this cases, I added some modifications trying to adapt the tutorial aim into an updated version of it. Hope you can take advantage of it!

1. Tutorial: BumperCourse from [Sentdex](https://twitter.com/sentdex?lang=es).  
    Modifications on step 11: 
    1.  Changed XML parsing to HTML parsing.
    2.  Changed step aim to: Parse [golang website](https://golang.org/) topbar links and show them into our webapp '/links'
    3.  Using new package: [package html](https://godoc.org/golang.org/x/net/html#example-Parse)
    4.  Removed structs: SitemapIndex and Location.
    5.  Added structs: Link.

## Go installation:
Go can be installed from their [Official Website](https://golang.org/).

## Bibliography:
- [Sentdex Go Course](https://www.youtube.com/watch?v=G3PvTWRIhZA): Course used to make my Go BumperCourse.
- [Go Styles](https://golang.org/doc/effective_go.html): Go Styles documentation.
- [Go Redirections](https://gist.github.com/hSATAC/5343225): How to redirect pages in Go.
- [Go Simple WebScraper](https://schier.co/blog/2015/04/26/a-simple-web-scraper-in-go.html): Simple webscraper idea.
- [Go HTML parsing](https://godoc.org/golang.org/x/net/html#example-Parse): HTML parsing package.
- [A Tour of Go](https://tour.golang.org/welcome/1): Go basic documentation and examples.
