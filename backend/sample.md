    // client := graphql.NewClient("http://localhost:8080/v1/graphql", nil)
    client := graphql.NewClient("http://localhost:8080/v1/graphql")

    // query := `
    // query MyQuery {
    // 	users_by_pk(id: "3337a965-4ca7-4e7e-9ee2-7b907f7e3351") {
    // 	  firstName
    // 	  lastName
    // 	}
    //   }
    // `

    secondQuery := `
    mutation {
    	insert_Tags(objects: {tagName: "fetira"}) {
    	  returning {
    		tagName
    	  }
    	}
      }

    `
    // "data": {
    // 	"insert_Tags": {
    // 	  "returning": [
    // 		{
    // 		  "tagName": "chechebsa"
    // 		}
    // 	  ]
    // 	}
    //   }

    type TagSchema struct {
    	Insert_Tags struct {
    		Returning []struct {
    			TagName string
    		}
    	}
    }

    type PostSchema struct {
    	Users_by_pk struct {
    		FirstName string
    		LastName  string
    	}
    }

    // request := graphql.NewRequest(query)
    secondRequest := graphql.NewRequest(secondQuery)

    // var response PostSchema
    var SecondRespose TagSchema
    err := client.Run(context.Background(), secondRequest, &SecondRespose)

    if err != nil {
    	fmt.Println("error", err)
    }

    fmt.Println("First name", SecondRespose.Insert_Tags.Returning[0].TagName)
