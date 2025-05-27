package main

import (
    "context"
    "errors"
    "log"
    "net/http"
    "os"
    "time"

    "github.com/99designs/gqlgen/graphql"
    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/handler/extension"
    "github.com/99designs/gqlgen/graphql/handler/transport"
    "github.com/99designs/gqlgen/graphql/playground"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/cors"
    "yourapp/graph"
    "yourapp/graph/generated"
)

const defaultPort = "8080"

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = defaultPort
    }

    router := chi.NewRouter()

    router.Use(cors.Handler(cors.Options{
        AllowedOrigins:   []string{"https://yourfrontend.com"},
        AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
        AllowCredentials: true,
        MaxAge:           300,
    }))

    router.Use(loggingMiddleware)
    router.Use(authMiddleware)

    srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
    srv.AddTransport(transport.POST{})
    srv.AddTransport(transport.Options{})
    srv.Use(extension.AutomaticPersistedQuery{})

    srv.SetErrorPresenter(func(ctx context.Context, err error) *graphql.Error {
        log.Printf("GraphQL error: %v", err)
        return graphql.DefaultErrorPresenter(ctx, errors.New("Internal server error"))
    })

    srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
        log.Printf("GraphQL panic: %+v", err)
        return errors.New("Unexpected server error")
    })

    router.Group(func(r chi.Router) {
        r.With(timeoutMiddleware(10 * time.Second)).Post("/query", srv.ServeHTTP)
    })

    if os.Getenv("ENV") != "production" {
        router.Get("/", playground.Handler("GraphQL playground", "/query"))
    }

    srvAddr := ":" + port
    log.Printf("🚀 Server ready at http://localhost%s", srvAddr)

    server := &http.Server{
        Addr:              srvAddr,
        Handler:           router,
        ReadTimeout:       15 * time.Second,
        WriteTimeout:      15 * time.Second,
        IdleTimeout:       60 * time.Second,
        ReadHeaderTimeout: 5 * time.Second,
    }

    log.Fatal(server.ListenAndServe())
}

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("[%s] %s %s", r.RemoteAddr, r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func timeoutMiddleware(duration time.Duration) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.TimeoutHandler(next, duration, "Request timed out")
    }
}
