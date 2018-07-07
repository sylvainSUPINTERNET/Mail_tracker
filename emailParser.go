package main

import (
	"os"
	"fmt"
	"github.com/anaskhan96/soup"
	"regexp"
)


func main() {
	mailList := seekEmails()
	if(len(mailList) == 0){
		fmt.Println(" No mails found.")
		os.Exit(1);
	} else {
		for x := range mailList {
			fmt.Println(mailList[x])
		}
		os.Exit(0);
	}
}


func seekEmails() []string{

	re := regexp.MustCompile(`\bmailto:(.*)(\?.*)?\b`) // regex to validate format mailto into href

	var emailsFound []string


	for i := range os.Args {

		//ignore first item (pwd)
		if(i != 0){
			resp, err := soup.Get(os.Args[i])
			if err != nil {
				os.Exit(1)
			}
			doc := soup.HTMLParse(resp)
			links := doc.FindAll("a")

			for _, link := range links {

				if(re.MatchString(link.Attrs()["href"])){
					emailsFound = append(emailsFound,link.Text())
				}
			}
		}
	}

	return emailsFound
}
