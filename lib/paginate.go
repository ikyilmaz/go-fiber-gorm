package lib

import (
	"fiber-rest-api/utils"
	"strconv"
)

func Paginate(page string, take string) (offset, limit int, err error) {
	pageInt, err := strconv.Atoi(page)

	if err != nil {
		return 0, 0, utils.BadRequest("query 'page' must be type int")
	}

	limitInt, err := strconv.Atoi(take)

	if err != nil {
		return 0, 0, utils.BadRequest("query 'limit' must be type int")
	}

	return (pageInt * limitInt) - limitInt, limitInt, nil
}
