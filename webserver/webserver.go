package webserver

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Xackery/eqemuconfig"
	"github.com/Xackery/shinshop/database"
	"github.com/Xackery/shinshop/webserver/rest"
	"github.com/Xackery/shinshop/webserver/rest/inventory"
	"github.com/Xackery/shinshop/webserver/template"
	"github.com/skratchdot/open-golang/open"
)

var isProduction = true

func Start(addr string) (err error) {

	config, err := eqemuconfig.GetConfig()
	if err != nil {
		log.Println("Error loading config:", err.Error())
		fmt.Println("Press enter to exit.")
		exit := ""
		fmt.Scan(&exit)
		os.Exit(1)
	}
	//Do a quick DB test
	_, err = database.Connect(config)
	if err != nil {
		log.Println("Error connecting to DB:", err.Error())
		fmt.Println("Press enter to exit.")
		exit := ""
		fmt.Scan(&exit)
		os.Exit(1)
	}

	template.LoadTemplates()

	http.HandleFunc("/", Index)
	http.HandleFunc("/item/", template.ItemIndex)
	http.HandleFunc("/character/", template.CharacterSearch)
	http.HandleFunc("/character/search/", template.CharacterSearch)
	http.HandleFunc("/character/inventory/", template.CharacterInventory)

	http.HandleFunc("/map/editor/", template.MapEditor)

	http.HandleFunc("/item/editor/", template.ItemEditor)

	http.HandleFunc("/rest/", rest.Index)
	http.HandleFunc("/rest/item/getbyid", rest.ItemGetById)
	http.HandleFunc("/rest/item/searchbyname", rest.ItemSearchByName)
	http.HandleFunc("/rest/inventory/getbycharacterid", rest.InventoryGetByCharacterId)
	http.HandleFunc("/rest/inventory/add", inventory.ActionAdd)
	http.HandleFunc("/rest/inventory/move", inventory.ActionMove)
	http.HandleFunc("/rest/inventory/remove", inventory.ActionRemove)
	http.HandleFunc("/rest/inventory/update", inventory.ActionUpdate)
	http.HandleFunc("/rest/map/getbyshortname/", rest.MapGetByShortname)

	//http.Handle("/", r)
	log.Println("Started Web Server on", addr)
	go openBrowser(strings.Replace(addr, "0.0.0.0", "127.0.0.1", 1))
	err = http.ListenAndServe(addr, nil)
	fmt.Println("Press enter to exit.")
	exit := ""
	fmt.Scan(&exit)
	os.Exit(1)
	return
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		if isProduction {
			http.FileServer(assetFS()).ServeHTTP(w, r)
		} else {
			http.FileServer(http.Dir("webserver/web/")).ServeHTTP(w, r)
		}
		return
	}
	template.Index(w, r)
	return
}

func openBrowser(addr string) {
	open.Run("http://" + addr + "/")
}
