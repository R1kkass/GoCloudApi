package middleware

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type client struct {
    limiter  *rate.Limiter
    lastSeen time.Time
}

var clients = make(map[string]*client)


func PerClientRateLimiter(next func(writer http.ResponseWriter, request *http.Request)) http.Handler {

    var (
        mu      sync.Mutex
    )

    go func() {
        for {
            time.Sleep(time.Minute)
            // Lock the mutex to protect this section from race conditions.
            mu.Lock()
            for ip, client := range clients {
                if time.Since(client.lastSeen) > 3*time.Minute {
                    delete(clients, ip)
                }
            }
            mu.Unlock()
        }
    }()

    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Extract the IP address from the request.
        ip, _, err := net.SplitHostPort(r.RemoteAddr)
        ip=ip+r.FormValue("hostname")+r.FormValue("pathname")+r.FormValue("stats_type")
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            return
        }
        // Lock the mutex to protect this section from race conditions.
        mu.Lock()
        if _, found := clients[ip]; !found {
            rt := rate.Every(1*time.Minute)
            clients[ip] = &client{limiter: rate.NewLimiter(rt, 5)}
        }
        clients[ip].lastSeen = time.Now()
        if !clients[ip].limiter.Allow() {
            mu.Unlock()

            return
        }
        mu.Unlock()
        next(w, r)
    })
}


func RateLimiter[T func(K), K any](w http.ResponseWriter, r *http.Request, next T, args K, count int, valueLimit string) {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	ip=ip + valueLimit
	if err != nil {
		return
	}
	if _, found := clients[ip]; !found {
		rt := rate.Every(1*time.Minute)
		clients[ip] = &client{limiter: rate.NewLimiter(rt, count)}
	}
	clients[ip].lastSeen = time.Now()

	if !clients[ip].limiter.Allow() {
		return
	}
	next(args)
}