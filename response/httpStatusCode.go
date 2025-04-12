package response

const (
	// ===== 2xxxx: Success, Validate, Common =====
	ErrCodeSuccess      = 20001
	ErrCodeBadRequest   = 20002
	ErrCodeParamInvalid = 20003

	// ===== 3xxxx: Auth & Security =====
	ErrInvalidToken     = 30001
	ErrExpiredToken     = 30002
	ErrUnauthorized     = 30003
	ErrPermissionDenied = 30004

	// ===== 4xxxx: System Errors =====
	ErrInternalServer        = 40001
	ErrDatabase              = 40002
	ErrServiceUnavailable    = 40003
	ErrTokenGenerationFailed = 40004

	// ===== 5xxxx: Business Logic - User =====
	ErrUserNotFound       = 50001
	ErrUserAlreadyExists  = 50002
	ErrInvalidCredentials = 50003
)

var msg = map[int]string{
	// ===== 2xxxx: Success, Validate, Common =====
	ErrCodeSuccess:      "Thành công",
	ErrCodeBadRequest:   "Yêu cầu không hợp lệ",
	ErrCodeParamInvalid: "Tham số không hợp lệ",

	// ===== 3xxxx: Auth & Security =====
	ErrInvalidToken:     "Token không hợp lệ",
	ErrExpiredToken:     "Token đã hết hạn",
	ErrUnauthorized:     "Chưa đăng nhập",
	ErrPermissionDenied: "Không có quyền truy cập",

	// ===== 4xxxx: System Errors =====
	ErrInternalServer:        "Lỗi hệ thống nội bộ",
	ErrDatabase:              "Lỗi truy vấn cơ sở dữ liệu",
	ErrServiceUnavailable:    "Dịch vụ hiện không sẵn sàng",
	ErrTokenGenerationFailed: "Tạo token không thành công",

	// ===== 5xxxx: Business Logic - User =====
	ErrUserNotFound:       "Không tìm thấy người dùng",
	ErrUserAlreadyExists:  "Người dùng đã tồn tại",
	ErrInvalidCredentials: "Tài khoản hoặc mật khẩu không đúng",
}
