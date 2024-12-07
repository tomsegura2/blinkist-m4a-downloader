package main

import (
	"log"

	"github.com/gocolly/colly"
)

// User credentials
var username = "username"
var password = "password"

func main() {
	// Initialize Colly collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.blinkist.com"),
		colly.Async(true),
	)

	// Fetch CSRF token
	var csrfToken string
	c.OnHTML("input[name='authenticity_token']", func(e *colly.HTMLElement) {
		csrfToken = e.Attr("value")
		log.Println("Fetched CSRF Token:", csrfToken)
	})

	// Visit login page
	err := c.Visit("https://www.blinkist.com/en/nc/login/")
	if err != nil {
		log.Fatal("Error visiting login page:", err)
	}

	// Perform login
	err = c.Post("https://www.blinkist.com/en/nc/login/", map[string]string{
		"login[email]":    username,
		"login[password]": password,
		"authenticity_token": csrfToken,
	})
	if err != nil {
		log.Fatal("Login failed:", err)
	}

	// Attach response logging
	c.OnResponse(func(r *colly.Response) {
		log.Printf("Response received. Status Code: %d\n", r.StatusCode)
	})

	// Define categories
	categories := []string{
		"https://www.blinkist.com/en/nc/categories/entrepreneurship-and-small-business-en/books",
		"https://www.blinkist.com/en/nc/categories/science-en/books",
		"https://www.blinkist.com/en/nc/categories/economics-en/books",
		"https://www.blinkist.com/en/nc/categories/corporate-culture-en/books",
		"https://www.blinkist.com/en/nc/categories/money-and-investments-en/books",
		"https://www.blinkist.com/en/nc/categories/relationships-and-parenting-en/books",
		"https://www.blinkist.com/en/nc/categories/parenting-en/books",
		"https://www.blinkist.com/en/nc/categories/education-en/books",
		"https://www.blinkist.com/en/nc/categories/society-and-culture-en/books",
		"https://www.blinkist.com/en/nc/categories/politics-and-society-en/books",
		"https://www.blinkist.com/en/nc/categories/health-and-fitness-en/books",
		"https://www.blinkist.com/en/nc/categories/biography-and-history-en/books",
		"https://www.blinkist.com/en/nc/categories/management-and-leadership-en/books",
		"https://www.blinkist.com/en/nc/categories/psychology-en/books",
		"https://www.blinkist.com/en/nc/categories/technology-and-the-future-en/books",
		"https://www.blinkist.com/en/nc/categories/nature-and-environment-en/books",
		"https://www.blinkist.com/en/nc/categories/philosophy-en/books",
		"https://www.blinkist.com/en/nc/categories/career-and-success-en/books",
		"https://www.blinkist.com/en/nc/categories/marketing-and-sales-en/books",
		"https://www.blinkist.com/en/nc/categories/personal-growth-and-self-improvement-en/books",
		"https://www.blinkist.com/en/nc/categories/communication-and-social-skills-en/books",
		"https://www.blinkist.com/en/nc/categories/motivation-and-inspiration-en/books",
		"https://www.blinkist.com/en/nc/categories/productivity-and-time-management-en/books",
		"https://www.blinkist.com/en/nc/categories/mindfulness-and-happiness-en/books",
		"https://www.blinkist.com/en/nc/categories/religion-and-spirituality-en/books",
		"https://www.blinkist.com/en/nc/categories/biography-and-memoir-en/books",
		"https://www.blinkist.com/en/nc/categories/creativity-en/books",
	}

	// Scrape each category
	for _, url := range categories {
		log.Println("Visiting category:", url)
		c.Visit(url)
	}

	// Wait for all async operations to complete
	c.Wait()
}
