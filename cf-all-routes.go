package main

import (
	"fmt"
	"github.com/krujos/cfcurl"
	"github.com/cloudfoundry/cli/plugin"
)

type AllRoutesPlugin struct{}

func (c *AllRoutesPlugin) Run(cliConnection plugin.CliConnection, args []string) {

	if args[0] == "all-routes" {

    c.getRoutes(cliConnection)

}

}

func (c *AllRoutesPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "all-routes",
		Version: plugin.VersionType{
			Major: 2,
			Minor: 0,
			Build: 1,
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
					Usage: "all-routes\n   cf all-routes",
				},
			},
		},
	}
}

//func (c *AllRoutesPlugin) getObjects (cliConnection) (apiURL){
func (c *AllRoutesPlugin) getRoutes(cliConnection plugin.CliConnection, args ...string) {

header:="hostname,domain_name,organization_name,space_name,app_name,path"
fmt.Println(header)

var nextURL interface{}
	nextURL = "/v2/routes"
	for nextURL != nil {

  json, _ := cfcurl.Curl(cliConnection, nextURL.(string))
  resources := toJSONArray(json["resources"])

	  for _, i := range resources {
	    res := toJSONObject(i)
	    entity := toJSONObject(res["entity"])
			host := entity["host"].(string)
			domain_url := entity["domain_url"].(string)
			space_url := entity["space_url"].(string)
			apps_url := entity["apps_url"].(string)
			path := entity["path"].(string)

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
			resources := toJSONArray(json["resources"])

			var app_name string
			for _, j := range resources {
			 res := toJSONObject(j)
			 entity := toJSONObject(res["entity"])
			 app_name = entity["name"].(string)
			}

			var record interface{}
			record = host+","+domain_name+","+organization_name+","+space_name+","+app_name+","+path
			fmt.Println(record)

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
