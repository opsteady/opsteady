package component

import "os"

// Clean removes any left-over files
func (c *DefaultComponent) Clean() {
	if err := os.RemoveAll(c.GlobalConfig.TmpFolder); err != nil {
		c.Logger.Fatal().Err(err).Msg("failed to clean tmp dir")
	}
}
