package request

import "time"

type NewRequest struct {
	Subject string    `form:"subject" binding:"required,min=2"`
	From    time.Time `form:"from" binding:"required,datetime" time_format:"2006-01-02"`
}
