package valueobject

type StatusParticipant string

const (
	StatusPresent StatusParticipant = "present"
	StatusAbsent  StatusParticipant = "absent"
)

// IsValidStatusParticipant memeriksa apakah status participant valid atau tidak
func IsValidStatusParticipant(status StatusParticipant) bool {
	switch status {
	case StatusPresent, StatusAbsent:
		return true
	default:
		return false
	}
}
