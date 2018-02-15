package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/urfave/cli"
)

func cmdlistMap(c *cli.Context) error {
	return listMap(c)
}

func listMap(c *cli.Context) error {
	id := setID(c)

	urlStr := fmt.Sprintf("%s/%s", URL, id)
	data := fetchData(urlStr, "GET")

	result, err := MapAPIRespParse(data)
	errorCheck(err)

	sorted := result.CurrentCidrs
	sort.Strings(sorted)

	if !isOutputFormatSupported(output) {
		log.Fatalf("Output you provided `%s` is not supported. We support 'raw' and 'apache'\n", output)
	}

	switch output {
	case "apache":
		if c.Bool("only-addresses") {
			join := strings.Join(sorted[:], " ")
			outputStr := fmt.Sprintf("# Akamai SiteShield\nRequire ip %s", join)
			fmt.Println(outputStr)
		} else {
			log.Println("'Apache' output format can be used only to show addresses")
			log.Printf("Please run '%s list map --only-addresses --output apache %s'\n", appName, id)
		}
	case "raw":
		if c.Bool("only-addresses") {
			printJSON(sorted)
		} else {
			printJSON(result)
		}
	}

	return nil
}