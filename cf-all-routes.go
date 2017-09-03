package main

import (
	"fmt"
	"github.com/krujos/cfcurl"
	"github.com/cloudfoundry/cli/plugin"
)

type AllRoutesPlugin struct{}

func (c *AllRoutesPlugin) Run(cliConnection plugin.CliConnection, args []string) {

	if args[0] == "all-routes" {

		var apiURL interface{}
		apiURL = "/v2/spaces"

		json, _ := cfcurl.Curl(cliConnection, apiURL.(string))
		resources := toJSONArray(json["resources"])

		for _, spaceIntf := range resources {
			space := toJSONObject(spaceIntf)
			metadata := toJSONObject(space["metadata"])

			apiURL="/v2/spaces/" +metadata["guid"].(string)

			json, _ :=cfcurl.Curl(cliConnection, apiURL.(string))
			entity := toJSONObject(json["entity"])
			routes_urls := entity["routes_url"]


			json1, _ := cfcurl.Curl(cliConnection, routes_urls.(string))
			resources := toJSONArray(json1["resources"])
			for _, routeIntf := range resources {
				route := toJSONObject(routeIntf)
				entity := toJSONObject(route["entity"])
				host := entity["host"]
				domainguid := entity["domain_guid"]


				apiURL="/v2/domains/" +domainguid.(string)

				json, _  := cfcurl.Curl(cliConnection, apiURL.(string))
				entity2 := toJSONObject(json["entity"])
				domain := entity2["name"]

				fmt.Print(host,".",domain,"\n")
			}

		}
}

}
func (c *AllRoutesPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "all-routes",
		Version: plugin.VersionType{
			Major: 1,
			Minor: 0,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "all-routes",
				HelpText: "cf all-routes",

				UsageDetails: plugin.Usage{
					Usage: "all-routesd\n   cf all-routes",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(AllRoutesPlugin))
}

func toJSONArray(obj interface{}) []interface{} {
	return obj.([]interface{})
}

func toJSONObject(obj interface{}) map[string]interface{} {
	return obj.(map[string]interface{})
}
