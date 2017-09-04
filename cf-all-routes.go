package main

import (
	"fmt"
	"github.com/krujos/cfcurl"
	"github.com/cloudfoundry/cli/plugin"
)

type AllRoutesPlugin struct{}

func (c *AllRoutesPlugin) Run(cliConnection plugin.CliConnection, args []string) {

	if args[0] == "all-routes" {

    c.getCurrentOrgAndSpace(cliConnection)

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

//func (c *AllRoutesPlugin) getObjects (cliConnection) (apiURL){
func (c *AllRoutesPlugin) getCurrentOrgAndSpace(cliConnection plugin.CliConnection, args ...string) {
  //var apiURL interface{}


  //items = append(items,"hostname,domain name,app")
var nextURL interface{}
items := []string{}
	nextURL = "/v2/routes"
	for nextURL != nil {
	//json, err := cfcurl.Curl(cliConnection, nextURL.(string))

  json, _ := cfcurl.Curl(cliConnection, nextURL.(string))
  resources := toJSONArray(json["resources"])

	  for _, spaceIntf := range resources {
	    space := toJSONObject(spaceIntf)
	    entity := toJSONObject(space["entity"])

			host := entity["host"].(string)
			domain_url := entity["domain_url"].(string)
			space_url := entity["space_url"].(string)
			apps_url := entity["apps_url"].(string)

			json, _ := cfcurl.Curl(cliConnection, domain_url)
			entity = toJSONObject(json["entity"])
			domain_name := entity["name"].(string)

			json, _ = cfcurl.Curl(cliConnection, space_url)
			entity = toJSONObject(json["entity"])
			space_name := entity["name"].(string)
			organization_url := entity["organization_url"].(string)

			json, _ = cfcurl.Curl(cliConnection, organization_url)
			entity = toJSONObject(json["entity"])
			organization_name := entity["name"].(string)

			json, _ = cfcurl.Curl(cliConnection, apps_url)
			resource := toJSONArray(json["resources"])

			var app_name string
			for _, spaceIntf := range resource {
			 space := toJSONObject(spaceIntf)
			 entity := toJSONObject(space["entity"])
			 app_name = entity["name"].(string)
			}

			var record interface{}
			record = host+","+domain_name+","+organization_name+","+space_name+","+app_name
			items = append(items,record.(string))
			for _,i := range items {
				fmt.Println(i)
			}
		}
	nextURL = json["next_url"]
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
