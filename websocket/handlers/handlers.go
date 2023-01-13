package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("E:\\projects\\transactions\\websocket\\html"),
	jet.InDevelopmentMode(),
)

var upgradeConn = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HomeWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	err := RenderPage(w, "index.jet", nil)
	if err != nil {
		log.Println(err)

	}
}

func RenderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {

	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}

	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
