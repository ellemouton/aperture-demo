package content

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/lightninglabs/aperture/pricesrpc"
)

var _ pricesrpc.PricesServer = (*Server)(nil)

func (s *Server) GetPrice(ctx context.Context,
	req *pricesrpc.GetPriceRequest) (*pricesrpc.GetPriceResponse, error) {

	if !strings.Contains(req.Path, "quote") {
		return nil, fmt.Errorf("no prices " +
			"for given path")
	}

	id, err := strconv.Atoi(strings.TrimLeft(req.Path, "/quote/"))
	if err != nil {
		return nil, fmt.Errorf("could not extract id " +
			"from the given path")
	}

	q, err := s.DB.GetQuote(id)
	if err != nil {
		return nil, err
	}

	return &pricesrpc.GetPriceResponse{
		Price: q.Price,
	}, nil
}
