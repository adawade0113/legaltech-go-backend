package models

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/gorm"
)

type MeetingStatus string

const (
    MeetingStatusScheduled MeetingStatus = "scheduled"
    MeetingStatusOngoing   MeetingStatus = "ongoing"
    MeetingStatusCompleted MeetingStatus = "completed"
    MeetingStatusCancelled MeetingStatus = "cancelled"
)

type MeetingMode string

const (
    MeetingModeVideo  MeetingMode = "video"
    MeetingModeOffice MeetingMode = "office"
    MeetingModePhone  MeetingMode = "phone"
)

type Meeting struct {
    ID              uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
    UserID          uuid.UUID      `json:"user_id" gorm:"type:uuid;not null;index"`
    CaseID          *uuid.UUID     `json:"case_id" gorm:"type:uuid;index"`
    ClientID        *uuid.UUID     `json:"client_id" gorm:"type:uuid;index"`
    Title           string         `json:"title" gorm:"not null"`
    RoomID          string         `json:"room_id" gorm:"uniqueIndex;not null"`
    MeetingURL      string         `json:"meeting_url"`
    ScheduledAt     time.Time      `json:"scheduled_at" gorm:"index"`
    ReminderAt      *time.Time     `json:"reminder_at"`
    DurationMinutes int            `json:"duration_minutes" gorm:"default:30"`
    Mode            MeetingMode    `json:"mode" gorm:"default:video"`
    Participants    string         `json:"participants" gorm:"type:text"`
    Notes           string         `json:"notes"`
    Status          MeetingStatus  `json:"status" gorm:"default:scheduled;index"`
    CreatedAt       time.Time      `json:"created_at"`
    UpdatedAt       time.Time      `json:"updated_at"`
    DeletedAt       gorm.DeletedAt `json:"-" gorm:"index"`
}

type MeetingRequest struct {
    CaseID          *uuid.UUID  `json:"case_id"`
    ClientID        *uuid.UUID  `json:"client_id"`
    Title           string      `json:"title" binding:"required"`
    ScheduledAt     time.Time   `json:"scheduled_at" binding:"required"`
    ReminderAt      *time.Time  `json:"reminder_at"`
    DurationMinutes int         `json:"duration_minutes"`
    Mode            MeetingMode `json:"mode"`
    Participants    []string    `json:"participants"`
    Notes           string      `json:"notes"`
}
