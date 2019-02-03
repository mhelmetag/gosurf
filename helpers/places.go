package helpers

import "github.com/mhelmetag/surfliner"

func Areas() ([]surfliner.Place, error) {
	c, err := surfliner.DefaultClient()
	if err != nil {
		return nil, err
	}

	as, err := c.Areas()
	if err != nil {
		return nil, err
	}

	return as, nil
}

func Regions(aID string) ([]surfliner.Place, error) {
	c, err := surfliner.DefaultClient()
	if err != nil {
		return nil, err
	}

	rs, err := c.Regions(aID)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func SubRegions(aID string, rID string) ([]surfliner.Place, error) {
	c, err := surfliner.DefaultClient()
	if err != nil {
		return nil, err
	}

	rs, err := c.SubRegions(aID, rID)
	if err != nil {
		return nil, err
	}

	return rs, nil
}

func ValidateSubRegion(aID string, rID string, srID string) error {
	c, err := surfliner.DefaultClient()
	if err != nil {
		return err
	}

	_, err = c.SubRegion(aID, rID, srID)

	return err
}
