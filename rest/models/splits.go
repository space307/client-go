package models

// ListSplitsParams is the set of parameters for the ListSplits method.
type ListSplitsParams struct {
	// Return the stock splits that contain this ticker.
	TickerEQ  *string `query:"ticker"`
	TickerLT  *string `query:"ticker.lt"`
	TickerLTE *string `query:"ticker.lte"`
	TickerGT  *string `query:"ticker.gt"`
	TickerGTE *string `query:"ticker.gte"`

	// Query by execution date with the format YYYY-MM-DD.
	ExecutionDateEQ  *Date `query:"execution_date"`
	ExecutionDateLT  *Date `query:"execution_date.lt"`
	ExecutionDateLTE *Date `query:"execution_date.lte"`
	ExecutionDateGT  *Date `query:"execution_date.gt"`
	ExecutionDateGTE *Date `query:"execution_date.gte"`

	// Query for reverse stock splits. A split ratio where split_from is greater than split_to represents a reverse
	// split. By default this filter is not used.
	ReverseSplit *bool `query:"reverse_split"`

	// Order results based on the sort field.
	Order *Order `query:"order"`

	// Limit the number of results returned, default is 10 and max is 1000.
	Limit *int `query:"limit"`

	// Sort field used for ordering.
	Sort *Sort `query:"sort"`
}

func (p ListSplitsParams) WithTicker(c Comparator, q string) *ListSplitsParams {
	if c == EQ {
		p.TickerEQ = &q
	} else if c == LT {
		p.TickerLT = &q
	} else if c == LTE {
		p.TickerLTE = &q
	} else if c == GT {
		p.TickerGT = &q
	} else if c == GTE {
		p.TickerGTE = &q
	}
	return &p
}

func (p ListSplitsParams) WithExecutionDate(c Comparator, q Date) *ListSplitsParams {
	if c == EQ {
		p.ExecutionDateEQ = &q
	} else if c == LT {
		p.ExecutionDateLT = &q
	} else if c == LTE {
		p.ExecutionDateLTE = &q
	} else if c == GT {
		p.ExecutionDateGT = &q
	} else if c == GTE {
		p.ExecutionDateGTE = &q
	}
	return &p
}

func (p ListSplitsParams) WithReverseSplit(q bool) *ListSplitsParams {
	p.ReverseSplit = &q
	return &p
}

func (p ListSplitsParams) WithOrder(q Order) *ListSplitsParams {
	p.Order = &q
	return &p
}

func (p ListSplitsParams) WithLimit(q int) *ListSplitsParams {
	p.Limit = &q
	return &p
}

func (p ListSplitsParams) WithSort(q Sort) *ListSplitsParams {
	p.Sort = &q
	return &p
}

// ListSplitsResponse is the response returned by the ListSplits method.
type ListSplitsResponse struct {
	BaseResponse
	Results []Split `json:"results,omitempty"`
}

// Split contains detailed information on a specified stock split.
type Split struct {
	ExecutionDate string `json:"execution_date,omitempty"`
	SplitFrom     int64  `json:"split_from,omitempty"`
	SplitTo       int64  `json:"split_to,omitempty"`
	Ticker        string `json:"ticker,omitempty"`
}
