package main

import (
	"io"
	"log"
	"math/big"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/semrush/zenrpc"
	"go.pyspa.org/brbundle"
	"go.pyspa.org/brbundle/brchi"
)

//go:generate go run github.com/semrush/zenrpc/zenrpc
//go:generate go run go.pyspa.org/brbundle/cmd/brbundle embedded dist

// APIService is a JSON RPC server interface
type APIService struct {
	zenrpc.Service
}

// CheckPrimeNumber returns input value is prime number or not
func (as APIService) CheckPrimeNumber(a int) bool {
	return big.NewInt(int64(a)).ProbablyPrime(0)
}

func main() {
	r := chi.NewRouter()
	rpc := zenrpc.NewServer(zenrpc.Options{
		TargetURL: "/api",
		ExposeSMD: true,
	})
	rpc.Register("", &APIService{})
	r.Handle("/api", rpc)
	r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello")
	})
	r.NotFound(brchi.Mount(brbundle.WebOption{
		SPAFallback: "index.html",
	}))
	log.Println("Open Server at :8888")
	log.Fatal(http.ListenAndServe("127.0.0.1:8888", r))
}
