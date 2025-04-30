package main

import (
	"fmt"
	"log"

	"github.com/playwright-community/playwright-go"
)

func main(){
	fmt.Println()
	pw,err := playwright.Run()
	defer pw.Stop()

	if err != nil{
		log.Fatal(err)
	}

	browser,err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		// Headless: playwright.Bool(false),
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()

	context,_ := browser.NewContext()
	defer context.Close()

	page, err := context.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	page.Goto("https://wild-oasis.arvindkhoisnam.com")


	aboutLink := page.GetByRole("link",playwright.PageGetByRoleOptions{
		Name: "About",
	})
	
	aboutLink.Click()
	page.WaitForURL("https://wild-oasis.arvindkhoisnam.com/about")

	if err := page.Locator("h1").First().WaitFor(); err != nil {
		log.Fatalf("h1 elements not found: %v", err)
	}

	headings,err := page.Locator("h1").All()
	if err != nil {
		fmt.Println("Could not load headings")
	}

	for i,heading := range headings {
		content,err := heading.TextContent()
		if err != nil {
			log.Fatalf("could not get text content for h%d",err)
		}

		fmt.Printf("Heading %d: %s \n",i+1,content)
	}

	paragraphs,err := page.Locator("p").All()
	if err != nil {
		log.Fatalf("could not get p elements: %v", err)
	}
	for i,p := range paragraphs {
		content, err := p.TextContent()
    if err != nil {
        log.Fatalf("could not get text content for p%d",err)
    }
	fmt.Printf("Paragraph %d: %s \n",i+1,content)
	}

	fmt.Println()
	fmt.Println("---END OF FILE---")
	fmt.Println("Press Enter to close the browser...")
	fmt.Scanln()
}