package forms

type (
    Login struct {
        Email    string `form:"email" json:"email" binding:"required"`
        Password string `form:"password" json:"password" binding:"required"`
    }

    SignUp struct {

    }

    ForgotPassword struct {

    }

    UpdateProfile struct {

    }
)