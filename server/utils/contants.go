package utils

const (
	HTTP_BAD_REQUEST                     int64 = 400
	HTTP_UNAUTHORIZED                    int64 = 401
	HTTP_PAYMENT_REQUIRED                int64 = 402
	HTTP_FORBIDDEN                       int64 = 403
	HTTP_NOT_FOUND                       int64 = 404
	HTTP_METHOD_NOT_ALLOWED              int64 = 405
	HTTP_NOT_ACCEPTABLE                  int64 = 406
	HTTP_PROXY_AUTHENTICATION_REQUIRED   int64 = 407
	HTTP_REQUEST_TIMEOUT                 int64 = 408
	HTTP_CONFLICT                        int64 = 409
	HTTP_GONE                            int64 = 410
	HTTP_LENGTH_REQUIRED                 int64 = 411
	HTTP_PRECONDITION_FAILED             int64 = 412
	HTTP_PAYLOAD_TOO_LARGE               int64 = 413
	HTTP_URI_TOO_LONG                    int64 = 414
	HTTP_UNSUPPORTED_MEDIA_TYPE          int64 = 415
	HTTP_RANGE_NOT_SATISFIABLE           int64 = 416
	HTTP_EXPECTATION_FAILED              int64 = 417
	HTTP_TEAPOT                          int64 = 418
	HTTP_MISDIRECTED_REQUEST             int64 = 421
	HTTP_UNPROCESSABLE_ENTITY            int64 = 422
	HTTP_LOCKED                          int64 = 423
	HTTP_FAILED_DEPENDENCY               int64 = 424
	HTTP_UPGRADE_REQUIRED                int64 = 426
	HTTP_PRECONDITION_REQUIRED           int64 = 428
	HTTP_TOO_MANY_REQUESTS               int64 = 429
	HTTP_REQUEST_HEADER_FIELDS_TOO_LARGE int64 = 431
	HTTP_UNAVAILABLE_FOR_LEGAL_REASONS   int64 = 451
	HTTP_INTERNAL_SERVER_ERROR           int64 = 500
	HTTP_NOT_IMPLEMENTED                 int64 = 501
	HTTP_BAD_GATEWAY                     int64 = 502
	HTTP_SERVICE_UNAVAILABLE             int64 = 503
	HTTP_GATEWAY_TIMEOUT                 int64 = 504
	HTTP_HTTP_VERSION_NOT_SUPPORTED      int64 = 505
	HTTP_VARIANT_ALSO_NEGOTIATES         int64 = 506
	HTTP_INSUFFICIENT_STORAGE            int64 = 507
	HTTP_LOOP_DETECTED                   int64 = 508
	HTTP_NOT_EXTENDED                    int64 = 510
	HTTP_NETWORK_AUTHENTICATION_REQUIRED int64 = 511
	HTTP_OK                              int64 = 200
	HTTP_NO_CONTENT                      int64 = 204
)

const (
	FAILURE string = "Failure"
	SUCCESS string = "Success"
)

const (
	LOGIN_SUCCESS                    string = "Login Successfull"
	LOGIN_FAILED                     string = "Login Failed"
	EMAIL_EXISTS                     string = "Email is already attached to another player"
	ACCESS_DENIED                    string = "Access Denied"
	INVALID_TOKEN                    string = "Token Absent or Invalid token"
	NOT_FOUND                        string = "Record not found"
	DATA_FETCH_SUCCESS               string = "Data Fetch Successfully"
	UPGRADE_LEVEL                    string = "Upgrade your level to unlock the car"
	NOT_ENOUGH_REPAIR_PARTS          string = "Not enough repair parts"
	NOT_ENOUGH_COINS                 string = "Not enough coins"
	GARAGE_BOUGHT_SUCESS             string = "Garage bought successfully"
	GARAGE_LIST_FETCHED              string = "Garage list fetched successfully"
	FAILED_TO_UPDATE                 string = "Failed to update in database"
	GARAGE_UPGRADED                  string = "Garage upgrade successfully"
	ADD_CAR_TO_GARAGE_FAILED         string = "Unable to add car to garage"
	PARTS_CANNOT_BE_UPGRADED         string = "Part cannot be upgarded more"
	CASH_LIMIT_EXCEEDED              string = "Cash limit exceeded"
	COINS_LIMIT_EXCEEDED             string = "Coins limit exceeded"
	STATS_ERROR                      string = "Stats error"
	EMAIL_UPDATED_SUCCESS            string = "Email updated successfully"
	CAR_SOLD_SUCCESS                 string = "Car sold sucessfully"
	CAR_BOUGHT_SUCESS                string = "Car bought successfully"
	UPGRADE_SUCCESS                  string = "Part upgraged successfully"
	LICENSE_PLATE_CUSTOMIZED_SUCCESS string = "License Plate updated sucessfuly"
	INTERIOR_CUSTOMIZED_SUCCESS      string = "Interior updated sucessfuly"
	WHEELS_CUSTOMIZED_SUCCESS        string = "Wheels updated succesfully"
	COLOR_CUSTOMIZED_SUCCESS         string = "Color updated succesfully"
	CAR_REPAIR_SUCCESS               string = "Car repaired successfully"
	UNAUTHORIZED                     string = "Player not authorized"
	CAR_ADDED_SUCCESS                string = "Car added to sucessfully"
	EQUIP_CORRECT_CAR                string = "Car need to be selected first"

	CAR_ALREAY_ALLOTED    string = "Car already alotted to others arena"
	CAR_LIMIT_REACHED     string = "Car Limit reached upgarde the garage to increse the limit"
	CAR_SELECETED_SUCCESS string = "Current car selected successfully"
	CAR_ALREADY_BOUGHT    string = "Car already bought"

	CAR_REPLACED_SUCCESS string = "Car replaced sucessfully"

	WON                       string = "You WON"
	LOSE                      string = "You LOST"
	UPGRADE_REACHED_MAX_LEVEL string = "Part reached to its max level"
)

const (
	UPGRADE_POWER      int64   = 2
	UPGRADE_SHIFT_TIME float64 = 0.1
	UPGRADE_GRIP       float64 = 1.0
)
