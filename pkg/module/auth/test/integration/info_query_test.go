package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/reynldi/modular-go-boilerplate/pkg/application"
	"github.com/reynldi/modular-go-boilerplate/pkg/datastore/postgres/seeds"
	"github.com/reynldi/modular-go-boilerplate/pkg/util"
	"github.com/stretchr/testify/require"
	"github.com/tidwall/gjson"
)

type graphQLRequest struct {
	Query     string                 `json:query`
	Variables map[string]interface{} `json:variables`
}

func TestIntegration_InfoQuery(t *testing.T) {
	util.LoadEnv()
	r := gin.Default()
	app := application.LoadApp(r)
	seedUser := seeds.SeedAuthUsers(app.Db).CreateOne()

	t.Run("Error invalid validation", func(t *testing.T) {
		q := map[string]string{
			"query": `
            { 
                infoQuery(input: {email: "mail.com"}) {
                    email
					isEmailVerified
					actionToken
                }
            }
        `,
		}

		jsonValue, _ := json.Marshal(q)
		req, _ := http.NewRequest("POST", "/auth/graph/query", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		errorMsg := gjson.Get(w.Body.String(), "errors")
		val := gjson.Get(errorMsg.Array()[0].String(), "message")
		require.Equal(t, "email must be a valid email address", val.Str)
	})

	t.Run("Error user not found", func(t *testing.T) {
		q := map[string]string{
			"query": `
            { 
                infoQuery(input: {email: "unknown@mail.com"}) {
                    email
					isEmailVerified
					actionToken
                }
            }
        `,
		}

		jsonValue, _ := json.Marshal(q)
		req, _ := http.NewRequest("POST", "/auth/graph/query", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		errorMsg := gjson.Get(w.Body.String(), "errors")
		val := gjson.Get(errorMsg.Array()[0].String(), "message")
		require.Equal(t, "User not found unknown@mail.com", val.Str)
	})

	t.Run("Success return Info Query Data", func(t *testing.T) {

		// graphql query
		q := `
		query($userEmail: String!) {
			infoQuery(input: { email: $userEmail }) {
			  isEmailVerified
			  email
			  actionToken
			}
		  }		  
		`

		// pass graphql variable
		variable := map[string]interface{}{
			"userEmail": seedUser.Email,
		}

		jsonValue, err := json.Marshal(graphQLRequest{Query: q, Variables: variable})
		if err != nil {
			panic(err)
		}

		req, _ := http.NewRequest("POST", "/auth/graph/query", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		fmt.Println(w.Body.String())
		valEmail := gjson.Get(w.Body.String(), "data.infoQuery.email")
		require.Equal(t, seedUser.Email, valEmail.Str)
	})
}
