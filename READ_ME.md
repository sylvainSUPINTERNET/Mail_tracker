
# Email tracker

 - Go version *1.10.3 (windows/amd64)*


Basic email parser for __severals URLs given__, return the list of emails found on these URLs.

## Dependencies

    $ go get github.com/anaskhan96/soup

*more infos for this package -> github.com/anaskhan96/soup*

## Run

    $ go run emailParser.go
    

## Exemple

  Error

    $ go run emailParser.go http://www.supiesland.net/ http://fr.le360.ma/

 Success

    $ go run emailParser.go http://www.supiesland.net/ http://www.thelin.net/laurent/labo/html/mailto.html

