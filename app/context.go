package app

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

type Context struct {
	Db *gorm.DB
}

func NewContext(db *gorm.DB, req *http.Request) *Context {
	return &Context{db}
}

type handlerFunc func(c *Context, rw http.ResponseWriter, req *http.Request) error

func (context *Context) Handler(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleError(w, r, func() error {
			/*user, err := app.authenticate(r, level)
			if err != nil {
				return err
			}*/
			return h(NewContext(context.Db, r), w, r)
		}())
	}
}
