package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/vote_demo/app"
)

// @contact.name   Vote API
// @contact.email  yllmis

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	app.Strat()

}
