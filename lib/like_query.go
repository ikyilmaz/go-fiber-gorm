package lib

import "gorm.io/gorm/clause"

func LikeQuery(column, query string, exact ...bool) clause.Like {
	if query == "" {
		return clause.Like{}
	}

	likeClause := clause.Like{Column: column}

	if len(exact) > 0 {
		likeClause.Value = query
	}

	likeClause.Value = "%" + query + "%"

	return likeClause
}
