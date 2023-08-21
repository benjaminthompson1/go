package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// Response struct represents the structure of the JSON data retrieved from the URL.
type Response struct {
	DFH0XCMNOperationResponse struct {
		CAReturnCode      int    `json:"ca_return_code"`
		CAResponseMessage string `json:"ca_response_message"`
		CAInquireSingle   struct {
			CASingleItem struct {
				CASnglItemRef     int    `json:"ca_sngl_item_ref"`
				InSnglStock       int    `json:"in_sngl_stock"`
				CASnglDescription string `json:"ca_sngl_description"`
				CASnglCost        string `json:"ca_sngl_cost"`
				CASnglDepartment  int    `json:"ca_sngl_department"`
				OnSnglOrder       int    `json:"on_sngl_order"`
			} `json:"ca_single_item"`
		} `json:"ca_inquire_single"`
	} `json:"DFH0XCMNOperationResponse"`
}

func main() {
	// Define the URL from which the data will be retrieved.
	url := "http://zcee.zpdt.duckdns.org/catalogManager/items/30"

	// Use the 'curl' command to fetch data from the specified URL.
	cmd := exec.Command("curl", "-s", url)
	output, err := cmd.Output()

	// If there's an error executing the curl command, print it and exit.
	if err != nil {
		fmt.Printf("Failed to execute command: %v\n", err)
		return
	}

	// Initialize a Response struct to store the unmarshaled data.
	var res Response

	// Attempt to unmarshal the JSON data into the 'res' struct.
	if err := json.Unmarshal(output, &res); err != nil {
		fmt.Printf("Failed to unmarshal response data: %v\n", err)
		return
	}

	// Extract the relevant information from the response and print it.
	item := res.DFH0XCMNOperationResponse.CAInquireSingle.CASingleItem
	fmt.Printf("Item Reference: %d\n", item.CASnglItemRef)
	fmt.Printf("Stock: %d\n", item.InSnglStock)
	fmt.Printf("Description: %s\n", item.CASnglDescription)
	fmt.Printf("Cost: %s\n", item.CASnglCost)
	fmt.Printf("Department: %d\n", item.CASnglDepartment)
	fmt.Printf("On Order: %d\n", item.OnSnglOrder)
}
