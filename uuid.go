package utils

import "github.com/google/uuid"

func NullUUIDToP(id uuid.NullUUID) *uuid.UUID {
	if id.Valid {
		return &id.UUID
	}
	return nil
}

func PToNullUUID(id *uuid.UUID) uuid.NullUUID {
	if id != nil {
		return uuid.NullUUID{
			UUID:  *id,
			Valid: true,
		}
	}
	return uuid.NullUUID{}
}
