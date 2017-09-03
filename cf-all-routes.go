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

		items := []string{}

		items = append(items,"hostname,domain name,app")

		json, _ := cfcurl.Curl(cliConnection, apiURL.(string))
		space_resources := toJSONArray(json["resources"])

		for _, spaceIntf := range space_resources {
			space := toJSONObject(spaceIntf)
			metadata := toJSONObject(space["metadata"])

			apiURL="/v2/spaces/" +metadata["guid"].(string)

			json, _ :=cfcurl.Curl(cliConnection, apiURL.(string))
			entity := toJSONObject(json["entity"])
			routes_urls := entity["routes_url"]



			json1, _ := cfcurl.Curl(cliConnection, routes_urls.(string))
			routes_resources := toJSONArray(json1["resources"])

			for _, routeIntf := range routes_resources {
				route := toJSONObject(routeIntf)
				entity := toJSONObject(route["entity"])
				host := entity["host"]
				domainguid := entity["domain_guid"]
				apps_url := entity["apps_url"]

				json3, _ := cfcurl.Curl(cliConnection, apps_url.(string))
				apps_url_resources := toJSONArray(json3["resources"])

				var app_name string
				for _, appsIntf := range apps_url_resources {
					apps := toJSONObject(appsIntf)
					entity := toJSONObject(apps["entity"])
					app_name = entity["name"].(string)
				}
				//fmt.Print(app_name)


				apiURL="/v2/domains/" +domainguid.(string)

				json, _  := cfcurl.Curl(cliConnection, apiURL.(string))
				entity2 := toJSONObject(json["entity"])
				domain := entity2["name"]
				//items = append(items)
				var record interface{}
				record = host.(string)+","+domain.(string)+","+app_name
				items = append(items,record.(string))

				//fmt.Print(host,",",domain,",",app_name,"\n")
			}

		}

		for _,i := range items {
			fmt.Println(i)
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
