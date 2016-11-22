package forms

type (
    Login struct {
        Email    string `form:"email" binding:"required"`
        Password string `form:"password" binding:"required"`
    }
)