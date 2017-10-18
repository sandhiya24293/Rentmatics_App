package Services

import (
	Db "Rentmatics_App/Common/DB/Mysql"
	Model "Rentmatics_App/Model"
	"encoding/json"
	"net/http"
)

func Userdata(w http.ResponseWriter, r *http.Request) {
	var User1 Model.User
	err := json.NewDecoder(r.Body).Decode(&User1)
	if err != nil {
		log.Error("Error - User Login", err)
	}
	Data := Db.InserUser(User1)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error - User Login", err)
	}
	cookie := &http.Cookie{
		Name:  "RentmaticsCookie",
		Value: Data.Loginid,
		Path:  "/",
	}
	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

func RentUserdata(w http.ResponseWriter, r *http.Request) {
	var Rentuser Model.RentUser
	err := json.NewDecoder(r.Body).Decode(&Rentuser)
	if err != nil {
		log.Error("Error - User Login", err)
	}
	Data := Db.InserRentUser(Rentuser)
	Senddata, err := json.Marshal(Data)

	if err != nil {
		log.Error("Error - User Login", err)
	}
	cookie := &http.Cookie{
		Name:  "RentmaticsAdminCookie",
		Value: Data.Username,
		Path:  "/",
	}
	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")

	w.Write(Senddata)

}

func RentUserlogout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "RentmaticsAdminCookie",
		Value: "NODATA",
		Path:  "/",
	}
	http.SetCookie(w, cookie)

}
func Userlogout(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:  "RentmaticsCookie",
		Value: "NODATA",
		Path:  "/",
	}
	http.SetCookie(w, cookie)

}
func Login(w http.ResponseWriter, r *http.Request) {
	var User Model.LoginUser
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		log.Error("Error - User Login", err)
	}

	Data := Db.GetUSer(User)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error - User Login", err)
	}
	if Data.Status == "Success" {

		cookie := &http.Cookie{
			Name:   "RentmaticsCookie",
			Domain: "develop.Rentmatics.com",
			Value:  Data.Loginid,
			Path:   "/",
		}
		var senddata Model.StatusResponse

		http.SetCookie(w, cookie)
		senddata.Status = "Success"
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Orgin", "*")

		w.Write(Senddata)

	} else {

		w.Write(Senddata)

	}
}

func RentLogin(w http.ResponseWriter, r *http.Request) {
	var User Model.LoginUser
	err := json.NewDecoder(r.Body).Decode(&User)
	if err != nil {
		log.Error("Error - User Login", err)
	}

	Data := Db.GetRentUSer(User)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error - User Login", err)
	}
	if Data.Role != "Invalid" {

		cookie := &http.Cookie{
			Name:   "RentmaticsAdminCookie",
			Domain: "develop.Rentmatics.com",
			Value:  Data.Username,
			Path:   "/",
		}
		var senddata Model.StatusResponse

		http.SetCookie(w, cookie)

		senddata.Status = "Success"
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Access-Control-Allow-Orgin", "*")

		w.Write(Senddata)

	} else {

		w.Write(Senddata)

	}
}

func Changepassword(w http.ResponseWriter, r *http.Request) {

	var User1 Model.Changepass
	err := json.NewDecoder(r.Body).Decode(&User1)
	if err != nil {
		log.Error("Error - Change password", err)
	}
	Data := Db.InsertChangepassword(User1)
	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error - Change password", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}

type Forgotpassword struct {
	Email string
}

func Forgot(w http.ResponseWriter, r *http.Request) {

	var Emailget Forgotpassword
	err := json.NewDecoder(r.Body).Decode(&Emailget)
	if err != nil {
		log.Error("Error - Change password", err)
	}
	Data := Db.InsertForgotspassword(Emailget.Email)

	Senddata, err := json.Marshal(Data)
	if err != nil {
		log.Error("Error - Change password", err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Access-Control-Allow-Orgin", "*")
	w.Write(Senddata)

}
