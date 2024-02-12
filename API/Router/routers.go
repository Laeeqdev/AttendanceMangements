package router

import (
	auth "github.com/Laeeqdev/AttendanceMangements/API/Auth"
	resthandler "github.com/Laeeqdev/AttendanceMangements/API/RestHandler"
	"github.com/gorilla/mux"
)

type Router interface {
	MyRouter() *mux.Router
}
type RouterImpl struct {
	userAuthHandler        auth.UserAuthHandler
	detailsHandler         resthandler.DetailsHandler
	principalHandler       resthandler.PrincipalHandler
	punchInPunchOutHandler resthandler.PunchInPunchOutHandler
}

func NewRouterImpl(userAuthHandler auth.UserAuthHandler,
	detailsHandler resthandler.DetailsHandler,
	principalHandler resthandler.PrincipalHandler,
	punchInPunchOutHandler resthandler.PunchInPunchOutHandler) *RouterImpl {
	return &RouterImpl{userAuthHandler: userAuthHandler, detailsHandler: detailsHandler, principalHandler: principalHandler, punchInPunchOutHandler: punchInPunchOutHandler}
}

var version string = "/v1"

func (impl *RouterImpl) MyRouter() *mux.Router {
	r1 := mux.NewRouter()
	r := r1.PathPrefix(version).Subrouter()
	r.HandleFunc("/login", impl.userAuthHandler.Login).Methods("POST")
	r.HandleFunc("/home", impl.userAuthHandler.Home).Methods("GET")
	r.HandleFunc("/adduser", impl.principalHandler.AddUser).Methods("POST")
	r.HandleFunc("/logout", impl.userAuthHandler.Logout).Methods("POST")
	r.HandleFunc("/punchin", impl.punchInPunchOutHandler.PunchInUser).Methods("POST")
	r.HandleFunc("/punchout", impl.punchInPunchOutHandler.PunchOutUser).Methods("POST")
	r.HandleFunc("/getteacherdetails", impl.detailsHandler.GetTeacherDetails).Methods("POST")
	r.HandleFunc("/refresh", impl.userAuthHandler.Refresh).Methods("GET")
	r.HandleFunc("/getstudentdetails", impl.detailsHandler.GetStudentDetails).Methods("POST")
	r.HandleFunc("/getstudentdetailsbyclass", impl.detailsHandler.GetStudentDetailsByClass).Methods("POST")
	r.HandleFunc("/getDetails", impl.detailsHandler.GetDetails).Methods("POST")
	return r
}
