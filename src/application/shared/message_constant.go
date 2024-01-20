package shared

import "net/http"

type StatusMessage struct {
	Code           int    `json:"code"`
	Message        string `json:"message"`
	HttpStatusCode int    `json:"http_status_code"`
}

func NewStatusMessage(code int, message string, httpStatusCode int) StatusMessage {
	return StatusMessage{Code: code, Message: message, HttpStatusCode: httpStatusCode}
}

var SUCCESS = NewStatusMessage(0, "success", http.StatusOK)
var SERVER_ERROR = NewStatusMessage(1, "Something went Wrong", http.StatusInternalServerError)
var INVALID_INPUT = NewStatusMessage(2, "Invalid Input, check your parameter", http.StatusBadRequest)
var UNAUTHORIZED = NewStatusMessage(3, "Unauthorized", http.StatusUnauthorized)
var INFO_NOTFOUND = NewStatusMessage(4, "Info Not Found", http.StatusNotFound)
var INVALID_PASSWORD = NewStatusMessage(5, "Please Try Again. Password Is Incorrect", http.StatusUnauthorized)
var INVALID_TOKEN = NewStatusMessage(6, "Invalid Token", http.StatusUnauthorized)
var MISSING_TOKEN = NewStatusMessage(7, "Missing Token", http.StatusBadRequest)
var INVALID_DEVICE_ID = NewStatusMessage(8, "Invalid Device ID", http.StatusBadRequest)
var INVALID_PLATFORM = NewStatusMessage(9, "Invalid Platform", http.StatusBadRequest)
var FAILED_LOGIN_QR = NewStatusMessage(10, "Failed Login Using QR", http.StatusUnauthorized)
var INVALID_REQUEST = NewStatusMessage(11, "Invalid Request", http.StatusBadRequest)
var INVALID_LOGIN_VID = NewStatusMessage(12, "Woops! Gonna sign in first! Only a click away and you can continue to enjoy RCTI+.", http.StatusUnauthorized)
var USER_NOTFOUND = NewStatusMessage(13, "User Not Found", http.StatusNotFound)
var EMAIL_NOTFOUND = NewStatusMessage(14, "Email Not Found", http.StatusNotFound)
var PHONE_NOTFOUND = NewStatusMessage(15, "Phone Number Not Found", http.StatusNotFound)
var ACCOUNT_NOTVERIFIED = NewStatusMessage(16, "Could not verify your account", http.StatusUnauthorized)
var ACCOUNT_ALREADY_EXIST = NewStatusMessage(17, "Please use another email/phonenumber", http.StatusConflict)
var INVALID_OTP = NewStatusMessage(18, "Invalid OTP", http.StatusUnauthorized)
var FAILED_SEND_EMAIL = NewStatusMessage(19, "Failed Send Email", http.StatusBadRequest)
var FAILED_SEND_SMS = NewStatusMessage(20, "Failed Send SMS", http.StatusBadRequest)
var EMAIL_BLACKLIST = NewStatusMessage(21, "Please use another email, this email domain is blacklisted", http.StatusBadRequest)
var PHONE_NUMBER_INCORRECT = NewStatusMessage(22, "Please Try Again Phone Number Is Incorrect", http.StatusBadRequest)
var ACCOUNT_NOT_ACTIVATED = NewStatusMessage(23, "Email/phone is not activated or not exist", http.StatusBadRequest)
var WARNING_ATTEMPTS_OTP = NewStatusMessage(0, "", http.StatusOK)
var ATTEMPTS_OTP_MAX = NewStatusMessage(25, "", http.StatusBadRequest)
var CAPTCHA_RESPONSE_EMPTY = NewStatusMessage(26, "Silahkan buka www.rctiplus.com di browser handphone anda untuk melakukan registrasi!", http.StatusBadRequest)
var INVALID_CAPTCHA = NewStatusMessage(27, "Invalid Captcha", http.StatusBadRequest)
var NICKNAME_ALREADY_EXIST = NewStatusMessage(28, "Your Nickname Has Been Taken", http.StatusConflict)
var NICKNAME_NOT_ELIGIBLE = NewStatusMessage(29, "You may change your username back after 14 days.", http.StatusConflict)
var SUCCESS_UPDATE_NICKNAME = NewStatusMessage(30, "Your Nickname Success Change", http.StatusOK)
var SUCCESS_UPDATE_PHONENUMBER = NewStatusMessage(31, "Your Phone Number Success Change", http.StatusOK)
var SUCCESS_UPDATE_EMAIL = NewStatusMessage(32, "Your Email Success Change", http.StatusOK)
var SUCCESS_UPDATE_DOB = NewStatusMessage(33, "Your Date Of Birthday Success Change", http.StatusOK)
var SUCCESS_UPDATE_DISPLAY_NAME = NewStatusMessage(34, "Your Display Name Success Change", http.StatusOK)
var SUCCESS_UPDATE_GENDER = NewStatusMessage(35, "Your Gender Success Change", http.StatusOK)
var SUCCESS_UPDATE_LOCATION = NewStatusMessage(36, "Your Location Success Change", http.StatusOK)
var ERROR_REDIS = NewStatusMessage(37, "Failed send to redis", http.StatusBadRequest)
var ERROR_MONGO = NewStatusMessage(38, "Failed get result from Mongo", http.StatusBadRequest)
var ERROR_ELK = NewStatusMessage(39, "Failed get result from Elastic", http.StatusBadRequest)
var (
	FAILED_COPY_DATA   = NewStatusMessage(3, "Failed to copy data", http.StatusUnprocessableEntity)
	FAILED_DELETE_DATA = NewStatusMessage(4, "Failed to delete data", http.StatusUnprocessableEntity)
)

func (x StatusMessage) GetMessage() string {
	return x.Message
}

func (x StatusMessage) GetCode() int {
	return x.Code
}
