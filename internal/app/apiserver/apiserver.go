package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/AlmasNurbayev/learn_go_crud/internal/app/store/sqlstore"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	srv := newServer(store, config)

	srv.logger.Info("starting server", config.BindAddr)
	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// type APIServer struct {
// 	config *Config
// 	logger *logrus.Logger
// 	router *mux.Router
// 	store  *sqlstore.Store
// }

// func New(config *Config) *APIServer {
// 	return &APIServer{
// 		config: config,
// 		logger: logrus.New(),
// 		router: mux.NewRouter(),
// 	}
// }

// func (s *APIServer) Start() error {

// 	if err := s.configureLogger(); err != nil {
// 		return err
// 	}

// 	if err := s.configureStore(); err != nil {
// 		return err
// 	}

// 	s.configureRouter()
// 	s.logger.Info("starting apiserver: ", s.config.BindAddr)

// 	return http.ListenAndServe(s.config.BindAddr, s.router)
// }

// func (s *APIServer) configureStore() error {
// 	st := sqlstore.New(s.config.Store)
// 	if err := st.Open(); err != nil {
// 		return err
// 	}
// 	s.store = st
// 	return nil
// }

// func (s *APIServer) configureRouter() {
// 	s.router.HandleFunc("/hello", s.handleHello())
// }

// func (s *APIServer) handleHello() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		io.WriteString(w, "Hello")
// 	}
// }
