package main

import "fmt"

func (a *analytics) handleEmailBounce(em email) error {
	errStatus := em.recipient.updateStatus(em.status)

	if errStatus != nil {
		return fmt.Errorf("error updating user status: %w", errStatus)
	}

	errTrack := a.track(em.status)

	if errTrack != nil {
		return fmt.Errorf("error tracking user bounce: %w", errTrack)
	}

	return nil
}
