package apiclient

import (
	"fmt" // "time"
	"net/url"
)

// Item ...
type Item struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	SubItem     AnotherItem `json:"sub_item"`
	ExampleBool bool        `json:"example_bool"`
	ExampleList []string    `json:"example_list"`
}

// GetItemsResponse ...
type GetItemsResponse struct {
	PageNumber         int    `json:"page_number"`
	PageSize           int    `json:"page_size"`
	TotalPages         int    `json:"total_pages"`
	TotalNumberOfItems int    `json:"total_number_of_items"`
	PageItems          []Item `json:"page_items"`
}

// AnotherItem Meep
type AnotherItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// GetItems : Accepts any number of ints, 1st int becomes Page Number, 2nd int becomes Page Size, Additional ints are ignored
func (c *Client) GetItems(args ...int) {
	// Not Implemented
	pageSize := 10
	pageNumber := 1
	if args != nil {
		pageNumber = args[0]
		if len(args) > 1 {
			pageSize = args[1]
		}
	}

	req, err := c.newRequest("GET", "/items", nil)
	//build querystring
	q := url.Values{}
	q.Add("page", fmt.Sprintf("%d", pageNumber))
	q.Add("page_size", fmt.Sprintf("%d", pageSize))

	req.URL.RawQuery = q.Encode()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		return
	}
	var resp GetItemsResponse
	_, err = c.do(req, &resp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.PageItems[0])
}
