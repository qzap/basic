   package main

      import (
     	"flag"
	"log"
   )

    var name = flag.String("name", "stranger", "your wonderful name")

    func main() {
flag.Parse()
    	log.Printf("Hello %s, Selamat datang di dunia perintah baris", *name)
}
