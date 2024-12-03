package listener

import (
	"fmt"

	"github.com/ink-paint/ink/service"
	"gorm.io/gorm"
)

type StartListener struct {
	db            *gorm.DB
	optionService service.OptionService
}

func NewStartListener(db *gorm.DB, optionService service.OptionService) {
	s := StartListener{
		db:            db,
		optionService: optionService,
	}
	fmt.Print(s)
}
