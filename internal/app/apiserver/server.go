package apiserver

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/AlmasNurbayev/learn_go_crud/internal/app/store"
	"github.com/AlmasNurbayev/learn_go_crud/other"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
	config *Config
}

func newServer(store store.Store, config *Config) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
		config: config,
	}
	s.configureLogger()
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.UserCreate()).Methods("Post")
	s.router.HandleFunc("/auth", s.UserAuth()).Methods("Post")
}

func (s *server) configureLogger() error {

	// dir, err := os.Getwd()
	// if err != nil {
	// 	//log.Fatal(err)
	// }
	// fmt.Println("getwd", dir)

	//fmt.Println("caller", logpath)
	// // from Executable Directory

	// ex, _ := os.Executable()
	// fmt.Println("Executable DIR:", filepath.Dir(ex))

	// cmd := exec.Command("go", "list", "-f", "{{.ImportPath}}")
	// output, _ := cmd.CombinedOutput()
	// fmt.Println("exec:", strings.TrimSpace(string(output)))
	//fmt.Println("logPath", main.logPath)

	logpath := other.GetLogPath()
	f, err := os.OpenFile(logpath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("logfile not found " + logpath)
		panic("logfile not found " + logpath)
	} else {
		mw := io.MultiWriter(os.Stdout, f) // вывод и в консоль и в файл
		s.logger.SetOutput(mw)             // если только в файл, то f
	}

	//defer f.Close()

	//println(s.config.LogLevel)
	level, err := logrus.ParseLevel("debug") // непонятно как подгружать config
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}
