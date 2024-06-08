package initializers

import "github.com/machinebox/graphql"

var CLIENT *graphql.Client

func ConnectToDB() {
	CLIENT = graphql.NewClient("http://localhost:8080/v1/graphql")
}
