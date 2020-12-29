package zoom

import (
	"net/http"
	"time"
)

type (
	// client
	Client struct {
		Transport http.RoundTripper
		Timeout   time.Duration
		endpoint  string
		config    *Config
	}
	Config struct {
		ApiKey    string
		ApiSecret string
		Enabled   bool
	}

	requestOpts struct {
		Client         *Client
		Method         string
		URLParameters  interface{}
		Path           string
		DataParameters interface{}
		Ret            interface{}
		// HeadResponse represents responses that don't have a body
		HeadResponse bool
	}

	// error
	errorContainer struct {
		*APIError
	}

	APIError struct {
		Code    int       `json:"code"`
		Message string    `json:"message"`
		Errors  ErrorList `json:"errors,omitempty"`
	}

	Error struct {
		Field   string `json:"field"`
		Message string `json:"message"`
	}

	ErrorList []Error

	// meeting
	CreateMeetingOptions struct {
		UserID         string            `json:"-"`
		Topic          string            `json:"topic,omitempty"`
		Type           MeetingType       `json:"type,omitempty"`
		StartTime      *time.Time        `json:"start_time,omitempty"`
		Duration       int               `json:"duration,omitempty"`
		Timezone       string            `json:"timezone,omitempty"`
		Password       string            `json:"password,omitempty"`
		Agenda         string            `json:"agenda,omitempty"`
		TrackingFields TrackingFieldList `json:"tracking_fields,omitempty"`
		Settings       MeetingSettings   `json:"settings,omitempty"`
	}

	DeleteMeetingOptions struct {
		MeetingID           int64  `url:"-"`
		OccurrenceID        string `url:"occurrence_id,omitempty"`
		ScheduleForReminder bool   `url:"schedule_for_reminder,omitempty"`
	}

	Type int

	MeetingType            int
	MeetingStatus          string
	ApprovalType           int
	RegistrationType       int
	Audio                  string
	AutoRecording          string
	GlobalDialInNumberType string
	RecurrenceType         int
	MonthlyWeek            int
	WeekDay                int

	Meeting struct {
		ID       int64  `json:"id"`
		StartURL string `json:"start_url"`
		JoinURL  string `json:"join_url"`
	}

	TrackingField struct {
		Field string `json:"field"`
		Value string `json:"value"`
	}

	TrackingFieldList []TrackingField

	MeetingSettings struct {
		HostVideo                    bool                   `json:"host_video,omitempty"`
		ParticipantVideo             bool                   `json:"participant_video,omitempty"`
		ChinaMeeting                 bool                   `json:"cn_meeting,omitempty"`
		IndiaMeeting                 bool                   `json:"in_meeting,omitempty"`
		JoinBeforeHost               bool                   `json:"join_before_host,omitempty"`
		MuteUponEntry                bool                   `json:"mute_upon_entry,omitempty"`
		Watermark                    bool                   `json:"watermark,omitempty"`
		UsePMI                       bool                   `json:"use_pmi,omitempty"`
		ApprovalType                 ApprovalType           `json:"approval_type,omitempty"`
		RegistrationType             RegistrationType       `json:"registration_type,omitempty"`
		Audio                        Audio                  `json:"audio,omitempty"`
		AutoRecording                AutoRecording          `json:"auto_recording,omitempty"`
		EnforceLogin                 bool                   `json:"enforce_login,omitempty"`
		EnforceLoginDomains          string                 `json:"enforce_login_domains,omitempty"`
		AlternativeHosts             string                 `json:"alternative_hosts,omitempty"`
		CloseRegistration            bool                   `json:"close_registration,omitempty"`
		WaitingRoom                  bool                   `json:"waiting_room,omitempty"`
		GlobalDialInNumbers          GlobalDialInNumberList `json:"global_dial_in_numbers,omitempty"`
		ContactName                  string                 `json:"contact_name,omitempty"`
		ContactEmail                 string                 `json:"contact_email,omitempty"`
		RegistrantsConfirmationEmail bool                   `json:"registrants_confirmation_email,omitempty"`
	}

	GlobalDialInNumberList []GlobalDialInNumber

	GlobalDialInNumber struct {
		Country     string                 `json:"country"`
		CountryName string                 `json:"country_name"`
		City        string                 `json:"city"`
		Number      string                 `json:"number"`
		Type        GlobalDialInNumberType `json:"type"`
	}

	Recurrence struct {
		Type           RecurrenceType `json:"type"`
		RepeatInterval int            `json:"repeat_interval"`
		WeeklyDays     string         `json:"weekly_days"`
		MonthlyDay     int            `json:"monthly_day"`
		MonthlyWeek    MonthlyWeek    `json:"monthly_week"`
		MonthlyWeekDay WeekDay        `json:"monthly_week_day"`
		EndTimes       int            `json:"end_times"`
		EndDateTime    *time.Time     `json:"end_date_time"`
	}

	// user
	CreateUserOptions struct {
		Action   Action   `json:"action,omitempty"`
		UserInfo UserInfo `json:"user_info,omitempty"`
	}

	GetUserOpts struct {
		EmailOrID string         `url:"-"`
		LoginType *UserLoginType `url:"login_type,omitempty"`
	}

	DeleteUserOptions struct {
		UserID string `url:"-"`
	}

	UserLoginType int

	Action string

	UserInfo struct {
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		Email     string `json:"email,omitempty"`
		Type      Type   `json:"type,omitempty"`
	}

	User struct {
		ID string `json:"id"`
	}

	// group
	AddUserToGroupOptions struct {
		GroupID string     `url:"-"`
		Members MemberList `json:"members"`
	}

	Member struct {
		ID    string `json:"id,omitempty"`
		Email string `json:"email,omitempty"`
	}

	MemberList []Member
)
