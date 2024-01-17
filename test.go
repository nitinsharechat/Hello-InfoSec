package main
import "net/http"
import "io"
import "fmt"
import "crypto/md5"
import "crypto/sha256"
import "golang.org/x/crypto/blake2s"
func handler (w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, r.URL.Query().Get("data"))
}
func main () {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}
func main () {
        h_md5 := md5.New()
        h_sha := sha256.New()
        h_blake2s, _ := blake2s.New256(nil)
        io.WriteString(h_md5, "Go Language Secure Coding Practices")
        io.WriteString(h_sha, "Go Language Secure Coding Practices")
        io.WriteString(h_blake2s, "Welcome to Go Language Secure Coding Practices")
        fmt.Printf("MD5        : %x\n", h_md5.Sum(nil))
        fmt.Printf("SHA256     : %x\n", h_sha.Sum(nil))
        fmt.Printf("Blake2s-256: %x\n", h_blake2s.Sum(nil))
}

// Output:
// MD5        : ea9321d8fb0ec6623319e49a634aad92
// SHA256     : ba4939528707d791242d1af175e580c584dc0681af8be2a4604a526e864449f6
// Blake2s-256: 1d65fa02df8a149c245e5854d980b38855fd2c78f2924ace9b64e8b21b3f2f82
ctx := context.Background()
customerId := r.URL.Query().Get("id")

row, _ := db.QueryContext(ctx, query)
