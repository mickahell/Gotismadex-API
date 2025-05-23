package router

import (
	"net/http"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/mickahell/Gotismadex/helpers"
	"github.com/mickahell/Gotismadex/logger"
)

var router *mux.Router
var secure *mux.Router
var server *http.Server

// Define our struct
type authenticationMiddleware struct {
	tokenUsers map[string]string
}

// Initialize it somewhere
func (amw *authenticationMiddleware) Populate(usersapi []string, tokensapi []string) {
	for induser, valuser := range usersapi {
		amw.tokenUsers[os.Getenv(tokensapi[induser])] = valuser
	}
}

// Middleware function, which will be called for each request
func (amw *authenticationMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-Session-Token")

		if user, found := amw.tokenUsers[token]; found {
			// We found the token in our map
			logger.GetLogger().LogInfo("auth", "user is connected : "+user)
			// Pass down the request to the next middleware (or final handler)
			next.ServeHTTP(w, r)
		} else {
			// Write an error and stop the handler chain
			http.Error(w, "Forbidden", http.StatusForbidden)
			logger.GetLogger().LogInfo("auth", "auth failed")
		}
	})
}

func InitializeRouter() {
	// StrictSlash is true => redirect /adherents/ to /adherents
	router = mux.NewRouter().StrictSlash(true)

	// swagger
	router.Handle(helpers.Swaggerpathflag, http.FileServer(http.Dir(".")))

	// documentation for developers
	opts := middleware.SwaggerUIOpts{SpecURL: helpers.Swaggerpathflag, Path: "swagger"}
	sh := middleware.SwaggerUI(opts, nil)
	router.Handle("/swagger", sh)

	// documentation for share
	opts1 := middleware.RedocOpts{SpecURL: helpers.Swaggerpathflag, Path: "docs"}
	sh1 := middleware.Redoc(opts1, nil)
	router.Handle("/docs", sh1)

	// authentication into /auth
	amw := authenticationMiddleware{tokenUsers: make(map[string]string)}
	amw.Populate(helpers.TheAppConfig().Usersapi, helpers.TheAppConfig().Tokensapi)
	secure = router.PathPrefix("/auth").Subrouter()
	secure.Use(amw.Middleware)

	// http server
	server = &http.Server{
		Addr:    ":" + helpers.TheAppConfig().PortApi,
		Handler: router,
	}
}

func GetRouter() *mux.Router {
	return router
}

func GetSecureRouter() *mux.Router {
	return secure
}

func GetServer() *http.Server {
	return server
}
