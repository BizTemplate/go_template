package core

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"viper_contract/src/common/utils"
	"viper_contract/src/infrastructure"
)

type RoutingStarter struct {
	config infrastructure.Config
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func (r *RoutingStarter) Start() {
	port := r.config.ServerConfig.Port
	valueUtils := utils.ValueUtils{}
	port = valueUtils.GetOrDefault(port, "5555")
	log.Println("Start HTTP Starter at Port ", port)
	http.HandleFunc("/", getRoot)

	http.ListenAndServe(":"+port, nil)
	r.initData()
}

func (r *RoutingStarter) Reload() {
	fmt.Printf("reload")

}

func (r *RoutingStarter) Ping() {
	fmt.Printf("ping")

}

func (r *RoutingStarter) initData() {
	fmt.Printf("init")
}
