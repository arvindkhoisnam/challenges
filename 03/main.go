package main

import (
	"fmt"
	"log"

	"math/rand"

	"github.com/playwright-community/playwright-go"
)

func main() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false), 
		Args: []string{
			"--disable-blink-features=AutomationControlled",  
		},
		SlowMo: playwright.Float(1000),
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

	if _, err = page.Goto("http://localhost:5173"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}

	randNumber := rand.Intn(100)
	if _,err := page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String(fmt.Sprintf("./temp/before-ss%d.jpg",randNumber)),
	}); err != nil {
		log.Fatalf("could not create screenshot: %v", err)
	}
	
	content,err := page.Locator("#content li").All()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("List of bands")
	for _,text := range content {
		band,_ := text.InnerText()
		fmt.Println(band)
	}

	var bands [3]string = [3]string{"Megadeth","Metallica","Lamb of God"}
	input := page.Locator("#text-input")
	for _,band := range bands {
		input.Fill(band)
		page.Locator("#btn").Click()
	}
	fmt.Println("---------------------------------------------")
	fmt.Println("List of updated bands")
	updatedContent,err := page.Locator("id=content").All()
	if err != nil {
		fmt.Println(err)
	}
	for _,text := range updatedContent {
		band,_ := text.InnerText()
		fmt.Println(band)
	}

	if _,err := page.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String(fmt.Sprintf("./temp/after-ss%d.jpg",randNumber)),
	}); err != nil {
		log.Fatalf("could not create screenshot: %v", err)
	}
	
	newPage, err := page.ExpectPopup(func() error {
		return page.Locator("#link").Click()
	})
	if err != nil {
		log.Fatalf("Failed to open new tab: %v", err)
	}

	iframe := newPage.FrameLocator("iframe[role='presentation']")
	err = iframe.Owner().WaitFor(playwright.LocatorWaitForOptions{
		Timeout: playwright.Float(2000),
	})
	if err != nil {
		fmt.Println("iframe not visible within timeout:", err)
	} else {
		fmt.Println("iframe is now visible!")
		signInLink := iframe.GetByRole("link")
		googlePermissionTab,_ := newPage.ExpectPopup(func() error {
			return signInLink.Click()
		})
		googlePermissionTab.Locator("#identifierId").Fill("arvindkhoisnam23@gmail.com")
		googlePermissionTab.Locator("#identifierNext").Click()
	}

	dialog := newPage.Locator("[jsname='haAclf']")
	err = dialog.WaitFor(playwright.LocatorWaitForOptions{
		Timeout: playwright.Float(2000),
	})
	if err != nil {
		fmt.Println("Dialog not visible within timeout:", err)
		} else {
			fmt.Println("Dialog is now visible!")
			signInLink := dialog.Locator("role=button[name='Sign in']")
			googlePermissionTab,_ := newPage.ExpectPopup(func() error {
				return signInLink.Click()
			})
			googlePermissionTab.Locator("#identifierId").Fill("arvindkhoisnam23@gmail.com")
			googlePermissionTab.Locator("#identifierNext").Click()
	}
	
	log.Println("Browser opened successfully!")
	fmt.Println("Press Enter to close the browser...")
    fmt.Scanln() 
}